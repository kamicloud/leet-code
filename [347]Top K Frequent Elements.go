//Given a non-empty array of integers, return the k most frequent elements. 
//
// Example 1: 
//
// 
//Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
// 
//
// 
// Example 2: 
//
// 
//Input: nums = [1], k = 1
//Output: [1] 
// 
//
// Note: 
//
// 
// You may assume k is always valid, 1 ≤ k ≤ number of unique elements. 
// Your algorithm's time complexity must be better than O(n log n), where n is the array's size. 
// 
// Related Topics Hash Table Heap

package main

//leetcode submit region begin(Prohibit modification and deletion)
type forSort [][]int

func (fs forSort) Swap(i int, j int) {
	fs[i], fs[j] = fs[j], fs[i]
}

func (fs forSort) Len() int {
	return len(fs)
}

func (fs forSort) Less(i int, j int) bool {
	return fs[i][0] < fs[j][0]
}

func topKFrequent(nums []int, k int) []int {
	ct := map[int]int{}

	for i := 0; i < len(nums); i++ {
		c := ct[nums[i]]
		c++
		ct[nums[i]] = c
	}

	var st [][]int

	for k, v := range ct {
		st = append(st, []int{v, k})
	}

	sort.Sort(forSort(st))

	var res []int
	max := math.MaxInt64
	for i := len(st) - 1; i >= 0; i-- {
		if max > st[i][0] {
			if k == 0 {
				return res
			}
			res = append(res, st[i][1])
			max  = st[i][0]
		} else {
			// equal
			res = append(res, st[i][1])
		}
		k--
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)
