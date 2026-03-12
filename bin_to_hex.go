package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	text := "1e hex"

	words := strings.Split(text, " ")

	for i := 0; i < len(words); i++ {
		if words[i] == "hex" {
			code, _ := strconv.ParseInt(words[i-1], 16, 64)
			words[i-1] = strconv.Itoa(int(code))
			words = append(words[:i], words[i+1:]...)
			i--
		}
		if words[i] == "(bin)" {
			code, _ := strconv.ParseInt(words[i-1], 2, 64)
			words[i-1] = strconv.Itoa(int(code))
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	result := strings.Join(words, " ")
	fmt.Println(result)

}
