package src

import (
	"os"
	"strings"
)

// The FileOpen function in fileopen.go is responsible for reading the content of a file and returning it as a string.
func FileOpen(filename string) string {
	f, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	// Normalize line endings to Unix-style
	content := strings.ReplaceAll(string(f), "\r\n", "\n")
	return content
}
