package mergeTwoList

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	//l1 = [6,3,2,1]

	a := ListNode{
		Val:  8,
		Next: nil,
	}

	b := ListNode{
		Val:  7,
		Next: &a,
	}

	c := ListNode{
		Val:  6,
		Next: &b,
	}

	d := ListNode{
		Val:  5,
		Next: &c,
	}
	l1 := &d

	//l2 = [6,3,2,1]
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
	l2 := &i

	/*	result := mergeTwoLists2(nil,&ListNode{0,nil})

		for result != nil{
			fmt.Println(result.Val)
			result = result.Next
		}*/
	/*

		result := mergeTwoLists2(l1,l2)

		for result != nil{
			fmt.Println(result.Val)
			result = result.Next
		}*/

	result := mergeTwoLists3(l1, l2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}

}
