package config

import (
	"testing"

	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
)

func TestFieldMatcher_DeserializedFromYaml_GivesExpectedMatchResult(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name          string
		yaml          string
		value string
		shouldMatch bool
	}{
		{"Literal", "field: foo", "foo", true},
		{"Literal, case insensitive", "field: foo", "Foo", true},
		{"Literal, no match", "field: foo", "Bar", false},
		{"Wildcard field", "field: f*", "Foo", true},
		{"Wildcard field, no match", "field: f*", "Bar", false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			var h matcherHost
			g.Expect( yaml.Unmarshal([]byte(c.yaml), &h)).To(Succeed())
			g.Expect(h.Field.Matches(c.value)).To(Equal(c.shouldMatch))
		})
	}
}

type matcherHost struct {
	Field FieldMatcher `yaml:'field'`
}
