package main

import (
	"fmt"
)

const LoginToken string = "sdfhsfveirgeuirb"

func main() {
	var username string = "Hisham"
	fmt.Println(username)
	fmt.Printf("Varialbe is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Varialbe is of type: %T \n", isLoggedIn)

	var score int = 786
	fmt.Println(score)
	fmt.Printf("Varialbe is of type: %T \n", score)

	var smallFloat float32 = 786.40455004
	fmt.Println(smallFloat)
	fmt.Printf("Varialbe is of type: %T \n", smallFloat)

	var web = "Hisdhsdf"
	fmt.Println(web)

	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Varialbe is of type: %T \n", LoginToken)
}
