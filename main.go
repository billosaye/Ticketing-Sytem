package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conference"
	const confrenceTickets = 34
	var remainingTickets = 16

	fmt.Printf("Welcome to our %v Booking Application \n", conferenceName)
	fmt.Println("We have a total of", confrenceTickets, "and", remainingTickets, "are still available")

	var firstName string
	var lastName string
	var userEmail string
	var userTickets int

	fmt.Println("Enter your name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&userEmail)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	fmt.Printf("Dear %s %s, thank you for booking %v tickets. You will receive a confirmation email at %s", firstName, lastName, userTickets, userEmail)

}
