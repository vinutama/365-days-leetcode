package main

import "fmt"

/*
739. Daily Temperatures

Given an array of integers temperatures represents the daily temperatures,
return an array answer such that answer[i] is the number of days you have to wait after the
ith day to get a warmer temperature.
If there is no future day for which this is possible, keep answer[i] == 0 instead.

Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
Example 2:

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
Example 3:

Input: temperatures = [30,60,90]
Output: [1,1,0]


Constraints:

1 <= temperatures.length <= 105
30 <= temperatures[i] <= 100

*/

// func dailyTemperatures(temperatures []int) []int {
// 	answers := make([]int, len(temperatures))
// 	for idx, temp := range temperatures {
// 		waiting := 0
// 		found := false
// 		for i := idx + 1; i < len(temperatures); i++ {
// 			waiting++
// 			if temperatures[i] > temp {
// 				found = true
// 				break
// 			}

// 			fmt.Println(waiting, "WAITING")
// 		}
// 		if found {
// 			answers[idx] = waiting
// 		} else {
// 			answers[idx] = 0
// 		}
// 	}

// 	return answers
// }

type Stack struct {
	idx  int
	temp int
}

func pop(s *[]Stack) Stack {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func dailyTemperatures(temperatures []int) []int {
	answers := make([]int, len(temperatures))
	stacks := make([]Stack, 0)
	for idx, temp := range temperatures {
		for len(stacks) != 0 && temp > stacks[len(stacks)-1].temp {
			prevStack := pop(&stacks)
			prevIdx := prevStack.idx
			answers[prevIdx] = idx - prevIdx
		}

		stacks = append(stacks, Stack{idx, temp})
	}
	return answers
}

func main() {
	// fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	// fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	// fmt.Println(dailyTemperatures([]int{30, 60, 90}))
	fmt.Println(dailyTemperatures([]int{55, 38, 53, 81, 61, 93, 97, 32, 43, 78}))
	//[3 1 1 2 1 1 0 0 0 0]
	//[3,1,1,2,1,1,0,1,1,0]
}

// currIdx 1
// idx 2

// currIdx 2
// idx 3

// currIdx 2
// idx 4

// currIdx 2
// idx 5

// currIdx 2
// idx 6

// currIdx 3
// idx 4
