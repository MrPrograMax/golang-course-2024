package main

import (
	"flag"
	"fmt"
	"github.com/kljensen/snowball"
	"strings"
)

func main() {
	sentence := flag.String("s", "", "Input a sentence (use \"*text*\")")
	flag.Parse()

	if *sentence == "" {
		fmt.Println("Use -s flag for correct work of app")
		return
	}

	words := strings.Split(*sentence, " ")
	var answer []string
	repetitiveWords := make(map[string]interface{})

	for i := 0; i < len(words); i++ {
		word, err := snowball.Stem(words[i], "english", true)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		exist := isContainsInStopList(word)
		if exist {
			continue
		}
		if strings.Contains(word, "'ll") {
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
