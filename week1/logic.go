package main

import (
	"fmt"
	"strconv"
)

func add(x int, y string) string {
	return strconv.Itoa(x) + y
}

func main() {
	x := 29
	y := " hello, victor"

	fmt.Println(add(x, y))
}
