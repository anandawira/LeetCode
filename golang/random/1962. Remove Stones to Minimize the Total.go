package random

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minStoneSum(piles []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	var sum int

	for _, pile := range piles {
		heap.Push(h, pile)
		sum += pile
	}

	for i := 0; i < k; i++ {
		popped := heap.Pop(h).(int)
		sum -= (popped / 2)
		heap.Push(h, popped-(popped/2))
	}

	return sum
}
