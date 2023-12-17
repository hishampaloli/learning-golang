package main

import "fmt"

func main() {
	fmt.Println("Welcom to pointer")

	// var ptr *int
	// fmt.Println(ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr + 2
	fmt.Println(myNumber)

}
