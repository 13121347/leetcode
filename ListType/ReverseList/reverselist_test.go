package ReverseList

import (
	"reflect"
	"testing"
)

func Test_reverseList1(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseList2(t *testing.T) {
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

	reverseList2(l1)
}
