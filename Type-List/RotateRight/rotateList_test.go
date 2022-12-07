package RotateRight

import (
	"fmt"
	"testing"
)

func Test_rotateRight(t *testing.T) {

	b := ListNode{
		Val:  2,
		Next: nil,
	}

	c := ListNode{
		Val:  1,
		Next: &b,
	}

	d := ListNode{
		Val:  0,
		Next: &c,
	}
	l1 := &d

	result := rotateRightListDouble(l1, 4)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
