type MinStack struct {
	stack []int
	p     int
	min   int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 10),
		p:     -1,
		min:   0,
	}
}

func (this *MinStack) Push(val int) {
	this.p++
	if this.p >= len(this.stack) {
		newSlice := make([]int, len(this.stack)+10)
		copy(newSlice, this.stack)
		this.stack = newSlice
	}
	if this.p == 0 || val < this.min {
		prevMin := this.min
		this.min = val
		if this.p > 0 {
			val = 2*val - prevMin
		}
	}
	this.stack[this.p] = val
}

func (this *MinStack) Pop() {
	if this.p < 0 {
		return
	}
	val := this.stack[this.p]
	this.p--
	if val < this.min {
		this.min = 2*this.min - val
	}
}

func (this *MinStack) Top() int {
	if this.p < 0 {
		return -1
	}
	val := this.stack[this.p]
	if val < this.min {
		return this.min
	}
	return val
}

func (this *MinStack) GetMin() int {
	return this.min
}