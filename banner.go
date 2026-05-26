package main

import (
	"bufio"
	"os"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		txt := scanner.Text()
		result = append(result, txt)

	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	mapbanner := make(map[rune][]string)
	for i := 32; i < 127; i++ {
		char := rune(i)
		start := (i - 32) * 9
		if start+9 > len(result) {
			break
		}
		completeline := result[start+1 : start+9]
		mapbanner[char] = completeline
	}
	return mapbanner, nil
}
