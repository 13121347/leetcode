package lastOrderTraversal

import (
	"fmt"
	"testing"
)

//  1
// 2  3
//4    6

func Test_lastOrderTraversal(t *testing.T) {
	testNodeTree := &TreeNode{
		Val: 1,
	}

	testNodeTree.Left = &TreeNode{
		Val: 2,
	}
	testNodeTree.Left.Left = &TreeNode{
		Val: 4,
	}

	testNodeTree.Right = &TreeNode{
		Val: 3,
	}
	testNodeTree.Right.Right = &TreeNode{
		Val: 6,
	}

	//suppose to be 4 2 6,3,1
	result := lastOrderTraversal(testNodeTree)
	fmt.Println(result)
}
