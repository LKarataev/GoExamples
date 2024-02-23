package SortPresents

import (
	"container/heap"
	"errors"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap []Present

func getNCoolestPresents(presents PresentHeap, n int) (PresentHeap, error) {
	if n < 0 || n > len(presents) {
		return nil, errors.New("invalid value for n")
	}
	h := &PresentHeap{}
	heap.Init(h)
	for _, p := range presents {
		heap.Push(h, p)
	}
	coolest := make(PresentHeap, n)
	for i := 0; i < n; i++ {
		coolest[i] = heap.Pop(h).(Present)
	}
	return coolest, nil
}

func (h PresentHeap) Len() int {
	return len(h)
}

func (h PresentHeap) Less(i, j int) bool {
	if h[i].Value == h[j].Value {
		return h[i].Size < h[j].Size
	}
	return h[i].Value > h[j].Value
}

func (h PresentHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PresentHeap) Push(x interface{}) {
	*h = append(*h, x.(Present))
}

func (h *PresentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
