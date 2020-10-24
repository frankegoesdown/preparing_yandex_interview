package main

import "math/rand"

func main() {

}

type RandomizedSet struct {
	m       map[int]int
	nums    []int
	lastIdx int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{m: make(map[int]int), nums: make([]int, 0, 0), lastIdx: -1}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}

	if this.lastIdx+1 >= len(this.nums) {
		this.nums = append(this.nums, val)
	} else {
		this.nums[this.lastIdx+1] = val
	}

	this.m[val] = this.lastIdx + 1
	this.lastIdx++
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}

	idx := this.m[val]
	delete(this.m, val)

	if idx != this.lastIdx {
		this.m[this.nums[this.lastIdx]] = idx
		this.nums[idx] = this.nums[this.lastIdx]
	}

	this.lastIdx--

	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(this.lastIdx+1)]
}
