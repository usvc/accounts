package main

import (
	"testing"
)

func TestCreatePasswordHash(t *testing.T) {
	testPassword := "hello"
	hash1, _ := utils.createPasswordHash(testPassword)
	hash2, _ := utils.createPasswordHash(testPassword)
	if hash1 == hash2 {
		t.Errorf("hashing of passphrase '%s' yielded the same hash '%s'", testPassword, hash1)
	}
}

func TestVerifyPasswordHash(t *testing.T) {
	testPassword := "hello"
	hash, _ := utils.createPasswordHash(testPassword)
	err := utils.verifyPasswordHash(testPassword, hash)
	if err != nil {
		t.Errorf("verification of passphrase '$s' failed with: %s", err)
	}
}
