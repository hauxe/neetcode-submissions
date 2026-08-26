type MinHeap struct {
	heap [][2]int
}

func (this *MinHeap) Len() int {
	return len(this.heap)
}

func (this *MinHeap) Add(point [2]int) {
	this.heap = append(this.heap, point)
	child := len(this.heap) - 1
	parent := (child - 1) / 2
	for child != parent && this.heap[parent][0] > this.heap[child][0] { // [0] not [1]
		this.heap[parent], this.heap[child] = this.heap[child], this.heap[parent]
		child, parent = parent, (parent-1)/2
	}
}

func (this *MinHeap) Pop() [2]int {
	if this.Len() == 0 {
		return [2]int{-1, -1}
	}
	top := this.heap[0]
	this.heap[0] = this.heap[len(this.heap)-1]
	this.heap = this.heap[:len(this.heap)-1]

	i := 0
	for {
		smallest := i
		left, right := i*2+1, i*2+2
		if left < len(this.heap) && this.heap[left][0] < this.heap[smallest][0] { // [0] not [1]
			smallest = left
		}
		if right < len(this.heap) && this.heap[right][0] < this.heap[smallest][0] { // [0] not [1]
			smallest = right
		}

		if i == smallest {
			break
		}
		this.heap[i], this.heap[smallest] = this.heap[smallest], this.heap[i]
		i = smallest
	}

	return top
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return 0
	}

	visited := make([]bool, n)
	pq := &MinHeap{}
	pq.Add([2]int{0, 0}) // (cost=0, node=0)

	totalCost := 0
	nodesConnected := 0

	for nodesConnected < n {
		top := pq.Pop()
		if top[0] == -1 { // empty
			break
		}
		cost, u := top[0], top[1]
		if visited[u] {
			continue
		}
		visited[u] = true
		totalCost += cost
		nodesConnected++

		for v := 0; v < n; v++ {
			if !visited[v] {
				d := abs(points[u][0]-points[v][0]) + abs(points[u][1]-points[v][1])
				pq.Add([2]int{d, v})
			}
		}
	}

	return totalCost
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}