type FreqStack struct {
	m      map[int]int
	stacks map[int][]int
	max    int
}

func Constructor() FreqStack {
	return FreqStack{
		m:      make(map[int]int),
		stacks: make(map[int][]int),
	}
}

func (this *FreqStack) Push(val int) {
	this.m[val]++
	this.stacks[this.m[val]] = append(this.stacks[this.m[val]], val)
	this.max = max(this.max, this.m[val])
}

func (this *FreqStack) Pop() int {
	n := len(this.stacks[this.max])
	top := this.stacks[this.max][n-1]
	this.stacks[this.max] = this.stacks[this.max][:n-1]
	if n == 1 {
		this.max--
	}
	this.m[top]--
	return top
}