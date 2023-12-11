package helper

import (
	"fmt"
	"strings"
	"time"
)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOfTicket uint
}

func GreetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("# Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("# We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("# Get your ticket here to attend")
}

func InputText(typeEntre string, value string) string {
	fmt.Printf("Enter your %v: ", typeEntre)
	fmt.Scan(&value)
	return value
}

func InputNumber(typeEntre string, value uint) uint {
	fmt.Printf("Enter number of %v: ", typeEntre)
	fmt.Scan(&value)
	return value
}

func Separator() {
	fmt.Println("")
	fmt.Println("###-----------------------###")
	fmt.Println("")
}

func EmpyLine() {
	fmt.Println()
}

func ConfirmationBook(firstName string, lastName string, userTickets uint, email string) {
	fmt.Printf("Thanks %v %v for booking %v tickets. You will receive a confirmation email at %v.\n\n", firstName, lastName, userTickets, email)
}

func AvailablesTickets(remainingTickets uint, conferenceName string) {
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func ValidateUserInputs(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func TicketsAvailables(remainingTickets uint) bool {
	return remainingTickets == 0
}

func Rest(a uint, b uint) uint {
	return a - b
}

func ListNamesGenerator(list []UserData, firstName string, lastName string, email string, usertTikects uint) []UserData {

	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		numberOfTicket: usertTikects,
	}
	return append(list, userData)
}

func ListFirstNamesGenerator(bookings []UserData) []string {
	listExit := []string{}
	for _, booking := range bookings {
		listExit = append(listExit, booking.firstName)
	}
	return listExit
}

func ShowFirstNamesBookings(firstNames []string) {
	fmt.Printf("These are all bookings: %v\n", firstNames)
}

func ByeBye() {
	fmt.Printf("Our conference is booked out. Come back next year.\n")
}

func AlertInvalidName() {
	fmt.Println("First name or last name you entered is too shot.")
}
func AlertInvalidEmail() {
	fmt.Println("Email address you entered doesn't contain @ sign.")
}
func AlertInvalidTicketNumber() {
	fmt.Println("Number of ticket you entered is invalid.")
}

func SendEmail(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(4 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("#\nSending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("####################")
}
