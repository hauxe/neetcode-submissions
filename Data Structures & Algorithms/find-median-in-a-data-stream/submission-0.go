
type MaxHeap struct {
	heap []int
}

func (this *MaxHeap) Push(val int) {
	this.heap = append(this.heap, val)
	child := len(this.heap) - 1
	parent := (child - 1) / 2
	for child != parent && this.heap[parent] < this.heap[child] {
		this.heap[child], this.heap[parent] = this.heap[parent], this.heap[child]
		child, parent = parent, (parent-1)/2
	}
}

func (this *MaxHeap) Pop() int {
	top := this.heap[0]
	this.heap[0] = this.heap[len(this.heap)-1]
	this.heap = this.heap[:len(this.heap)-1]
	i := 0
	for {
		largest := i
		left, right := i*2+1, i*2+2
		if left < len(this.heap) && this.heap[left] > this.heap[largest] {
			largest = left
		}
		if right < len(this.heap) && this.heap[right] > this.heap[largest] {
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

type MinHeap struct {
	heap []int
}

func (this *MinHeap) Push(val int) {
	this.heap = append(this.heap, val)
	child := len(this.heap) - 1
	parent := (child - 1) / 2
	for child != parent && this.heap[parent] > this.heap[child] {
		this.heap[child], this.heap[parent] = this.heap[parent], this.heap[child]
		child, parent = parent, (parent-1)/2
	}
}

func (this *MinHeap) Pop() int {
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
		this.heap[i], this.heap[smallest] = this.heap[smallest], this.heap[i]
		i = smallest
	}

	return top
}

type MedianFinder struct {
	minHeap *MinHeap
	maxHeap *MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: &MinHeap{},
		maxHeap: &MaxHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.minHeap.heap) == 0 || num > this.minHeap.heap[0] {
		this.minHeap.Push(num)
	} else {
		this.maxHeap.Push(num)
	}
	if len(this.maxHeap.heap)-len(this.minHeap.heap) > 1 {
		top := this.maxHeap.Pop()
		this.minHeap.Push(top)
	}
	if len(this.minHeap.heap)-len(this.maxHeap.heap) > 1 {
		top := this.minHeap.Pop()
		this.maxHeap.Push(top)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.maxHeap.heap) == 0 && len(this.minHeap.heap) == 0 {
		return 0
	}
	if len(this.maxHeap.heap) == len(this.minHeap.heap) {
		return float64(this.maxHeap.heap[0]+this.minHeap.heap[0]) / 2
	}
	if len(this.maxHeap.heap) > len(this.minHeap.heap) {
		return float64(this.maxHeap.heap[0])
	}
	return float64(this.minHeap.heap[0])
}