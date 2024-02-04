package dugtrio_test

import (
	"stray/pkg/dugtrio"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCrypto(t *testing.T) {
	srcPassword1 := "password1"
	hashedPassword1, err := dugtrio.BCryptGenerateFromPassword(srcPassword1)
	assert.Nil(t, err)
	t.Logf("hashed Password1: %s", string(hashedPassword1))

	err = dugtrio.BCryptCompareHashAndPassword(srcPassword1, hashedPassword1)
	assert.NoError(t, err)
	err = dugtrio.BCryptCompareHashAndPassword("bad-password", hashedPassword1)
	assert.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error(),
		"It's an invalid Password")
	t.Logf("compare error: %s", err.Error())

	hashedPassword2, err := dugtrio.BCryptGenerateFromPassword(srcPassword1)
	assert.NoError(t, err)
	assert.NotEqual(t, string(hashedPassword1), string(hashedPassword2),
		"bcrypt should return 2 different hashed result")
	t.Logf("hashed Password2: %s", string(hashedPassword2))

	err = dugtrio.BCryptCompareHashAndPassword(srcPassword1, hashedPassword2)
	assert.NoError(t, err)
	err = dugtrio.BCryptCompareHashAndPassword("bad-password", hashedPassword2)
	assert.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error(),
		"It's an invalid Password")
	t.Logf("compare error: %s", err.Error())
}
