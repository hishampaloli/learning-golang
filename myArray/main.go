package main

import "fmt"

func main() {
	fmt.Println("welocom")

	var fruitList [4]string

	fruitList[0] = "hisham"
	fruitList[1] = "paloli"
	fruitList[3] = "dsdsd"

	fmt.Println("Fruit list is :", fruitList)
	fmt.Print(len(fruitList))

	var vegList = [6]string{"pota", "beasn", "masdg"}
	fmt.Println("Vegy", vegList)
}