package main

import "github.com/akito0107/xunit"

type TestCaseTest struct {
	Name string
	test *xunit.WasRun
}

func (t *TestCaseTest) SetUp() {
	t.test = xunit.NewWasRun("TestMethod")
}

func (t *TestCaseTest) TestTemplateMethod() {
	xunit.Run(t.test)
	xunit.Assert("setUp testMethod " == t.test.Log)
}

func main() {
	xunit.Run(&TestCaseTest{Name: "TestTemplateMethod"})
}
