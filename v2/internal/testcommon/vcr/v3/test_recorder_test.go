/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v3

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
)

func TestRecorder_WhenRecordingAndRecordingDoesNotExist_MakesRecording(t *testing.T) {
	// NB: Can't run tests using Setenv() in parallel
	t.Setenv("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	t.Setenv("AZURE_TENANT_ID", "00000000-0000-0000-0000-000000000000")

	g := NewGomegaWithT(t)

	cfg := config.Values{}
	cassetteName := "recordings/" + t.Name()

	// Test prerequisite: The recording must not already exist
	// We delete the recording at the end of the test, so it shouldn't ever be committed
	// But if something goes awry, and we don't clean up, we need to flag the presence of
	// this file as a failure.
	// We're noisy about it (instead of just deleting the file proactively) in order to
	// ensure the dev knows the file shouldn't be committed.
	exists, err := vcr.CassetteFileExists(cassetteName)
	g.Expect(err).To(BeNil())
	g.Expect(exists).To(BeFalse())

	// Ensure we clean up the cassette file at the end of the test
	cassetteFile := vcr.CassetteFileName(cassetteName)
	defer func() {
		g.Expect(os.Remove(cassetteFile)).To(Succeed())
	}()

	// Create our TestRecorder and ensure it's recording
	recorder, err := NewTestRecorder(cassetteName, cfg, logr.Discard())
	g.Expect(err).To(BeNil())
	g.Expect(recorder.IsReplaying()).To(BeFalse())

	// Stand up a fake HTTP server:
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Hello World"))
	}))

	client := recorder.CreateClient(t)

	// Make sure we can get a response from the internet
	//nolint:noctx
	resp, err := client.Get(server.URL)
	g.Expect(err).To(BeNil())
	defer resp.Body.Close()

	// Ensure the HTTP response is as expected
	g.Expect(resp.StatusCode).To(Equal(http.StatusOK))

	// Ensure the body is not empty
	body, err := io.ReadAll(resp.Body)
	g.Expect(err).To(BeNil())
	g.Expect(body).NotTo(HaveLen(0))

	// Stop the recorder
	err = recorder.Stop()
	g.Expect(err).To(BeNil())

	// Verify we created a recording
	exists, err = vcr.CassetteFileExists(cassetteName)
	g.Expect(err).To(BeNil())
	g.Expect(exists).To(BeTrue())
}

func TestRecorder_WhenRecordingAndRecordingExists_DoesPlayback(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	//
	// Rerecording this test can be a challenge, due to
	// (a) a prerequisite check ensuring the cassette exists
	// (b) a similar check ensuring we're in replay mode, and
	// (c) the requirement for selected environment variables to be present.
	//
	// The easiest way to rerecord this test is to:
	//   (i) Delete the cassette file
	//  (ii) Comment out the prerequisite checks
	// (iii) Define the required environment variables
	//  (iv) go test -run TestRecorderV3_WhenRecordingAndRecordingExists_DoesPlayback
	//

	cfg := config.Values{}
	cassetteName := "recordings/" + t.Name()

	// Test prerequisite: The recording must already exist
	exists, err := vcr.CassetteFileExists(cassetteName)
	g.Expect(err).To(BeNil())
	g.Expect(exists).To(BeTrue())

	// Create our TestRecorder and ensure it's recording
	recorder, err := NewTestRecorder(cassetteName, cfg, logr.Discard())
	g.Expect(err).To(BeNil())
	g.Expect(recorder.IsReplaying()).To(BeTrue())

	url := "https://www.bing.com"
	client := recorder.CreateClient(t)

	// Make sure we can get a response from the internet
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	g.Expect(err).To(BeNil())

	resp, err := client.Do(req)
	g.Expect(err).To(BeNil())
	defer resp.Body.Close()

	// Ensure the body is not empty
	body, err := io.ReadAll(resp.Body)
	g.Expect(err).To(BeNil())
	g.Expect(body).NotTo(HaveLen(0))

	// Stop the recorder
	err = recorder.Stop()
	g.Expect(err).To(BeNil())
}
