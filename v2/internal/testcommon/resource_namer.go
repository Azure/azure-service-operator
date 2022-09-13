/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"hash/fnv"
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
// If the original ResourceNamer was entirely random (mode == ResourceNamerModeRandom),
// the returned namer is not actually based on the test name and is instead still entirely random
func (n ResourceNamer) WithTestName(testName string) ResourceNamer {
	// Short circuit for the case we're supposed to be totally random, as we already are
	if n.mode == ResourceNamerModeRandom {
		return n
	}

	return n.NewResourceNamer(testName)
}

// NewResourceNamer returns a ResourceNamer that generates random
// suffixes based upon the test name
func (rnc ResourceNameConfig) NewResourceNamer(testName string) ResourceNamer {
	var r *rand.Rand
	if rnc.mode == ResourceNamerModeRandom {
		//nolint:gosec // do not want cryptographic randomness here
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	} else {
		hasher := fnv.New64()
		n, err := hasher.Write([]byte(testName))
		if n != len(testName) || err != nil {
			panic("failed to write hash")
		}

		seed := hasher.Sum64()
		//nolint:gosec // do not want cryptographic randomness here
		r = rand.New(rand.NewSource(int64(seed)))
	}

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

func (n ResourceNamer) GeneratePassword() string {
	return n.GeneratePasswordOfLength(n.randomChars)
}

func (n ResourceNamer) GeneratePasswordOfLength(length int) string {
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^&*()")
	// This pass + <content> + pass pattern is to make it so that matchers can reliably find and prune
	// generated passwords from the recordings. If you change it make sure to change the passwordMatcher
	// in test_context.go as well.
	return "pass" + n.makeRandomStringOfLength(length, runes) + "pass"
}

func (n ResourceNamer) GenerateUUID() (uuid.UUID, error) {
	return uuid.NewRandomFromReader(n.rand)
}
