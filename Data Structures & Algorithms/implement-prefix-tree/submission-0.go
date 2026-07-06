type TreeNode struct {
	Children [26]*TreeNode
	IsEnd    bool
}

type PrefixTree struct {
	Root *TreeNode
}

func Constructor() PrefixTree {
	return PrefixTree{
		Root: &TreeNode{},
	}
}

func (this *PrefixTree) Insert(word string) {
	node := this.Root
	for _, c := range word {
		idx := int(c - 'a')
		if node.Children[idx] == nil {
			node.Children[idx] = &TreeNode{}
		}
		node = node.Children[idx]
	}
	node.IsEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	node := this.Root
	for _, c := range word {
		idx := int(c - 'a')
		if node.Children[idx] == nil {
			return false
		}
		node = node.Children[idx]
	}
	return node.IsEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	node := this.Root
	for _, c := range prefix {
		idx := int(c - 'a')
		if node.Children[idx] == nil {
			return false
		}
		node = node.Children[idx]
	}
	return true
}
