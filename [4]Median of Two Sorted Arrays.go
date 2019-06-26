//There are two sorted arrays nums1 and nums2 of size m and n respectively. 
//
// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)). 
//
// You may assume nums1 and nums2 cannot be both empty. 
//
// Example 1: 
//
// 
//nums1 = [1, 3]
//nums2 = [2]
//
//The median is 2.0
// 
//
// Example 2: 
//
// 
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//The median is (2 + 3)/2 = 2.5
// 
//

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1)+len(nums2) == 1 || len(nums1)+len(nums2) == 2 {

		return findMedianSortedArray(append(nums1, nums2...))
	}

	if len(nums1) == 0 {
		return findMedianSortedArray(nums2)
	}
	if len(nums2) == 0 {
		return findMedianSortedArray(nums1)
	}

	if nums1[0] < nums2[0] {
		nums1 = nums1[1:]
	} else {
		nums2 = nums2[1:]
	}
	if len(nums2) == 0 || (len(nums1) > 0 && nums1[len(nums1)-1] > nums2[len(nums2)-1]) {
		nums1 = nums1[:len(nums1)-1]
	} else {
		nums2 = nums2[:len(nums2)-1]
	}

	return findMedianSortedArrays(nums1, nums2)
}

func findMedianSortedArray(nums []int) float64 {
	if len(nums)%2 == 1 {
		return float64(nums[len(nums)/2])
	}

	return float64(nums[len(nums)/2]+nums[len(nums)/2-1]) / 2.0
}
