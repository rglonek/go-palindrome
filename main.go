package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	word := flag.String("word", "", "the word to check")
	flag.Parse()
	if word == nil || *word == "" {
		log.Fatal("The `--word' parameter must be defined.")
	}
	if find(*word) {
		fmt.Println("The word is a palindrome")
		return
	}
	fmt.Println("The word is NOT a palindrome")
}

// strips out any non-alpha-numeric letters, converts all upper-case to lower-case, and then
// check if the resulting word is a palindrome; returns true if yes; otherwise returns false
func find(word string) bool {
	// ignore all non-alpha-numeric chars
	// convert upper-case to lower-case
	// compare
	loffset := 0
	roffset := len(word) - 1
	var l1, l2 rune
	for {
		l1, loffset = findNextLetter(&word, loffset)
		if loffset == -1 {
			// finished, no more letters
			return true
		}
		l2, roffset = findPreviousLetter(&word, roffset)
		if roffset == -1 {
			// finished, no more letters
			return true
		}
		if loffset >= roffset {
			// we overlapped, so we have gone too far
			return true
		}
		if l1 != l2 {
			return false
		}
		loffset++
		roffset--
		if loffset >= roffset {
			// we overlapped, so we have gone too far
			return true
		}
	}
}

func findNextLetter(word *string, offset int) (letter rune, newOffset int) {
	for i := offset; i < len(*word); i++ {
		letter, newOffset := findNextPreviousDo(word, i)
		if newOffset == -1 {
			continue
		}
		return letter, newOffset
	}
	return rune(0), -1
}

func findPreviousLetter(word *string, offset int) (letter rune, newOffset int) {
	for i := offset; i < len(*word); i-- {
		letter, newOffset := findNextPreviousDo(word, i)
		if newOffset == -1 {
			continue
		}
		return letter, newOffset
	}
	return rune(0), -1
}

func findNextPreviousDo(word *string, offset int) (letter rune, newOffset int) {
	letter = rune((*word)[offset])
	if letter >= 'a' && letter <= 'z' {
		return letter, offset
	}
	if letter >= 'A' && letter <= 'Z' {
		return letter + 32, offset
	}
	if letter >= '0' && letter <= '9' {
		return letter, offset
	}
	return rune(0), -1
}

func findx(word string) bool {
	r := regexp.MustCompile(`[a-zA-Z0-9]+`)
	res := strings.ToLower(strings.Join(r.FindAllString(word, -1), ""))
	for i := 0; i <= len(res)/2; i++ {
		if res[i] != res[len(res)-i-1] {
			return false
		}
	}
	return true
}

func findy(word string) bool {
	res := []rune{}
	for _, l := range word {
		if l >= 'a' && l <= 'z' {
			res = append(res, l)
		}
		if l >= 'A' && l <= 'Z' {
			res = append(res, l+32)
		}
		if l >= '0' && l <= '9' {
			res = append(res, l)
		}
	}
	for i := 0; i <= len(res)/2; i++ {
		if res[i] != res[len(res)-i-1] {
			return false
		}
	}
	return true
}

func findz(word string) bool {
	res := make([]rune, 0, len(word))
	for _, l := range word {
		if l >= 'a' && l <= 'z' {
			res = append(res, l)
		}
		if l >= 'A' && l <= 'Z' {
			res = append(res, l+32)
		}
		if l >= '0' && l <= '9' {
			res = append(res, l)
		}
	}
	for i := 0; i <= len(res)/2; i++ {
		if res[i] != res[len(res)-i-1] {
			return false
		}
	}
	return true
}
