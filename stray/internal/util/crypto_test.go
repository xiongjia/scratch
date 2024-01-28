package util_test

import (
	"stray/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrypto(t *testing.T) {
	srcPassword1 := "password1"
	hashedPassword1, err := util.BCryptGenerateFromPassword(srcPassword1)
	assert.Nil(t, err)
	t.Logf("hashed Password1: %s", string(hashedPassword1))

	err = util.BCryptCompareHashAndPassword(srcPassword1, hashedPassword1)
	assert.Nil(t, err)
	err = util.BCryptCompareHashAndPassword("bad-password", hashedPassword1)
	assert.NotNil(t, err)
	t.Logf("compare error: %s", err.Error())

	hashedPassword2, err := util.BCryptGenerateFromPassword(srcPassword1)
	assert.Nil(t, err)
	assert.NotEqual(t, string(hashedPassword1), string(hashedPassword2),
		"bcrypt should return 2 different hashed result")
	t.Logf("hashed Password2: %s", string(hashedPassword2))

	err = util.BCryptCompareHashAndPassword(srcPassword1, hashedPassword2)
	assert.Nil(t, err)
	err = util.BCryptCompareHashAndPassword("bad-password", hashedPassword2)
	assert.NotNil(t, err)
	t.Logf("compare error: %s", err.Error())
}
