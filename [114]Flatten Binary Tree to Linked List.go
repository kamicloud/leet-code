//Given a binary tree, flatten it to a linked list in-place. 
//
// For example, given the following tree: 
//
// 
//    1
//   / \
//  2   5
// / \   \
//3   4   6
// 
//
// The flattened tree should look like: 
//
// 
//1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6
// 
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	var left *TreeNode = root.Left
	var right *TreeNode = root.Right

	flatten(root.Left)
	root.Left = nil
	flatten(root.Right)

	root.Right = left

	for {
		if root.Right == nil {
			break
		}

		root = root.Right
	}

	root.Right = right
}