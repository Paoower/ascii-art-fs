package main

import (
	src "ascii-art/src"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Checking if the correct number of arguments is provided
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Please put a valid argument.")
		return
	}
	// Get the input string and banner style from command line arguments
	input := os.Args[1]
	style := "standard"
	if len(os.Args) == 3 {
		style = os.Args[2]
	}

	// Use correct Banner files
	bannerFile, err := src.GetBannerFile(style)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Store line and get line of input
	lines := make([]string, 0)
	words := strings.Split(input, "\\n")

	for _, word := range words {
		if word == "" {
			lines = append(lines, "")
		} else {
			lines = append(lines, src.GetWord(word, bannerFile)...)
		}

	}

	for _, line := range lines {
		fmt.Println(line)
	}

}
