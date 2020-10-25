package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	var subs [256]bool
	i, j, result := 0, 0, 0
	max := len(s)
	for i < max && j < max {
		if !subs[s[j]] {
			subs[int8(s[j])] = true
			j++
			if result > j-i {
				result = result
			} else {
				result = j - i
			}
		} else {
			subs[int8(s[i])] = false
			i++
		}
	}

	return result
}

func lengthOfLongestSubstring2(s string) int {
	var chPosition [256]int // [0, 0, 0, ...]
	maxLength, substringLen, lastRepeatPos := 0, 0, 0

	for i := 0; i < len(s); i++ {
		if pos := chPosition[s[i]]; pos > 0 {
			// record current substring length
			fmt.Printf("maxlen %d \n", substringLen)
			maxLength = max(substringLen, maxLength)
			// update characters position
			chPosition[s[i]] = i + 1
			// update last repeat character position
			lastRepeatPos = max(pos, lastRepeatPos)

			// update the current substring from last repeat character
			substringLen = i + 1 - lastRepeatPos
		} else {
			substringLen++
			chPosition[s[i]] = i + 1
		}
		fmt.Println(i)
		fmt.Println(string(s[i]))
		fmt.Println(maxLength)
		fmt.Println(substringLen)
		fmt.Println(chPosition)
		fmt.Println("==============")
	}

	return max(maxLength, substringLen)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
