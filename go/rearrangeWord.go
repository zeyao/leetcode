package main

//https://leetcode.com/problems/rearrange-words-in-a-sentence/
import (
	"sort"
	"strings"
)

func arrangeWords(text string) string {
	words := strings.Split(strings.ToLower(text), " ")
	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	//用 sort.SliceStable 方法，sort.SliceStable 可以保持元素顺序和之前切片中的一样。
	// 用 sort.Slice会失败
	ans := strings.Join(words, " ")
	arr := []byte(ans)
	arr[0] = arr[0] - 32
	return string(arr)
}
