package helper

import (
	"fmt"
	"strings"
)

func ValidateUserInputs(firstName string, lastName string, email string) bool {
	if !strings.Contains(email,"@") {
		fmt.Println("Please enter the correct email")	
		return false		
	} else if len(firstName) <= 2 || len(lastName) <= 2 {
		fmt.Println("Please enter the correct name")	
		return false		
	}
	return true
}
