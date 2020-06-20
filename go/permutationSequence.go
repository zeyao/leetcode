package main

import (
	"strconv"
	"strings"
)

func main() {

}

func getPermutation(n int, k int) string {
	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, i+1)
	}
	for i := 1; i < k; i++ {
		nextPermutation(arr)
	}
	var sb strings.Builder
	for i := range arr {
		sb.WriteString(strconv.Itoa(arr[i]))
	}
	return sb.String()
}

func nextPermutation(arr []int) {
	//14 987653
	i := len(arr) - 1
	for i >= 0 && arr[i-1] >= arr[i] {
		i--
	}
	i--
	if i >= 0 {
		j := len(arr) - 1
		for j >= 0 && arr[j] <= arr[i] {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	reverse(i+1, arr)
}

func reverse(i int, arr []int) {
	left := i
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
