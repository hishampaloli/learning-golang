package main

import (
	"fmt"
	"strings"
)

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


func validateUserInputs(firstName string, lastName string, email string) bool {
	if !strings.Contains(email,"@") {
		fmt.Println("Please enter the correct email")	
		return false		
	} else if len(firstName) <= 2 || len(lastName) <= 2 {
		fmt.Println("Please enter the correct name")	
		return false		
	}
	return true
}

func validateUserTickets(userTickets int)bool{
	if uint(userTickets) > remainingTickets{
		fmt.Printf("Sorry, only %v tickets left\n", remainingTickets)
		return false
	}
	return true
}


func printBooking(){
	firstNames := []string{}
	for _, booking:= range bookings{
		var firstName = strings.Fields(booking)[0]
		firstNames = append(firstNames, firstName)	
	}
	fmt.Printf("This are the attentees coming for the conferece := %v \n",firstNames)
}


func main(){
	

	greetUser()

	var userName string
	var lastName string
	var email string
	var userTickets int

	for {
		
		getUserInput(&userName, &lastName, &email, &userTickets)

		if !validateUserInputs(userName,lastName,email) || !validateUserTickets(userTickets) {
			continue
		}

		remainingTickets = remainingTickets - uint(userTickets)
	
		printBooking()

		if int(remainingTickets) <= 0 {
			fmt.Println("Woof, you were the last to book, All slots are complted")
			break
		}
	}

	
	// main().
}

