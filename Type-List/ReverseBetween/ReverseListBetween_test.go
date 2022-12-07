package ReverseBetween

import (
	"fmt"
	"testing"
)

func Test_reverseBetween(t *testing.T) {

	c := ListNode{
		Val:  5,
		Next: nil,
	}

	d := ListNode{
		Val:  3,
		Next: &c,
	}
	l1 := &d

	result := reverseBetween(l1, 1, 2)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
