package src

import (
	"strings"
)

const LETTER_HEIGHT = 8

// The GetLetter function in getletter.go is responsible for retrieving the ASCII representation of a single character from a banner.
func GetLetter(content string, ascii int) string {
	if ascii == 32 {
		s := ""
		for i := 0; i < 8; i++ {
			if i != 7 {
				s += "    " + "\n"
				continue
			}
			s += "    "
		}
		return s
	}

	str := ""
	lines := strings.Split(content, "\n")

	place := ascii - 31
	times := (place - 1) * LETTER_HEIGHT
	beginning := (ascii - 30) + times

	for i := beginning; i < beginning+LETTER_HEIGHT; i++ {
		if i != (beginning+LETTER_HEIGHT)-1 {
			str += lines[i-1] + "\n"
		} else {
			str += lines[i-1]
		}
	}

	return str
}
