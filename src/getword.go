package src

import (
	"strings"
	"fmt"
)

func GetWord(input string, bannerFile string) []string {
	lines := make([]string, 8)
	content := FileOpen(bannerFile)
	for _, char := range input {
		if !(char >= 32 && char <= 126) {
			fmt.Println("character not allowed")
			return []string{}
		}
		c := strings.Split(GetLetter(content, int(char)), "\n")
		for i := 0; i < len(lines); i++ {
			lines[i] += c[i]
		}
	}
	return lines
}



