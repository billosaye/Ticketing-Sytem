package main

import (
	"booking-app/helper"
	"fmt"
)

// Package-level (global) variables
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0) // Slice of maps to store booking details

// Main function
func main() {
	greetUsers()

	for {
		firstName, lastName, userEmail, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber :=
			helper.ValidateUserInput(firstName, lastName, userEmail, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets -= userTickets

			var userData = make(map[string]string)
			userData["firstName"] = firstName
			userData["lastName"] = lastName
			userData["email"] = userEmail
			userData["numberOfTickets"] = fmt.Sprintf("%v", userTickets)

			bookings = append(bookings, userData)

			fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n",
				firstName, lastName, userTickets, userEmail)
			fmt.Printf("%v tickets remaining for %v.\n\n", remainingTickets, conferenceName)

			firstNames := getFirstNames()
			fmt.Printf("Bookings so far: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("\nOur conference is fully booked. Come back next year! 🎉")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name must be at least 2 characters long.")
			}
			if !isValidEmail {
				fmt.Println("Email address must contain the '@' symbol.")
			}
			if !isValidTicketNumber {
				fmt.Printf("Invalid ticket number. Must be greater than 0 and less than or equal to %v.\n", remainingTickets)
			}
			fmt.Println("Your input is invalid. Please try again.")
		}
	}
}

// Function to print the initial greeting
func greetUsers() {
	fmt.Printf("Hello Users, Welcome to our %v Booking Application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are still available.\n\n", conferenceTickets, remainingTickets)
}

// Function to extract first names from bookings
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

// Function to get input from the user
func getUserInput() (string, string, string, uint) {
	var firstName, lastName, userEmail string
	var userTickets uint

	fmt.Println("\nEnter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&userEmail)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets
}
