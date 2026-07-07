type TreeNode struct {
	Children [26]*TreeNode
	IsEnd    bool
}

type WordDictionary struct {
	Root *TreeNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		Root: &TreeNode{},
	}
}

func (this *WordDictionary) AddWord(word string) {
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

func (this *WordDictionary) Search(word string) bool {
	return this.doSearch(word, this.Root)
}
func (this *WordDictionary) doSearch(word string, root *TreeNode) bool {
	for i, c := range word {
		if c == '.' {
			for _, root := range root.Children {
				if root != nil && this.doSearch(word[i+1:], root) {
					return true
				}
			}
			return false
		}
		idx := int(c - 'a')
		if root.Children[idx] == nil {
			return false
		}
		root = root.Children[idx]
	}
	return root.IsEnd
}