package trie

type Trie struct {
	end      bool
	children map[rune]*Trie
}

func NewTrie() Trie {
	return Trie{}
}

func (t *Trie) Insert(s string) {
	node := t
	for _, char := range s {
		if node.children == nil {
			node.children = make(map[rune]*Trie)
		}

		child, ok := node.children[char]
		if !ok {
			child = &Trie{}
			node.children[char] = child
		}

		node = child
	}
	node.end = true
}

func (t *Trie) Search(s string) bool {
	node := t
	for _, char := range s {
		child, ok := node.children[char]
		if !ok {
			return false
		}
		node = child
	}
	return node.end
}
