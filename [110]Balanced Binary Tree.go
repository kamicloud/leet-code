//Given a binary tree, determine if it is height-balanced. 
//
// For this problem, a height-balanced binary tree is defined as: 
//
// 
// a binary tree in which the depth of the two subtrees of every node never differ by more than 1. 
// 
//
// Example 1: 
//
// Given the following tree [3,9,20,null,null,15,7]: 
//
// 
//    3
//   / \
//  9  20
//    /  \
//   15   7 
//
// Return true. 
// 
//Example 2: 
//
// Given the following tree [1,2,2,3,3,null,null,4,4]: 
//
// 
//       1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
// 
//
// Return false. 
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	maxl, bl := deep(root.Left)
	maxr, br := deep(root.Right)

	return int(math.Abs(float64(maxr)-float64(maxl))) <= 1 && bl && br
}

func deep(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	var max = math.MinInt32

	maxl, bl := deep(root.Left)
	maxr, lr := deep(root.Right)

	all := []int{maxl, maxr}
	for i := 0; i < 2; i++ {
		if max < all[i] {
			max = all[i]
		}
	}

	return max + 1, int(math.Abs(float64(maxr)-float64(maxl))) <= 1 && bl && lr
}
