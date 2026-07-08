type TreeNode struct {
	Children [26]*TreeNode
	Word     string
	IsEnd    bool
}

type PrefixTree struct {
	Root *TreeNode
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
	node.Word = word
	node.IsEnd = true
}

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 {
		return nil
	}
	trie := PrefixTree{
		Root: &TreeNode{},
	}
	for _, word := range words {
		trie.Insert(word)
	}
	// trie.Root.Print()
	results := map[string]struct{}{}
	for row := range board {
		for col := range board[row] {
			findWord(board, row, col, trie.Root, results)
		}
	}
	list := make([]string, 0, len(results))
	for k := range results {
		list = append(list, k)
	}
	return list
}

var directions = [][]int{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

func findWord(board [][]byte, row int, col int, root *TreeNode, results map[string]struct{}) {
	idx := board[row][col] - 'a'
	if root.Children[idx] == nil {
		return
	}
	if root.Children[idx].IsEnd {
		results[root.Children[idx].Word] = struct{}{}
	}
	c := board[row][col]
	board[row][col] = '#'
	for _, direction := range directions {
		newRow, newCol := row+direction[0], col+direction[1]
		if newRow >= 0 && newRow < len(board) && newCol >= 0 && newCol < len(board[0]) && board[newRow][newCol] != '#' {
			findWord(board, newRow, newCol, root.Children[idx], results)
		}
	}
	board[row][col] = c
}