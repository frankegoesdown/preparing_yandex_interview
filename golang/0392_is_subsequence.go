package main

func main() {

}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	for a, b := 0, 0; b < len(t); b++ {
		if s[a] == t[b] {
			a++
		}
		if a == len(s) {
			return true
		}
	}
	return false
}

func isSubsequence2(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	for i := 0; i < len(t); i++ {
		if t[i] == s[0] {
			return isSubsequence(s[1:], t[i+1:])
		}
	}

	return false
}
