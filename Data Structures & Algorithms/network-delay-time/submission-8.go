type MinHeap struct {
	heap [][3]int // [from, to, dist]
}

func (this *MinHeap) Add(val [3]int) {
	this.heap = append(this.heap, val)
	child := len(this.heap) - 1
	parent := (child - 1) / 2
	for child != parent && this.heap[parent][2] > this.heap[child][2] {
		this.heap[child], this.heap[parent] = this.heap[parent], this.heap[child]
		child, parent = parent, (parent-1)/2
	}
}

func (this *MinHeap) Pop() (top [3]int) {
	if len(this.heap) == 0 {
		return top
	}
	top = this.heap[0]
	this.heap[0] = this.heap[len(this.heap)-1]
	this.heap = this.heap[:len(this.heap)-1]

	i := 0
	for {
		smallest := i
		left, right := smallest*2+1, smallest*2+2
		if left < len(this.heap) && this.heap[left][2] < this.heap[smallest][2] {
			smallest = left
		}
		if right < len(this.heap) && this.heap[right][2] < this.heap[smallest][2] {
			smallest = right
		}
		if smallest == i {
			break
		}
		this.heap[i], this.heap[smallest] = this.heap[smallest], this.heap[i]
		i = smallest
	}
	return
}

func (this *MinHeap) Len() int {
	return len(this.heap)
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][][2]int, n+1)
	for _, t := range times {
		graph[t[0]] = append(graph[t[0]], [2]int{t[1], t[2]})
	}

	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[k] = 0

	visited := make([]bool, n+1)

	h := &MinHeap{}
	h.Add([3]int{k, k, 0})

	for h.Len() > 0 {
		cur := h.Pop()
		node, d := cur[1], cur[2]
		if visited[node] {
			continue
		}
		visited[node] = true

		for _, nb := range graph[node] {
			next, w := nb[0], nb[1]
			if !visited[next] && d+w < dist[next] {
				dist[next] = d + w
				h.Add([3]int{node, next, dist[next]})
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt {
			return -1
		}
		ans = max(ans, dist[i])
	}
	return ans
}