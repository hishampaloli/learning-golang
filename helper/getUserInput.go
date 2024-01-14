package helper

import "fmt"

func GetUserInput(userName *string, lastName *string, email *string, userTickets *int) {
	userInputMessages := []string{"Enter you first name: ", "Enter you last name: ", "Enter you email: ", "Please enter the number of tickets to purchase"}
	userInputs := []interface{}{userName, lastName, email, userTickets}

	for index, message := range userInputMessages {
		fmt.Println(message)
		fmt.Scan(userInputs[index])
	}
}