type Node struct {
	Prev, Next *Node
	Val        int
	Key        int
}

type LRUCache struct {
	m          map[int]*Node
	size       int
	capacity   int
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:        make(map[int]*Node, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if this.m[key] != nil {
		this.move(this.m[key])
		return this.m[key].Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node := this.m[key]
	if node == nil {
		if this.size == this.capacity {
			node = this.tail
			delete(this.m, node.Key)
			node.Key, node.Val = key, value
			this.m[key] = node
			this.move(node)
			return
		} else {
			node = &Node{
				Val: value,
				Key: key,
			}
		}
		this.m[key] = node
		this.size++
		if this.head == nil {
			this.head, this.tail = node, node
		} else {
			this.head, this.head.Prev, node.Prev, node.Next = node, node, nil, this.head
		}
		return
	}
	node.Val = value
	this.move(node)
}

func (this *LRUCache) move(node *Node) {
	if node == this.head {
		return
	}
	if node == this.tail {
		this.head, this.head.Prev, node.Prev, node.Next, node.Prev.Next, this.tail = node, node, nil, this.head, node.Next, node.Prev
		return
	}
	this.head, this.head.Prev, node.Prev, node.Next, node.Prev.Next, node.Next.Prev = node, node, nil, this.head, node.Next, node.Prev
}