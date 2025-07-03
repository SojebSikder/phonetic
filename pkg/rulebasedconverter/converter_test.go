package rulebasedconverter

import (
	"testing"
)

func TestTransliterate(t *testing.T) {
	trie := BuildTrieFromMaps()

	tests := []struct {
		word     string
		expected string
	}{
		{"ami bangla valObasi", "আমি বাংলা ভালোবাসি"},
		{"sojib sOjib", "সজিব সোজিব"},
		{"ami ekjon valO manuSh", "আমি একজন ভালো মানুষ"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.word, func(t *testing.T) {
			result := Transliterate(test.word, trie)
			if result != test.expected {
				t.Errorf("Transliterate(%q) = %q; want %q", test.word, result, test.expected)
			}
		})
	}
}
