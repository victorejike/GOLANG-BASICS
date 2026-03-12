package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	result := add(30, 40)
	myslice := []int{20, 30, 40, 60}
	myslice1 := []int{190, 169, 189}
	myslice2 := append(myslice, myslice1...)
	num := add(myslice2[6], result)

	fmt.Println(num)
	fmt.Println(result)
	fmt.Printf("num %d\n", len(myslice2))

}
