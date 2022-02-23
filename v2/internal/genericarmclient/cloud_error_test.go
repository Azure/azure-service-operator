/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestCloudError_WhenLoadedFromJSON_ReturnsExpectedResults(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Name            string
		JSON            string
		ExpectedCode    string
		ExpectedMessage string
		ExpectedTarget  string
	}{
		{
			Name:            "ARM style",
			JSON:            `{"error": { "code": "broken", "message": "It's dead, Jim", "target": "Janus VI" } }`,
			ExpectedCode:    "broken",
			ExpectedMessage: "It's dead, Jim",
			ExpectedTarget:  "Janus VI",
		},
		{
			Name:            "ARM style, code missing",
			JSON:            `{"error": { "message": "It's dead, Jim", "target": "Janus VI" } }`,
			ExpectedCode:    "UnknownError",
			ExpectedMessage: "It's dead, Jim",
			ExpectedTarget:  "Janus VI",
		},
		{
			Name:            "Simple style",
			JSON:            `{ "code": "broken", "message": "It's dead, Jim", "target": "Janus VI" }`,
			ExpectedCode:    "broken",
			ExpectedMessage: "It's dead, Jim",
			ExpectedTarget:  "Janus VI",
		},
		{
			Name:            "Simple style, code missing",
			JSON:            `{ "message": "It's dead, Jim", "target": "Janus VI" }`,
			ExpectedCode:    "UnknownError",
			ExpectedMessage: "It's dead, Jim",
			ExpectedTarget:  "Janus VI",
		},
		{
			Name:            "CosmosDB Actual",
			JSON:            `{"code":"BadRequest","message":"The requested operation cannot be performed because the database account asotestdbiosaow is in the process of being created."}`,
			ExpectedCode:    "BadRequest",
			ExpectedMessage: "The requested operation cannot be performed because the database account asotestdbiosaow is in the process of being created.",
			ExpectedTarget:  "",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			var cloudError CloudError
			g := NewGomegaWithT(t)
			g.Expect(json.Unmarshal([]byte(c.JSON), &cloudError)).To(Succeed())

			g.Expect(cloudError.Code()).To(Equal(c.ExpectedCode))
			g.Expect(cloudError.Message()).To(Equal(c.ExpectedMessage))
			g.Expect(cloudError.Target()).To(Equal(c.ExpectedTarget))
		})
	}
}
