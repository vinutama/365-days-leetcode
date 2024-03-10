package main

import (
	"fmt"
	"math"
)

/*
231. Power of Two
Given an integer n, return true if it is a power of two. Otherwise, return false.

An integer n is a power of two, if there exists an integer x such that n == 2x.

Example 1:

Input: n = 1
Output: true
Explanation: 20 = 1
Example 2:

Input: n = 16
Output: true
Explanation: 24 = 16
Example 3:

Input: n = 3
Output: false


Constraints:

-231 <= n <= 231 - 1

Follow up: Could you solve it without loops/recursion?
*/

func isPowerOfTwo(n int) bool {
	a := math.Log2(float64(n))
	b := math.Floor(math.Log2(float64(n)))
	fmt.Println(a, b)
	if n > 0 && a == b {
		return true
	}
	return false
}

// func isPowerOfTwo(n int) bool {
// 	if n == 1 || n == 2 {
// 		return true
// 	}
// 	if n%2 != 0 {
// 		return false
// 	}

// 	for true {
// 		n /= 2
// 		if n%2 != 0 {
// 			return false
// 		} else if n == 2 {
// 			return true
// 		}
// 	}
// 	return false
// }

func main() {
	fmt.Println(isPowerOfTwo(1))
}
