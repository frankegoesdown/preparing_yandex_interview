package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
}

func compress(chars []byte) int {
	read, anchor, write := 0, 0, 0
	for read < len(chars) {
		fmt.Println(anchor)
		for anchor < len(chars) && chars[read] == chars[anchor] {
			anchor++
		}
		fmt.Println(write)
		chars[write] = chars[read]
		write++
		if anchor-read > 1 {
			for _, i := range []byte(strconv.Itoa(anchor - read)) {
				fmt.Println(i)
				chars[write] = i
				write++
			}
		}
		read = anchor
		fmt.Println(read)
		fmt.Println("========")
	}
	return write
}
