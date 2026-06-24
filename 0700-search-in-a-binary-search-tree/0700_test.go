package searchinabinarysearchtree

import (
	"reflect"
	"testing"
)

// constructBinaryTree builds a binary tree from a level-order (BFS) slice,
// matching the array format used by LeetCode.
func constructBinaryTree(treeArray []int) *TreeNode {
	if len(treeArray) == 0 {
		return nil
	}
	root := &TreeNode{Val: treeArray[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(treeArray) {
		node := queue[0]
		queue = queue[1:]
		if i < len(treeArray) {
			node.Left = &TreeNode{Val: treeArray[i]}
			queue = append(queue, node.Left)
			i++
		}
		if i < len(treeArray) {
			node.Right = &TreeNode{Val: treeArray[i]}
			queue = append(queue, node.Right)
			i++
		}
	}
	return root
}

func TestCase1(t *testing.T) {
	root := constructBinaryTree([]int{4, 2, 7, 1, 3})
	expected := constructBinaryTree([]int{2, 1, 3})
	result := searchBST(root, 2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error! Result: %v, Expected: %v", result, expected)
	}
}

func TestCase2(t *testing.T) {
	root := constructBinaryTree([]int{4, 2, 7, 1, 3})
	result := searchBST(root, 5)
	if result != nil {
		t.Errorf("Error! Expected nil, got: %v", result)
	}
}
