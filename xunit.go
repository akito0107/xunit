package xunit

import "reflect"

type TestMethod func()

type TestCase struct {
}

type WasRun struct {
	TestCase
	name   string
	wasRun bool
}

func (w *WasRun) TestMethod() {
	w.wasRun = true
}

func (w *WasRun) Run() {
	method := reflect.ValueOf(w).MethodByName(w.name)
	method.Call([]reflect.Value{})
}
