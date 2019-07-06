//Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root). 
//
// 
//For example: 
//Given binary tree [3,9,20,null,null,15,7], 
// 
//    3
//   / \
//  9  20
//    /  \
//   15   7
// 
// 
// 
//return its bottom-up level order traversal as: 
// 
//[
//  [15,7],
//  [9,20],
//  [3]
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


func levelOrderBottom(root *TreeNode) [][]int {
	var res = make([][]int, 0)

	var next = make([]*TreeNode, 0)
	next = append(next, root)
	if root == nil {
		return res
	}
	for {
		var temp2 []int
		temp2, next = levelGet(next)
		if len(temp2) > 0 {
			res = append(res, temp2)
		}
		if len(next) == 0 {
			break
		}
	}
	return reverse(res)
}
func reverse(input [][] int) [][]int {
	for i := 0; i < len(input) / 2; i++ {
		input[len(input) - i - 1], input[i] = input[i], input[len(input) - i - 1]
	}
	return input
}

func levelGet(level []*TreeNode) (res []int, next []*TreeNode) {
	for i := 0; i < len(level); i++ {
		if level[i] != nil {
			res = append(res, level[i].Val)
			if level[i].Left != nil {
				next = append(next, level[i].Left)
			}
			if level[i].Right != nil {
				next = append(next, level[i].Right)
			}
		}
	}
	return
}