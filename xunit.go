package xunit

import "reflect"

type TestMethod func()

type TestCase struct {
	name string
}

func (t *TestCase) Run() {
	method := reflect.ValueOf(t).MethodByName(t.name)
	method.Call([]reflect.Value{})
}

type WasRun struct {
	TestCase
	wasRun bool
}

func NewWasRun(name string) *WasRun {
	return &WasRun{
		TestCase: TestCase{name: name},
	}
}

func (w *WasRun) TestMethod() {
	w.wasRun = true
}
