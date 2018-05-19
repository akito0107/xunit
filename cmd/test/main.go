package main

import "github.com/akito0107/xunit"

type TestCaseTest struct {
	Name string
	test *xunit.WasRun
}

func (t *TestCaseTest) SetUp() {
	t.test = xunit.NewWasRun("TestMethod")
}

func (t *TestCaseTest) TestRunning() {
	xunit.Run(t.test)
	xunit.Assert(t.test.WasRun)
}

func (t *TestCaseTest) TestSetUp() {
	xunit.Run(t.test)
	xunit.Assert(t.test.WasSetUp)
}

func main() {
	xunit.Run(&TestCaseTest{Name: "TestRunning"})
	xunit.Run(&TestCaseTest{Name: "TestSetUp"})
}
