package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
}
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	ans := ""
	ansLength := 0

	for i := 0; i < len(s); i++ {
		for j := len(s); j > i+ansLength; j-- {
			subString := s[i:j]
			fmt.Println(i)
			fmt.Println(j)
			fmt.Println(ansLength)
			fmt.Println(subString)
			if isPalindrome(subString) && len(subString) > ansLength {
				ans = subString
				ansLength = len(subString)
			}
		}
	}

	return ans
}

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}

	return true
}
