package main

import (
	"fmt"
	"log"

	"github.com/sojebsikder/phonetic/internal/trie"
	"github.com/sojebsikder/phonetic/rulebasedconverter"
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
	rulebasedConv := rulebasedconverter.NewConverter()
	fmt.Println(rulebasedConv.Convert("ami bangla valObasi"))   // আমি বাংলা ভালোবাসি
	fmt.Println(rulebasedConv.Convert("sojib sOjib"))           // সজিব সোজিব
	fmt.Println(rulebasedConv.Convert("ami ekjon valO manuSh")) // আমি একজন ভালো মানুষ
}
