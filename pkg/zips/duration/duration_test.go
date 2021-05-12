/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package duration

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/onsi/gomega"
)

type (
	DurationTest struct {
		ADuration *ISO8601 `json:"aduration,omitempty"`
		Name      string   `json:"name,omitempty"`
	}
)

func TestDuration_String(t *testing.T) {
	cases := []struct {
		Duration time.Duration
		Expected string
	}{
		{
			Duration: 30 * time.Second,
			Expected: "PT30S",
		},
		{
			Duration: 30*time.Second + 4*time.Millisecond,
			Expected: "PT30.004S",
		},
		{
			Duration: 987654*time.Minute + 30*time.Second + 4*time.Millisecond,
			Expected: "P1Y10M2W6DT20H54M30.004S",
		},
		{
			Duration: time.Hour * 24 * 7,
			Expected: "P1W",
		},
		{
			Duration: time.Hour * 24 * 2,
			Expected: "P2D",
		},
		{
			Duration: time.Hour*24*2 + time.Hour,
			Expected: "P2DT1H",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Expected, func(t *testing.T) {
			d := ISO8601{Duration: c.Duration}
			g := gomega.NewGomegaWithT(t)
			g.Expect(d.String()).To(gomega.Equal(c.Expected))
		})
	}
}

func TestDuration_JSON(t *testing.T) {
	cases := []struct {
		Name   string
		JSON   string
		Struct DurationTest
	}{
		{
			Name: "duration specified",
			JSON: `{"aduration":"PT30.05S","name":"foo"}`,
			Struct: DurationTest{
				ADuration: &ISO8601{
					Duration: 30050 * time.Millisecond,
				},
				Name: "foo",
			},
		},
		{
			Name: "duration not specified",
			JSON: `{"name":"foo"}`,
			Struct: DurationTest{
				ADuration: nil,
				Name:      "foo",
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			var testStruct DurationTest
			g := gomega.NewGomegaWithT(t)
			g.Expect(json.Unmarshal([]byte(c.JSON), &testStruct))
			g.Expect(testStruct).To(gomega.Equal(c.Struct))
			bits, err := json.Marshal(testStruct)
			g.Expect(err).ToNot(gomega.HaveOccurred())
			g.Expect(string(bits)).To(gomega.Equal(c.JSON))
		})
	}
}
