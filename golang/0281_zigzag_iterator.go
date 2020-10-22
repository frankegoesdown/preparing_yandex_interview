package main

func main() {

}

type ZigzagIterator struct {
	arrs [][]int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
	return &ZigzagIterator{
		arrs: [][]int{v1, v2},
	}
}

func (this *ZigzagIterator) next() int {
	if !this.hasNext() {
		return -1
	}

	top := this.arrs[0][0]
	this.arrs[0] = this.arrs[0][1:]

	if len(this.arrs) > 1 { // swap to end
		this.arrs[0], this.arrs[1] = this.arrs[1], this.arrs[0]
	}

	return top
}

func (this *ZigzagIterator) hasNext() bool {
	for len(this.arrs) > 0 && len(this.arrs[0]) == 0 {
		this.arrs = this.arrs[1:] // remove empty
	}

	return len(this.arrs) != 0
}
