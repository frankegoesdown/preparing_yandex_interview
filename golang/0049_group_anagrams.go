package main

import "sort"

func main() {

}

type sortByte []byte

func (s sortByte) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortByte) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByte) Len() int {
	return len(s)
}

func SortString(str string) string {
	chars := []byte(str)
	sort.Sort(sortByte(chars))
	return string(chars)
}

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	if len(strs) == 0 {
		return result
	}

	dict := make(map[string]int)
	for _, s := range strs {
		sorted := SortString(s)
		if idx, ok := dict[sorted]; ok {
			result[idx] = append(result[idx], s)
		} else {
			dict[sorted] = len(result)
			result = append(result, []string{s})
		}
	}

	return result
}
