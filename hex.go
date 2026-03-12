package main

import (
	"fmt"
	"strconv"
)

func HexToDecimal(hex string) int64 {
	decimal, _ := strconv.ParseInt(hex, 32, 64)
	return decimal
}

func main() {
	result := HexToDecimal("1e")
	fmt.Println(result)
	fmt.Printf("result %d\n", result)
}
