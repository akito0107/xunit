package xunit

import (
	"fmt"
	"reflect"
)

func Run(i interface{}, result *TestResult) *TestResult {
	switch v := i.(type) {
	case *TestSuite:
		for _, t := range v.Tests {
			result = run(t, result)
		}
	default:
		return run(i, result)
	}
	return result
}

func run(i interface{}, res *TestResult) *TestResult {
	res.TestStarted()

	setupper, ok := i.(SetUpper)
	if ok {
		setupper.SetUp()
	}

	n := reflect.ValueOf(i).Elem().FieldByName("Name").Interface()
	name, _ := n.(string)

	func() {
		defer func() {
			if err := recover(); err != nil {
				res.TestFailed()
			}
		}()
		method := reflect.ValueOf(i).MethodByName(name)
		method.Call([]reflect.Value{})
	}()

	downer, ok := i.(TearDowner)
	if ok {
		downer.TearDown()
	}

	return res
}

type SetUpper interface {
	SetUp()
}

type TearDowner interface {
	TearDown()
}

type WasRun struct {
	Name string
	Log  string
}

func NewWasRun(name string) *WasRun {
	return &WasRun{
		Name: name,
	}
}

func (w *WasRun) SetUp() {
	w.Log = "setUp "
}

func (w *WasRun) TestMethod() {
	w.Log = w.Log + "testMethod "
}

func (w *WasRun) TestBrokenMethod() {
	panic("test")
}

func (w *WasRun) TearDown() {
	w.Log = w.Log + "tearDown "
}

var _ SetUpper = &WasRun{}
var _ TearDowner = &WasRun{}

type TestResult struct {
	RunCount   int
	ErrorCount int
}

func NewTestResult() *TestResult {
	return &TestResult{
		RunCount:   0,
		ErrorCount: 0,
	}
}

func (t *TestResult) TestStarted() {
	t.RunCount += 1
}

func (t *TestResult) TestFailed() {
	t.ErrorCount += 1
}

func (t *TestResult) Summary() string {
	return fmt.Sprintf("%d run, %d failed", t.RunCount, t.ErrorCount)
}

type TestSuite struct {
	Tests []interface{}
}

func (t *TestSuite) Add(i interface{}) {
	t.Tests = append(t.Tests, i)
}
