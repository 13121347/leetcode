package RemoveNode

import (
	"fmt"
	"testing"
)

func Test_removeElements(t *testing.T) {
	a3 := ListNode{
		Val:  6,
		Next: nil,
	}

	a2 := ListNode{
		Val:  6,
		Next: &a3,
	}
	a := ListNode{
		Val:  6,
		Next: &a2,
	}

	b := ListNode{
		Val:  6,
		Next: &a,
	}

	c := ListNode{
		Val:  6,
		Next: &b,
	}

	d := ListNode{
		Val:  6,
		Next: &c,
	}
	l1 := &d
	newList := removeElements(l1, 6)
	for newList != nil {
		fmt.Print(newList.Val)
		newList = newList.Next
	}
}
