package main

// 208. 实现 Trie (前缀树)
type Trie struct {
	isEnd    bool
	children [26]*Trie
}

func TrieConstructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, w := range word {
		n := w - 'a'
		if node.children[n] == nil {
			node.children[n] = &Trie{}
		}
		node = node.children[n]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, w := range word {
		n := w - 'a'
		if node.children[n] == nil {
			return false
		}
		node = node.children[n]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, w := range prefix {
		n := w - 'a'
		if node.children[n] == nil {
			return false
		}
		node = node.children[n]
	}
	return true
}

func main() {

}
