package main

import (
	"fmt"
	"strconv"
	"strings"
)

// processText applies (low), (up), (cap) transformations with optional counts.
func processText(input string) string {
	// Split the input into parts
	parts := strings.Fields(input)
	var result []string

	for _, part := range parts {
		if strings.HasPrefix(part, "(low") {
			count := extractCount(part, "(low")
			result = append(result, strings.ToLower(input)[0:count])
		} else if strings.HasPrefix(part, "(up") {
			count := extractCount(part, "(up")
			result = append(result, strings.ToUpper(input)[0:count])
		} else if strings.HasPrefix(part, "(cap") {
			count := extractCount(part, "(cap")
			result = append(result, strings.Title(strings.ToLower(input))[0:count])
		} else {
			result = append(result, part)
		}
	}

	return strings.Join(result, " ")
}

// extractCount extracts the count from the transformation command.
func extractCount(part string, prefix string) int {
	countStr := strings.TrimSuffix(strings.TrimPrefix(part, prefix), ")")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return len(part) // Default to the length of the part if count is invalid
	}
	return count
}

func main() {
	input := "(low5) (up3) (cap4) Hello World"
	output := processText(input)
	fmt.Println(output)
}
