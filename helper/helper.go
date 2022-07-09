package helper

import (
	"booking/types"
	"fmt"
	"strings"
)

func GreetUsers(conferenceName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func PrintFirstNames(bookings map[string]types.User) {
	firstNames := []string{}
	for _, user := range bookings {
		firstNames = append(firstNames, user.FirstName)
	}
	fmt.Printf("These are all of our bookings: %v\n", firstNames)
}

func ValidateUserInput(user types.User, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(user.FirstName) >= 2 && len(user.LastName) >= 2
	isValidEmail := strings.Contains(user.Email, "@")
	isValidUserTickets := user.NumberOfTickets > 0 && user.NumberOfTickets <= remainingTickets
	return isValidName, isValidEmail, isValidUserTickets
}

