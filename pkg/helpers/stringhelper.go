// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package helpers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/sethvargo/go-password/password"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	passwordLength  = 16
	passwordChars   = lowerAlphaChars + upperAlphaChars + numberChars + specialChars
	lowerAlphaChars = "abcdefghijklmnopqrstuvwxyz"
	upperAlphaChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars     = "0123456789"
	specialChars    = "!@#$%^&*"
)

var seededRand = NewSeeded()

// NewPassword generates a strong, random password
// stolen from osba-azure
func NewPassword() string {
	b := make([]byte, passwordLength)
	// Passwords need to include at least one character from each of the three
	// groups. To ensure that, we'll fill each of the first three []byte elements
	// with a random character from a specific group.
	b[0] = lowerAlphaChars[seededRand.Intn(len(lowerAlphaChars))]
	b[1] = upperAlphaChars[seededRand.Intn(len(upperAlphaChars))]
	b[2] = numberChars[seededRand.Intn(len(numberChars))]
	// The remainder of the characters can be completely random and drawn from
	// all three character groups.
	for i := 3; i < passwordLength; i++ {
		b[i] = passwordChars[seededRand.Intn(len(passwordChars))]
	}
	// For good measure, shuffle the elements of the entire []byte so that
	// the 0 character isn't predicatably lowercase, etc...
	for i := range b {
		j := seededRand.Intn(len(b))
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

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
	return randomStringWithCharset(length, lowerAlphaChars)
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
	inBytes, _ := json.Marshal(i)
	h.Write(inBytes)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// ReplaceAny replaces any instance of the strings passes in the chars slice
// replacing a backslash is problematic so it will require 4 eg []string{"\\\\"}
func ReplaceAny(s string, chars []string) string {
	reg := regexp.MustCompile(fmt.Sprintf(`(%s)`, strings.Join(chars, "|")))
	return reg.ReplaceAllString(s, ".")
}
