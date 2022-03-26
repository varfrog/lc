package trie208

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie_Search(t1 *testing.T) {
	trie := Constructor()
	trie.Insert("app")
	trie.Insert("apple")
	trie.Insert("jam")
	trie.Insert("jumbo")

	assert.True(t1, trie.Search("app"))
	assert.True(t1, trie.Search("apple"))
	assert.True(t1, trie.Search("jam"))
	assert.True(t1, trie.Search("jumbo"))

	assert.False(t1, trie.Search("ap"))
	assert.False(t1, trie.Search("jumb"))
}

func TestTrie_StartsWith(t1 *testing.T) {
	trie := Constructor()
	trie.Insert("app")
	trie.Insert("apple")
	trie.Insert("jam")
	trie.Insert("jumbo")

	assert.True(t1, trie.StartsWith("app"))
	assert.True(t1, trie.StartsWith("apple"))
	assert.True(t1, trie.StartsWith("jam"))
	assert.True(t1, trie.StartsWith("jumbo"))

	assert.True(t1, trie.StartsWith("ap"))
	assert.True(t1, trie.StartsWith("jumb"))
}
