package xunit

import (
	"testing"
	"fmt"
)

func TestWasRun(t *testing.T) {
	test := &WasRun{}
	fmt.Println(test.wasRun)
	test.testMethod()
	fmt.Println(test.wasRun)
}
