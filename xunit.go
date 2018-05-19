package xunit

type WasRun struct {
	name   string
	wasRun bool
}

func (w *WasRun) testMethod() {
	w.wasRun = true
}
