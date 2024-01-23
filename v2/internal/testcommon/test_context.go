/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	cassettev1 "github.com/dnaeon/go-vcr/cassette"
	recorderv1 "github.com/dnaeon/go-vcr/recorder"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/metrics"
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
	AzureClientRecorder   testRecorder
	AzureClient           *genericarmclient.GenericClient
	AzureSubscription     string
	AzureTenant           string
	AzureBillingInvoiceID string
	AzureMatch            *ARMMatcher
	HttpClient            *http.Client
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

const DummyBillingId = "/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/0000-0000-000-000/invoiceSections/0000-0000-000-000"

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
	details, err := createRecorder(cassetteName, cfg, tc.RecordReplay)
	if err != nil {
		return PerTestContext{}, errors.Wrapf(err, "creating recorder")
	}

	// Use the recorder-specific CFG, which will force URLs and AADAuthorityHost (among other things) to default
	// values so that the recordings look the same regardless of which cloud you ran them in
	cfg = details.Cfg()

	// To Go SDK client reuses HTTP clients among instances by default. We add handlers to the HTTP client based on
	// the specific test in question, which means that clients cannot be reused.
	// We explicitly create a new http.Client so that the recording from one test doesn't
	// get used for all other parallel tests.
	httpClient := &http.Client{
		Transport: details.CreateRoundTripper(t),
	}

	var globalARMClient *genericarmclient.GenericClient
	options := &genericarmclient.GenericClientOptions{
		Metrics:    metrics.NewARMClientMetrics(),
		HttpClient: httpClient,
	}
	globalARMClient, err = genericarmclient.NewGenericClient(cfg.Cloud(), details.Creds(), options)
	if err != nil {
		return PerTestContext{}, errors.Wrapf(err, "failed to create generic ARM client")
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
				err = ensureCassetteFileExists(cassetteName)
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
		AzureSubscription:     details.Ids().subscriptionID,
		AzureTenant:           details.Ids().tenantID,
		AzureBillingInvoiceID: details.Ids().billingInvoiceID,
		AzureMatch:            NewARMMatcher(globalARMClient),
		AzureClientRecorder:   details,
		HttpClient:            httpClient,
		TestName:              t.Name(),
		Namespace:             createTestNamespaceName(t),
		Ctx:                   context,
	}, nil
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

func ensureCassetteFileExists(cassetteName string) error {
	filename := cassetteName + ".yaml"
	_, err := os.Stat(filename)
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return err
	}
	f, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0o644)
	if err != nil {
		return errors.Wrapf(err, "creating empty cassette %q", filename)
	}
	if err := f.Close(); err != nil {
		return errors.Wrapf(err, "failed to close empty cassette %q", filename)
	}
	return nil
}

type recorderDetails struct {
	cassetteName string
	creds        azcore.TokenCredential
	ids          AzureIDs
	recorder     *recorderv1.Recorder
	cfg          config.Values
}

// Cfg returns the available configuration for the test
func (r recorderDetails) Cfg() config.Values {
	return r.cfg
}

// Creds returns Azure credentials when running for real
func (r recorderDetails) Creds() azcore.TokenCredential {
	return r.creds
}

// Ids returns the available Azure resource IDs for the test
func (r recorderDetails) Ids() AzureIDs {
	return r.ids
}

// Stop recording
func (r recorderDetails) Stop() error {
	return r.recorder.Stop()
}

// IsReplaying returns true if we're replaying a recorded test, false if we're recording a new test
func (r recorderDetails) IsReplaying() bool {
	return r.recorder.Mode() == recorderv1.ModeReplaying
}

// CreateRoundTripper creates a client RoundTripper that can be used to record or replay HTTP requests.
// t is a reference to the test currently executing.
// TODO: Remove the reference to t to reduce coupling
func (r recorderDetails) CreateRoundTripper(t *testing.T) http.RoundTripper {
	return addCountHeader(translateErrorsV1(r.recorder, r.cassetteName, t))
}

