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
go run . convert "amar nam sojib"
```

## Installation

```bash
go get -u github.com/SojebSikder/phonetic
```

## Usage

```go
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
```

## Tests

```bash
go test ./...
# with benchmark
go test ./... -bench=.
```
