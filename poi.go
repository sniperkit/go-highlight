/*
Sniperkit-Bot
- Status: analyzed
*/

package highlight

import (
	"container/heap"
	"log"
)

type poi struct {
	i         int
	start     bool
	class     string
	highlight *highlight
}

type poiHeap []poi

func (h poiHeap) Len() int { return len(h) }
func (h poiHeap) Less(i, j int) bool {
	if h[i].i == h[j].i {
		return h[i].start && !h[j].start
	}
	return -h[i].i < -h[j].i
}
func (h poiHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *poiHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(poi))
}

func (h *poiHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *poiHeap) Remove(p poi) {
	for i, p2 := range *h {
		if p2.highlight == p.highlight {
			heap.Remove(h, i)
			return
		}
	}
	log.Fatalf("heap, %+v p %+v", *h, p)
}

func (h *poiHeap) Peek() poi {
	return (*h)[0]
}
