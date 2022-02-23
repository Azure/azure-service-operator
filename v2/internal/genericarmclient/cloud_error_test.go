package genericarmclient

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestCloudError_WhenLoadedFromJSON_ReturnsExpectedResults(t *testing.T) {

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
			Name:            "Simple style",
			JSON:            `{ "code": "broken", "message": "It's dead, Jim", "target": "Janus VI" }`,
			ExpectedCode:    "broken",
			ExpectedMessage: "It's dead, Jim",
			ExpectedTarget:  "Janus VI",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			var cloudError CloudError
			g := NewGomegaWithT(t)
			g.Expect(json.Unmarshal([]byte(c.JSON), &cloudError)).To(Succeed())

			g.Expect(cloudError.ErrorCode()).To(Equal(c.ExpectedCode))
			g.Expect(cloudError.ErrorMessage()).To(Equal(c.ExpectedMessage))
			g.Expect(cloudError.ErrorTarget()).To(Equal(c.ExpectedTarget))
		})
	}
}
