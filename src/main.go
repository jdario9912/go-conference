// 2:02

package main

import (
	"desde-cero/src/helper"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := make([]helper.UserData, 0)  
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var noTicketsRemaining bool
	firstNames := []string{}

	helper.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		helper.EmpyLine()

		firstName = helper.InputText("first name", firstName)
		lastName = helper.InputText("last name", lastName)
		email = helper.InputText("email", email)
		userTickets = helper.InputNumber("tickets", userTickets)

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = helper.Rest(remainingTickets, userTickets)

			bookings = helper.ListNamesGenerator(bookings, firstName, lastName, email, userTickets)

			helper.EmpyLine()

			helper.ConfirmationBook(firstName, lastName, userTickets, email)

			wg.Add(1)
			go helper.SendEmail(userTickets, firstName, lastName, email)

			helper.AvailablesTickets(remainingTickets, conferenceName)

			firstNames = helper.ListFirstNamesGenerator(bookings)

			helper.EmpyLine()

			helper.ShowFirstNamesBookings(firstNames)

			helper.Separator()

			noTicketsRemaining = helper.TicketsAvailables(remainingTickets)

			if noTicketsRemaining {

				helper.EmpyLine()
				helper.ByeBye()

				break
			}
		} else {

			if !isValidName {
				helper.AlertInvalidName()
			}
			if !isValidEmail {
				helper.AlertInvalidEmail()
			}
			if !isValidTicketNumber {
				helper.AlertInvalidTicketNumber()
			}
		}
	}
	wg.Done()
}
