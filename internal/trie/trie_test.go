package trie

import (
	"testing"
)

func TestAddAndSearchWord(t *testing.T) {
	trie := NewTrie()

	words := []string{"apple", "apply", "app", "banana"}
	for _, word := range words {
		trie.AddWord(word)
	}

	tests := []struct {
		word     string
		expected bool
	}{
		{"apple", true},
		{"apply", true},
		{"app", true},
		{"appl", false},
		{"banana", true},
		{"bananas", false},
		{"bat", false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.word, func(t *testing.T) {
			result := trie.Search(test.word)
			if result != test.expected {
				t.Errorf("Search(%q) = %v; want %v", test.word, result, test.expected)
			}
		})
	}
}
