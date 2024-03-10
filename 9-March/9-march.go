package main

import "fmt"

/*
2540. Minimum Common Value

Given two integer arrays nums1 and nums2, sorted in non-decreasing order, return the minimum integer common to both arrays. If there is no common integer amongst nums1 and nums2, return -1.

Note that an integer is said to be common to nums1 and nums2 if both arrays have at least one occurrence of that integer.



Example 1:

Input: nums1 = [1,2,3], nums2 = [2,4]
Output: 2
Explanation: The smallest element common to both arrays is 2, so we return 2.
Example 2:

Input: nums1 = [1,2,3,6], nums2 = [2,3,4,5]
Output: 2
Explanation: There are two common elements in the array 2 and 3 out of which 2 is the smallest, so 2 is returned.


Constraints:

1 <= nums1.length, nums2.length <= 105
1 <= nums1[i], nums2[j] <= 109
Both nums1 and nums2 are sorted in non-decreasing order.
*/

func getCommon(nums1 []int, nums2 []int) int {
	idx1 := 0
	idx2 := 0
	for idx1 < len(nums1) && idx2 < len(nums2) {
		val1 := nums1[idx1]
		val2 := nums2[idx2]

		if val1 < val2 {
			idx1++
		} else if val2 < val1 {
			idx2++
		} else {
			return val1
		}
	}
	return -1
}

func main() {
	fmt.Println(getCommon([]int{1, 2, 3}, []int{2, 3}))
}
