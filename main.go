package main

import "fmt"

func main(){
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welocme to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available",conferenceTickets, remainingTickets)

	var userName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter you first name: ")
	fmt.Scan(&userName)

	fmt.Println("Enter you last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter you email: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you need: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - uint(userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets for %v, the details will be mailed to %v \n", userName,lastName,userTickets,conferenceName,email)

	fmt.Printf("%v are still remaining for you frients to join", remainingTickets)
	// main().
}

