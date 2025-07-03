# Rule-Based Phonetic to Bangla Converter

This Go package provides a rule-based transliteration engine that converts phonetic Latin-based text into Bangla script using predefined mappings and trie-based prefix matching for efficient parsing.

## Features

- Converts Latin phonetic input into Bangla characters
- Uses a trie data structure for efficient longest-prefix matching
- Supports both standalone vowels and matras (vowel signs)
- Handles consonants, vowels, and complex clusters

## Demo

For testing, run the following command:

```bash
go run ./cmd/phoneticdemo
```

## Usage

```go
package main

import (
    "fmt"
    "log"
    "sojebsikder/phonetic/internal/trie"
    "sojebsikder/phonetic/rulebasedconverter"
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
    fmt.Println(rulebasedConv.Transliterate("ami bangla valObasi"))   // আমি বাংলা ভালোবাসি
    fmt.Println(rulebasedConv.Transliterate("sojib sOjib"))           // সজিব সোজিব
    fmt.Println(rulebasedConv.Transliterate("ami ekjon valO manuSh")) //আমি একজন ভালো মানুষ
}
```
