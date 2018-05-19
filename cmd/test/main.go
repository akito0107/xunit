package main

import (
	"github.com/akito0107/xunit"
)

type TestCaseTest struct {
	Name string
	test *xunit.WasRun
}

func (t *TestCaseTest) TestTemplateMethod() {
	test := xunit.NewWasRun("TestMethod")
	xunit.Run(test)
	xunit.Assert("setUp testMethod tearDown " == test.Log)
}

func (t *TestCaseTest) TestResult() {
	test := xunit.NewWasRun("TestMethod")
	result := xunit.Run(test)
	xunit.Assert("1 run, 0 failed" == result.Summary())
}

func main() {
	xunit.Run(&TestCaseTest{Name: "TestTemplateMethod"})
	xunit.Run(&TestCaseTest{Name: "TestResult"})
}
