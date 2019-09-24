//Given a binary tree, return all root-to-leaf paths. 
//
// Note: A leaf is a node with no children. 
//
// Example: 
//
// 
//Input:
//
//   1
// /   \
//2     3
// \
//  5
//
//Output: ["1->2->5", "1->3"]
//
//Explanation: All root-to-leaf paths are: 1->2->5, 1->3
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	left := binaryTreePaths(root.Left)
	right := binaryTreePaths(root.Right)

	all := append(left, right...)

	if len(all) == 0 {
		return []string{strconv.Itoa(root.Val)}
	}

	for i := 0; i < len(all); i++ {
		all[i] = strconv.Itoa(root.Val) + "->" + all[i]
	}

	return all
}