package phonenumber

import (
	"errors"
	"fmt"
)

// Number cleans and validates a phone number, returning the 10-digit string.
func Number(phoneNumber string) (string, error) {
	var digits []byte
	for i := 0; i < len(phoneNumber); i++ {
		c := phoneNumber[i]
		if c >= '0' && c <= '9' {
			digits = append(digits, c)
		}
	}

	// Handle 11-digit numbers with country code '1'
	if len(digits) == 11 {
		if digits[0] == '1' {
			digits = digits[1:]
		} else {
			return "", errors.New("11 digits must start with 1")
		}
	}

	// Must be exactly 10 digits after stripping country code
	if len(digits) != 10 {
		return "", errors.New("invalid number of digits")
	}

	// Area code (first 3 digits): first digit must be 2-9
	if digits[0] == '0' || digits[0] == '1' {
		return "", errors.New("area code cannot start with 0 or 1")
	}

	// Exchange code (next 3 digits): first digit must be 2-9
	if digits[3] == '0' || digits[3] == '1' {
		return "", errors.New("exchange code cannot start with 0 or 1")
	}

	return string(digits), nil
}

// AreaCode extracts the area code from a phone number.
func AreaCode(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return num[:3], nil
}

// Format formats a phone number as (XXX) XXX-XXXX.
func Format(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", num[:3], num[3:6], num[6:]), nil
}