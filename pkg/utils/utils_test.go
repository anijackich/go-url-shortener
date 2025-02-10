package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/anijackich/go-url-shortener/pkg/utils"
)

type testCase struct {
	alphabet string
	length   int
}

func containsSymbol(a string, b rune) bool {
	for _, c := range a {
		if c == b {
			return true
		}
	}
	return false
}

func assertConsistsOf(t *testing.T, a string, b string) {
	for _, c := range a {
		if !containsSymbol(b, c) {
			assert.True(t, false, "\"%s\" contains \"%c\" not from \"%s\"", a, c, b)
			return
		}
	}
	assert.True(t, true)
}

func TestGenerateRandomStringValid(t *testing.T) {
	validCases := []testCase{
		{"abc123", 10},
		{"+-!@#$%^&*()[]{}", 5},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_", 25},
	}

	for _, c := range validCases {
		randomString, err := utils.GenerateRandomString(c.alphabet, c.length)

		assert.NoError(t, err)
		{
			assert.NotEmpty(t, randomString, "String is empty")
			assert.Len(t, randomString, c.length, "\"%s\" length is not equal %d", randomString, c.length)
			assertConsistsOf(t, randomString, c.alphabet)
		}
	}
}

func TestGenerateRandomStringInvalid(t *testing.T) {
	invalidCases := []testCase{
		{"", 5},
		{"abc123", -5},
		{"", 0},
	}

	for _, c := range invalidCases {
		_, err := utils.GenerateRandomString(c.alphabet, c.length)

		if assert.Error(t, err) {
			if c.length < 1 {
				assert.Equal(t, utils.ErrInvalidLength, err)
			} else if len(c.alphabet) < 1 {
				assert.Equal(t, utils.ErrEmptyAlphabet, err)
			}
		}
	}
}
