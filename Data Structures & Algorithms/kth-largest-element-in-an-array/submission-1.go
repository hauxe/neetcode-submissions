type KHeap struct {
	heap []int
	k    int
}

func (this *KHeap) Push(val int) {
	this.heap = append(this.heap, val)
	child, parent := len(this.heap)-1, (len(this.heap)-2)/2
	for child != parent && this.heap[parent] > this.heap[child] {
		this.heap[child], this.heap[parent] = this.heap[parent], this.heap[child]
		child, parent = parent, (parent-1)/2
	}
}

func (this *KHeap) Pop() int {
	top := this.heap[0]
	this.heap[0] = this.heap[len(this.heap)-1]
	this.heap = this.heap[:len(this.heap)-1]
	i := 0
	for {
		smallest := i
		left, right := i*2+1, i*2+2
		if left < len(this.heap) && this.heap[left] < this.heap[smallest] {
			smallest = left
		}

		if right < len(this.heap) && this.heap[right] < this.heap[smallest] {
			smallest = right
		}

		if smallest == i {
			break
		}
		this.heap[smallest], this.heap[i] = this.heap[i], this.heap[smallest]
		i = smallest
	}

	return top
}

func (this *KHeap) Add(val int) {
	if len(this.heap) < this.k {
		this.Push(val)
	} else if this.heap[0] < val {
		this.Pop()
		this.Push(val)
	}
}

func findKthLargest(nums []int, k int) int {
	kHeap := &KHeap{
		heap: make([]int, 0, k),
		k:    k,
	}
	for _, num := range nums {
		kHeap.Add(num)
	}
	return kHeap.heap[0]
}