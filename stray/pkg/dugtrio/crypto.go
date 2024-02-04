package dugtrio

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	defaultBCryptCost int = bcrypt.DefaultCost
)

func BCryptGenerateFromPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), defaultBCryptCost)
}

func BCryptCompareHashAndPassword(password string, hashedPassword []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
}
