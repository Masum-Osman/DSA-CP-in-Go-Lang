package main

import "fmt"

const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{
		root: &Node{},
	}

	return result
}

func (t *Trie) InsertWord(w string) {
	wordLength := len(w)
	currentNode := t.root

	fmt.Println(wordLength, currentNode)
}

func main() {
	goTrie := InitTrie()
	fmt.Println(testTrie)
}
