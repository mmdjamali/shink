package utils

import (
	"net/mail"
	"net/url"
)

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

func IsValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}
