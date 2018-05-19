package main

import "github.com/akito0107/xunit"

type TestCaseTest struct {
	Name string
}

func (t *TestCaseTest) TestRunning() {
	test := xunit.NewWasRun("TestMethod")
	xunit.AssertNot(test.WasRun)
	xunit.Run(test)
	xunit.Assert(test.WasRun)
}

func main() {
	t1 := &TestCaseTest{"TestRunning"}
	xunit.Run(t1)
}
