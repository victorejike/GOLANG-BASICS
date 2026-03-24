package main

import (
	"fmt"
)

func myFunction(x int, y string) (result int, txt string) {
	result = x + x
	txt = y + "world"
	return

}

func main() {
	a, b := myFunction(5, "hello")
	fmt.Println(a, b)
}
