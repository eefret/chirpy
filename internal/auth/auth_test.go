package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

// TestMakeJWT ensures that a token can be created and validated successfully.
func TestMakeJWT(t *testing.T) {
	tokenSecret := "supersecret"
	userID := uuid.New()
	expiresIn := time.Minute

	token, err := MakeJWT(userID, tokenSecret, expiresIn)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedUserID, err := ValidateJWT(token, tokenSecret)
	assert.NoError(t, err)
	assert.Equal(t, userID, parsedUserID)
}

// TestValidateJWTInvalidToken ensures that an invalid token is rejected.
func TestValidateJWTInvalidToken(t *testing.T) {
	tokenSecret := "supersecret"
	invalidToken := "invalid.token.here"

	_, err := ValidateJWT(invalidToken, tokenSecret)
	assert.Error(t, err)
}

// TestValidateJWTExpiredToken ensures that expired tokens are rejected.
func TestValidateJWTExpiredToken(t *testing.T) {
	tokenSecret := "supersecret"
	userID := uuid.New()
	expiredTime := -time.Minute // Token expired 1 minute ago

	token, err := MakeJWT(userID, tokenSecret, expiredTime)
	assert.NoError(t, err)

	_, err = ValidateJWT(token, tokenSecret)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "token is expired")
}
