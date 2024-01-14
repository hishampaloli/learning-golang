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

func main(){
	

	greetUser()
	
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

