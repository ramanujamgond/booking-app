package main

// import (
// 	"booking-app/helper"
// 	"fmt"
// 	"strings"
// )

import (
	"fmt"
	"strings"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// var bookings = []string{}
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	// var conferenceName string = "Go Conference"
	// or
	// conferenceName := "Go Conference"
	// const conferenceTickets int = 50
	// var remainingTickets uint = 50
	// var bookings = [50]string{}
	// or
	// var bookings [50]string
	// var bookings = []string // slice
	// var bookings = []string{} // alternative way for creating slice
	// or
	// bookings := []string{}

	// fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	greetUsers()

	// fmt.Printf("\nWelcome to %v booking application\n", conferenceName)
	// fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	// fmt.Println("Get your tickets here to attend")

	// for remainingTickets > 0 && len(bookings) < 50 {
	// or
	for remainingTickets > 0 && len(bookings) < 50 {
		// var firstName string
		// var lastName string
		// var email string
		// var userTickets uint
		// ask user for their name

		// fmt.Println(remainingTickets)  // prints the value of the varibale
		// fmt.Println(&remainingTickets) // print the memory location of the variable and here the pointer is used to print the memory

		// fmt.Print("\nEnter your first name: ")
		// fmt.Scan(&firstName)

		// fmt.Print("Enter your last name: ")
		// fmt.Scan(&lastName)

		// fmt.Print("Enter your email address: ")
		// fmt.Scan(&email)

		// fmt.Print("Enter number of tickets: ")
		// fmt.Scan(&userTickets)

		// isValidCity := city != "Singapore" && city != "London"
		// isValidCity := city == "Singapore" || city == "London"

		// if userTickets > remainingTickets {
		// 	fmt.Printf("We only have %v tickets remaining, so you can't book %v rickets", remainingTickets, userTickets)
		// 	continue
		// }

		firstName, lastName, email, userTickets := getUserInput()

		// isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			// bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("\n The whole array: %v", bookings)
			// fmt.Printf("\n The whole slice: %v", bookings)
			// fmt.Printf("\n The first value: %v", bookings[0])
			// fmt.Printf("\n Array Type: %T", bookings)
			// fmt.Printf("\n Slice Type: %T", bookings)
			// fmt.Printf("\n Array length: %v", len(bookings))
			// fmt.Printf("\n Slice length: %v", len(bookings))

			// fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v", firstName, lastName, userTickets, email)
			// fmt.Printf("\n%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// firstNames := []string{}

			// for index, booking := range bookings {
			// 	var names = strings.Fields(booking)
			// 	firstNames = append(firstNames, names[0])
			// }
			// or use blank identifier
			// to ignore a variable you don't want to use
			// so with go you need to make unused variables explicit
			// for _, booking := range bookings {
			// 	var names = strings.Fields(booking)
			// 	firstNames = append(firstNames, names[0])
			// }

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstName()

			// fmt.Printf("These are all our bookings %v\n", bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			fmt.Printf("\n")

			// var noTickets bool = remainingTickets == 0
			// or
			noTickets := remainingTickets == 0

			if noTickets {
				// end program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			// fmt.Printf("\nWe only have %v tickets remaining, so you can't book %v rickets", remainingTickets, userTickets)
			// fmt.Print("Your input data is invalid, try again!")
			if !isValidName {
				fmt.Println("\nFirst name or last name you entered is too short!")
			}

			if !isValidEmail {
				fmt.Println("\nEmail address you entered doesn't conat @ sign!")
			}

			if !isValidTicketNumber {
				fmt.Println("\nNumber of tickets you entered is invalid!")
			}
			// continue
			fmt.Printf("\n")
		}
	}
}

func greetUsers() {
	// fmt.Println("Welcome to our conference")
	fmt.Printf("\nWelcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// firstNames = append(firstNames, names[0])
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	// map is data type in golang for storing key value pair
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// bookings = append(bookings, userData)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v", firstName, lastName, userTickets, email)
	fmt.Printf("\n%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending tickets:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##################")
}
