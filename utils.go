package main

import (
	"golang.org/x/crypto/bcrypt"
)

type UtilityFunctions struct{}

var utils = UtilityFunctions{}

func (*UtilityFunctions) createPasswordHash(password string) (string, error) {
	inBytes := []byte(password)
	hashedBytes, err := bcrypt.GenerateFromPassword(inBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes[:]), nil
}

func (*UtilityFunctions) verifyPasswordHash(hash string, password string) error {
	inBytes := []byte(password)
	hashedBytes := []byte(hash)
	return bcrypt.CompareHashAndPassword(inBytes, hashedBytes)
}
