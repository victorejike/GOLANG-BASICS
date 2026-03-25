package main

import(
	"fmt"
	"unicode"
)

func AlphaCount(s string) int {
	victor := 0

	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsSpace(v) {
			victor++
		}
	}
	return victor
	
}

func main(){
	fmt.Println(AlphaCount("hello wolrd 4"))
}