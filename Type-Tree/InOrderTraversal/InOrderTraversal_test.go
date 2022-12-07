package InOrderTraversal

import (
	"fmt"
	"testing"
)

//  1
// 2  3
//4    6

func Test_inOrderTraversal(t *testing.T) {

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

	result := inOrderTraversal(testNodeTree)
	fmt.Println(result)
}
