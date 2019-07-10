//Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array. 
//
// Note: 
//
// 
// The number of elements initialized in nums1 and nums2 are m and n respectively. 
// You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2. 
// 
//
// Example: 
//
// 
//Input:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//Output: [1,2,2,3,5,6]
// 
//


func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, p := m-1, n-1, len(nums1)-1
	if len(nums2) == 0 {
		return
	}
	if len(nums1) == 1 && len(nums2) == 1 {
		nums1[0] = nums2[0]
		return
	}
	for {
		if p < 0 {
			break
		}

		if i >= 0 && (j < 0 || nums1[i] > nums2[j]) {
			nums1[p] = nums1[i]
			i--
		} else {
			nums1[p] = nums2[j]
			j--
		}
		p--

	}
}
