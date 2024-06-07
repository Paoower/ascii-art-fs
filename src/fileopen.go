package src

import "os"

// The FileOpen function in fileopen.go is responsible for reading the content of a file and returning it as a string.
func FileOpen(filename string) string {
	f, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(f)
}
