package main

import (
	"booking/helper"
	"booking/types"
	"fmt"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = conferenceTickets
var bookings = make(map[string]types.User, 0)

func main() {

	helper.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		user := types.User{}
		fmt.Println("Enter your first name:")
		fmt.Scan(&user.FirstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&user.LastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&user.Email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&user.NumberOfTickets)

		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(
			user,
			remainingTickets,
		)

		if !isValidName || !isValidEmail || !isValidUserTickets {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email does not contain @ sign")
			}
			if !isValidUserTickets {
				fmt.Println("Number of tickets is invalid")
			}
			fmt.Println("Your user data is invalid. Please try again")
			continue
		}

		remainingTickets -= user.NumberOfTickets
		bookings[user.GetFullName()] = user

		fmt.Printf(
			"Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
			user.FirstName,
			user.LastName,
			user.NumberOfTickets,
			user.Email,
		)

		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		helper.PrintFirstNames(bookings)
	}
}
