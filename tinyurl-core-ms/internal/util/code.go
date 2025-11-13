package util

import (
	crand "crypto/rand"
	"fmt"
	mrand "math/rand"
	"time"
)

const (
	codeCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	codeLength  = 8
)

// Generates a random short code of fixed length of 8 using alphanumeric characters.
func GenerateCode() string {
	b := make([]byte, codeLength)
	if _, err := crand.Read(b); err != nil {
		// fallback to math/rand if crypto/rand fails
		r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
		for i := range b {
			b[i] = codeCharset[r.Intn(len(codeCharset))]
		}
		return string(b)
	}

	for i := range b {
		b[i] = codeCharset[int(b[i])%len(codeCharset)]
	}
	return string(b)
}

// Validates if the given code is a valid short code of fixed length 8.
func ValidateCode(code string) error {
	if code == "" {
		return fmt.Errorf("given code is empty")
	}

	if len(code) != codeLength {
		return fmt.Errorf("code length is different than expected (%d)", codeLength)
	}

	for i := 0; i < codeLength; i++ {
		c := code[i]
		if !((c >= 'a' && c <= 'z') ||
			(c >= 'A' && c <= 'Z') ||
			(c >= '0' && c <= '9')) {
			return fmt.Errorf("code contains invalid character: %c", c)
		}
	}

	return nil
}