func createRecorder(cassetteName string, cfg config.Values, recordReplay bool) (testRecorder, error) {
	var err error
	var r *recorderv1.Recorder
	if recordReplay {
		r, err = recorderv1.New(cassetteName)
	} else {
		r, err = recorderv1.NewAsMode(cassetteName, recorderv1.ModeDisabled, nil)
	}

	if err != nil {
		return recorderDetails{}, errors.Wrapf(err, "creating recorder")
	}

	var creds azcore.TokenCredential
	var azureIDs AzureIDs
	if r.Mode() == recorderv1.ModeRecording ||
		r.Mode() == recorderv1.ModeDisabled {
		// if we are recording, we need auth
		creds, azureIDs, err = getCreds()
		if err != nil {
			return recorderDetails{}, err
		}
	} else {
		// if we are replaying, we won't need auth
		// and we use a dummy subscription ID/tenant ID
		creds = MockTokenCredential{}
		azureIDs.tenantID = uuid.Nil.String()
		azureIDs.subscriptionID = uuid.Nil.String()
		azureIDs.billingInvoiceID = DummyBillingId
		// Force these values to be the default
		cfg.ResourceManagerEndpoint = config.DefaultEndpoint
		cfg.ResourceManagerAudience = config.DefaultAudience
		cfg.AzureAuthorityHost = config.DefaultAADAuthorityHost
	}

	// check body as well as URL/Method (copied from go-vcr documentation)
	r.SetMatcher(func(r *http.Request, i cassettev1.Request) bool {
		if !cassettev1.DefaultMatcher(r, i) {
			return false
		}

		// verify custom request count header (see counting_roundtripper.go)
		if r.Header.Get(COUNT_HEADER) != i.Headers.Get(COUNT_HEADER) {
			return false
		}

		if r.Body == nil {
			return i.Body == ""
		}

		var b bytes.Buffer
		if _, err := b.ReadFrom(r.Body); err != nil {
			panic(err)
		}

		r.Body = io.NopCloser(&b)
		return b.String() == "" || hideRecordingData(b.String()) == i.Body
	})

	r.AddSaveFilter(func(i *cassettev1.Interaction) error {
		// rewrite all request/response fields to hide the real subscription ID
		// this is *not* a security measure but intended to make the tests updateable from
		// any subscription, so a contributor can update the tests against their own sub.
		hide := func(s string, id string, replacement string) string {
			return strings.ReplaceAll(s, id, replacement)
		}

		// Note that this changes the cassette in-place so there's no return needed
		hideCassetteString := func(cas *cassettev1.Interaction, id string, replacement string) {
			i.Request.Body = strings.ReplaceAll(cas.Request.Body, id, replacement)
			i.Response.Body = strings.ReplaceAll(cas.Response.Body, id, replacement)
			i.Request.URL = strings.ReplaceAll(cas.Request.URL, id, replacement)
		}

		// Hide the subscription ID
		hideCassetteString(i, azureIDs.subscriptionID, uuid.Nil.String())
		// Hide the tenant ID
		hideCassetteString(i, azureIDs.tenantID, uuid.Nil.String())
		// Hide the billing ID
		if azureIDs.billingInvoiceID != "" {
			hideCassetteString(i, azureIDs.billingInvoiceID, DummyBillingId)
		}

		// Hiding other sensitive fields
		i.Request.Body = hideRecordingData(i.Request.Body)
		i.Response.Body = hideRecordingData(i.Response.Body)
		i.Request.URL = hideURLData(i.Request.URL)

		for _, values := range i.Request.Headers {
			for i := range values {
				values[i] = hide(values[i], azureIDs.subscriptionID, uuid.Nil.String())
				values[i] = hide(values[i], azureIDs.tenantID, uuid.Nil.String())
				if azureIDs.billingInvoiceID != "" {
					values[i] = hide(values[i], azureIDs.billingInvoiceID, DummyBillingId)
				}
			}
		}

		for key, values := range i.Response.Headers {
			for i := range values {
				values[i] = hide(values[i], azureIDs.subscriptionID, uuid.Nil.String())
				values[i] = hide(values[i], azureIDs.tenantID, uuid.Nil.String())
				if azureIDs.billingInvoiceID != "" {
					values[i] = hide(values[i], azureIDs.billingInvoiceID, DummyBillingId)
				}
			}
			// Hide the base request URL in the AzureOperation and Location headers
			if key == genericarmclient.AsyncOperationHeader || key == genericarmclient.LocationHeader {
				for i := range values {
					values[i] = hideBaseRequestURL(values[i])
				}
			}
		}

		for _, header := range requestHeadersToRemove {
			delete(i.Request.Headers, header)
		}

		for _, header := range responseHeadersToRemove {
			delete(i.Response.Headers, header)
		}

		return nil
	})

	return recorderDetails{
		cassetteName: cassetteName,
		creds:        creds,
		ids:          azureIDs,
		recorder:     r,
		cfg:          cfg,
	}, nil
}

