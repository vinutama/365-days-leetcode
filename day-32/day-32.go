package main

import (
	"fmt"
	"sort"
)

/*
2966. Divide Array Into Arrays With Max Difference

You are given an integer array nums of size n and a positive integer k.

Divide the array into one or more arrays of size 3 satisfying the following conditions:

Each element of nums should be in exactly one array.
The difference between any two elements in one array is less than or equal to k.
Return a 2D array containing all the arrays. If it is impossible to satisfy the conditions, return an empty array. And if there are multiple answers, return any of them.

Example 1:

Input: nums = [1,3,4,8,7,9,3,5,1], k = 2
Output: [[1,1,3],[3,4,5],[7,8,9]]
Explanation: We can divide the array into the following arrays: [1,1,3], [3,4,5] and [7,8,9].
The difference between any two elements in each array is less than or equal to 2.
Note that the order of elements is not important.
Example 2:

Input: nums = [1,3,3,2,7,3], k = 3
Output: []
Explanation: It is not possible to divide the array satisfying all the conditions.


Constraints:
n == nums.length
1 <= n <= 105
n is a multiple of 3.
1 <= nums[i] <= 105
1 <= k <= 105
*/

func divideArray(nums []int, k int) [][]int {
	result := make([][]int, len(nums)/3)
	sort.Slice(nums[:], func(i, j int) bool {
		return nums[i] < nums[j]
	})

	subArrayIdx := 0
	for i := len(nums) - 1; i >= 0; i -= 3 {
		if nums[i]-nums[i-1] <= k && nums[i]-nums[i-2] <= k && nums[i-1]-nums[i-2] <= k {
			result[subArrayIdx] = nums[i-2 : i+1]
			subArrayIdx++
		} else {
			return [][]int{}
		}

	}
	return result
}

func main() {
	fmt.Println(divideArray([]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2))
	fmt.Println(divideArray([]int{1, 3, 3, 2, 7, 3}, 3))
	fmt.Println(divideArray([]int{6, 10, 5, 12, 7, 11, 6, 6, 12, 12, 11, 7}, 2))
}
