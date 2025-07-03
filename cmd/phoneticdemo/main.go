package main

import (
	"fmt"
	"log"
	"sojebsikder/phonetic/internal/trie"
	"sojebsikder/phonetic/pkg/rulebasedconverter"
)

func main() {
	// Suggestions
	t, err := trie.LoadFromFile("data/en.txt")
	if err != nil {
		log.Fatalf("Failed to load dictionary: %v", err)
	}

	words := []string{"sojeb", "sikder", "shikder", "appl", "ball"}
	for _, word := range words {
		if t.Search(word) {
			fmt.Printf("%q is spelled correctly.\n", word)
		} else {
			fmt.Printf("%q not found. Suggestions: %v\n", word, t.Suggestions(word))
		}
	}

	// Transliteration
	trie := rulebasedconverter.BuildTrieFromMaps()
	fmt.Println(rulebasedconverter.Transliterate("ami bangla valObasi", trie))   // আমি বাংলা ভালোবাসি
	fmt.Println(rulebasedconverter.Transliterate("sojib sOjib", trie))           // সজিব সোজিব
	fmt.Println(rulebasedconverter.Transliterate("ami ekjon valO manuSh", trie)) //আমি একজন ভালো মানুষ
}
