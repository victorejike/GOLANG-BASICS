package main

import "fmt"

func add(x int, y int) int {
	return x + y

}

func main() {
	var x = 10
	var y = 20
	// what i have learn is that if  you want to assign a bool you dont have to chose for it

	if x < y {
		//this will tell  the program what to output
		fmt.Println("x is lesser")
		//this check for if the first dont work it is going to take the second
	} else {
		fmt.Println("y is greater")
	}
	sumre := add(x, y)
	fmt.Println("this the totall after add", sumre)
}