var requestHeadersToRemove = []string{
	// remove all Authorization headers from stored requests
	"Authorization",

	// Not needed, adds to diff churn:
	"User-Agent",
}

var responseHeadersToRemove = []string{
	// Request IDs
	"X-Ms-Arm-Service-Request-Id",
	"X-Ms-Correlation-Request-Id",
	"X-Ms-Request-Id",
	"X-Ms-Routing-Request-Id",
	"X-Ms-Client-Request-Id",
	"Client-Request-Id",

	// Quota limits
	"X-Ms-Ratelimit-Remaining-Subscription-Deletes",
	"X-Ms-Ratelimit-Remaining-Subscription-Reads",
	"X-Ms-Ratelimit-Remaining-Subscription-Writes",

	// Not needed, adds to diff churn
	"Date",
}

var (
	dateMatcher     = regexp.MustCompile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(.\d+)?Z`)
	sshKeyMatcher   = regexp.MustCompile("ssh-rsa [0-9a-zA-Z+/=]+")
	passwordMatcher = regexp.MustCompile("\"pass[^\"]*?pass\"")

	// keyMatcher matches any valid base64 value with at least 10 sets of 4 bytes of data that ends in = or ==.
	// Both storage account keys and Redis account keys are longer than that and end in = or ==. Note that technically
	// base64 values need not end in == or =, but allowing for that in the match will flag tons of false positives as
	// any text (including long URLs) have strings of characters that meet this requirement. There are other base64 values
	// in the payloads (such as operationResults URLs for polling async operations for some services) that seem to use
	// very long base64 strings as well.
	keyMatcher = regexp.MustCompile("(?:[A-Za-z0-9+/]{4}){10,}(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)")

	// kubeConfigMatcher specifically matches base64 data returned by the AKS get keys API
	kubeConfigMatcher = regexp.MustCompile(`"value": "[a-zA-Z0-9+/]+={0,2}"`)

	// baseURLMatcher matches the base part of a URL
	baseURLMatcher = regexp.MustCompile(`^https://[^/]+/`)

	// customKeyMatcher is used to match 'key' or 'Key' followed by the base64 patterns without '=' padding.
	customKeyMatcher  = regexp.MustCompile(`"([a-z]+)?[K-k]ey":"[a-zA-Z0-9+/]+"`)
	customKeyReplacer = regexp.MustCompile(`"(?:[A-Za-z0-9+/]{4}){10,}(?:[A-Za-z0-9+/]{4}|[A-Za-z0-9+/])"`)
)

// hideDates replaces all ISO8601 datetimes with a fixed value
// this lets us match requests that may contain time-sensitive information (timestamps, etc)
func hideDates(s string) string {
	return dateMatcher.ReplaceAllLiteralString(s, "2001-02-03T04:05:06Z") // this should be recognizable/parseable as a fake date
}

// hideSSHKeys hides anything that looks like SSH keys
func hideSSHKeys(s string) string {
	return sshKeyMatcher.ReplaceAllLiteralString(s, "ssh-rsa {KEY}")
}

// hidePasswords hides anything that looks like a generated password
func hidePasswords(s string) string {
	return passwordMatcher.ReplaceAllLiteralString(s, "\"{PASSWORD}\"")
}

func hideKeys(s string) string {
	return keyMatcher.ReplaceAllLiteralString(s, "{KEY}")
}

func hideKubeConfigs(s string) string {
	return kubeConfigMatcher.ReplaceAllLiteralString(s, `"value": "IA=="`) // Have to replace with valid base64 data, so replace with " "
}

func hideBaseRequestURL(s string) string {
	return baseURLMatcher.ReplaceAllLiteralString(s, `https://management.azure.com/`)
}

func hideCustomKeys(s string) string {
	return customKeyMatcher.ReplaceAllStringFunc(s, func(matched string) string {
		return customKeyReplacer.ReplaceAllString(matched, `"{KEY}"`)
	})
}

func hideRecordingData(s string) string {
	result := hideDates(s)
	result = hideSSHKeys(result)
	result = hidePasswords(result)
	result = hideKubeConfigs(result)
	result = hideKeys(result)
	result = hideCustomKeys(result)

	return result
}

func hideURLData(s string) string {
	return hideBaseRequestURL(s)
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
