//Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary). 
//
// You may assume that the intervals were initially sorted according to their start times. 
//
// Example 1: 
//
// 
//Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
//Output: [[1,5],[6,9]]
// 
//
// Example 2: 
//
// 
//Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//Output: [[1,2],[3,10],[12,16]]
//Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10]. 
//
// NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature. 
// Related Topics Array Sort

package main



//leetcode submit region begin(Prohibit modification and deletion)

type forSort [][]int

func (fs forSort) Len() int {
	return len(fs)
}

func (fs forSort) Swap(i, j int) {
	fs[i], fs[j] = fs[j], fs[i]
}
func (fs forSort) Less(i, j int) bool {
	return fs[i][0] < fs[j][0]
}


func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)

	sort.Sort(forSort(intervals))

	if len(intervals) == 0 {
		return intervals
	}

	max := intervals[0][1]

	for i := 0; i < len(intervals); i++ {
		if len(intervals) - 1 == i {
			return intervals
		}

		if intervals[i + 1][0] <= max {
			if max < intervals[i + 1][1] {
				intervals[i][1] = intervals[i + 1][1]
				max = intervals[i + 1][1]
			}
			if i + 2 <= len(intervals) {
				intervals = append(intervals[0:i+1], intervals[i+2:]...)
				i--
			}
		} else {
			max = intervals[i + 1][1]
		}
	}

	return intervals
}
//leetcode submit region end(Prohibit modification and deletion)
