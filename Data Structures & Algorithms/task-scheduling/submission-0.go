type KHeap struct {
	heap  []int
	tasks [26]int
	k     int
}

func (this *KHeap) Value(index int) int {
	if index < 0 {
		return 0
	}
	return this.tasks[index]
}

func (this *KHeap) Push(index int) {
	this.heap = append(this.heap, index)
	child, parent := len(this.heap)-1, (len(this.heap)-2)/2
	for child != parent && this.Value(this.heap[parent]) < this.Value(this.heap[child]) {
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
		if left < len(this.heap) && this.Value(this.heap[left]) > this.Value(this.heap[largest]) {
			largest = left
		}
		if right < len(this.heap) && this.Value(this.heap[right]) > this.Value(this.heap[largest]) {
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

type Queue struct {
	queue [][2]int
}

func (this *Queue) Push(val [2]int) {
	this.queue = append(this.queue, val)
}

func (this *Queue) Top(time int) int {
	if len(this.queue) == 0 || this.queue[0][1]-time > 0 {
		return -1
	}
	index := this.queue[0][0]
	this.queue = this.queue[1:]
	return index
}

func leastInterval(tasks []byte, n int) int {
	m := [26]int{}
	for _, task := range tasks {
		m[int(task-'A')]++
	}
	kHeap := &KHeap{
		heap:  []int{},
		k:     -1, // no limit on total tasks
		tasks: m,
	}
	delayQueue := &Queue{}
	for index, count := range m {
		if count > 0 {
			kHeap.Push(index)
		}
	}
	fmt.Println(kHeap.heap, m)
	cycles := 0
	for len(kHeap.heap) > 0 || len(delayQueue.queue) > 0 {
		topIndex := -1
		if len(kHeap.heap) > 0 {
			topIndex = kHeap.Pop()
			kHeap.tasks[topIndex]--
		}
		cycles++
		if repushIndex := delayQueue.Top(cycles); repushIndex >= 0 {
			kHeap.Push(repushIndex)
		}
		if topIndex >= 0 && kHeap.tasks[topIndex] > 0 {
			delayQueue.Push([2]int{topIndex, cycles + n})
		}
	}
	return cycles
}