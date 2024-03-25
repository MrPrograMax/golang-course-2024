package main

import (
	"flag"
	"fmt"
	"github.com/kljensen/snowball"
	"strings"
)

func main() {
	badWords := map[string]interface{}{
		"of":  nil,
		"a":   nil,
		"the": nil,
		"as":  nil,
		"are": nil,
		"is":  nil,

		"i":          nil,
		"he":         nil,
		"she":        nil,
		"it":         nil,
		"you":        nil,
		"we":         nil,
		"they":       nil,
		"me":         nil,
		"him":        nil,
		"her":        nil,
		"us":         nil,
		"them":       nil,
		"my":         nil,
		"your":       nil,
		"our":        nil,
		"their":      nil,
		"mine":       nil,
		"his":        nil,
		"myself":     nil,
		"yourself":   nil,
		"himself":    nil,
		"herself":    nil,
		"itself":     nil,
		"ourselves":  nil,
		"yourselves": nil,
		"themselves": nil,

		"will": nil,
	}

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
		strings.ToLower(word)

		_, exist := badWords[word]
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
