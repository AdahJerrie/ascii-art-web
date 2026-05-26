package main

import "strings"

func Render(input string, banner map[rune][]string) []string {
	result := []string{}

	for row := 1; row <= 8; row++ {
		var rended strings.Builder
		for _, char := range input {
			render := banner[char][row]
			rended.WriteString(render)
		}
		result = append(result, rended.String())
	}
	return result
}
