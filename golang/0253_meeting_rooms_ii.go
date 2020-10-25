package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *minHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func minMeetingRooms(ivs [][]int) int {

	if len(ivs) == 0 {
		return 0
	}

	sort.Slice(ivs, func(i, j int) bool {
		return ivs[i][0] < ivs[j][0]
	})

	h := minHeap{ivs[0][1]}

	for _, iv := range ivs[1:] {
		if h[0] <= iv[0] {
			heap.Pop(&h)
		}
		heap.Push(&h, iv[1])
	}

	return h.Len()

}
