package InsertSort

import (
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	//l2 = [6,3,2,1]
	f := ListNode{
		Val:  1,
		Next: nil,
	}

	g := ListNode{
		Val:  2,
		Next: &f,
	}

	h := ListNode{
		Val:  3,
		Next: &g,
	}

	k := ListNode{
		Val:  6,
		Next: &h,
	}

	l2 := &k

	insertionSortList(l2)
}
