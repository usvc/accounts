package main

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var PasswordMinimumLength = 8
var PasswordMandatorySpecialCharacters = true
var PasswordMandatoryNumbers = true

type UtilityFunctions struct{}

var utils = UtilityFunctions{}

func (*UtilityFunctions) CreatePasswordHash(password string) (string, error) {
	inBytes := []byte(password)
	hashedBytes, err := bcrypt.GenerateFromPassword(inBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes[:]), nil
}

func (*UtilityFunctions) VerifyPasswordHash(hash string, password string) error {
	inBytes := []byte(password)
	hashedBytes := []byte(hash)
	return bcrypt.CompareHashAndPassword(inBytes, hashedBytes)
}

// ValidateEmail returns an error if the :email is not valid
func (utility *UtilityFunctions) ValidateEmail(email string) error {
	emailParts := strings.Split(email, "@")
	if len(emailParts) != 2 {
		return errors.New("E_EMAIL_INVALID")
	}
	localPart := emailParts[0]
	if err := utility.validateEmailLocalPart(localPart); err != nil {
		return err
	}
	domainPart := emailParts[1]
	if err := utility.validateEmailDomainPart(domainPart); err != nil {
		return err
	}
	return nil
}

// ValidatePassword returns an error if the :password is not valid
func (utility *UtilityFunctions) ValidatePassword(password string) error {
	specialCharacters := "`~!@#$%^&*()_+-={}[]\\|;:'\",<.>/?"
	if len(password) < PasswordMinimumLength {
		return errors.New("E_PASSWORD_TOO_SHORT")
	} else if !strings.ContainsAny(password, specialCharacters) {
		return errors.New("E_PASSWORD_NO_SPECIAL_CHARACTERS")
	} else if !strings.ContainsAny(password, "1234567890") {
		return errors.New("E_PASSWORD_NO_NUMBERS")
	}
	return nil
}

func (*UtilityFunctions) validateEmailLocalPart(localPart string) error {
	specialCharacters := "!#$%^&*()=\\|;:\"'<>,/?[]{}`~"
	invalidStartingAndEndingCharacters := "-_+."
	if len(localPart) > 63 {
		return errors.New("E_EMAIL_LOCALPART_TOO_LONG")
	}
	if strings.ContainsAny(localPart, specialCharacters) {
		return errors.New("E_EMAIL_SPECIAL_CHARS")
	}
	if strings.LastIndexAny(localPart, invalidStartingAndEndingCharacters) == len(localPart)-1 {
		return errors.New("E_EMAIL_INVALID_LOCALPART_POSTFIX")
	}
	if strings.IndexAny(localPart, invalidStartingAndEndingCharacters) == 0 {
		return errors.New("E_EMAIL_INVALID_LOCALPART_PREFIX")
	}
	return nil
}

func (UtilityFunctions) validateEmailDomainPart(domainPart string) error {
	specialCharacters := "!#$%^&*()=\\|;:\"'<>,/?[]{}`~+_"
	invalidStartingAndEndingCharacters := "-"
	if len(domainPart) > 253 {
		return errors.New("E_EMAIL_DOMAIN_TOO_LONG")
	}
	domainParts := strings.Split(domainPart, ".")
	if len(domainParts) < 2 {
		return errors.New("E_EMAIL_DOMAINPART_INVALID")
	}
	for _, currentDomainPart := range domainParts {
		if strings.ContainsAny(currentDomainPart, specialCharacters) {
			return errors.New("E_EMAIL_SPECIAL_CHARS")
		}
		if strings.LastIndexAny(currentDomainPart, invalidStartingAndEndingCharacters) == len(currentDomainPart)-1 {
			return errors.New("E_EMAIL_INVALID_DOMAINPART_POSTFIX")
		}
		if strings.IndexAny(currentDomainPart, invalidStartingAndEndingCharacters) == 0 {
			return errors.New("E_EMAIL_INVALID_DOMINPART_PREFIX")
		}
	}
	return nil
}
