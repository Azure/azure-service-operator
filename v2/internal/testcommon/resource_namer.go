/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	crypto "crypto/rand"
	"hash/fnv"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

type ResourceNameConfig struct {
	runes       []rune
	prefix      string
	randomChars int
	separator   string
	mode        ResourceNamerMode
}

type ResourceNamerMode string

const (
	ResourceNamerModeRandomBasedOnTestName = ResourceNamerMode("basedOnTestName")
	ResourceNamerModeRandom                = ResourceNamerMode("random")
)

type ResourceNamer struct {
	ResourceNameConfig
	rand *rand.Rand
}

// WithTestName returns a new ResourceNamer configured based on the provided test name.
// The mode of the original resource namer is preserved.
func (n ResourceNamer) WithTestName(testName string) ResourceNamer {
	return n.NewResourceNamer(testName)
}

// NewResourceNamer returns a ResourceNamer that generates random
// suffixes based upon the test name
func (rnc ResourceNameConfig) NewResourceNamer(testName string) ResourceNamer {
	// Calculate an initial seed based on the test name
	hasher := fnv.New64()
	n, err := hasher.Write([]byte(testName))
	if n != len(testName) || err != nil {
		panic("failed to write hash")
	}

	//nolint:gosec // don't care about int overflow
	seed := int64(hasher.Sum64())

	// This seed is enough to get the same sequence of "random" values every time.
	// If we want a different sequence every time, include time as part of the seed too
	if rnc.mode == ResourceNamerModeRandom {
		t := time.Now().UnixNano()
		seed = seed ^ t
	}

	src := rand.NewSource(seed)
	//nolint:gosec // do not need/want cryptographic randomness here
	r := rand.New(src)

	return ResourceNamer{
		ResourceNameConfig: rnc,
		rand:               r,
	}
}

func NewResourceNameConfig(prefix string, separator string, randomChars int, mode ResourceNamerMode) *ResourceNameConfig {
	return &ResourceNameConfig{
		runes:       []rune("abcdefghijklmnopqrstuvwxyz"),
		prefix:      prefix,
		randomChars: randomChars,
		separator:   separator,
		mode:        mode,
	}
}

// WithSeparator returns a copy of the ResourceNamer with the given separator
func (n ResourceNamer) WithSeparator(separator string) ResourceNamer {
	n.separator = separator
	return n
}

// WithNumRandomChars returns a copy of the ResourceNamer which will generate names with
// a random string of the given length included
func (n ResourceNamer) WithNumRandomChars(num int) ResourceNamer {
	n.randomChars = num
	return n
}

func (n ResourceNamer) makeRandomStringOfLength(num int, runes []rune) string {
	result := make([]rune, num)
	for i := range result {
		result[i] = runes[n.rand.Intn(len(runes))]
	}

	return string(result)
}

func (n ResourceNamer) makeSecureStringOfLength(num int, runes []rune) string {
	result := make([]rune, num)
	for i := range result {
		// Ignoring error here since big.NewInt does not return `0 <= x < length`
		idx, _ := crypto.Int(crypto.Reader, big.NewInt(int64(len(runes))))

		result[i] = runes[idx.Int64()]
	}

	return string(result)
}

func (n ResourceNamer) generateName(prefix string, num int) string {
	result := n.makeRandomStringOfLength(num, n.runes)

	var s []string
	// TODO: Do we need a check here for if both result and prefix are empty?
	if result == "" {
		s = []string{n.prefix, prefix}
	} else if prefix == "" {
		s = []string{n.prefix, result}
	} else {
		s = []string{n.prefix, prefix, result}
	}

	return strings.Join(s, n.separator)
}

func (n ResourceNamer) GenerateName(prefix string) string {
	return n.generateName(prefix, n.randomChars)
}

func (n ResourceNamer) GenerateNameOfLength(length int) string {
	return n.generateName("", length)
}

func (n ResourceNamer) GeneratePassword() string {
	return n.GeneratePasswordOfLength(n.randomChars)
}

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^*()._-")

// GeneratePasswordOfLength generates and returns a non-deterministic password.
// This method does not use any seed value, so the returned password is never stable.
func (n ResourceNamer) GeneratePasswordOfLength(length int) string {
	// This pass + <content> + pass pattern is to make it so that matchers can reliably find and prune
	// generated passwords from the recordings. If you change it make sure to change the passwordMatcher
	// in test_context.go as well.
	return "pass" + n.makeSecureStringOfLength(length, runes) + "pass"
}

// GenerateSecretOfLength generates and returns a non-deterministic secret.
// This method does not use any seed value, so the returned password is never stable.
func (n ResourceNamer) GenerateSecretOfLength(length int) string {
	return n.makeSecureStringOfLength(length, runes)
}

func (n ResourceNamer) GenerateUUID() (uuid.UUID, error) {
	return uuid.NewRandomFromReader(n.rand)
}
