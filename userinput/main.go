package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welocme user"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating:")

	input, err := reader.ReadString('\n')
	fmt.Println("Thanls gor rating, ",input)
	
}