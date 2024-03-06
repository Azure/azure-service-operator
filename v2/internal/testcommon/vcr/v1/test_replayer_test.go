/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	"context"
	"io"
	"net/http"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/config"
)

func TestReplayer_WhenRecordingExists_ReturnsResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := config.Values{}
	cassetteName := "recordings/" + t.Name()
	replayer, err := NewTestPlayer(cassetteName, cfg)
	g.Expect(err).To(BeNil())

	url := "https://www.bing.com"
	client := replayer.CreateClient(t)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	g.Expect(err).To(BeNil())

	resp, err := client.Do(req)
	g.Expect(err).To(BeNil())
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	g.Expect(err).To(BeNil())
	g.Expect(body).NotTo(HaveLen(0))

	g.Expect(replayer.Stop()).To(Succeed())
}
