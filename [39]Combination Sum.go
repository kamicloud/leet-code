//Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target. 
//
// The same repeated number may be chosen from candidates unlimited number of times. 
//
// Note: 
//
// 
// All numbers (including target) will be positive integers. 
// The solution set must not contain duplicate combinations. 
// 
//
// Example 1: 
//
// 
//Input: candidates = [2,3,6,7], target = 7,
//A solution set is:
//[
//  [7],
//  [2,2,3]
//]
// 
//
// Example 2: 
//
// 
//Input: candidates = [2,3,5], target = 8,
//A solution set is:
//[
//  [2,2,2,2],
//  [2,3,3],
//  [3,5]
//]
// 
//

func combinationSum(candidates []int, target int) [][]int {
	for i := 0; i < len(candidates); i++ {
		for j := i + 1; j < len(candidates); j++ {
			if candidates[i] > candidates[j] {
				candidates[i], candidates[j] = candidates[j], candidates[i]
			}
		}
	}

	return doCombine(candidates, target)
}

func doCombine(candidates []int, target int) [][]int {

	var res [][]int

	if len(candidates) == 0 || target < 0 {
		return res
	}

	if candidates[0] == target {
		return [][]int{{target}}
	}

	for i := 0; i < len(candidates); i++ {
		var ress [][]int
		if candidates[i] == target {
			res = append(res, []int{target})
		} else {
			ress = doCombine(candidates[i:], target-candidates[i])

			for j := 0; j < len(ress); j++ {
				res = append(res, append([]int{candidates[i]}, ress[j]...))
			}
		}
	}

	return res
}
