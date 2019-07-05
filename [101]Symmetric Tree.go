//Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center). 
//
// For example, this binary tree [1,2,2,3,4,4,3] is symmetric: 
//
// 
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3
// 
//
// 
//
// But the following [1,2,2,null,3,null,3] is not: 
//
// 
//    1
//   / \
//  2   2
//   \   \
//   3    3
// 
//
// 
//
// Note: 
//Bonus points if you could solve it both recursively and iteratively. 
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTwo(root.Left, root.Right)
}

func isSymmetricTwo(left *TreeNode, right *TreeNode) bool {
	if left == nil {
		return right == nil
	}
	if right == nil || left.Val != right.Val {
		return false
	}

	return isSymmetricTwo(left.Left, right.Right) && isSymmetricTwo(left.Right, right.Left)
}