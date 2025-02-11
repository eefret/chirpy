package auth

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := "testPassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		t.Fatal("hashed password does not match original password")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "testPassword"
	hashedPassword, _ := HashPassword(password)

	if err := CheckPasswordHash(password, hashedPassword); err != nil {
		t.Fatal("expected no error, got", err)
	}

	invalidPassword := "wrongPassword"
	if err := CheckPasswordHash(invalidPassword, hashedPassword); err == nil {
		t.Fatal("expected an error for wrong password, got nil")
	}
}
