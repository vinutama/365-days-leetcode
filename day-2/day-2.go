package main

import "fmt"

/*
You are given an integer array nums. You need to create a 2D array from nums satisfying the following conditions:

The 2D array should contain only the elements of the array nums.
Each row in the 2D array contains distinct integers.
The number of rows in the 2D array should be minimal.
Return the resulting array. If there are multiple answers, return any of them.

Note that the 2D array can have a different number of elements on each row.


Input: nums = [1,3,4,1,2,3,1]
Output: [[1,3,4,2],[1,3],[1]]
Explanation: We can create a 2D array that contains the following rows:
- 1,3,4,2
- 1,3
- 1
All elements of nums were used, and each row of the 2D array contains distinct integers, so it is a valid answer.
It can be shown that we cannot have less than 3 rows in a valid array.


Input: nums = [1,2,3,4]
Output: [[4,3,2,1]]
Explanation: All elements of the array are distinct, so we can keep all of them in the first row of the 2D array.

*/

func findMatrix(nums []int) [][]int {

	mapperFreq := make(map[int]int)
	var size int

	for _, num := range nums {
		if _, found := mapperFreq[num]; found {
			mapperFreq[num]++
		} else {
			mapperFreq[num] = 1
		}

		if mapperFreq[num] > size {
			size = mapperFreq[num]
		}
	}

	result := make([][]int, size)

	for num, freq := range mapperFreq {
		n := 0
		for n < freq {
			result[n] = append(result[n], num)
			n++
		}
	}

	return result
}

func main() {
	fmt.Println(findMatrix([]int{1, 3, 4, 1, 2, 3, 1}))
	fmt.Println(findMatrix([]int{1, 2, 3, 4}))
}
