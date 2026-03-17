package main

import (
	"fmt"
	"strconv"
)

func ConvertBase(h string, base int) (string, error) {
	n, err := strconv.ParseInt(h, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(n, 10), nil
}
func main() {
	fmt.Println(ConvertBase("1E", 16))  // 30
	fmt.Println(ConvertBase("10", 8))   // 8
	fmt.Println(ConvertBase("1010", 2)) // 10
	fmt.Println(ConvertBase("ff", 16))  // 255
	fmt.Println(ConvertBase("77", 10))  // 77
}
