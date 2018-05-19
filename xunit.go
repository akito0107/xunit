package xunit

import "reflect"

type TestMethod func()

func Run(t TestCase) {
	method := reflect.ValueOf(t).MethodByName(t.Name())
	method.Call([]reflect.Value{})
}

type TestCase interface {
	Name() string
}

type WasRun struct {
	name   string
	wasRun bool
}

func NewWasRun(name string) *WasRun {
	return &WasRun{
		name: name,
	}
}

func (w *WasRun) TestMethod() {
	w.wasRun = true
}

func (w *WasRun) Name() string {
	return w.name
}
