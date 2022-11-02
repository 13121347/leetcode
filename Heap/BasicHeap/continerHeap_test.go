package BasicHeap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestIntHeap_Pop(t *testing.T) {
	h := &IntHeap{3, 7, 2, 5, 1}
	heap.Init(h)
	for h.Len() > 0 {
		fmt.Println(h.Pop())
	}

}
