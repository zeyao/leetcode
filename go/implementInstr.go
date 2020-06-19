package main

func main() {

}

func strStr(h string, n string) int {
	base := 256
	mod := 1000000007
	if len(n) == 0 {
		return 0
	}
	if len(h) < len(n) {
		return -1
	}
	hash := 0
	targetHash := 0
	pow := 1

	for i := 0; i < len(n); i++ {
		hash = (base*hash + int(h[i])) % mod
		targetHash = (base*targetHash + int(n[i])) % mod
		pow = pow * base % mod
	}
	if hash == targetHash {
		return 0
	}
	for i := len(n); i < len(h); i++ {
		hash = (hash*base + int(h[i]) - int(h[i-len(n)])*pow%mod) % mod
		if hash == targetHash {
			return i - len(n) + 1
		}
	}
	return -1
}
