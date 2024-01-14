package helper

import "fmt"

func ValidateUserTickets(userTickets int, remainingTickets uint) bool {
	if uint(userTickets) > remainingTickets {
		fmt.Printf("Sorry, only %v tickets left\n", remainingTickets)
		return false
	}
	return true
}
