package domain

import (
	"regexp"
)

const (
	emailRegexPattern = `^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`
)

type Email struct {
	address string
}

func NewEmail(address string) (Email, error) {
	if !regexp.MustCompile(emailRegexPattern).MatchString(address) {
		return Email{}, &ValidationError{Msg: "invalid email address"}
	}
	return Email{address: address}, nil
}

func (e Email) Address() string {
	return e.address
}
