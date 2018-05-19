package main

import "github.com/akito0107/xunit"

type TestCaseTest struct {
	name string
}

func (t *TestCaseTest) Name() string {
	return t.name
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
