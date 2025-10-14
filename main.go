package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

// Package-level (global) variables
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings []string

// Main function
func main() {
	greetUsers()

	for {
		// --- Get User Input ---
		firstName, lastName, userEmail, userTickets := getUserInput()
		// ----------------------

		// --- Call the Validation Function ---
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, userEmail, userTickets, remainingTickets)
		// ------------------------------------

		if isValidName && isValidEmail && isValidTicketNumber {
			// Successful booking
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n",
				firstName, lastName, userTickets, userEmail)
			fmt.Printf("%v tickets remaining for %v.\n\n", remainingTickets, conferenceName)

			firstNames := getFirstNames()
			fmt.Printf("Bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("\nOur conference is fully booked. Come back next year! 🎉")
				break
			}
		} else {
			// Detailed error messages
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

// Function to print the initial greeting (uses global variables)
func greetUsers() {
	fmt.Printf("Hello Users, Welcome to our %v Booking Application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are still available.\n\n", conferenceTickets, remainingTickets)
}

// Function to extract first names from the global bookings slice
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		if len(names) > 0 {
			firstNames = append(firstNames, names[0])
		}
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
