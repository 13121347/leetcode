package AddTwoNumber

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	//l1 = [6,3,2,1]

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
	l2 := i

	/*	nilresult := AddTwoNumbers(nil,nil)
		fmt.Println(nilresult)
		fmt.Println("***************************************")
		l1nilresult := AddTwoNumbers(nil,&l2)
		for l1nilresult != nil {
			fmt.Println(l1nilresult.Val)
			l1nilresult = l1nilresult.Next
		}
		fmt.Println("***************************************")
		l2nilresult := AddTwoNumbers(&l1,nil)
		for l2nilresult != nil {
			fmt.Println(l2nilresult.Val)
			l2nilresult = l2nilresult.Next
		}*/
	fmt.Println("***************************************")
	normalResult := AddTwoNumbers(&l1, &l2)
	for normalResult != nil {
		fmt.Println(normalResult.Val)
		normalResult = normalResult.Next
	}
}
