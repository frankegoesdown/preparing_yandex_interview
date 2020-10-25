package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}

func checkInclusion(s1 string, s2 string) bool {
	wLen := len(s1)
	chars1, chars2 := [26]int{}, [26]int{}
	for i := 0; i < len(s1); i++ {
		chars1[int(s1[i]-'a')]++
	}
	for i := 0; i < len(s2); i++ {
		chars2[int(s2[i]-'a')]++
		if i-wLen >= 0 {
			chars2[int(s2[i-wLen]-'a')]--
		}

		if compare(chars1, chars2) {
			return true
		}
	}

	return false
}

func compare(chars1, chars2 [26]int) bool {
	for i := 0; i < 26; i++ {
		if chars1[i] != chars2[i] {
			return false
		}
	}
	return true
}
