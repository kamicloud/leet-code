//Given a binary tree, return the inorder traversal of its nodes' values. 
//
// Example: 
//
// 
//Input: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//Output: [1,3,2] 
//
// Follow up: Recursive solution is trivial, could you do it iteratively? 
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int

	stack = append(stack, root)


	for {
		if root == nil || len(stack) == 0 {
			return res
		}
		poparr := stack[len(stack) - 1:]

		pop := poparr[0]


		temp := stack[0: len(stack) - 1]
		if pop.Left == nil && pop.Right == nil {
			res = append(res, pop.Val)
			stack = temp
			continue
		}
		if pop.Right != nil {
			temp = append(temp, pop.Right)
			pop.Right = nil
		}
		temp = append(temp, pop)
		if pop.Left != nil {
			temp = append(temp, pop.Left)
			pop.Left = nil
		}

		stack = temp
	}
}