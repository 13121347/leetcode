package ReverseKGroup

import (
	"fmt"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	f := ListNode{
		Val:  4,
		Next: nil,
	}

	g := ListNode{
		Val:  3,
		Next: &f,
	}

	h := ListNode{
		Val:  2,
		Next: &g,
	}

	i := ListNode{
		Val:  1,
		Next: &h,
	}
	head := &i //1-2-3-4
	tail := &g //3-4

	myReverse(head, tail)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
