type Heap struct {
	heap []int
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

func (this *Heap) Pop() int {
	top := this.heap[0]
	this.heap[0] = this.heap[len(this.heap)-1]
	this.heap = this.heap[:len(this.heap)-1]
	i := 0
	for {
		largest := i
		left := largest*2 + 1
		right := largest*2 + 2
		if left < len(this.heap) && this.heap[largest] < this.heap[left] {
			largest = left
		}
		if right < len(this.heap) && this.heap[largest] < this.heap[right] {
			largest = right
		}
		if largest == i {
			break
		}
		this.heap[i], this.heap[largest] = this.heap[largest], this.heap[i]
		i = largest
	}
	return top
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
		x, y := h.Pop(), h.Pop()
		if x != y {
			h.Push(x - y)
		}
		fmt.Println("smashed", h.heap)
	}
	if len(h.heap) == 0 {
		return 0
	}
	return h.heap[0]
}