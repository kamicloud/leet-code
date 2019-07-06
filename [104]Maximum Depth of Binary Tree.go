//Given a binary tree, find its maximum depth. 
//
// The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node. 
//
// Note: A leaf is a node with no children. 
//
// Example: 
//
// Given binary tree [3,9,20,null,null,15,7], 
//
// 
//    3
//   / \
//  9  20
//    /  \
//   15   7 
//
// return its depth = 3. 
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	var ll = 0
	var rl = 0
	if root.Left != nil {
		ll = 1 + maxDepth(root.Left)
	}
	if root.Right != nil {
		rl = 1 + maxDepth(root.Right)
	}

	if ll > rl {
		return ll
	}

	return rl
}