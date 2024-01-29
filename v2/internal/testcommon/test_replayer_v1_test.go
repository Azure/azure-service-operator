/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"io"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/config"
)

func TestReplayerV1_WhenRecordingExists_ReturnsResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := config.Values{}
	cassetteName := "recordings/" + t.Name()
	replayer, err := newTestPlayerV1(cassetteName, cfg)
	g.Expect(err).To(BeNil())

	defer replayer.Stop()

	url := "https://www.bing.com"
	client := replayer.CreateClient(t)

	resp, err := client.Get(url)
	g.Expect(err).To(BeNil())

	body, err := io.ReadAll(resp.Body)
	g.Expect(err).To(BeNil())
	g.Expect(body).NotTo(HaveLen(0))
}