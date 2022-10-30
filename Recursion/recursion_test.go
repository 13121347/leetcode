package Recursion

import (
	"fmt"
	"testing"
)

func Test_recursion(t *testing.T) {

	var testFunc func(idx int)
	testFunc = func(idx int) {
		if idx >= 4 {
			return
		}

		for i := idx; i < 3; i++ {
			fmt.Println("index=", idx, " i=", i)
			testFunc(i + 1)
		}
	}
	testFunc(0)
}
