package main

import "strings"

func BuildArt(input string, banner map[rune][]string) string {
	if input == "" {
		return input
	}

	split := Split(input)

	var result strings.Builder

	for i, words := range split {
		if words == "" {
			if i < len(split) {
				continue
			}
			result.WriteString("\n")
		}
		rend := Render(words, banner)
		for _, word := range rend {
			result.WriteString(word)
			result.WriteString("\n")
		}
	}
	return result.String()
}
