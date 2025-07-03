package rulebasedconverter

import (
	"testing"
)

func TestTransliterate(t *testing.T) {
	rulebasedConv := NewConverter()

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
			result := rulebasedConv.Transliterate(test.word)
			if result != test.expected {
				t.Errorf("Transliterate(%q) = %q; want %q", test.word, result, test.expected)
			}
		})
	}
}
