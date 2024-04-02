package main

import (
	"flag"
)

func main() {
	var sentence string
	flag.StringVar(&sentence, "s", "", "Input a sentence (use \"text\")")
	flag.Parse()

}
