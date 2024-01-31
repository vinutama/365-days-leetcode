package main

import (
	"fmt"
	"strconv"
)

/*
150. Evaluate Reverse Polish Notation

You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the expression.

Note that:

The valid operators are '+', '-', '*', and '/'.
Each operand may be an integer or another expression.
The division between two integers always truncates toward zero.
There will not be any division by zero.
The input represents a valid arithmetic expression in a reverse polish notation.
The answer and all the intermediate calculations can be represented in a 32-bit integer.

Example 1:

Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9
Example 2:

Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6
Example 3:

Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22


Constraints:

1 <= tokens.length <= 104
tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].

*/

func pop(s *[]int) int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, val := range tokens {
		if val == "+" {
			num1 := pop(&stack)
			num2 := pop(&stack)
			stack = append(stack, num1+num2)
		} else if val == "-" {
			num1 := pop(&stack)
			num2 := pop(&stack)
			stack = append(stack, num2-num1)
		} else if val == "/" {
			num1 := pop(&stack)
			num2 := pop(&stack)
			stack = append(stack, num2/num1)
		} else if val == "*" {
			num1 := pop(&stack)
			num2 := pop(&stack)
			stack = append(stack, num2*num1)
		} else {
			num, err := strconv.Atoi(val)
			if err != nil {
				break
			}
			stack = append(stack, num)
		}
	}

	if len(stack) == 0 {
		return 0
	} else {
		return stack[len(stack)-1]
	}
}

func main() {
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))

}
