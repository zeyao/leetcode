package main

//https://leetcode.com/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	dic := make(map[string]bool)
	for i := range wordDict {
		dic[wordDict[i]] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			_, ok := dic[s[j:i]]
			if dp[j] && ok {
				dp[i] = true
			}
		}
	}
	return dp[len(dp)-1]

}
