package PartitionList

import (
	"fmt"
	"testing"
)

func Test_partition(t *testing.T) {

	a3 := ListNode{
		Val:  2,
		Next: nil,
	}

	a2 := ListNode{
		Val:  5,
		Next: &a3,
	}
	a := ListNode{
		Val:  2,
		Next: &a2,
	}

	b := ListNode{
		Val:  3,
		Next: &a,
	}

	c := ListNode{
		Val:  4,
		Next: &b,
	}

	d := ListNode{
		Val:  1,
		Next: &c,
	}
	l1 := &d

	result := partition(l1, 3)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

}
