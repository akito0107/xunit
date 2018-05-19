package xunit

type WasRun struct {
	name   string
	wasRun bool
}

func (w *WasRun) TestMethod() {
	w.wasRun = true
}

func (w *WasRun) Run() {
	w.TestMethod()
}
