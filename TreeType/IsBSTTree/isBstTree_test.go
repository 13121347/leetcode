package IsBSTTree

import (
	"fmt"
	"testing"
)

func Test_isValidBST(t *testing.T) {

	t7 := &TreeNode{7, nil, nil}
	t2 := &TreeNode{2, nil, nil}
	t33 := &TreeNode{3, nil, nil}
	t4 := &TreeNode{4, nil, nil}
	t3 := &TreeNode{3, t2, t4}
	t6 := &TreeNode{6, t33, t7}
	t5 := &TreeNode{5, t3, t6}

	fmt.Println(isValidBST2(t5))
}
