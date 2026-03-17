package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// processText applies (low), (up), (cap) transformations with optional counts.
func processText(input string) string {
	re := regexp.MustCompile(`\((low|up|cap)(?:,\s*(\d+))?\)`)
	words := strings.Fields(input)
	matches := re.FindAllStringSubmatchIndex(input, -1)

	// Track word positions for transformations
	wordPos := 0
	for _, match := range matches {
		cmd := input[match[2]:match[3]]
		count := 1
		if match[4] != -1 && match[5] != -1 {
			c, err := strconv.Atoi(input[match[4]:match[5]])
			if err == nil {
				count = c
			}
		}
		// Find the previous 'count' words
		start := wordPos - count
		if start < 0 {
			start = 0
		}
		for i := start; i < wordPos; i++ {
			switch cmd {
			case "low":
				words[i] = strings.ToLower(words[i])
			case "up":
				words[i] = strings.ToUpper(words[i])
			case "cap":
				words[i] = strings.Title(strings.ToLower(words[i]))
			}
		}
		wordPos++
	}
	// Remove transformation commands from input
	clean := re.ReplaceAllString(input, "")
	return strings.Join(strings.Fields(clean), " ")
}

func main() {
	fmt.Println(processText("Hello World (low)"))
	fmt.Println(processText("This is a Test (up, 2)"))
	fmt.Println(processText("capitalize this (cap, 3)"))
}
