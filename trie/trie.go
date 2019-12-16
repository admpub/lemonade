package trie

type Trie struct {
	Root *Node
}

func New() *Trie {
	var r rune

	return &Trie{Root: NewNode(r)}
}

func (trie *Trie) Add(word string) {
	if len(word) <=0 {
		return
	}

	chars := []rune(word)
	node := trie.Root

	for _, char := range chars {
		node = node.put(char)
	}

	node.End = true
}

func (trie *Trie) Scan(word string) bool {
	chars := []rune(word)
	node := trie.Root

	for _, char := range chars {
		found := node.get(char)

		if found == nil {
			continue
		}
		node = found

		if node.End {
			return true
		}
	}

	return false
}

func (trie *Trie) Search(word string) (matches []string) {
	if len(word) == 0 {
		return
	}

	chars := []rune(word)
	length := len(chars)

	for pointer := 0; pointer < length; pointer++ {
		node := trie.Root
		matchFlag := 0
		flag := false

		for i := pointer; i < length; i++ {
			child := node.get(chars[i])

			if child == nil {
				break
			}

			node = child
			matchFlag++

			if node.End {
				flag = true
				continue
			}
		}

		if !flag {
			matchFlag = 0
		}
		if matchFlag == 0 {
			continue
		}

		matches = append(matches, string(chars[pointer:pointer+matchFlag]))
		pointer += matchFlag - 1
	}

	return unique(matches)
}

func unique(elements []string) (result []string) {
	encountered := map[string]bool{}

	for v:= range elements {
		encountered[elements[v]] = true
	}

	for key := range encountered {
		result = append(result, key)
	}

	return
}