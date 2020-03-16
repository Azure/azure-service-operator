// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package helpers

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"

	"github.com/sethvargo/go-password/password"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

func randomStringWithCharset(length int, charset string) string {
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return randomStringWithCharset(length, charset)
}

// IsBeingDeleted returns true if the current object is being deleted from the API server.
func IsBeingDeleted(o metav1.Object) bool {
	return !o.GetDeletionTimestamp().IsZero()
}

// GenerateRandomUsername - helper function to generate random username for sql server
func GenerateRandomUsername(n int, numOfDigits int) (string, error) {

	// Generate a username that is n characters long, with n/2 digits and 0 symbols (not allowed),
	// allowing only lower case letters (upper case not allowed), and disallowing repeat characters.
	res, err := password.Generate(n, numOfDigits, 0, true, false)
	if err != nil {
		return "", err
	}

	return res, nil
}

// GenerateRandomPassword - helper function to generate random password for sql server
func GenerateRandomPassword(n int) (string, error) {

	// Math - Generate a password where: 1/3 of the # of chars are digits, 1/3 of the # of chars are symbols,
	// and the remaining 1/3 is a mix of upper- and lower-case letters
	digits := n / 3
	symbols := n / 3

	// Generate a password that is n characters long, with # of digits and symbols described above,
	// allowing upper and lower case letters, and disallowing repeat characters.
	res, err := password.Generate(n, digits, symbols, false, false)
	if err != nil {
		return "", err
	}

	return res, nil
}

// RemoveNonAlphaNumeric removes all runes that are not letters or digits
func RemoveNonAlphaNumeric(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// PadRightWithRandom pads a string up to a maxLen with random characters
func FillWithRandom(s string, maxLen int) string {
	diff := maxLen - len(s)
	if diff <= 0 {
		return s
	}
	return s + RandomString(diff)
}

// Hash256 hashes the i argument to a sha265 string
func Hash256(i interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", i)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
