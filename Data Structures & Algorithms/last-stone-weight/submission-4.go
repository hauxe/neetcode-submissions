type Heap struct {
	heap []int
	k    int
}

func (this *Heap) Push(val int) {
	this.heap = append(this.heap, val)
	child := len(this.heap) - 1
	parent := (child - 1) / 2
	for parent != child && this.heap[parent] < this.heap[child] {
		this.heap[child], this.heap[parent] = this.heap[parent], this.heap[child]
		child, parent = parent, (parent-1)/2
	}
}

func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return -1
	}
	h := &Heap{
		heap: make([]int, 0, len(stones)),
	}
	for _, stone := range stones {
		h.Push(stone)
	}
	fmt.Println(h.heap)
	for len(h.heap) > 1 {
		if len(h.heap) == 2 {
			h.heap[0] = h.heap[0] - h.heap[1]
			h.heap = h.heap[:1]
			break
		}
		if h.heap[1] < h.heap[2] {
			h.heap[1], h.heap[2] = h.heap[2], h.heap[1]
		}
		h.heap[1] = h.heap[0] - h.heap[1]
		k := &Heap{
			heap: make([]int, 0, len(h.heap)-1),
		}
		for i := 1; i < len(h.heap); i++ {
			k.Push(h.heap[i])
		}
		h = k

		fmt.Println("smashed", h.heap)
	}
	return h.heap[0]
}