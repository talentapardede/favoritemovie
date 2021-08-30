package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findFirstStringInBracket("semua murid (semua guru)"))
}

func findFirstStringInBracket(str string) (res string) {
	if len(str) <= 0 {
		return res
	}

	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound < 0 {
		return res
	}

	// fmt.Println("indexFirstBracketFound: ", indexFirstBracketFound)

	runes := []rune(str)
	// fmt.Println("runes: ", runes)

	wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
	// fmt.Println("wordsAfterFirstBracket: ", wordsAfterFirstBracket)

	indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
	if indexClosingBracketFound < 0 {
		return res
	}
	// fmt.Println("indexClosingBracketFound: ", indexClosingBracketFound)

	newRunes := []rune(wordsAfterFirstBracket)
	// fmt.Println("newRunes: ", newRunes)

	return string(newRunes[1 : indexClosingBracketFound-1])
}
