//Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target. 
//
// Each number in candidates may only be used once in the combination. 
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
//Input: candidates = [10,1,2,7,6,1,5], target = 8,
//A solution set is:
//[
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
//]
// 
//
// Example 2: 
//
// 
//Input: candidates = [2,5,2,1,2], target = 5,
//A solution set is:
//[
//  [1,2,2],
//  [5]
//]
// 
//

func combinationSum2(candidates []int, target int) [][]int {
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
			ress = doCombine(candidates[i + 1:], target-candidates[i])

			for j := 0; j < len(ress); j++ {
				res = append(res, append([]int{candidates[i]}, ress[j]...))
			}
		}

		for {
			if i + 1 >= len(candidates) || candidates[i] != candidates[i + 1] {
				break
			}
			i++
		}
	}

	return res
}
