package SwapPairs

import (
	"fmt"
	"testing"
)

func Test_swapPairs(t *testing.T) {

	a := ListNode{
		Val:  1,
		Next: nil,
	}

	b := ListNode{
		Val:  2,
		Next: &a,
	}

	c := ListNode{
		Val:  3,
		Next: &b,
	}

	d := ListNode{
		Val:  4,
		Next: &c,
	}
	l1 := &d

	result := swapPairs(l1)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

}
