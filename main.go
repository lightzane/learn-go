package main

// Note:
// booking-app
// is the custom module_name that we created from the command:
// go mod init <module_name>
// see file: go.mod

import (
	"booking-app/questioner"          // "<module_name>/<package_name>" --> notice: the filename is the same with the package_name
	IDUtil "booking-app/shared/utils" // package_name "<module_name>/path/to/other/folder" --> We need to specify package_name since the package_name is not the same as the filename
	"fmt"
	"strings"
	"sync"
	"time"
)

// Global variables
var bookingName = "Air Flight Booking"

const bookingTickets uint = 50

// var bookings = []string{}
var bookings = make([]UserData, 0) // initialize with 0 size

// var remainingTickets = 50
//
// is the same as
//
// remainingTickets := 50
//
// But this is only for "var" variables, and not for constants
// This is also should only be declared within functions
var remainingTickets = bookingTickets

// Create a custom type that is called "UserData"
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	fmt.Println(IDUtil.ObjectID())

	for {

		greetUsers()

		firstName, lastName, email, userTickets := questioner.GetUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			// takes time to execute and will block the next line of codes
			// sendTicket(userTickets, firstName, lastName, email)

			wg.Add(1) // Adds a number of threads to wait for so that the program will not immediately exit
			// change to "Concurrency" to be executed on a seperate thread
			go sendTicket(userTickets, firstName, lastName, email) // will not block the next line of codes

			firstNames := getFirstNames()

			fmt.Println(firstNames)

			if remainingTickets == 0 {
				fmt.Println("\nAll tickets are sold out!")
				break // out of the loop
			}

		} else {

			fmt.Println()

			if !isValidName {
				fmt.Println("First name or last name is too short")
			}

			if !isValidEmail {
				fmt.Println("Email is invalid")
			}

			if !isValidTicketNumber {
				fmt.Println("Number of ticket you entered is invalid")
			}

		}
	}

	// Before main thread exits:
	wg.Wait()

}

func greetUsers() {
	// %v is a placeholder format that returns the `value` of the specified variable
	// see other available formats: https://pkg.go.dev/fmt
	fmt.Printf("\nWelcome to %v application\n", bookingName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", bookingTickets, remainingTickets)
	fmt.Println("Get your tickets here")
}

func getFirstNames() []string {

	firstNames := []string{}

	// Loop through each items in `bookings`
	// when `index` is not needed, use `_`
	// `range` allows us to iterate over a list
	for _, booking := range bookings {
		names := strings.Fields(booking.firstName) // Fields will split the string into an array seperated by spaces
		firstNames = append(firstNames, names[0])
	}

	return firstNames

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData) // `append()` pushes items to the array

	fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, bookingName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	// Simulate a delay
	time.Sleep(15 * time.Second)

	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("\n##########################")
	fmt.Printf("Sending ticket:\n %v \nto email address: %v\n", ticket, email)
	fmt.Println("##########################")

	wg.Done()

}
