package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sojebsikder/phonetic/internal/trie"
	"github.com/sojebsikder/phonetic/rulebasedconverter"

	_ "embed"
)

//go:embed data/bn.txt
var bnData []byte

var version = "0.0.1"
var appName = "phonetic"

func showUsage() {
	fmt.Printf("Usage:\n")
	fmt.Printf("  %s convert <text>\n\n", appName)
	fmt.Printf("  %s help\n", appName)
	fmt.Printf("  %s version\n", appName)
}

func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "convert":
		if len(os.Args) < 3 {
			fmt.Println("Error: No text provided for conversion.")
			os.Exit(1)
		}
		text := os.Args[2]

		rulebasedConv := rulebasedconverter.NewConverter()
		result := rulebasedConv.Convert(text)
		fmt.Println(result)
	case "suggest":
		if len(os.Args) < 3 {
			fmt.Println("Error: No text provided for suggestion.")
			os.Exit(1)
		}
		word := os.Args[2]

		t, err := trie.LoadFromBytes(bnData)
		if err != nil {
			log.Fatalf("Failed to load dictionary: %v", err)
		}

		if t.Search(word) {
			fmt.Printf("%q is spelled correctly.\n", word)
		} else {
			fmt.Printf("%q not found. Suggestions: %v\n", word, t.Suggestions(word))
		}

		fmt.Println(t.Suggestions(word))

	case "help":
		showUsage()
	case "version":
		fmt.Println("phonetic version " + version)
	default:
		fmt.Println("Unknown command:", cmd)
		fmt.Println("Use 'phonetic help' to see available commands.")
		os.Exit(1)
	}
}
