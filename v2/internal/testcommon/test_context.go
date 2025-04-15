/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"golang.org/x/crypto/ssh"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
)

// Use WestUS2 as some things (such as VM quota) are hard to get in West US.
var DefaultTestRegion = "westus2" // Could make this an env variable if we wanted

type TestContext struct {
	AzureRegion  *string
	NameConfig   *ResourceNameConfig
	RecordReplay bool
}

type PerTestContext struct {
	TestContext
	T                     *testing.T
	logger                logr.Logger
	AzureClientRecorder   vcr.Interface
	AzureClient           *genericarmclient.GenericClient
	AzureSubscription     string
	AzureTenant           string
	AzureBillingInvoiceID string
	AzureMatch            *ARMMatcher
	HTTPClient            *http.Client
	Namer                 ResourceNamer
	NoSpaceNamer          ResourceNamer
	TestName              string
	Namespace             string
	Ctx                   context.Context
	// CountsTowardsParallelLimits true means that the envtest (if any) started for this test pass counts towards the limit of
	// concurrent envtests running at once. If this is false, it doesn't count towards the limit.
	CountsTowardsParallelLimits bool
}

// There are two prefixes here because each represents a resource kind with a distinct lifecycle.
// If you modify these, make sure to modify the cleanup-azure-resources target in the Taskfile too

// ResourcePrefix is for resources which are used in the record/replay test infrastructure using envtest.
// These tests are not expected to be run live in parallel. In parallel runs should be done from the recordings.
// As such, for cleanup, we can delete any resources
// with this prefix without fear of disrupting an existing test pass. Again, this is ok because there's a single Github Action
// that runs with MaxParallelism == 1.
const ResourcePrefix = "asotest"

// LiveResourcePrefix is the prefix reserved for resources used in tests that cannot be recorded. These resources must support
// parallel runs (can be triggered by multiple parallel PRs), so deleting existing resources any time a new run starts is not allowed.
// Instead, deletion should be time-based - delete any leaked resources older than X hours. These are generally run against
// either a real cluster or a kind cluster.
const LiveResourcePrefix = "asolivetest"

func NewTestContext(
	region string,
	recordReplay bool,
	nameConfig *ResourceNameConfig,
) TestContext {
	return TestContext{
		AzureRegion:  &region,
		RecordReplay: recordReplay,
		NameConfig:   nameConfig,
	}
}

func (tc TestContext) ForTest(t *testing.T, cfg config.Values) (PerTestContext, error) {
	logger := NewTestLogger(t)

	cassetteName := "recordings/" + t.Name()
	details, err := createTestRecorder(cassetteName, cfg, tc.RecordReplay, logger)
	if err != nil {
		return PerTestContext{}, eris.Wrapf(err, "creating recorder")
	}

	// Use the recorder-specific CFG, which will force URLs and AADAuthorityHost (among other things) to default
	// values so that the recordings look the same regardless of which cloud you ran them in
	cfg = details.Cfg()

	// To Go SDK client reuses HTTP clients among instances by default. We add handlers to the HTTP client based on
	// the specific test in question, which means that clients cannot be reused.
	// We explicitly create a new http.Client so that the recording from one test doesn't
	// get used for all other parallel tests.
	httpClient := details.CreateClient(t)

	var globalARMClient *genericarmclient.GenericClient
	options := &genericarmclient.GenericClientOptions{
		Metrics:    metrics.NewARMClientMetrics(),
		HTTPClient: httpClient,
	}
	globalARMClient, err = genericarmclient.NewGenericClient(cfg.Cloud(), details.Creds(), options)
	if err != nil {
		return PerTestContext{}, eris.Wrapf(err, "failed to create generic ARM client")
	}

	t.Cleanup(func() {
		if !t.Failed() {
			err := details.Stop()
			if err != nil {
				// cleanup function should not error-out
				logger.Error(err, "unable to stop ARM client recorder")
				t.Fail()
			}
			if tc.RecordReplay {
				logger.Info("saving ARM client recorder")
				// If the cassette file doesn't exist at this point
				// then the operator must not have made any ARM
				// requests. Create an empty cassette to record that
				// so subsequent replay tests can run without needing
				// credentials.
				err = vcr.EnsureCassetteFileExists(cassetteName)
				if err != nil {
					logger.Error(err, "ensuring cassette file exists")
					t.Fail()
				}
			}
		}
	})

	namer := tc.NameConfig.NewResourceNamer(t.Name())

	context := context.Background() // we could consider using context.WithTimeout(OperationTimeout()) here

	return PerTestContext{
		TestContext:           tc,
		T:                     t,
		logger:                logger,
		Namer:                 namer,
		NoSpaceNamer:          namer.WithSeparator(""),
		AzureClient:           globalARMClient,
		AzureSubscription:     details.IDs().SubscriptionID,
		AzureTenant:           details.IDs().TenantID,
		AzureBillingInvoiceID: details.IDs().BillingInvoiceID,
		AzureMatch:            NewARMMatcher(globalARMClient),
		AzureClientRecorder:   details,
		HTTPClient:            httpClient,
		TestName:              t.Name(),
		Namespace:             createTestNamespaceName(t),
		Ctx:                   context,
	}, nil
}

// WithLiteralRedaction method is used to add literal redaction values based on the requirements of each test.
// The method takes in redactValue to be `redacted` and replacement value to which the value should be `replaced`.
func (tc PerTestContext) WithLiteralRedaction(redactionValue string, replacementValue string) {
	tc.AzureClientRecorder.AddLiteralRedaction(redactionValue, replacementValue)
}

// WithRegexpRedaction method is used to add regexp redaction values based on the requirements of each test.
// The method takes in regexp pattern to be `redacted` and replacement value to which the value should be `replaced`.
func (tc PerTestContext) WithRegexpRedaction(pattern string, replacementValue string) {
	tc.AzureClientRecorder.AddRegexpRedaction(pattern, replacementValue)
}

var replaceRegex = regexp.MustCompile("[./_]+")

func createTestNamespaceName(t *testing.T) string {
	const maxLen = 63
	const disambigLength = 5

	result := "aso-" + replaceRegex.ReplaceAllString(strings.ToLower(t.Name()), "-")
	if len(result) <= maxLen {
		return result
	}

	// we need to trim it but also add disambiguation; append a shortened hash
	hash := sha256.Sum256([]byte(result))
	hashStr := hex.EncodeToString(hash[:])
	disambig := hashStr[0:disambigLength]
	result = result[0:(maxLen - disambigLength - 1 /* hyphen */)]
	return result + "-" + disambig
}

func (tc PerTestContext) NewTestResourceGroup() *resources.ResourceGroup {
	return &resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name: tc.Namer.GenerateName("rg"),
		},
		Spec: resources.ResourceGroup_Spec{
			Location: tc.AzureRegion,
			Tags:     CreateTestResourceGroupDefaultTags(),
		},
	}
}

// GenerateSSHKey generates an SSH key.
func (tc PerTestContext) GenerateSSHKey(size int) (*string, error) {
	// Note: If we ever want to make sure that the SSH keys are the same between
	// test runs, we can base it off of a hash of subscription ID. Right now since
	// we just replace the SSH key in the recordings regardless of what the value is
	// there's no need for uniformity between runs though.

	key, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return nil, err
	}

	err = key.Validate()
	if err != nil {
		return nil, err
	}

	sshPublicKey, err := ssh.NewPublicKey(&key.PublicKey)
	if err != nil {
		return nil, err
	}

	resultBytes := ssh.MarshalAuthorizedKey(sshPublicKey)
	result := string(resultBytes)

	return &result, nil
}
