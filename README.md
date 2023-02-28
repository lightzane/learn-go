# Learn Go

Learning GO language (https://go.dev/)

## Youtube Reference

- https://www.youtube.com/watch?v=yyUHQIec83I

## Contents:

1. [Setup GO Project](#start-go)
2. [Hello World](#create-main-entry-point)
3. [Variables](#variables)
4. [Print and Format](#print-and-format)
5. [Data Types](#data-types)
6. [Pointers](#pointers)
7. [Arrays and Slices](#arrays-and-slices)
8. [Loops](#loops)
9. [Conditions](#conditions)
10. [Switch Statements](#switch-statements)
11. [Functions](#functions)
    - [Pointers and Functions](#pointers-and-functions)
      - [Passing a Copy of Object](#passing-a-copy-of-object)
      - [Passing Same Reference of Object](#passing-same-reference-of-the-object)
12. [Split Code](#split-code-in-different-files-and-packages)
    - [Using Same Package](#using-same-package)
    - [Using Different Package](#using-different-package)
13. [Maps](#maps)
14. [Structs (OOP)](#structs)
15. [Goroutines - Multithreading](#goroutines)

## Start GO

Create a folder named `booking-app` and initialize a **Go module** inside of that folder:

```
go mod init booking-app
```

> `go mod init <module_name>`

It should generate the following file:

```
module booking-app

go 1.20
```

## Create Main Entry Point

Create `main.go` file

```go
// main.go

package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World")
}
```

## Variables

```go
// main.go
func main() {
    var bookingName = "Air Flight Booking"
	const bookingTickets = 50

	//
	// var remainingTickets = 50
	//
	// is the same as
	//
	// remainingTickets := 50
	//
	// But this is only for "var" variables, and not for constants
    // This is also should only be declared within functions

	remainingTickets := bookingTickets
}
```

## Print and Format

```go
fmt.Printf("Welcome to %v application\n", bookingName)
fmt.Printf("We have total of %v tickets and %v are still available.\n", bookingTickets, remainingTickets)
fmt.Println("Get your tickets here")
```

**Placeholders**

`%v` - returns the **value** of the specified variable
`%T` - returns the **type** of the specified variable
`%x` - returns the **hexadecimal value** of the specified variable (in lowercase)
`%X` - same as above but in uppercase

See other available formats: https://pkg.go.dev/fmt

## Data Types

```go

var userName string
var userTickets int

userName = "Lightzane"
userTickets = 2

```

## Pointers

Each variable has a **value**. The values are stored in a **memory**. To access these memory, we can make use of **pointers**. Same concept in `C/C++` languages

```go
var tickets = 50
fmt.Println(tickets) // outputs => 50
fmt.Println(&tickets) // outputs => 0xc000018160
```

Pointers are commonly used when asking for a user input in `cli` terminal environment through `fmt.Scan()`

Pointers are also used when we want to pass the same reference of the object into a function as a parameter instead of making a copy of the same object. [See Example Here](#pointers-and-functions)

```go
type Cookies struct {}
```

```go
// Ask for a user input
var firstName string

fmt.Println("Enter first name:")
fmt.Scan(&firstName)
fmt.Printf("User input is: %v", firstName)
```

## Arrays and Slices

**Initialize Array with Fixed length and starting values**

```go
// Create an Array
// with a max of 50 length
// and initializes with 3 items/size
var fruits = [50]string{"apple","banana","cherry"}
```

**Declare Array with Fixed length**

```go
// Define an Array
// with a max of 50 length
var fruits [50]string

// and add a value later...
fruits[0] = "apple"
fruits[1] = "banana"
```

**Slices** or Dynamic Arrays

Automatically expands in size.

```go
// Define a slice
var fruits []string

// or
// fruits := []string{}

fruits = append(fruits, "apple", "banana")
```

> `append` is a built-in function in **Go** and also same as `.push()` in **Javascript**

**Creating an Array from String splits**

```go
fruits := "apple banana cherry"

fruitsList := strings.Fields(fruits) // splits the string into words (seperated by spaces)
```

## Loops

```go
for {
    fmt.Println("Infinite loop")
}

// equivalent to above
for true {
    fmt.Println("Infinite loop")
}

// With fixed number
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// Looping in a list
fruits := []string{"apple", "banana", "cherry"}

// `range` allows us to iterate over a list
for index, fruit := range fruits {
    fmt.Printf("%v: %v", index, fruit)
}

// when index is not needed:
for _, fruit := range fruits {
    fmt.Printf("Fruit name: %v", fruit)
}
```

## Conditions

```go
x, y := 5, 10

if y > x {
    fmt.Println("True")
} else if x == y {
    fmt.Println("Equal")
} else {
    fmt.Println("False")
}
```

## Switch Statements

```go
city := "London"

switch city {

	case "New York":
		fmt.Println("You selected USA")

	case "Singapore":
		fmt.Println("You selected Singapore")

	case "London":
		fmt.Println("You selected England")

	default:
		fmt.Println("No valid city")

}
```

## Functions

```go
func add(x int, y int) int {
	return x + y
}

func multipleReturnValues() (string, bool, int) {
    return "hello", true, 25
}
```

### Pointers and Functions

Pointers are also used when we want to pass the same reference of the object into a function as a parameter instead of making a copy of the same object.

Consider having this object structure

```go
package main

import "fmt"

type Cookies struct {
	name string
}
```

#### Passing a Copy of Object

```go
func updateName(cookies Cookies) {
    // cookies - is now a copy of the object passed
	cookies.name = "Oatmeal"
	fmt.Println("updateName", cookies)
}

func main() {

    myCookies := Cookies{
        name: "Chocolate Chip",
    }

    updateName(myCookies)

    fmt.Println("main", myCookies)
}
```

**Output**

```
updateName {Oatmeal}
main {Chocolate Chip}
```

Notice the `main` does NOT have the updated value!

#### Passing Same Reference of the Object

```go
// adding "*" within the parameter will reference to the same memory / pointer
// will also tell `Go` NOT to create a copy of the object passed
// instead, modify the same reference to that object
func updateName(cookies *Cookies) {
    // cookies - will refer to the same object that was passed
	cookies.name = "Oatmeal"
	fmt.Println("updateName", cookies)
}

func main() {

    myCookies := Cookies{
        name: "Chocolate Chip",
    }

    updateName(&myCookies) // must also pass a Pointer

    fmt.Println("main", myCookies)
}
```

**Output**

```
updateName &{Oatmeal}
main {Oatmeal}
```

Notice the `main` already have the updated value!

## Split Code in Different Files and Packages

### Using same Package

```go
// main.go
package main

func main() {
    otherFunction()
}
```

```go
// other-file.go
package main

import "fmt"

func otherFunction() {
    fmt.Println("Printing...!")
}
```

### Using Different Package

`go.mod` file contains the following

```
module booking-app
```

```go
// main.go
package main

import (
    "booking-app/questioner" // if the filename is the same with the package name
    IDUtil "booking-app/shared/utils" // if the filename is different with the package name
)

func main() {
    questioner.PublicFunction()
    IDUtil.ObjectID()
}
```

**questioner.go**

```go
// questioner/questioner.go
package questioner

import "fmt"

func PublicFunction() {
    fmt.Println("Function name written in PascalCase will exports this function and accessible to public")
}

func privateFunction() {
    // will not be accessible outside this package
}
```

**id.util.go**

```go
// shared/utils/id.util.go
package IDUtil

func ObjectID() {

}
```

## Maps

A map is also a custom data type. It is also a collection of **key/value** pairs

But all the values **cannot** have a mixed data types.

```go
// Create a map for a user
var userData = make(map[string]string)
userData["username"] = "Lightzane"
```

## Structs

Structures defines key value pairs that allows a mixed data types

```go
// Create a custom type that is called "UserData"
type UserData struct {
    firstName string
    lastName string
    email string
    numberOfTickets uint
}

var userDataList = make([]UserData, 0) // initialize with empty list

userData := UserData{
    firstName:       "firstName",
    lastName:        "lastName",
    email:           "email@email.com",
    numberOfTickets: 5,
}

userDataList = append(userDataList, userData)
```

## Goroutines

**Concurrency** (multithreading concept) in GO is easy!

```go
// main.go
package main

import (
    "time"
    "fmt"
)

func main() {
    // executes on a seperate thread by adding "go" keyword
    go sendTicket()
    fmt.Println("This line will IMMEDIATELY executed")

    // will take time
    sendTicket() // blocks the next line until it is finished
    fmt.Println("This line will not get immediately executed")
}

func sendTicket() {
    // simulate delay
    time.Sleep(10 * time.Second)
    fmt.Println("Done!")
}
```

Sure, the program will not wait anymore on those functions that we added on another thread.

But when all the codes in the main thread are done.

Then the program will immediately exit even without the other threads are done.

So to handle this, we need to wait for these threads.

```go
// main.go
package main

import (
    "time"
    "fmt"
    "sync"
)

var wg = sync.WaitGroup{}

func main() {

    // Adds a counter to wait
    // usually should be the same on how many "go" keyword is added
    wg.Add(1)

    // executes on a seperate thread by adding "go" keyword
    go sendTicket()
    fmt.Println("This line will IMMEDIATELY executed")

    // Before main thread exits, wait for other threads
    wg.Wait()
}

func sendTicket() {
    // simulate delay
    time.Sleep(10 * time.Second)
    fmt.Println("Done!")

    wg.Done()
}
```
