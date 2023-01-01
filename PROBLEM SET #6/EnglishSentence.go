package main

import (
	"fmt"
	"strings"
	"unicode"
)

func checkSentenceForSpace(sentence string) {
	//	if sentence == "," {
	//newSentence := strings.SplitAfterN(sentence, ", ", 2)
	//strings.SplitAfter(sentence, " ")
	fmt.Printf("Sentence is : %#v\n", strings.FieldsFunc(sentence, func(r rune) bool {
		return !unicode.IsLetter(r)
	}))
}

func main() {
	//	var sentence string
	//fmt.Scanf("%s", &sentence)

	sentence := "Q1 months are January,February,March"
	fmt.Println("Enter string : ", sentence)

	checkSentenceForSpace(sentence)
}
