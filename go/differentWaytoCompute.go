package main

//https://leetcode.com/problems/different-ways-to-add-parentheses/
//分治 模板题

import "strconv"

func diffWaysToCompute(input string) []int {
	res := make([]int, 0)
	num, err := strconv.Atoi(input)
	if err == nil {
		res = append(res, num)
		return res
	}
	for i := range input {
		if input[i] == '+' || input[i] == '*' || input[i] == '-' {
			l := input[0:i]
			r := input[i+1:]
			left := diffWaysToCompute(l)
			right := diffWaysToCompute(r)
			for x := range left {
				for y := range right {
					if input[i] == '+' {
						res = append(res, left[x]+right[y])
					} else if input[i] == '-' {
						res = append(res, left[x]-right[y])
					} else {
						res = append(res, left[x]*right[y])
					}
				}
			}
		}
	}
	return res

}
