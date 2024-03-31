package Trie

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

// Constructor Initialize your data structure here
func Constructor() Trie {
	return Trie{}
}

// Insert Inserts a word into the trie
func (this *Trie) Insert(word string) {
	node := this
	for _, v := range word {
		v -= 'a'
		if node.next[v] == nil {
			node.next[v] = &Trie{}
		}
		node = node.next[v]
	}
	node.isEnd = true
}

// Search Returns if the word is in the trie.
func (this *Trie) Search(word string) bool {
	node := this
	for _, v := range word {
		if node = node.next[v-'a']; node == nil {
			return false
		}
	}
	return node.isEnd
}

// StartsWith Returns if there is any word in the trie that starts with the given prefix.
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, v := range prefix {
		if node = node.next[v-'a']; node == nil {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
