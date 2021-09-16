/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_CrlfTransformer_GivenInput_ReturnsExpectedOutput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
		buffer   int
	}{
		{"Plain string", "hello world", "hello world", 100},
		{"Windows EoLn", "hello\r\nworld\r\n", "hello\r\nworld\r\n", 100},
		{"Linux EoLn", "hello\nworld\n", "hello\r\nworld\r\n", 100},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)
			transformer := &crlfTransformer{}

			buffer := make([]byte, c.buffer)
			written, read, err := transformer.Transform(buffer, []byte(c.input), false)
			g.Expect(err).To(Succeed())
			g.Expect(read).To(Equal(len(c.input)))

			actual := string(buffer[0:written])
			g.Expect(actual).To(Equal(c.expected))
		})
	}
}
