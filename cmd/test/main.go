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
	result := xunit.NewTestResult()
	xunit.Run(test, result)
	xunit.Assert("setUp testMethod tearDown " == test.Log)
}

func (t *TestCaseTest) TestResult() {
	test := xunit.NewWasRun("TestMethod")
	result := xunit.NewTestResult()
	xunit.Run(test, result)
	xunit.Assert("1 run, 0 failed" == result.Summary())
}

func (t *TestCaseTest) TestFailedResult() {
	test := xunit.NewWasRun("TestBrokenMethod")
	result := xunit.NewTestResult()
	xunit.Run(test, result)
	xunit.Assert("1 run, 1 failed" == result.Summary())
}

func (t *TestCaseTest) TestFailedResultFormatting() {
	res := xunit.NewTestResult()
	res.TestStarted()
	res.TestFailed()
	xunit.Assert("1 run, 1 failed" == res.Summary())
}

func (t *TestCaseTest) TestSuite() {
	suite := &xunit.TestSuite{}
	suite.Add(xunit.NewWasRun("TestMethod"))
	suite.Add(xunit.NewWasRun("TestBrokenMethod"))
	result := xunit.NewTestResult()
	xunit.Run(suite, result)
	xunit.Assert("2 run, 1 failed" == result.Summary())
}

func main() {
	s := &xunit.TestSuite{}
	s.Add(&TestCaseTest{Name: "TestTemplateMethod"})
	s.Add(&TestCaseTest{Name: "TestResult"})
	s.Add(&TestCaseTest{Name: "TestFailedResult"})
	s.Add(&TestCaseTest{Name: "TestFailedResultFormatting"})
	s.Add(&TestCaseTest{Name: "TestSuite"})
	res := xunit.NewTestResult()
	xunit.Run(s, res)
	fmt.Println(res.Summary())
}
