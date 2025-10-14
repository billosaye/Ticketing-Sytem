package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	remainingTickets := conferenceTickets

	fmt.Printf("Welcome to our %v Booking Application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are still available.\n\n", conferenceTickets, remainingTickets)

	var bookings []string

	for {
		var firstName, lastName, userEmail string
		var userTickets int

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&userEmail)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		// 1. Corrected strings.Contains usage and logic
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(userEmail, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// --- Input Validation ---
		if !isValidName {
			fmt.Println("First name or last name must be at least 2 characters long.")
		}
		if !isValidEmail {
			fmt.Println("Email address must contain the '@' symbol.")
		}
		if !isValidTicketNumber {
			fmt.Printf("Invalid ticket number. Must be greater than 0 and less than or equal to the remaining %v tickets.\n", remainingTickets)
		}

		isValidInput := isValidName && isValidEmail && isValidTicketNumber

		if !isValidInput {
			fmt.Println("Your input is invalid. Please try again.")
			continue // Start the loop over
		}
		// --- End Input Validation ---

		// The ticket availability check is now partially handled by isValidTicketNumber,
		// but we keep the dedicated check as good practice for user feedback clarity.
		if userTickets > remainingTickets {
			fmt.Printf("Sorry, we only have %v tickets remaining. You cannot book %v tickets.\n\n", remainingTickets, userTickets)
			continue // ask again
		}

		remainingTickets -= userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, userEmail)
		fmt.Printf("%v tickets remaining for %v.\n\n", remainingTickets, conferenceName)

		// Print first names only to be cleaner for demonstration
		firstNames := []string{}
		for _, booking := range bookings {
			// Get the first word (first name) of the full name string
			names := strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("Bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("\nOur conference is fully booked. Come back next year! 🎉")
			break
		}
	}
}
