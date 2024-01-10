package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"apple", "tom"}

	fruitList = append(fruitList, "Mango", "Bandana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 343
	highScores[1] = 34443
	highScores[2] = 3433
	highScores[3] = 34

	highScores = append(highScores, 555,34,533434)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println((sort.IntsAreSorted(highScores)))

	highScores = append(highScores[:2], highScores[2+1:]...)

}