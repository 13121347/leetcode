package IntersectionLIst

import (
	"fmt"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	a := ListNode{
		Val:  1,
		Next: nil,
	}

	b := ListNode{
		Val:  1,
		Next: &a,
	}

	c := ListNode{
		Val:  1,
		Next: &b,
	}

	d := ListNode{
		Val:  1,
		Next: &c,
	}
	l1 := d

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

	i := ListNode{
		Val:  6,
		Next: &h,
	}

	j := ListNode{
		Val:  6,
		Next: &i,
	}

	k := ListNode{
		Val:  6,
		Next: &j,
	}

	l2 := k

	getIntersectionNode(&l1, &l2)
}

func Test_getIntersectionNode2(t *testing.T) {
	result := getIntersectionNode(&ListNode{1, nil}, &ListNode{1, nil})
	fmt.Println(result.Val)
}
