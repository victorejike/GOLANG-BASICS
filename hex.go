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

func ConvertBase(h string, base int) (string, error) {
	n, err := strconv.ParseInt(h, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(n, 10), nil
}
