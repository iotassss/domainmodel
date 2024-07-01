package domain

import (
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

const (
	MinPasswordLength = 8
)

type Credential struct {
	userUUID     UUID
	passwordHash string
	// TODO: Saltを追加する
}

func NewCredential(userUUID UUID) *Credential {
	return &Credential{
		userUUID: userUUID,
	}
}

func (c *Credential) UserUUID() UUID {
	return c.userUUID
}

func (c *Credential) PasswordHash() string {
	return c.passwordHash
}

// caution: this method is not to be used in case of user registration
func (c *Credential) SetPasswordHash(passwordHash string) {
	c.passwordHash = passwordHash
}

// use this method only in case of user registration
func (c *Credential) SetPassword(plainPassword string) error {
	if len(plainPassword) < MinPasswordLength {
		return &ValidationError{Msg: "password must be at least 8 characters long"}
	}
	if !isValidPasswordComplexity(plainPassword) {
		return &ValidationError{Msg: "password must contain uppercase, lowercase, numbers, and special characters"}
	}
	hash, err := hashPassword(plainPassword)
	if err != nil {
		return &ServerError{Msg: "failed to hash password"}
	}

	c.passwordHash = hash

	return nil
}

func (p *Credential) Verify(plainPassword string) bool {
	return checkPasswordHash(plainPassword, p.passwordHash)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func isValidPasswordComplexity(password string) bool {
	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}
