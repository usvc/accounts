package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var EmailLocalpartMaxLength = 63
var EmailDomainpartMaxLength = 253
var PasswordMinimumLength = 8
var PasswordMandatorySpecialCharacters = true
var PasswordMandatoryNumbers = true
var UsernameMinLength = 4
var UsernameMaxLength = 64

var UtilsErrorEmail = "E_EMAIL_INVALID"

type ValidationError struct {
	Code    string
	Message string
}

func (self *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", self.Code, self.Message)
}

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
		return errors.New(UtilsErrorEmail)
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
		return &ValidationError{
			Code:    "E_PASSWORD_TOO_SHORT",
			Message: fmt.Sprintf("password should be at least of length %v", PasswordMinimumLength),
		}
	} else if !strings.ContainsAny(password, specialCharacters) {
		return &ValidationError{
			Code:    "E_PASSWORD_NO_SPECIAL_CHARACTERS",
			Message: "password should contain at least one special character",
		}
	} else if !strings.ContainsAny(password, "1234567890") {
		return &ValidationError{
			Code:    "E_PASSWORD_NO_NUMBERS",
			Message: "password should contain at least one numerical character",
		}
	}
	return nil
}

func (*UtilityFunctions) ValidateUsername(username string) error {
	test, err := regexp.Compile(`^[a-zA-Z0-9]+[a-zA-Z0-9_\-\.]*[a-zA-Z0-9]$`)
	if err != nil {
		panic(err)
	}
	if len(username) < UsernameMinLength {
		return &ValidationError{
			Code:    "E_USERNAME_TOO_SHORT",
			Message: fmt.Sprintf("username ('%v') should be more than %v characters", username, UsernameMinLength-1),
		}
	}
	if len(username) > UsernameMaxLength {
		return &ValidationError{
			Code:    "E_USERNAME_TOO_LONG",
			Message: fmt.Sprintf("username ('%v') should be less than %v characters", username, UsernameMaxLength),
		}
	}
	if !test.MatchString(username) {
		return &ValidationError{
			Code:    "E_USERNAME_INVALID_CHARACTERS",
			Message: fmt.Sprintf("username ('%v') should be alpha-numeric and cannot start/end with any of [-, _, .]", username),
		}
	}
	return nil
}

func (*UtilityFunctions) validateEmailLocalPart(localPart string) error {
	specialCharacters := "!#$%^&*()=\\|;:\"'<>,/?[]{}`~"
	invalidStartingAndEndingCharacters := "-_+."
	if len(localPart) > EmailLocalpartMaxLength {
		return &ValidationError{
			Code:    "E_EMAIL_LOCALPART_TOO_LONG",
			Message: fmt.Sprintf("email username ('%v') should be less than %v characters", localPart, EmailLocalpartMaxLength),
		}
	}
	if strings.ContainsAny(localPart, specialCharacters) {
		return &ValidationError{
			Code:    "E_EMAIL_LOCALPART_SPECIAL_CHARS",
			Message: fmt.Sprintf("email username ('%v') should not contain any special characters", localPart),
		}
	}
	if strings.LastIndexAny(localPart, invalidStartingAndEndingCharacters) == len(localPart)-1 {
		return &ValidationError{
			Code:    "E_EMAIL_LOCALPART_INVALID_POSTFIX",
			Message: fmt.Sprintf("email username ('%v') should not end with any of %v", localPart, invalidStartingAndEndingCharacters),
		}
	}
	if strings.IndexAny(localPart, invalidStartingAndEndingCharacters) == 0 {
		return &ValidationError{
			Code:    "E_EMAIL_LOCALPART_INVALID_PREFIX",
			Message: fmt.Sprintf("email username ('%v') should not end with any of %v", localPart, invalidStartingAndEndingCharacters),
		}
	}
	return nil
}

func (UtilityFunctions) validateEmailDomainPart(domainPart string) error {
	specialCharacters := "!#$%^&*()=\\|;:\"'<>,/?[]{}`~+_"
	invalidStartingAndEndingCharacters := "-"
	if len(domainPart) > EmailDomainpartMaxLength {
		return &ValidationError{
			Code:    "E_EMAIL_DOMAIN_TOO_LONG",
			Message: fmt.Sprintf("email domain ('%v') should not be more than length %v", domainPart, EmailDomainpartMaxLength),
		}
	}
	domainParts := strings.Split(domainPart, ".")
	if len(domainParts) < 2 {
		return &ValidationError{
			Code:    "E_EMAIL_DOMAINPART_INVALID",
			Message: fmt.Sprintf("email domain ('%v') should not be of a TLD", domainPart),
		}
	}
	for _, domainPartSubsection := range domainParts {
		if strings.ContainsAny(domainPartSubsection, specialCharacters) {
			return &ValidationError{
				Code:    "E_EMAIL_SPECIAL_CHARS",
				Message: fmt.Sprintf("email domain part ('%v') should not contain any special characters (%v)", domainPartSubsection, specialCharacters),
			}
		}
		if strings.LastIndexAny(domainPartSubsection, invalidStartingAndEndingCharacters) == len(domainPartSubsection)-1 {
			return &ValidationError{
				Code:    "E_EMAIL_INVALID_DOMAINPART_POSTFIX",
				Message: fmt.Sprintf("email domain part ('%v') should not end with any of %v", domainPartSubsection, invalidStartingAndEndingCharacters),
			}
		}
		if strings.IndexAny(domainPartSubsection, invalidStartingAndEndingCharacters) == 0 {
			return &ValidationError{
				Code:    "E_EMAIL_INVALID_DOMINPART_PREFIX",
				Message: fmt.Sprintf("email domain part ('%v') should not begin with any of %v", domainPartSubsection, invalidStartingAndEndingCharacters),
			}
		}
	}
	return nil
}
