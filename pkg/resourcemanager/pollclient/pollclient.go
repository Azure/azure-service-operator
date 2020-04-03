// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package pollclient

import (
	"context"
	"net/http"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
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
func NewPollClient() PollClient {
	return NewPollClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
}

// NewPollClientWithBaseURI returns a paramterized client
func NewPollClientWithBaseURI(baseURI string, subscriptionID string) PollClient {
	c := PollClient{NewWithBaseURI(baseURI, subscriptionID)}
	a, _ := iam.GetResourceManagementAuthorizer()
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

// PollRespons models the expected response from the poll url
type PollRespons struct {
	autorest.Response `json:"-"`
	Name              string             `json:"name,omitempty"`
	Status            string             `json:"status,omitempty"`
	Error             azure.ServiceError `json:"error,omitempty"`
}

// Get takes a context and a polling url and performs a Get request on the url
func (client PollClient) Get(ctx context.Context, pollURL string) (result PollRespons, err error) {
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
func (client PollClient) GetResponder(resp *http.Response) (result PollRespons, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
