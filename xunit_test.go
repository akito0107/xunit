package xunit

import (
	"fmt"
	"testing"
)

func TestWasRun(t *testing.T) {
	test := &WasRun{name: "TestMethod"}
	fmt.Println(test.wasRun)
	test.Run()
	fmt.Println(test.wasRun)
}
