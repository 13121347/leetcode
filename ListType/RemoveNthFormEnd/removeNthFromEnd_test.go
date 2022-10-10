package RemoveNthFormEnd

import (
	"fmt"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	a := ListNode{
		Val:  4,
		Next: nil,
	}

	b := ListNode{
		Val:  3,
		Next: &a,
	}

	c := ListNode{
		Val:  2,
		Next: &b,
	}

	d := ListNode{
		Val:  1,
		Next: &c,
	}
	l1 := &d

	result := removeNthFromEndTwoPtr(l1, 4)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
