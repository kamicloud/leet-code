//Find the sum of all left leaves in a given binary tree. 
//
// Example:
// 
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//There are two left leaves in the binary tree, with values 9 and 15 respectively. Return 24.
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
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
    if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
    	return root.Left.Val + sumOfLeftLeaves(root.Right)
	}

	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
//leetcode submit region end(Prohibit modification and deletion)
