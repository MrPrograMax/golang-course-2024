package main

import (
	"flag"
	"fmt"
	"github.com/kljensen/snowball"
	"strings"
	"unicode"
)

const English = "english"

func main() {
	var sentence string
	flag.StringVar(&sentence, "s", "", "Input a sentence (use \"text\")")
	flag.Parse()

	if sentence == "" {
		fmt.Println("Use -s flag for correct work of app")
		return
	}

	sep := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '\''
	}

	words := strings.FieldsFunc(sentence, sep)
	var answer []string
	repetitiveWords := make(map[string]interface{})

	for _, word := range words {
		word = strings.ToLower(word)

		if skipFlag := isContainsLiteral(word); skipFlag {
			continue
		}

		word, err := snowball.Stem(word, English, true)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		if exist := isStopWord(word); exist {
			continue
		}

		_, exist := repetitiveWords[word]
		if !exist {
			repetitiveWords[word] = nil
		} else {
			continue
		}

		answer = append(answer, word)
	}

	fmt.Println(strings.Join(answer, " "))
}
