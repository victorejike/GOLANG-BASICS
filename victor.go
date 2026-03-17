package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func formatPunctuation(text string) string {
	// Handle grouped punctuations first (e.g., ..., !?, !!, etc.)
	grouped := regexp.MustCompile(`([.]{3}|[!?]{2,})`)
	text = grouped.ReplaceAllStringFunc(text, func(m string) string {
		return strings.TrimSpace(m)
	})

	// Handle single punctuation marks
	re := regexp.MustCompile(`\s*([.,!?:;])\s*`)
	text = re.ReplaceAllString(text, "$1 ")

	// Remove extra space after grouped punctuations
	text = regexp.MustCompile(`([.]{3}|[!?]{2,})\s+`).ReplaceAllString(text, "$1 ")

	// Remove space before punctuation
	text = regexp.MustCompile(`\s+([.,!?:;])`).ReplaceAllString(text, "$1")

	// Remove space after punctuation at end of line
	text = regexp.MustCompile(`([.,!?:;])\s+$`).ReplaceAllString(text, "$1")

	// Trim leading/trailing spaces
	return strings.TrimSpace(text)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(formatPunctuation(line))
	}
}