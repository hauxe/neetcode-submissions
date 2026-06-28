type MinStack struct {
	stack []int
	p     int
	min   int
	begin bool
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 10),
		p:     -1,
		min:   0,
		begin: true,
	}
}

func (this *MinStack) Push(val int) {
	if this.p == len(this.stack)-1 {
		olstack := this.stack
		this.stack = make([]int, len(olstack)+10)
		copy(this.stack, olstack)
	}
	if this.begin {
		this.begin = false
		this.min = val
	} else if val < this.min {
		newMin := 2*val - this.min
		this.min, val = val, newMin
	}
	this.p++
	this.stack[this.p] = val
}

func (this *MinStack) Pop() {
	if this.p < 0 {
		return
	}
	val := this.stack[this.p]
	if val < this.min {
		// detected previous min
		this.min = 2*this.min - val
	}
	this.p--
	if this.p < 0 {
		this.begin = true
	}
}

func (this *MinStack) Top() int {
	if this.p < 0 {
		return 0
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