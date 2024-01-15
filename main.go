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
	bookings = make([]UserData,0)
)

type UserData struct {
	userName string
	numberOfTickets uint
}



func greetUser(){
	fmt.Printf("Welocme to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n",conferenceTickets, remainingTickets)
}


func PrintBookings(bookings []UserData){
	firstNames := []string{}
	for _, booking:= range bookings{
		firstNames = append(firstNames, booking.userName)	
	}
	fmt.Printf("This are the attentees coming for the conferece := %v \n",firstNames)
}

func handleUpdateTickets(remainingTickets *uint, bookings *[]UserData, userTickets uint, userName string) {
	*remainingTickets -= userTickets
	var userData = UserData {
		userName: userName,
		numberOfTickets: userTickets,
	}
	*bookings = append(*bookings, userData)
	
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
		fmt.Println(bookings)
		PrintBookings(bookings)
		if int(remainingTickets) <= 0 {
			fmt.Println("Woof, you were the last to book, All slots are complted")
			break
		}
	}
	
}

