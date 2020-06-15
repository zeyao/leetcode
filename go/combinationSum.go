package main

func main() {
	/**
		& * 定义了 go的指针
		go 不像 java 传递对象的引用，
		go 传递的是值
		go 不是 面向对象， slice这里更像是java String 的传递，而不是list collection的传递，
		需要等同于list的使用方法，需要指定指针

	**/
}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	dfs(target, &res, candidates, []int{}, 0)
	return res
}

func dfs(target int, res *[][]int, candidates []int, innerList []int, start int) {
	if target == 0 {
		*res = append(*res, append([]int{}, innerList...))
		return
	}
	if target < 0 {
		return
	}
	for i := start; i < len(candidates); i++ {
		dfs(target-candidates[i], res, candidates, append(innerList, candidates[i]), i)
	}
}
