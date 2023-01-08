package helpers

import (
	"strings"
)

func ValidateUserInput(firstName string, email string, userTickets uint, RemaningTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= RemaningTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
