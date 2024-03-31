package Trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")

	t.Logf("trie.Search(\"apple\"): %v", trie.Search("apple"))

	t.Logf("trie.Search(\"app\"): %v", trie.Search("app"))

	trie.StartsWith("app")
	trie.Insert("app")
	t.Logf("trie.Search(\"app\"): %v", trie.Search("app"))
}
