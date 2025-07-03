package trie

import (
	"bufio"
	"os"
	"strings"
)

type Node struct {
	children     map[rune]*Node
	completeWord string
}

type Trie struct {
	root *Node
}

func NewNode() *Node {
	return &Node{
		children:     map[rune]*Node{},
		completeWord: "",
	}
}

func NewTrie() *Trie {
	return &Trie{root: NewNode()}
}

func (t *Trie) AddWord(word string) {
	node := t.root

	for _, c := range word {
		_, ok := node.children[c]
		if !ok {
			node.children[c] = NewNode()
		}
		node = node.children[c]
	}
	node.completeWord = word
}

func (t *Trie) Search(word string) bool {
	node := t.root

	for _, c := range word {
		_, ok := node.children[c]
		if !ok {
			return false
		}
		node = node.children[c]
	}
	return node.completeWord != ""
}

func (t *Trie) Suggestions(word string) []string {
	results := []string{}
	var dfs func(Node *Node, path string)
	dfs = func(node *Node, path string) {
		if node.completeWord != "" {
			results = append(results, path)
		}
		for ch, child := range node.children {
			dfs(child, path+string(ch))
		}
	}

	node := t.root
	for _, ch := range strings.ToLower(word) {
		if _, exists := node.children[ch]; !exists {
			return results
		}
		node = node.children[ch]
	}
	dfs(node, strings.ToLower(word))
	return results
}

// MatchLongestPrefix matches the longest phonetic prefix from input[i:].
func (t *Trie) MatchLongestPrefix(input []rune, start int) (match string, length int) {
	node := t.root
	var result string

	for i := start; i < len(input); i++ {
		ch := input[i]
		if node.children[ch] == nil {
			break
		}
		node = node.children[ch]
		if node.completeWord != "" {
			result = node.completeWord
			length = i - start + 1
		}
	}

	return result, length
}

func LoadFromFile(path string) (*Trie, error) {
	trie := NewTrie()

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trie.AddWord(scanner.Text())
	}
	return trie, scanner.Err()

}
