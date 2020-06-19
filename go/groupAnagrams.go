package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := {"avbv","aba","bab"}
	fmt.Println(groupAnagrams(arr));
}

func groupAnagrams(strs []string) [][]string {
	dic := make(map[string][]string)
	for i := range strs {
		arr := []byte(strs[i])
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		key := string(arr)
		_, ok := dic[key]
		if !ok {
			dic[key] = make([]string, 0)
		}
		dic[key] = append(dic[key], strs[i])
	}
	res := make([][]string, 0)
	for _, v := range dic {
		res = append(res, v)
	}
	return res
}


