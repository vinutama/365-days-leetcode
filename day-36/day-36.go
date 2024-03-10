package main

import "fmt"

/*
387. First Unqiue character in a String

Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

Example 1:

Input: s = "leetcode"
Output: 0
Example 2:

Input: s = "loveleetcode"
Output: 2
Example 3:

Input: s = "aabb"
Output: -1

Constraints:

1 <= s.length <= 105
s consists of only lowercase English letters.
*/

func firstUniqChar(s string) int {
	countMapper := make(map[rune]int)
	for _, x := range s {
		countMapper[x]++
	}

	for idx, char := range s {
		if countMapper[char] == 1 {
			return idx
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
}
