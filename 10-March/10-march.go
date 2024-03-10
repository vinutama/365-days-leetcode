package main

import "fmt"

/*
349. Intersection of Two Arrays

Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.


Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.


Constraints:

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000

*/

func intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	mapIntersect := make(map[int]bool)

	for _, v := range nums1 {
		mapIntersect[v] = true
	}
	for _, v := range nums2 {
		if mapIntersect[v] {
			result = append(result, v)
			mapIntersect[v] = false
		}
	}
	return result

}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}
