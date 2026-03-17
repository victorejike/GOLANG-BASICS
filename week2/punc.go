package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func fixPunctuation(s string) string {
	// Remove spaces before punctuation
	re := regexp.MustCompile(`\s*([.,!?:;])`)
	s = re.ReplaceAllString(s, "$1")
	// Ensure exactly one space after punctuation, unless at end of line
	re2 := regexp.MustCompile(`([.,!?:;])([^\s])`)
	s = re2.ReplaceAllString(s, "$1 $2")
	return s
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	reLow := regexp.MustCompile(`(\b\w+\b)\s*\(low\)`)
	for scanner.Scan() {
		line := scanner.Text()
		// Handle (low, n), (up, n), (cap, n)
		reMulti := regexp.MustCompile(`\b((?:\w+\s+){0,10}\w+)\s*\((low|up|cap),\s*(\d+)\)`)
		line = reMulti.ReplaceAllStringFunc(line, func(match string) string {
			submatches := reMulti.FindStringSubmatch(match)
			if len(submatches) == 4 {
				words := strings.Fields(submatches[1])
				n := 0
				fmt.Sscanf(submatches[3], "%d", &n)
				if n > len(words) {
					n = len(words)
				}
				switch submatches[2] {
				case "low":
					for i := len(words) - n; i < len(words); i++ {
						words[i] = strings.ToLower(words[i])
					}
				case "up":
					for i := len(words) - n; i < len(words); i++ {
						words[i] = strings.ToUpper(words[i])
					}
				case "cap":
					for i := len(words) - n; i < len(words); i++ {
						if len(words[i]) > 0 {
							words[i] = strings.ToUpper(words[i][:1]) + strings.ToLower(words[i][1:])
						}
					}
				}
				return strings.Join(words, " ")
			}
			return match
		})
		// Replace every "<word> (low)" with "<word in lowercase>"
		// Replace every "<word> (low)" with "<word in lowercase>"
		line = reLow.ReplaceAllStringFunc(line, func(match string) string {
			submatches := reLow.FindStringSubmatch(match)
			if len(submatches) > 1 {
				return strings.ToLower(submatches[1])
			}
			return match
		})
		fmt.Println(fixPunctuation(line))
	}
}