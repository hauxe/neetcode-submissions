type StockSpanner struct {
	stack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	val := [2]int{price, 1}
	for !this.Empty() {
		top := this.Top()
		if top[0] > price {
			break
		}
		this.Pop()
		val[1] += top[1]
	}
	this.Push(val)
	return val[1]
}

func (this *StockSpanner) Top() [2]int {
	return this.stack[len(this.stack)-1]
}

func (this *StockSpanner) Pop() [2]int {
	top := this.Top()
	this.stack = this.stack[:len(this.stack)-1]
	return top
}

func (this *StockSpanner) Push(val [2]int) {
	this.stack = append(this.stack, val)
}

func (this *StockSpanner) Empty() bool {
	return len(this.stack) == 0
}