package words

import (
	"fmt"
	"github.com/kljensen/snowball"
	"strings"
	"unicode"
)

func isStopWord(word string) bool {
	switch word {
	case "a", "about", "above", "after", "again", "against", "all", "am", "an",
		"and", "any", "are", "as", "at", "be", "because", "been", "before",
		"being", "below", "between", "both", "but", "by", "can", "did", "do",
		"does", "doing", "don", "down", "during", "each", "few", "for", "from",
		"further", "had", "has", "have", "having", "he", "her", "here", "hers",
		"herself", "him", "himself", "his", "how", "i", "if", "in", "into", "is",
		"it", "its", "itself", "just", "me", "more", "most", "my", "myself",
		"no", "nor", "not", "now", "of", "off", "on", "once", "only", "or",
		"other", "our", "ours", "ourselves", "out", "over", "own", "s", "same",
		"she", "should", "so", "some", "such", "t", "than", "that", "the", "their",
		"theirs", "them", "themselves", "then", "there", "these", "they",
		"this", "those", "through", "to", "too", "under", "until", "up",
		"very", "was", "we", "were", "what", "when", "where", "which", "while",
		"who", "whom", "why", "will", "with", "you", "your", "yours", "yourself",
		"yourselves":
		return true
	}
	return false
}

func isContainLiteral(word string) bool {
	return strings.Contains(word, "'")
}

const English = "english"

func Normalize(sentence string) ([]string, error) {
	words := strings.FieldsFunc(sentence, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && string(c) != "'"
	})

	var answer []string
	repetitiveWords := make(map[string]interface{})

	for _, word := range words {
		word = strings.ToLower(word)

		if skipFlag := isContainLiteral(word); skipFlag {
			continue
		}

		word, err := snowball.Stem(word, English, true)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return nil, err
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

	return answer, nil
}
