package main

import "fmt"

func main() {
	// var conferenceName string = "Go Conference"
	// or
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// var bookings = [50]string{}
	// or
	// var bookings [50]string
	// var bookings = []string // slice
	// var bookings = []string{} // alternative way for creating slice
	// or
	bookings := []string{}

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name

		// fmt.Println(remainingTickets)  // prints the value of the varibale
		// fmt.Println(&remainingTickets) // print the memory location of the variable and here the pointer is used to print the memory

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		// fmt.Printf("\n The whole array: %v", bookings)
		// fmt.Printf("\n The whole slice: %v", bookings)
		// fmt.Printf("\n The first value: %v", bookings[0])
		// fmt.Printf("\n Array Type: %T", bookings)
		// fmt.Printf("\n Slice Type: %T", bookings)
		// fmt.Printf("\n Array length: %v", len(bookings))
		// fmt.Printf("\n Slice length: %v", len(bookings))

		fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v", firstName, lastName, userTickets, email)
		fmt.Printf("\n%v tickets remaining for %v\n", remainingTickets, conferenceName)

		fmt.Printf("These are all our bookings %v\n", bookings)
		fmt.Printf("\n")
	}
}
