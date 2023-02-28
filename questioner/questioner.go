package questioner

import "fmt"

// The function name should be in `PascalCase` for it to be exported and be used as public functions
func GetUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("\nEnter first name:")
	fmt.Scan(&firstName) // points to the memory of the variable to allow user input

	fmt.Println("Enter last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter email name:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}
