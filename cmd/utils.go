package cmd

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers          = "0123456789"
	specialChars     = "!@#$%^&*()-_+=<>?"
)

func generatePassword(length int, includes string) (string, error) {
	password := make([]byte, length)
	for i := range password {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(includes))))
		if err != nil {
			return "", err
		}
		password[i] = includes[index.Int64()]
	}
	return string(password), nil
}

func generateIncludes(isUpperCase, isLowerCase, isNumbers, isSpecialChar bool) string {
	var includes string

	if isLowerCase {
		includes += lowercaseLetters
	}
	if isUpperCase {
		includes += uppercaseLetters
	}
	if isNumbers {
		includes += numbers
	}
	if isSpecialChar {
		includes += specialChars
	}

	return includes
}

func CheckUpperCase(password string) bool {
	for _, c := range password {
		if c >= 'A' && c <= 'Z' {
			return true
		}
	}
	return false
}

func CheckLowerCase(password string) bool {
	for _, c := range password {
		if c >= 'a' && c <= 'z' {
			return true
		}
	}
	return false
}

func CheckNumbers(password string) bool {
	for _, c := range password {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func CheckSpecialCharacters(password string) bool {
	for _, c := range password {
		if strings.ContainsRune(specialChars, c) {
			return true
		}
	}
	return false
}

func CheckPasswordStrength(password string) bool {
	if len(password) >= 12 && CheckUpperCase(password) &&
		CheckLowerCase(password) && CheckSpecialCharacters(password) &&
		CheckNumbers(password) {
		return true
	}
	return false
}
