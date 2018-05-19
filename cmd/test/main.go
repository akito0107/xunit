package main

import "github.com/akito0107/xunit"

type TestCaseTest struct {
	Name string
	test *xunit.WasRun
}

func (t *TestCaseTest) TestTemplateMethod() {
	test := xunit.NewWasRun("TestMethod")
	xunit.Run(test)
	xunit.Assert("setUp testMethod " == test.Log)
}

func main() {
	xunit.Run(&TestCaseTest{Name: "TestTemplateMethod"})
}
