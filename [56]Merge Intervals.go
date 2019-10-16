//Given a collection of intervals, merge all overlapping intervals. 
//
// Example 1: 
//
// 
//Input: [[1,3],[2,6],[8,10],[15,18]]
//Output: [[1,6],[8,10],[15,18]]
//Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
// 
//
// Example 2: 
//
// 
//Input: [[1,4],[4,5]]
//Output: [[1,5]]
//Explanation: Intervals [1,4] and [4,5] are considered overlapping. 
//
// NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature. 
// Related Topics Array Sort

package main

//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	sort.Sort(ForSort(intervals))
	var res [][]int
	var current []int
	biggest := math.MinInt64
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > biggest {
			res = append(res, intervals[i])
			current = intervals[i]
		} else if intervals[i][1] > biggest {
			// 右侧大
			current[1] = intervals[i][1]
		}
		if intervals[i][1] > biggest {
			biggest = intervals[i][1]
		}
	}

	return res
}

type ForSort [][]int

func (fs ForSort) Len() int {
	return len(fs)
}

func (fs ForSort) Swap(i, j int) {
	fs[i], fs[j] = fs[j], fs[i]
}

func (fs ForSort) Less(i, j int) bool {
	return fs[i][0] < fs[j][0]
}


//leetcode submit region end(Prohibit modification and deletion)
