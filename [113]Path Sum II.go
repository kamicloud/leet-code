//Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum. 
//
// Note: A leaf is a node with no children. 
//
// Example: 
//
// Given the below binary tree and sum = 22, 
//
// 
//      5
//     / \
//    4   8
//   /   / \
//  11  13  4
// /  \    / \
//7    2  5   1
// 
//
// Return: 
//
// 
//[
//   [5,4,11,2],
//   [5,8,4,5]
//]
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

func pathSum(root *TreeNode, sum int) [][]int {
	var temp [][]int

	if root == nil {
		return temp
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return [][]int{[]int{root.Val}}
	}

	temp = append(temp, pathSum(root.Left, sum - root.Val)...)
	temp = append(temp, pathSum(root.Right, sum - root.Val)...)


	for i := 0; i < len(temp); i++ {
		temp[i] = append([]int{root.Val}, temp[i]...)
	}
	return temp
}