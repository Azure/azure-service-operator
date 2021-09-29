/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"hash/fnv"
	"math/rand"
	"strings"
)

type ResourceNameConfig struct {
	runes       []rune
	prefix      string
	randomChars int
	separator   string
}

type ResourceNamer struct {
	ResourceNameConfig
	rand *rand.Rand
}

// NewResourceNamer returns a ResourceNamer that generates random
// suffixes based upon the test name
func (rnc ResourceNameConfig) NewResourceNamer(testName string) ResourceNamer {
	hasher := fnv.New64()
	n, err := hasher.Write([]byte(testName))
	if n != len(testName) || err != nil {
		panic("failed to write hash")
	}

	seed := hasher.Sum64()
	return ResourceNamer{
		ResourceNameConfig: rnc,
		//nolint:gosec // do not want cryptographic randomness here
		rand: rand.New(rand.NewSource(int64(seed))),
	}
}

func NewResourceNameConfig(prefix string, separator string, randomChars int) *ResourceNameConfig {
	return &ResourceNameConfig{
		runes:       []rune("abcdefghijklmnopqrstuvwxyz"),
		prefix:      prefix,
		randomChars: randomChars,
		separator:   separator,
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

func (n ResourceNamer) makeRandomStringOfLength(num int) string {
	result := make([]rune, num)
	for i := range result {
		result[i] = n.runes[n.rand.Intn(len(n.runes))]
	}

	return string(result)
}

func (n ResourceNamer) generateName(prefix string, num int) string {
	result := n.makeRandomStringOfLength(num)

	var s []string
	if prefix != "" {
		s = []string{n.prefix, prefix, result}
	} else {
		s = []string{n.prefix, result}
	}

	return strings.Join(s, n.separator)
}

func (n ResourceNamer) GenerateName(prefix string) string {
	return n.generateName(prefix, n.randomChars)
}

func (n ResourceNamer) GeneratePassword() string {
	// This pass + <content> + pass pattern is to make it so that matchers can reliably find and prune
	// generated passwords from the recordings. If you change it make sure to change the passwordMatcher
	// in test_context.go as well.
	return "pass" + n.makeRandomStringOfLength(n.randomChars) + "pass"
}
