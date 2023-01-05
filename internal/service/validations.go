package service

import (
	"errors"
	"net/mail"
	"unicode"
)

var (
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
)

func checkPassword(password string) error {
	var (
		minLen    = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)
	if len(password) > 7 {
		minLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}
	if minLen && hasLower && hasNumber && hasUpper {
		return nil
	}
	return ErrInvalidPassword
}

func checkUsername(username string) error {
	for _, ch := range username {
		if ch < 33 || ch > 126 {
			return ErrInvalidUsername
		}
	}
	return nil
}

func checkEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return ErrInvalidEmail
	}
	return nil
}
