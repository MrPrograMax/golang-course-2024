package main

import "strings"

var stopList = map[string]interface{}{
	"of":  nil,
	"a":   nil,
	"the": nil,
	"as":  nil,
	"are": nil,
	"is":  nil,
	"to":  nil,

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

func isContainsInStopList(word string) bool {
	_, exist := stopList[word]
	return exist
}

var literalList = []string{
	"'t",
	"'ve",
	"'re",
	"'re",
	"'ll",
}

func isWordContainsLiteral(word string) bool {
	for _, literal := range literalList {
		if strings.Contains(word, literal) {
			return true

		}
	}

	return false
}
