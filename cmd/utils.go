package cmd

import (
	"crypto/rand"
	"math/big"
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
	const (
		lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
		uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numbers          = "0123456789"
		specialChars     = "!@#$%^&*()-_+=<>?"
	)

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
