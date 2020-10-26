package main

func main() {

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chars := [26]int{}
	for i := 0; i < len(s); i++ {
		chars[int(s[i]-'a')]++
		chars[int(t[i]-'a')]--
	}
	for i := 0; i < 26; i++ {
		if chars[i] != 0 {
			return false
		}
	}

	return true
}
