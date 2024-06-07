package main_test

import (
	src "ascii-art/src"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	correct := ` _    _          _   _          | |  | |        | | | |         | |__| |   ___  | | | |   ___   |  __  |  / _ \ | | | |  / _ \  | |  | | |  __/ | | | | | (_) | |_|  |_|  \___| |_| |_|  \___/                                                                  `

	input := "Hello"
	output := ""

	lines := make([]string, 0)
	words := strings.Split(input, "\\n")

	for _, word := range words {
		lines = append(lines, src.GetWord(word, "banners/standard.txt")...)
	}

	for _, line := range lines {
		output += line
	}

	if output != correct {
		t.Errorf("Program Output: %v \n", output)
		t.Errorf("Correct Output: %v \n", correct)
	}
}

func TestHELLO(t *testing.T) {
	correct := ` _    _   ______   _        _         ____   | |  | | |  ____| | |      | |       / __ \  | |__| | | |__    | |      | |      | |  | | |  __  | |  __|   | |      | |      | |  | | | |  | | | |____  | |____  | |____  | |__| | |_|  |_| |______| |______| |______|  \____/                                                                                            `

	input := "HELLO"
	output := ""

	lines := make([]string, 0)
	words := strings.Split(input, "\\n")

	for _, word := range words {
		lines = append(lines, src.GetWord(word, "banners/standard.txt")...)
	}

	for _, line := range lines {
		output += line
	}

	if output != correct {
		t.Errorf("Program Output: %v \n", output)
		t.Errorf("Correct Output: %v \n", correct)
	}
}

func TestHeLloHuMaN(t *testing.T) {
	correct := " _    _          _        _               _    _           __  __           _   _  | |  | |        | |      | |             | |  | |         |  \\/  |         | \\ | | | |__| |   ___  | |      | |   ___       | |__| |  _   _  | \\  / |   __ _  |  \\| | |  __  |  / _ \\ | |      | |  / _ \\      |  __  | | | | | | |\\/| |  / _` | | . ` | | |  | | |  __/ | |____  | | | (_) |     | |  | | | |_| | | |  | | | (_| | | |\\  | |_|  |_|  \\___| |______| |_|  \\___/      |_|  |_|  \\__,_| |_|  |_|  \\__,_| |_| \\_|                                                                                                                                                                       "

	input := "HeLlo HuMaN"
	output := ""

	lines := make([]string, 0)
	words := strings.Split(input, "\\n")

	for _, word := range words {
		lines = append(lines, src.GetWord(word, "banners/standard.txt")...)
	}

	for _, line := range lines {
		output += line
	}

	if output != correct {
		t.Errorf("Program Output: %v \n", output)
		t.Errorf("Correct Output: %v \n", correct)
	}
}

func Test1Hello2There(t *testing.T) {
	correct := "     _    _          _   _                       _______   _                            _  | |  | |        | | | |              ____   |__   __| | |                          / | | |__| |   ___  | | | |   ___       |___ \\     | |    | |__     ___   _ __    ___  | | |  __  |  / _ \\ | | | |  / _ \\        __) |    | |    |  _ \\   / _ \\ | '__|  / _ \\ | | | |  | | |  __/ | | | | | (_) |      / __/     | |    | | | | |  __/ | |    |  __/ |_| |_|  |_|  \\___| |_| |_|  \\___/      |_____|    |_|    |_| |_|  \\___| |_|     \\___|                                                                                                                                                                               "

	input := "1Hello 2There"
	output := ""

	lines := make([]string, 0)
	words := strings.Split(input, "\\n")

	for _, word := range words {
		lines = append(lines, src.GetWord(word, "banners/standard.txt")...)
	}

	for _, line := range lines {
		output += line
	}

	if output != correct {
		t.Errorf("Program Output: %v \n", output)
		t.Errorf("Correct Output: %v \n", correct)
	}
}
