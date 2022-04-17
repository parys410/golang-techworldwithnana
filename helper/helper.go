package helper

import (
	"fmt"
	"strings"
)

func ValidateDataInput(remainingTicket uint8) (string, string, uint8) {
	var firstName string
	var lastName string
	var userName string
	var email string
	var userTicket uint8

	for {
		fmt.Printf("\nInput your First Name: ")
		fmt.Scan(&firstName)
		
		if len(firstName) > 2 {
			break
		} else {
			fmt.Printf("Please enter a valid first name. Re-enter your first name: \n")
			continue
		}
	}

	for {
		fmt.Printf("Input your Last Name: ")
		fmt.Scan(&lastName)
		
		if len(lastName) > 2 {
			break
		} else {
			fmt.Printf("Please enter a valid last name. Re-enter your last name: \n")
			continue
		}
	}

	for {
		fmt.Printf("Input your Email: ")
		fmt.Scan(&email)
		
		if strings.Contains(email, "@") {
			break
		} else {
			fmt.Printf("Please enter a valid email. Re-enter your email: \n")
			continue
		}
	}

	for {
		fmt.Printf("Enter the number of tickets: ")
		fmt.Scan(&userTicket)

		if userTicket <= remainingTicket {
			break
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTicket, userTicket)
			continue
		}
	}

	userName = firstName + " " + lastName

	return userName, email, userTicket
}