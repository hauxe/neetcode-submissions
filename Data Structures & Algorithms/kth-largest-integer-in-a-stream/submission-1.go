type KthLargest struct {
	heap []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	s := KthLargest{
		heap: make([]int, 0, k),
		k:    k,
	}
	for _, n := range nums {
		s.Add(n)
	}
	return s
}

func (this *KthLargest) Push(val int) {
	this.heap = append(this.heap, val)
	l := len(this.heap) - 1
	parent := (l - 1) / 2
	for parent != l && this.heap[parent] > val {
		this.heap[parent], this.heap[l] = this.heap[l], this.heap[parent]
		l, parent = parent, (parent-1)/2
	}
	// fmt.Println("push", this.heap)
}

func (this *KthLargest) Pop() {
	this.heap[0], this.heap[len(this.heap)-1] = this.heap[len(this.heap)-1], this.heap[0]
	this.heap = this.heap[:len(this.heap)-1]
	i := 0
	for {
		smallest := i
		left := i*2 + 1
		right := i*2 + 2
		if left < len(this.heap) && this.heap[smallest] > this.heap[left] {
			smallest = left
		}
		if right < len(this.heap) && this.heap[smallest] > this.heap[right] {
			smallest = right
		}
		if smallest == i {
			// found
			break
		}
		this.heap[i], this.heap[smallest] = this.heap[smallest], this.heap[i]
		i = smallest
	}
	// fmt.Println("pop", this.heap)
}

func (this *KthLargest) Add(val int) int {
	if len(this.heap) < this.k {
		this.Push(val)
	} else if val > this.heap[0] {
		this.Pop()
		this.Push(val)
		// fmt.Println("processing", val)
	}
	return this.heap[0]
}