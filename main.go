package main

import (
	"booking-app/helpers"
	"fmt"
	"sync"
	"time"
)

const conferanceTickets uint = 50

var conferanceName string = "Go Conference"
var RemaningTickets uint = conferanceTickets

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

var email string
var firstName string
var userTickets uint

type UserData struct {
	firstName string
	// lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	for {
		if RemaningTickets == 0 {
			fmt.Println("Our conference is booked out.Come back next year.")
			break
		}
		greetUsers(conferanceName, conferanceTickets, RemaningTickets)
		wg.Add(1)
		getUserInputs()
		go sendTicket(userTickets, firstName, email)
		//wg.Wait()

		isValidEmail, isValidName, isValidTicketNumber := helpers.ValidateUserInput(firstName, email, userTickets, RemaningTickets)
		if isValidEmail && isValidName && isValidTicketNumber {

			if userTickets > RemaningTickets {
				fmt.Printf("You cant buy more than remeaning tickets: %v \n", RemaningTickets)
				continue
			}

			firstNames := getFirstNames()
			fmt.Println(firstNames)
			sellTickets()
			RemaningTickets = RemaningTickets - userTickets

		} else {
			fmt.Println("Wrong input pls try again")
		}

	}
	wg.Wait()

}

func greetUsers(confName string, confTicket uint, remeaningTickets uint) {
	fmt.Printf("conferanceTickets is %T, RemaningTickets is %T, conferanceName is %T\n", confTicket, remeaningTickets, confName)
	fmt.Printf("Welcome to %v booking appliacation\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still avilable.\n", confTicket, remeaningTickets)
	fmt.Printf("Get your tickets here to attend\n")
}

func getUserInputs() {
	fmt.Println("Enter your name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your email adress: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
}
func sellTickets() {
	// var userData = make(map[UserData])
	// userData["firstName"] = firstName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	var userData = UserData{
		firstName:       firstName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recive a confirmation email at %v \n", firstName, email, userTickets, RemaningTickets)
	fmt.Printf(" %v tickets remaning for %v \n", RemaningTickets, conferanceName)

	fmt.Printf("Theese are all our bookings: %v\n", bookings)
}
func checkCity(city string) {
	switch city {
	case "Ankara", "Kars":
		//sa
	case "Istanbul":
		//as
	case "Konya":

	default:
		fmt.Println("No valid city selected!")

	}

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames
}

func sendTicket(userTicket uint, name string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v", userTicket, name)
	fmt.Println("###########")
	fmt.Printf("Sending ticket:\n %v \nto email adress %v\n", ticket, email)
	fmt.Println("###########")
	wg.Done()
}
