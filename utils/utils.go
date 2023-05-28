package utils

import "net/mail"

// IsEmailValid checks if the email provided is valid by regex.
func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
