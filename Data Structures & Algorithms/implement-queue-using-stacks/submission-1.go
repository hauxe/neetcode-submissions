type MyQueue struct {
	main *MyStack
	push *MyStack
}

func Constructor() MyQueue {
	return MyQueue{
		main: &MyStack{},
		push: &MyStack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.push.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.main.Empty() {
		for !this.push.Empty() {
			this.main.Push(this.push.Pop())
		}
	}
	return this.main.Pop()
}

func (this *MyQueue) Peek() int {
	if this.main.Empty() {
		for !this.push.Empty() {
			this.main.Push(this.push.Pop())
		}
	}
	return this.main.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.main.Empty() && this.push.Empty()
}

type MyStack struct {
	stack []int
}

func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyStack) Pop() int {
	x := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return x
}

func (this *MyStack) Peek() int {
	return this.stack[len(this.stack)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.stack) == 0
}

func (this *MyStack) Size() int {
	return len(this.stack)
}