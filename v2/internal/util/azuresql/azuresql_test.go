/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package azuresql

import (
	"testing"
	"time"

	"github.com/microsoft/go-mssqldb/msdsn"
	. "github.com/onsi/gomega"
)

func TestNewConfigTreatsCredentialsAsValues(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	const (
		serverAddress = "victim.database.windows.net"
		database      = "victimdb"
		username      = "expected-admin;server=attacker.example;user id=attacker"
		password      = "p@ss;word='value';server=attacker.example"
	)

	config, err := newConfig(serverAddress, database, ServerPort, username, password)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(config.Host).To(Equal(serverAddress))
	g.Expect(config.Port).To(Equal(uint64(ServerPort)))
	g.Expect(config.Database).To(Equal(database))
	g.Expect(config.User).To(Equal(username))
	g.Expect(config.Password).To(Equal(password))
	g.Expect(config.Encryption).To(Equal(msdsn.Encryption(msdsn.EncryptionRequired)))
	g.Expect(config.TrustServerCertificate).To(BeFalse())
	g.Expect(config.ConnTimeout).To(Equal(30 * time.Second))
	g.Expect(config.Protocols).To(Equal([]string{"tcp"}))
	g.Expect(config.TLSConfig).ToNot(BeNil())
	g.Expect(config.TLSConfig.ServerName).To(Equal(serverAddress))
	g.Expect(config.TLSConfig.InsecureSkipVerify).To(BeFalse())
}

func TestFindBadChars(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		input       string
		expectError bool
	}{
		{
			name:        "clean string passes",
			input:       "validPassword123!@#",
			expectError: false,
		},
		{
			name:        "single quote fails",
			input:       "pass'word",
			expectError: true,
		},
		{
			name:        "double quote fails",
			input:       "pass\"word",
			expectError: true,
		},
		{
			name:        "semicolon fails",
			input:       "pass;word",
			expectError: true,
		},
		{
			name:        "double dash fails",
			input:       "pass--word",
			expectError: true,
		},
		{
			name:        "block comment fails",
			input:       "pass/*word",
			expectError: true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			err := findBadChars(c.input)
			if c.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).ToNot(HaveOccurred())
			}
		})
	}
}

func TestEscapeStringLiteral(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no special chars",
			input:    "simplepassword",
			expected: "'simplepassword'",
		},
		{
			name:     "single quote is doubled",
			input:    "pass'word",
			expected: "'pass''word'",
		},
		{
			name:     "semicolons are preserved",
			input:    "pass;word",
			expected: "'pass;word'",
		},
		{
			name:     "double dashes are preserved",
			input:    "pass--word",
			expected: "'pass--word'",
		},
		{
			name:     "complex special chars",
			input:    "p@ss;w'rd--/*test",
			expected: "'p@ss;w''rd--/*test'",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := escapeStringLiteral(c.input)
			g.Expect(result).To(Equal(c.expected))
		})
	}
}

func TestEscapeBracketIdentifier(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no special chars",
			input:    "username",
			expected: "[username]",
		},
		{
			name:     "closing bracket is doubled",
			input:    "user]name",
			expected: "[user]]name]",
		},
		{
			name:     "multiple closing brackets",
			input:    "user]]name]x",
			expected: "[user]]]]name]]x]",
		},
		{
			name:     "AAD username with at sign",
			input:    "bob@contoso.com",
			expected: "[bob@contoso.com]",
		},
		{
			name:     "other special chars are preserved",
			input:    "user;name--test",
			expected: "[user;name--test]",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := escapeBracketIdentifier(c.input)
			g.Expect(result).To(Equal(c.expected))
		})
	}
}
