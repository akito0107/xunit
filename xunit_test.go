package xunit

import (
	"fmt"
	"testing"
)

func TestWasRun(t *testing.T) {
	test := &WasRun{name: "testMethod"}
	fmt.Println(test.wasRun)
	test.testMethod()
	fmt.Println(test.wasRun)
}
