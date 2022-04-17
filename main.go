package main

import (
	"BookingTicketV1/helper"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type UserData struct {
	userName string
	email string
	numberOfTickets uint8
}

var wg = sync.WaitGroup{}

const conferenceTicket uint8 = 50
var conferenceName string = "Go Conference"
var remainingTicket uint8 = conferenceTicket
var bookings = make([]UserData, 0)

func main() {
	greetUsers()
	userName, userEmail, userTicket := helper.ValidateDataInput(remainingTicket)
	bookTicket(userName, userEmail, userTicket)

	wg.Add(1)
	go sendTicket(userTicket, userName, userEmail)

	noTicketsRemaining := remainingTicket <= 0
	if noTicketsRemaining {
		// End Program
		fmt.Printf("\n\nOur conference is booked out. Come back next year.\n\n")
		// break
	}
	
	// fmt.Printf("\n\nDo you want to continue? (y/n) ")
	// var input string
	// fmt.Scan(&input)

	// if(string(input) == "y"){
	// 	continue
	// } else {
	// 	break
	// }

	wg.Wait()
}

func greetUsers() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available. \n", conferenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend")
}

func getBookingsNames() []string {
	hideNames := []string{}

	for _, booking := range bookings {
		var tempName string

		// Map using booking["userName"] is just like array in JS
		// Struct using booking.userName is just like object in JS

		var splittedNames = strings.Fields(booking.userName)
		firstName := splittedNames[0]
		lastName := splittedNames[1]
		
		for index, name := range firstName {
			if (index > 2) {
				tempName = tempName + "*"
			} else {
				tempName = tempName + string(name)
			}
		}
		for index, name := range lastName {
			if (index > 2) {
				tempName = tempName + "*"
			} else {
				tempName = tempName + string(name)
			}
		}

		hideNames = append(hideNames, tempName)
	}

	return hideNames
}

func bookTicket(userName string, userEmail string, userTicket uint8) {
	remainingTicket = remainingTicket - userTicket

	// var myslice []string
	// var mymap map[string]string

	// Create a map for a user
	// var userData = make(map[string]string)
	// userData["userName"] = userName
	// userData["email"] = userEmail
	// userData["tickets"] = strconv.FormatUint(uint64(userTicket), 10)
	
	var userData = UserData {
		userName: userName,
		email: userEmail,
		numberOfTickets: userTicket,
	}
	
	bookings = append(bookings, userData)

	fmt.Printf("\n\nThank you %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, userTicket, userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTicket, conferenceName)
	fmt.Printf("These are all our bookings: %v\n", getBookingsNames())
	fmt.Printf("These are all our data: %v\n", bookings)
}

func sendTicket(userTicket uint8, userName string, userEmail string) {
	time.Sleep(15 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v", userTicket, userName)
	fmt.Printf("\n\n#############################################\n")
	fmt.Printf("Sending ticket:\n%v\nto email address %v\n", ticket, userEmail)
	fmt.Printf("#############################################\n\n")
	wg.Done()
}