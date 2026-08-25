type MinHeap struct {
	heap []string
}

func (this *MinHeap) Add(dest string) {
	this.heap = append(this.heap, dest)
	child := len(this.heap) - 1
	parent := (child - 1) / 2
	for child != parent && this.heap[parent] > this.heap[child] {
		this.heap[parent], this.heap[child] = this.heap[child], this.heap[parent]
		child, parent = parent, (parent-1)/2
	}
}

func (this *MinHeap) Len() int {
	return len(this.heap)
}

func (this *MinHeap) Pop() string {
	if this.Len() == 0 {
		return ""
	}
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

		if i == smallest {
			break
		}
		this.heap[i], this.heap[smallest] = this.heap[smallest], this.heap[i]
		i = smallest
	}

	return top
}

func findItinerary(tickets [][]string) []string {
	graph := make(map[string]*MinHeap)
	for _, t := range tickets {
		if graph[t[0]] == nil {
			graph[t[0]] = &MinHeap{}
		}
		graph[t[0]].Add(t[1])
	}

	var res []string
	var dfs func(u string)
	dfs = func(u string) {
		for graph[u] != nil && graph[u].Len() > 0 {
			v := graph[u].Pop()
			dfs(v)
		}
		res = append(res, u) // post-order
	}

	dfs("JFK")

	// reverse res
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}