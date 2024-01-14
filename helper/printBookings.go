package helper

import (
	"fmt"
	"strings"
)

func PrintBookings(bookings []string){
	firstNames := []string{}
	for _, booking:= range bookings{
		var firstName = strings.Fields(booking)[0]
		firstNames = append(firstNames, firstName)	
	}
	fmt.Printf("This are the attentees coming for the conferece := %v \n",firstNames)
}