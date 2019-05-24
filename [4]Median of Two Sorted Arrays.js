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

var getRight = (nums, cutCount) => {
    return nums.length - cutCount;
}

var getMids = (nums) => {
    var offset = parseInt(nums.length / 2) - 1;

    var cut = offset >= 0 ? offset : 0;
    return nums.slice(cut, nums.length - cut)
}

var getMid = (nums1, nums2) => {
    var mids = getMids(nums1).concat(getMids(nums2))
    mids.sort((a, b) => {
        return a>b;
    });
    mids = getMids(mids);
    if (mids.length % 2) {
        return mids[parseInt(mids.length / 2)]
    } else {
        return (mids[parseInt(mids.length / 2) - 1] + mids[parseInt(mids.length / 2)]) / 2;
    }
}

/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function(nums1, nums2) {
    if (nums1.length < nums2.length) {
        var temp = nums2;
        nums2 = nums1;
        nums1 = temp;
    }

    if (nums1.length < 2 || nums2.length < 2) {
        return getMid(nums1, nums2)
    }


    var cutCount = parseInt(nums2.length / 2)

    var left = cutCount - 1;
    if (nums1[left] < nums2[left]) {
        nums1 = nums1.slice(cutCount, nums1.length)
    } else {
        nums2 = nums2.slice(cutCount, nums2.length)
    }
    var r1 = nums1[getRight(nums1, cutCount)];
    var r2 = nums2[getRight(nums2, cutCount)];
    if (r1 > r2) {
        nums1 = nums1.slice(0, nums1.length - cutCount);
    } else {
        nums2 = nums2.slice(0, nums2.length - cutCount);

    }
    return findMedianSortedArrays(nums1, nums2);
};