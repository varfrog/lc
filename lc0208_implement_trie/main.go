package trie208

import "fmt"

type Trie struct {
	isEnd    bool // end of word (not leaf node!)
	children [26]*Trie
}

func Constructor() Trie {
	return Trie{
		isEnd:    false,
		children: [26]*Trie{},
	}
}

func (t *Trie) Insert(word string) {
	runes := []rune(word)
	curTrie := t
	for _, r := range runes {
		childTrie := curTrie.children[r-97]
		if childTrie == nil {
			newTrie := Constructor()
			curTrie.children[r-97] = &newTrie
			curTrie = &newTrie
		} else {
			curTrie = childTrie
		}
	}
	curTrie.isEnd = true
}

func (t *Trie) Search(word string) bool {
	runes := []rune(word)
	curTrie := t
	for _, r := range runes {
		childTrie := curTrie.children[r-97]
		if childTrie == nil {
			return false
		}
		curTrie = childTrie
	}

	return curTrie.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	runes := []rune(prefix)
	curTrie := t
	for _, r := range runes {
		childTrie := curTrie.children[r-97]
		if childTrie == nil {
			return false
		}
		curTrie = childTrie
	}

	return true
}

func main() {
	trie := Constructor()
	trie.Insert("le")
	trie.Insert("lex")

	trie.Insert("a")
	trie.Insert("abc")

	fmt.Println("done")
}
