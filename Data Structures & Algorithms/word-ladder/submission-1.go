func ladderLength(beginWord string, endWord string, wordList []string) int {
	graph := make(map[string][][2]string)
	for _, word := range wordList {
		words := breakdown(word)
		if len(words) == 1 {
			graph[words[0]] = append(graph[words[0]], [2]string{words[0], word})
		} else {
			for j, stub := range words {
				for _, w := range words[0:j] {
					graph[stub] = append(graph[stub], [2]string{w, word})
				}
				for _, w := range words[j+1:] {
					graph[stub] = append(graph[stub], [2]string{w, word})
				}
			}
		}
	}
	fmt.Println(graph)
	queue := make([][2]string, 0)
	words := breakdown(beginWord)
	for _, word := range words {
		queue = append(queue, [2]string{word, beginWord})
	}
	level := 1
	visited := make(map[string]bool)
	fmt.Println(queue)
	for len(queue) > 0 {
		next := make([][2]string, 0)
		for _, word := range queue {
			if word[1] == endWord {
				return level
			}
			visited[word[1]] = true
			for _, vertice := range graph[word[0]] {
				if visited[vertice[1]] {
					continue
				}

				next = append(next, vertice)
			}
		}
		level++
		queue = next
	}
	return 0
}

func breakdown(word string) (result []string) {
	if len(word) == 1 {
		return []string{"*"}
	}
	for i := range word {
		stub := make([]byte, len(word))
		stub = append(stub, word[0:i]...)
		stub = append(stub, '*')
		stub = append(stub, word[i+1:]...)
		result = append(result, string(stub))
	}
	return
}