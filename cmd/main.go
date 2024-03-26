package main

import (
	"flag"
	"fmt"
	"github.com/kljensen/snowball"
	"strings"
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

	words := strings.Split(sentence, " ")
	var answer []string
	repetitiveWords := make(map[string]interface{})

	for _, word := range words {
		word = strings.ToLower(word)
		word, err := snowball.Stem(word, English, true)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		exist := isContainsInStopList(word)
		if exist {
			continue
		}

		skipFlag := isWordContainsLiteral(word)
		if skipFlag {
			continue
		}

		_, exist = repetitiveWords[word]
		if !exist {
			repetitiveWords[word] = nil
		} else {
			continue
		}

		answer = append(answer, word)
	}

	fmt.Println(strings.Join(answer, " "))
}
