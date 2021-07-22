// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package pollclient

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"

	"github.com/Azure/azure-service-operator/api"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

const (
	LongRunningOperationPollStatusFailed    = "Failed"
	LongRunningOperationPollStatusSucceeded = "Succeeded"
	LongRunningOperationPollStatusCancelled = "Cancelled"
)

const fqdn = "github.com/Azure/azure-service-operator/pollingclient"

// BaseClient was modeled off some of the other Baseclients in the go sdk and contains an autorest client
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// PollClient inherits from the autorest client and has the methods needed to handle GETs to the polling url
type PollClient struct {
	BaseClient
}

// NewPollClient returns a client using hte env values from config
func NewPollClient(creds config.Credentials) PollClient {
	return NewPollClientWithBaseURI(config.BaseURI(), creds)
}

// NewPollClientWithBaseURI returns a paramterized client
func NewPollClientWithBaseURI(baseURI string, creds config.Credentials) PollClient {
	c := PollClient{NewWithBaseURI(baseURI, creds.SubscriptionID())}
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	c.Authorizer = a
	c.AddToUserAgent(config.UserAgent())
	return c
}

// NewWithBaseURI creates an instance of the BaseClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(config.UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// PollResponse models the expected response from the poll url
type PollResponse struct {
	autorest.Response `json:"-"`
	Name              string             `json:"name,omitempty"`
	Status            string             `json:"status,omitempty"`
	Error             azure.ServiceError `json:"error,omitempty"`
}

// Get takes a context and a polling url and performs a Get request on the url
func (client PollClient) Get(ctx context.Context, pollURL string) (result PollResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PollClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, pollURL)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.PollClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.PollClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.PollClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PollClient) GetPreparer(ctx context.Context, pollURL string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(pollURL))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PollClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PollClient) GetResponder(resp *http.Response) (result PollResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

type LongRunningOperationPollResult string

const (
	PollResultNoPollingNeeded       = LongRunningOperationPollResult("noPollingNeeded")
	PollResultCompletedSuccessfully = LongRunningOperationPollResult("completedSuccessfully")
	PollResultTryAgainLater         = LongRunningOperationPollResult("tryAgainLater")
	PollResultBadRequest            = LongRunningOperationPollResult("badRequest")
)

func (client PollClient) PollLongRunningOperationIfNeededV1Alpha1(ctx context.Context, status *v1alpha1.ASOStatus, kind api.PollingURLKind) (LongRunningOperationPollResult, error) {
	wrapper := v1beta1.ASOStatus(*status)
	result, err := client.PollLongRunningOperationIfNeeded(ctx, &wrapper, kind)

	// Propagate changes from wrapper to original type
	status.PollingURL = wrapper.PollingURL
	status.PollingURLKind = wrapper.PollingURLKind
	status.Message = wrapper.Message

	return result, err
}

func (client PollClient) PollLongRunningOperationIfNeededV1Alpha2(ctx context.Context, status *v1alpha2.ASOStatus, kind api.PollingURLKind) (LongRunningOperationPollResult, error) {
	wrapper := v1beta1.ASOStatus(*status)
	result, err := client.PollLongRunningOperationIfNeeded(ctx, &wrapper, kind)

	// Propagate changes from wrapper to original type
	status.PollingURL = wrapper.PollingURL
	status.PollingURLKind = wrapper.PollingURLKind
	status.Message = wrapper.Message

	return result, err
}

func (client PollClient) PollLongRunningOperationIfNeeded(ctx context.Context, status *v1beta1.ASOStatus, kind api.PollingURLKind) (LongRunningOperationPollResult, error) {
	// Before we attempt to issue a new update, check if there is a previously ongoing update
	if status.PollingURL == "" {
		return PollResultNoPollingNeeded, nil
	}

	// If there is a URL but it's the wrong kind then clear the old URL and return NoPollingNeeded
	if status.PollingURLKind != nil && *status.PollingURLKind != kind {
		status.ClearPollingURL()
		return PollResultNoPollingNeeded, nil
	}

	res, err := client.Get(ctx, status.PollingURL)
	pollErr := errhelp.NewAzureError(err)
	if pollErr != nil {
		if pollErr.Type == errhelp.OperationIdNotFound {
			// Something happened to our OperationId, just clear things out and try again
			status.ClearPollingURL()
		}
		return PollResultTryAgainLater, err
	}

	if res.Status == LongRunningOperationPollStatusFailed {
		status.Message = res.Error.Error()
		// There can be intermediate errors and various other things that cause requests to fail, so we need to try again.
		status.ClearPollingURL()

		if res.Error.Code == errhelp.BadRequest {
			return PollResultBadRequest, nil
		}

		return PollResultTryAgainLater, nil
	}

	// TODO: May need a notion of fatal error here too

	if res.Status == "InProgress" || res.Status == "Enqueued" || res.Status == "Dequeued" {
		// We're waiting for an async op... keep waiting
		return PollResultTryAgainLater, nil
	}

	// Previous operation was a success, clear polling URL and continue
	if res.Status == LongRunningOperationPollStatusSucceeded {
		status.ClearPollingURL()

		return PollResultCompletedSuccessfully, nil
	}

	// TODO: Unsure if this should be continue or tryagainlater. In the existing code it's continue
	// TODO: which is why I've made it that here
	return PollResultCompletedSuccessfully, nil
}
