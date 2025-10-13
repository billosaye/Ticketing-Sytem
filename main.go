package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets int = conferenceTickets

	fmt.Printf("Welcome to our %v Booking Application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var userEmail string
	var userTickets int

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&userEmail)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("Tickets remaining: %v\n", remainingTickets)
}
