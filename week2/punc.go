package main

import (
	"fmt"
)

func punc(s string) bool {
	for _, v := range s {
		if v == '!' || v == ',' || v == '.' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(punc("!Nnayi Ejike, is a good mentor."))
	fmt.Println(punc("Nnayi Ejike is a good mentor"))
}

// Every instance of the punctuations ., ,, !, ?, : and 
// ; should be close to the previous word and with space apart from the
//  next one. (Ex: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there,
//   and then BAMM!!").
