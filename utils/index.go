package utils

import "net/mail"

type CustomError struct {
	Message string
}

func (ce *CustomError) Error() string {
	return ce.Message
}

func IsValidEmail(e string) bool {
	_, err := mail.ParseAddress(e)
	return err == nil
}
