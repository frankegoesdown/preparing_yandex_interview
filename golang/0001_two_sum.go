package main

import "fmt"

func main() {
    fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
 
    index := make(map[int]int, len(nums))
    
    for idx, i := range nums {
 
        if j, ok := index[target-i]; ok {
        
            return []int{j, idx}
        
        }
        
        index[i] = idx
    } 

    return nil
}