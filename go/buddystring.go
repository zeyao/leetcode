package main

//https://leetcode.com/problems/buddy-strings/

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	dic := make(map[byte]int)
	if A == B {
		//at least one char should have more than one count so we can switch
		for i := range A {
			dic[A[i]]++
			if dic[A[i]] > 1 {
				return true
			}
		}

	} else {
		dif1 := -1
		dif2 := -1
		for i := range A {
			if A[i] != B[i] {
				if dif1 == -1 {
					dif1 = i
				} else if dif2 == -1 {
					dif2 = i
				} else {
					return false
				}
			}
		}
		if dif2 == -1 {
			return false
		}
		if A[dif1] == B[dif2] && A[dif2] == B[dif1] {
			return true
		}
	}
	return false

}
