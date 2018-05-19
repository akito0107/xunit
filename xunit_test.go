package xunit

import (
	"fmt"
	"testing"
)

func TestWasRun(t *testing.T) {
	test := NewWasRun("TestMethod")
	fmt.Println(test.wasRun)
	Run(test)
	fmt.Println(test.wasRun)
}
