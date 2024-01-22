package main

import (
	"fmt"
	"math"
)

/*
931. Minimum Falling Path Sum

Given an n x n array of integers matrix, return the minimum sum of any falling path through matrix.

A falling path starts at any element in the first row and chooses the element in the next row that is either directly below or diagonally left/right. Specifically, the next element from position (row, col) will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1).

Example 1:


Input: matrix = [[2,1,3],[6,5,4],[7,8,9]]
Output: 13
Explanation: There are two falling paths with a minimum sum as shown.
Example 2:


Input: matrix = [[-19,57],[-40,-5]]
Output: -59
Explanation: The falling path with a minimum sum is shown.


Constraints:

n == matrix.length == matrix[i].length
1 <= n <= 100
-100 <= matrix[i][j] <= 100

*/

func minFallingPathSum(matrix [][]int) int {
	startArr := matrix[0]

	res := math.MinInt

	res = startArr[0]
	res = 2
	res = -99
	res = -89
	res = 0

	// return result
	return res
}

func main() {
	// fmt.Println(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))                                                // 13
	// fmt.Println(minFallingPathSum([][]int{{-19, 57}, {-40, -5}}))                                                           // -59
	fmt.Println(minFallingPathSum([][]int{{100, -42, -46, -41}, {31, 97, 10, -10}, {-58, -51, 82, 89}, {51, 81, 69, -51}})) // -36
}
