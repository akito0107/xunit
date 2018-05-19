package main

import (
	"fmt"

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

func (t *TestCaseTest) TestFailedResult() {
	test := xunit.NewWasRun("TestBrokenMethod")
	result := xunit.Run(test)
	xunit.Assert("1 run, 1 failed" == result.Summary())
}

func (t *TestCaseTest) TestFailedResultFormatting() {
	res := xunit.NewTestResult()
	res.TestStarted()
	res.TestFailed()
	xunit.Assert("1 run, 1 failed" == res.Summary())
}

func main() {
	fmt.Println(xunit.Run(&TestCaseTest{Name: "TestTemplateMethod"}).Summary())
	fmt.Println(xunit.Run(&TestCaseTest{Name: "TestResult"}).Summary())
	fmt.Println(xunit.Run(&TestCaseTest{Name: "TestFailedResult"}).Summary())
	fmt.Println(xunit.Run(&TestCaseTest{Name: "TestFailedResultFormatting"}).Summary())
}
