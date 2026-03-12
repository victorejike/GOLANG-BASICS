package main

import "fmt"

func main() {
	// i used the append function to add the elements of myslice2 to myslicee1 and store the result in myslice3
	myslicee1 := []int{10, 20, 30, 40}
	myslice2 := []int{50, 60, 70}
	// the append function takes two arguments: the first is the slice to which we want to append, and the second is the slice whose elements we want to append. The ... operator is used to unpack the elements of myslice2 so that they can be appended individually to myslicee1.
	myslice3 := append(myslicee1, myslice2...)
	//the output of my program will be the content of myslice3 added togther
	fmt.Printf("myslice3=%v\n", myslice3)
	//the length of myslice3 is that it add the myslice1 and myslice2 and check the length of the result
	fmt.Printf("length=%d\n", len(myslice3))
	fmt.Printf("capacity=%d\n", cap(myslice3))
}
