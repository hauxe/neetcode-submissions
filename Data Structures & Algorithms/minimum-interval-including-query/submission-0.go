type KHeap struct {
	heap      []int
	intervals [][]int
}

func (this *KHeap) Value(index int) int {
	return this.intervals[index][1] - this.intervals[index][0] + 1
}

func (this *KHeap) End(index int) int {
	return this.intervals[index][1]
}

func (this *KHeap) Push(index int) {
	this.heap = append(this.heap, index)
	child := len(this.heap) - 1
	parent := (child - 1) / 2
	for child != parent && this.Value(this.heap[parent]) > this.Value(this.heap[child]) {
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
		smallest := i
		left, right := i*2+1, i*2+2
		if left < len(this.heap) && this.Value(this.heap[left]) < this.Value(this.heap[smallest]) {
			smallest = left
		}
		if right < len(this.heap) && this.Value(this.heap[right]) < this.Value(this.heap[smallest]) {
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

func (this *KHeap) Top() int {
	if len(this.heap) == 0 {
		return -1
	}
	return this.heap[0]
}

func minInterval(intervals [][]int, queries []int) []int {
	result := make([]int, len(queries))
	m := make(map[int][]int, len(queries))
	for i := range queries {
		m[queries[i]] = append(m[queries[i]], i)
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	sort.Slice(queries, func(i, j int) bool {
		return queries[i] < queries[j]
	})
	kHeap := &KHeap{
		intervals: intervals,
	}
	pos := 0
	for i := range queries {
		l := -1
		for ; pos < len(intervals); pos++ {
			if queries[i] < intervals[pos][0] {
				break
			}
			kHeap.Push(pos)
		}
		for index := kHeap.Top(); index >= 0; index = kHeap.Top() {
			if kHeap.End(index) < queries[i] {
				kHeap.Pop()
				continue
			}
			l = kHeap.Value(index)
			break
		}
		for _, j := range m[queries[i]] {
			result[j] = l
		}
	}
	return result
}