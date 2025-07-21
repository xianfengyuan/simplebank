package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullname = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLen int, maxLen int) error {
	n := len(value)
	if n < minLen || n > maxLen {
		return fmt.Errorf("must contain from %d-%d characters", minLen, maxLen)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return  err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("must contain lower letters, digits, underscore")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return  err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email")
	}
	return nil
}

func ValidateFullname(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return  err
	}
	if !isValidFullname(value) {
		return fmt.Errorf("must contain letters, space")
	}
	return nil
}
