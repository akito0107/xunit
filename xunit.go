package xunit

import "reflect"

func Run(t TestCase) {
	method := reflect.ValueOf(t).MethodByName(t.Name())
	method.Call([]reflect.Value{})
}

type TestCase interface {
	Name() string
}

type WasRun struct {
	name   string
	WasRun bool
}

func NewWasRun(name string) *WasRun {
	return &WasRun{
		name:   name,
		WasRun: false,
	}
}

func (w *WasRun) TestMethod() {
	w.WasRun = true
}

func (w *WasRun) Name() string {
	return w.name
}
