package main

import (
	"booking-app/helper"
	"fmt"
)

const (
    conferenceName      = "Go Conference"
    conferenceTickets   = 10
)

var (
    remainingTickets = uint(conferenceTickets)
    bookings         []string
)

func greetUser(){
	fmt.Printf("Welocme to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",conferenceTickets, remainingTickets)
}


func handleUpdateTickets(remainingTickets *uint, bookings *[]string, userTickets uint, userName string) {
	*remainingTickets -= userTickets
	*bookings = append(*bookings, userName)
	
	fmt.Printf("Thank you %v  for booking %v tickets, the details will be mailed to you \n", userName,userTickets)
	fmt.Printf("%v are still remaining for you frients to join\n", *remainingTickets)
}

func main(){
	
	greetUser()
	
	for {
		var (
			userName string
			lastName string
			email string
			userTickets int
		)
	
		helper.GetUserInput(&userName, &lastName, &email, &userTickets)

		if ! helper.ValidateUserInputs(userName,lastName,email) || !helper.ValidateUserTickets(userTickets, remainingTickets) {
			continue
		}		

		handleUpdateTickets(&remainingTickets,&bookings,uint(userTickets), userName + " " + lastName)
		helper.PrintBookings(bookings)
		if int(remainingTickets) <= 0 {
			fmt.Println("Woof, you were the last to book, All slots are complted")
			break
		}
	}
	
}

