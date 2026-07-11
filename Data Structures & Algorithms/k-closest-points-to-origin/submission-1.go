type KHeap struct {
	heap   []int
	points [][]int
	k      int
}

func (this *KHeap) GetVal(index int) int {
	point := this.points[index]
	return point[0]*point[0] + point[1]*point[1]
}

func (this *KHeap) Push(idx int) {
	this.heap = append(this.heap, idx)
	child, parent := len(this.heap)-1, (len(this.heap)-2)/2
	for child != parent && this.GetVal(this.heap[parent]) < this.GetVal(this.heap[child]) {
		this.heap[parent], this.heap[child] = this.heap[child], this.heap[parent]
		child, parent = parent, (parent-1)/2
	}
}

func (this *KHeap) Pop() int {
	top := this.heap[0]
	this.heap[0] = this.heap[len(this.heap)-1]
	this.heap = this.heap[:len(this.heap)-1]
	i := 0
	for {
		largest := i
		left, right := i*2+1, i*2+2
		if left < len(this.heap) && this.GetVal(this.heap[left]) > this.GetVal(this.heap[largest]) {
			largest = left
		}
		if right < len(this.heap) && this.GetVal(this.heap[right]) > this.GetVal(this.heap[largest]) {
			largest = right
		}
		if largest == i {
			break
		}
		this.heap[largest], this.heap[i] = this.heap[i], this.heap[largest]
		i = largest
	}

	return top
}

func (this *KHeap) Add(index int) {
	if len(this.heap) < this.k {
		this.Push(index)
	} else if this.GetVal(index) < this.GetVal(this.heap[0]) {
		this.Pop()
		this.Push(index)
	}
}

func kClosest(points [][]int, k int) [][]int {
	kHeap := &KHeap{
		heap:   make([]int, 0, k),
		points: points,
		k:      k,
	}
	for idx := range points {
		kHeap.Add(idx)
	}
	// fmt.Println(kHeap.heap)
	results := make([][]int, k)
	for i := 0; i < k; i++ {
		results[i] = points[kHeap.heap[i]]
	}
	return results
}