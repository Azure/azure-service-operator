/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"math/rand"
	"strings"
	"time"
)

type ResourceNamer struct {
	rand        *rand.Rand
	runes       []rune
	prefix      string
	randomChars int
	separator   string
}

func NewResourceNamer(prefix string, separator string, randomChars int) *ResourceNamer {
	return &ResourceNamer{
		// nolint: do not want cryptographic randomness here
		rand:        rand.New(rand.NewSource(time.Now().UnixNano())),
		runes:       []rune("abcdefghijklmnopqrstuvwxyz"),
		prefix:      prefix,
		randomChars: randomChars,
		separator:   separator,
	}
}

func (n *ResourceNamer) WithSeparator(separator string) *ResourceNamer {
	return NewResourceNamer(n.prefix, separator, n.randomChars)
}

func (n *ResourceNamer) generateName(prefix string, num int) string {
	result := make([]rune, num)
	for i := 0; i < num; i++ {
		result[i] = n.runes[n.rand.Intn(len(n.runes))]
	}

	var s []string
	if prefix != "" {
		s = []string{n.prefix, prefix, string(result)}
	} else {
		s = []string{n.prefix, string(result)}
	}

	return strings.Join(s, n.separator)
}

func (n *ResourceNamer) GenerateName(prefix string) string {
	return n.generateName(prefix, n.randomChars)
}
