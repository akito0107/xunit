package xunit

import "reflect"

func Run(i interface{}) {
	n := reflect.ValueOf(i).Elem().FieldByName("Name").Interface()
	name, _ := n.(string)
	method := reflect.ValueOf(i).MethodByName(name)
	method.Call([]reflect.Value{})
}

type TestCase interface {
	Name() string
}

type WasRun struct {
	Name   string
	WasRun bool
}

func NewWasRun(name string) *WasRun {
	return &WasRun{
		Name:   name,
		WasRun: false,
	}
}

func (w *WasRun) TestMethod() {
	w.WasRun = true
}
