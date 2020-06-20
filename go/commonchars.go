package main

//https://leetcode.com/problems/find-common-characters/
func main() {

}

func commonChars(A []string) []string {
	dic := make([][26]int, 0)
	for i := range A {
		var count [26]int
		str := A[i]
		for j := range str {
			count[str[j]-'a']++
		}
		dic = append(dic, count)
	}

	res := make([]string, 0)
	for i := 0; i < 26; i++ {
		common := 500
		for j := range dic {
			arr := dic[j]
			if arr[i] < common {
				common = arr[i]
			}
		}
		for k := 0; k < common; k++ {
			res = append(res, string('a'+i))
		}
	}
	return res
}
