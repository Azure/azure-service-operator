/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package creds

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ResourceGroupTagsFromEnvironment_givenSetting_returnsExpectedTags(t *testing.T) {
	cases := map[string]struct {
		setting  string
		expected map[string]string
		valid    bool
	}{
		"Unset": {
			setting:  "",
			expected: nil,
			valid:    true,
		},
		"Empty object": {
			setting:  `{}`,
			expected: map[string]string{},
			valid:    true,
		},
		"Single tag": {
			setting: `{"Owner":"someone"}`,
			expected: map[string]string{
				"Owner": "someone",
			},
			valid: true,
		},
		"Several tags": {
			setting: `{"Owner":"someone","Project":"something"}`,
			expected: map[string]string{
				"Owner":   "someone",
				"Project": "something",
			},
			valid: true,
		},
		// A tag a form of KEY=VALUE,KEY=VALUE couldn't carry
		"Value containing a comma and significant spaces": {
			setting: `{"Owner":" Last, First "}`,
			expected: map[string]string{
				"Owner": " Last, First ",
			},
			valid: true,
		},
		"Empty value": {
			setting: `{"Owner":""}`,
			expected: map[string]string{
				"Owner": "",
			},
			valid: true,
		},
		"Not an object": {
			setting: `Owner=someone`,
			valid:   false,
		},
		"Null": {
			setting: `null`,
			valid:   false,
		},
		"Values that aren't strings": {
			setting: `{"Owner":1}`,
			valid:   false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			g := NewGomegaWithT(t)
			t.Setenv(TestResourceGroupTagsVar, c.setting)

			tags, err := ResourceGroupTagsFromEnvironment()
			if !c.valid {
				g.Expect(err).To(HaveOccurred())
				return
			}

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(tags).To(Equal(c.expected))
		})
	}
}
