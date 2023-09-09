package utils

import (
	"fmt"
	"net/mail"
	"net/url"
	"os"
	"strings"
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

func IsDomainDiffrent(url string) bool {
	// basically this functions removes all the commonly found
	// prefixes from URL such as http, https, www
	// then checks of the remaining string is the DOMAIN itself
	if url == os.Getenv("DOMAIN") {
		return false
	}
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]
	newURL = strings.Replace(newURL, " ", "", -1)

	fmt.Println(newURL)
	fmt.Println(os.Getenv("DOMAIN"))

	return newURL != os.Getenv("DOMAIN")
}
