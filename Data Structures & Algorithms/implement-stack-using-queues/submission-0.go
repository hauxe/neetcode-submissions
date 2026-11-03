type MyStack struct {
	main *MyQueue
	push *MyQueue
}

func Constructor() MyStack {
	return MyStack{
		main: &MyQueue{},
		push: &MyQueue{},
	}
}

func (this *MyStack) Push(x int) {
	this.push.Push(x)
	for !this.main.Empty() {
		this.push.Push(this.main.Pop())
	}
	this.main, this.push = this.push, this.main
}

func (this *MyStack) Pop() int {
	return this.main.Pop()
}

func (this *MyStack) Top() int {
	return this.main.Top()
}

func (this *MyStack) Empty() bool {
	return this.main.Empty()
}

type MyQueue struct {
	queue []int
}

func (q *MyQueue) Push(x int) {
	q.queue = append(q.queue, x)
}
func (q *MyQueue) Top() int {
	return q.queue[0]
}

func (q *MyQueue) Pop() int {
	top := q.queue[0]
	q.queue = q.queue[1:]
	return top
}
func (q *MyQueue) Empty() bool {
	return len(q.queue) == 0
}