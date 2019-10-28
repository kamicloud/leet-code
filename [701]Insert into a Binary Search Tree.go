//Given the root node of a binary search tree (BST) and a value to be inserted into the tree, insert the value into the BST. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST. 
//
// Note that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them. 
//
// For example, 
//
// 
//Given the tree:
//        4
//       / \
//      2   7
//     / \
//    1   3
//And the value to insert: 5
// 
//
// You can return this binary search tree: 
//
// 
//         4
//       /   \
//      2     7
//     / \   /
//    1   3 5
// 
//
// This tree is also valid: 
//
// 
//         5
//       /   \
//      2     7
//     / \   
//    1   3
//         \
//          4
// 
// Related Topics Tree


package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
    	return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	if root.Val > val {
		// left
		if root.Left == nil {
			root.Left = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
		} else {
			root.Left = insertIntoBST(root.Left, val)
		}
	} else {
		// right
		if root.Right == nil {
			root.Right = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
		} else {
			root.Right = insertIntoBST(root.Right, val)
		}
	}
	return root
}
//leetcode submit region end(Prohibit modification and deletion)
