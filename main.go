package main

import "fmt"

const conferenceName string = "Go Conference"
const conferenceTickets int = 10
var remainingTickets uint = uint(conferenceTickets)
var bookings []string

func greetUser(){
	fmt.Printf("Welocme to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",conferenceTickets, remainingTickets)

}

func getUserInput(userName *string, lastName *string, email *string, userTickets *int){
	userInputMessages := []string {"Enter you first name: ", "Enter you last name: ", "Enter you email: ", "Please enter the number of tickets to purchase"}
	userInputs := []interface{}{userName, lastName, email, userTickets}

	for index, message := range userInputMessages {
		fmt.Println(message)
		fmt.Scan(userInputs[index])
	}
}


func main(){
	

	greetUser()

	var userName string
	var lastName string
	var email string
	var userTickets int

	for {
		
		getUserInput(&userName, &lastName, &email, &userTickets)

		
		remainingTickets = remainingTickets - uint(userTickets)
	
		fmt.Printf("Thank you %v %v for booking %v tickets for %v, the details will be mailed to %v \n", userName,lastName,userTickets,conferenceName,email)
	
		fmt.Printf("%v are still remaining for you frients to join", remainingTickets)


		if int(remainingTickets) <= 0 {
			fmt.Println("Woof, you were the last to book, All slots are complted")
			break
		}
	}

	
	// main().
}

