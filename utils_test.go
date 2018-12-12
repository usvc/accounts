package main

import (
	"testing"
)

func TestCreatePasswordHash(t *testing.T) {
	testPassword := "hello"
	hash1, _ := utils.CreatePasswordHash(testPassword)
	hash2, _ := utils.CreatePasswordHash(testPassword)
	if hash1 == hash2 {
		t.Errorf("hashing of passphrase '%s' yielded the same hash '%s'", testPassword, hash1)
	}
}

func TestVerifyPasswordHash(t *testing.T) {
	testPassword := "hello"
	hash, _ := utils.CreatePasswordHash(testPassword)
	err := utils.VerifyPasswordHash(testPassword, hash)
	if err != nil {
		t.Errorf("verification of passphrase '$s' failed with: %s", err)
	}
}

func TestValidateEmailValidCases(t *testing.T) {
	validEmails := []string{
		"user@domain.com",          // normal af email #1
		"user@domain.net",          // normal af email #2
		"user@domain.org",          // normal af email #3
		"user@domain.com.sg",       // normal af email #4
		"1171824681273@domain.com", // email
		"user1999@domain.com",      // with postfixed numbers
		"1999user@domain.com",      // with prefixed numbers
		"user.1999@domain.com",     // with dot
		"user_1999@domain.com",     // with underscore
		"abcdefghijklmnop" +
			"abcdefghijklmnop" +
			"abcdefghijklmnop" +
			"abcdefghijklmno" +
			"@domain.com", // superlong local part (63 characters)
		"a@" +
			"abcdefghijklmnopqrstuvwxyz0123456789" +
			"abcdefghijklmnopqrstuvwxyz0123456789" +
			"abcdefghijklmnopqrstuvwxyz0123456789" +
			"abcdefghijklmnopqrstuvwxyz0123456789" +
			"abcdefghijklmnopqrstuvwxyz0123456789" +
			"abcdefghijklmnopqrstuvwxyz0123456789" +
			"abcdefghijklmnopqrstuvwxyz0123456.com", // superlong domain part (253 characters)
	}
	var failedEmails []string
	for _, validEmail := range validEmails {
		if err := utils.ValidateEmail(validEmail); err != nil {
			failedEmails = append(failedEmails, validEmail)
		}
	}
	if len(failedEmails) > 0 {
		t.Errorf("valid emails marked as invalid: %v", failedEmails)
	}
}

func TestValidateEmailInvalidCases(t *testing.T) {
	invalidEmails := []string{
		"user!@domain.com", // wonky symbols
		"user@@domain.com", // wonky symbols
		"use#r@domain.com", // wonky symbols
		"use$r@domain.com", // wonky symbols
		"use%r@domain.com", // wonky symbols
		"use^r@domain.com", // wonky symbols
		"use&r@domain.com", // wonky symbols
		"use*r@domain.com", // wonky symbols
		"use(r@domain.com", // wonky symbols
		"use)r@domain.com", // wonky symbols
		"use=r@domain.com", // wonky symbols
		"user-@domain.com", // wonky postfixed symbols
		"user_@domain.com", // wonky postfixed symbols
		"user+@domain.com", // wonky postfixed symbols
		"user.@domain.com", // wonky postfixed symbols
		"-user@domain.com", // wonky prefixed symbols
		"_user@domain.com", // wonky prefixed symbols
		"+user@domain.com", // wonky prefixed symbols
		".user@domain.com", // wonky prefixed symbols
		"user@domain",      // wonky domain
		"user@domain.",     // wonky domain postfixed symbol
		"user@dom_ain.com", // wonky domain symbol
		"user@.domain.com", // wonky domain prefixed symbol
		"abcdefghijklmnop" +
			"abcdefghijklmnop" +
			"abcdefghijklmnop" +
			"abcdefghijklmnop" +
			"@domain.com", // superlong local part (63 characters)
		"a@" +
			"abcdefghijklmnopqrstuvwxyz0123456789" +
			"abcdefghijklmnopqrstuvwxyz0123456789" +
			"abcdefghijklmnopqrstuvwxyz0123456789" +
			"abcdefghijklmnopqrstuvwxyz0123456789" +
			"abcdefghijklmnopqrstuvwxyz0123456789" +
			"abcdefghijklmnopqrstuvwxyz0123456789" +
			"abcdefghijklmnopqrstuvwxyz01234567.com", // superlong domain part (253 characters)
	}
	var failedEmails []string
	for _, validEmail := range invalidEmails {
		if err := utils.ValidateEmail(validEmail); err == nil {
			failedEmails = append(failedEmails, validEmail)
		}
	}
	if len(failedEmails) > 0 {
		t.Errorf("invalid emails marked as valid: %v", failedEmails)
	}
}

func TestValidateValidPassword(t *testing.T) {
	validPasswords := []string{
		"abcdef1!",
		"abc1!def",
		"1!abcdef",
	}
	var failedPasswords []string
	for _, validPassword := range validPasswords {
		if err := utils.ValidatePassword(validPassword); err != nil {
			failedPasswords = append(failedPasswords, validPassword)
		}
	}
	if len(failedPasswords) > 0 {
		t.Errorf("valid passwords marked as invalid: %v", failedPasswords)
	}
}

func TestValidateInvalidPassword(t *testing.T) {
	invalidPasswords := []string{
		"abcde1!",
		"abcdefgh",
		"abcdefg1",
		"abcde111",
		"abcdefg!",
		"abcde!!!",
	}
	var failedPasswords []string
	for _, invalidPassword := range invalidPasswords {
		if err := utils.ValidatePassword(invalidPassword); err == nil {
			failedPasswords = append(failedPasswords, invalidPassword)
		}
	}
	if len(failedPasswords) > 0 {
		t.Errorf("invalid passwords marked as valid: %v", failedPasswords)
	}
}

func TestValidateValidUsernames(t *testing.T) {
	validUsernames := []string{
		"1234",
		"abcd",
		"ab-cd",
		"ab_cd",
		"ab.cd",
		"abcdef12.3",
		"1.23abcdef",
	}
	var failedUsernames []string
	for _, validUsername := range validUsernames {
		if err := utils.ValidateUsername(validUsername); err != nil {
			failedUsernames = append(failedUsernames, validUsername)
		}
	}
	if len(failedUsernames) > 0 {
		t.Errorf("valid usernames marked as invalid: %v", failedUsernames)
	}
}

func TestValidateInvalidUsernames(t *testing.T) {
	invalidUsernames := []string{
		"abc",
		".abc",
		"abc.",
		"_abc",
		"abc_",
		"-abc",
		"abc-",
	}
	var failedUsernames []string
	for _, invalidUsername := range invalidUsernames {
		if err := utils.ValidateUsername(invalidUsername); err == nil {
			failedUsernames = append(failedUsernames, invalidUsername)
		}
	}
	if len(failedUsernames) > 0 {
		t.Errorf("invalid usernames marked as valid: %v", failedUsernames)
	}
}
