package xunit

import "reflect"

func Run(i interface{}) {
	setupper, ok := i.(SetUpper)
	if ok {
		setupper.SetUp()
	}

	n := reflect.ValueOf(i).Elem().FieldByName("Name").Interface()
	name, _ := n.(string)
	method := reflect.ValueOf(i).MethodByName(name)
	method.Call([]reflect.Value{})
}

type SetUpper interface {
	SetUp()
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
