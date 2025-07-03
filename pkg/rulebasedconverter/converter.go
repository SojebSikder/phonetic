package rulebasedconverter

import (
	"sojebsikder/phonetic/internal/trie"
	"strings"
)

// Maps for phonetic to Bangla
// Consonants map: phonetic consonant → Bangla consonant
var consonants = map[string]string{
	"k": "ক", "kh": "খ", "g": "গ", "gh": "ঘ", "Ng": "ঙ", "c": "চ", "ch": "ছ", "j": "জ", "jh": "ঝ", "NG": "ঞ",
	"T": "ট", "Th": "ঠ", "D": "ড", "Dh": "ঢ", "N": "ণ", "t": "ত", "th": "থ", "d": "দ", "dh": "ধ", "n": "ন",
	"p": "প", "ph": "ফ", "f": "ফ", "b": "ব", "bh": "ভ", "v": "ভ", "m": "ম", "z": "য", "r": "র", "l": "ল", "sh": "শ", "S": "শ", "Sh": "ষ",
	"s": "স", "h": "হ", "R": "ড়", "RH": "ঢ়", "y": "য়", "Y": "য়", "ng": "ং", ":": "ঃ", "^": "ঁ", "J": "জ",
}

// Vowels (standalone) map: phonetic vowel → Bangla vowel letter
var vowels = map[string]string{
	"a": "আ", "aa": "আ", "i": "ই", "ii": "ঈ", "u": "উ", "uu": "ঊ",
	"e": "এ", "oi": "ঐ", "o": "অ", "ou": "ঔ",
	"O": "ও", "Ou": "ঔ",
}

// Matras map: phonetic vowel → vowel sign (matra)
var matras = map[string]string{
	"a": "া", "i": "ি", "ii": "ী", "u": "ু", "uu": "ূ",
	"e": "ে", "oi": "ৈ", "o": "", "O": "ো", "ou": "ৌ", "Ou": "ৌ",
}

func isVowel(s string) bool {
	_, ok := vowels[s]
	return ok
}

func isConsonant(s string) bool {
	_, ok := consonants[s]
	return ok
}

func BuildTrieFromMaps() *trie.Trie {
	t := trie.NewTrie()
	for key := range consonants {
		t.AddWord(key)
	}
	for key := range vowels {
		t.AddWord(key)
	}
	return t
}

// Transliterate converts a phonetic input string to Bangla script using matra logic.
func Transliterate(input string, t *trie.Trie) string {
	runes := []rune(input)
	var output strings.Builder

	lastWasConsonant := false

	for i := 0; i < len(runes); {
		match, length := t.MatchLongestPrefix(runes, i)
		if match == "" {
			// fallback: output original char
			output.WriteRune(runes[i])
			lastWasConsonant = false
			i++
			continue
		}

		if isConsonant(match) {
			// Add consonant
			output.WriteString(consonants[match])
			lastWasConsonant = true
			i += length
		} else if isVowel(match) {
			if lastWasConsonant {
				// Attach matra to previous consonant
				output.WriteString(matras[match])
			} else {
				// standalone vowel letter
				output.WriteString(vowels[match])
			}
			lastWasConsonant = false
			i += length
		} else {
			// unknown, output raw
			output.WriteString(match)
			lastWasConsonant = false
			i += length
		}
	}

	return output.String()
}
