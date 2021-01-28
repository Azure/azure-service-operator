/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package armclient

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/devigned/tab"
)

type Client struct {
	autorest.Client
	Host string
}

const UserAgent = "k8sinfra-generated"

// NewClient creates a new raw client
func NewClient(authorizer autorest.Authorizer) *Client {

	autorestClient := autorest.NewClientWithUserAgent(UserAgent)
	// Disable retries by default
	autorestClient.RetryAttempts = 0
	autorestClient.Authorizer = authorizer

	c := &Client{
		Client: autorestClient,
		Host:   azure.PublicCloud.ResourceManagerEndpoint, // TODO: We need to support other endpoints
	}

	return c
}

// WithExponentialRetries creates a new client with exponential retries configured and returns it
func (c *Client) WithExponentialRetries(attempts int, backoff time.Duration, maxBackoff time.Duration) *Client {
	// Copy the client
	result := *c
	result.SendDecorators = nil
	// Deep copy the send decorators
	result.SendDecorators = append(result.SendDecorators, c.SendDecorators...)

	// There's no place to set a backoff cap on the actual client?
	result.RetryAttempts = attempts
	result.RetryDuration = backoff

	result.SendDecorators = append(
		result.SendDecorators,
		autorest.DoRetryForStatusCodesWithCap(
			result.RetryAttempts,
			result.RetryDuration,
			maxBackoff,
			autorest.StatusCodesForRetry...))

	return &result
}

// PutDeployment creates or updates a deployment in Azure, and updates the given Deployment
// with the current deployment state.
func (c *Client) PutDeployment(ctx context.Context, deployment *Deployment) error {
	entityPath, err := deployment.GetEntityPath()
	if err != nil {
		return err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json"),
		autorest.WithJSON(deployment))

	req, err := c.newRequest(ctx, http.MethodPut, entityPath)
	if err != nil {
		tab.For(ctx).Error(err)
		return err
	}

	req, err = preparer.Prepare(req)
	if err != nil {
		tab.For(ctx).Error(err)
		return err
	}

	// The linter below doesn't realize that the response is closed in the course of
	// the autorest.Respond call below, suppressing the false positive.
	// nolint:bodyclose
	resp, err := c.Send(req)

	if err != nil {
		tab.For(ctx).Error(err)
		return err
	}

	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(deployment),
		autorest.ByClosing())
	if err != nil {
		tab.For(ctx).Error(err)
		return err
	}

	return nil
}

var zeroDuration time.Duration = 0

func (c *Client) GetResource(ctx context.Context, resourceID string, resource interface{}) (time.Duration, error) {

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json"))

	req, err := c.newRequest(ctx, http.MethodGet, resourceID)
	if err != nil {
		return zeroDuration, err
	}

	req, err = preparer.Prepare(req)
	if err != nil {
		tab.For(ctx).Error(err)
		return zeroDuration, err
	}

	// The linter below doesn't realize that the response is closed in the course of
	// the autorest.Respond call below, suppressing the false positive.
	// nolint:bodyclose
	resp, err := c.Send(req)
	if err != nil {
		tab.For(ctx).Error(err)
		return zeroDuration, err
	}

	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(resource),
		autorest.ByClosing())

	retryAfter := getRetryAfter(resp)

	if err != nil {
		tab.For(ctx).Error(err)
		return retryAfter, err
	}

	return retryAfter, nil
}

func getRetryAfter(resp *http.Response) time.Duration {
	if retryAfterStr := resp.Header.Get("Retry-After"); retryAfterStr != "" {
		if retryAfterVal, parseErr := strconv.ParseInt(retryAfterStr, 10, 64); parseErr == nil {
			return time.Duration(retryAfterVal) * time.Second
		}

		if retryAfterTime, parseErr := parseHttpDate(retryAfterStr); parseErr == nil {
			result := time.Until(retryAfterTime)
			if result > 0 {
				return result
			}
		}
	}

	return 0
}

func parseHttpDate(s string) (time.Time, error) {
	if t, err := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", s); err == nil {
		return t, nil
	} else if t, err = time.Parse("Monday, 02-Jan-06 15:04:05 MST", s); err == nil {
		return t, nil
	} else if t, err = time.Parse("Mon Jan  2 15:04:05 2006", s); err == nil {
		return t, nil
	}

	return time.Time{}, errors.New("unable to parse date")
}

// DeleteResource will make an HTTP DELETE call to the resourceId and attempt to fill the resource with the response.
// If the body of the response is empty, the resource will be nil.
func (c *Client) DeleteResource(ctx context.Context, resourceID string, resource interface{}) (time.Duration, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json"))

	req, err := c.newRequest(ctx, http.MethodDelete, resourceID)
	if err != nil {
		return zeroDuration, err
	}

	req, err = preparer.Prepare(req)
	if err != nil {
		tab.For(ctx).Error(err)
		return zeroDuration, err
	}

	// The linter below doesn't realize that the response is closed in the course of
	// the autorest.Respond call below, suppressing the false positive.
	// nolint:bodyclose
	resp, err := c.Send(req)

	if err != nil {
		tab.For(ctx).Error(err)
		return zeroDuration, err
	}

	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(resource),
		autorest.ByClosing())

	retryAfter := getRetryAfter(resp)

	if err != nil {
		if IsNotFound(err) {
			// you asked it to be gone, well, it is.
			return zeroDuration, nil /* no need to retry */
		}

		tab.For(ctx).Error(err)
		return retryAfter, err
	}

	return retryAfter, nil
}

func (c *Client) newRequest(ctx context.Context, method string, entityPath string) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, c.Host+strings.TrimPrefix(entityPath, "/"), nil)
}

func IsNotFound(err error) bool {
	var typedError *azure.RequestError
	if errors.As(err, &typedError) {
		if typedError.Response != nil && typedError.Response.StatusCode == 404 {
			return true
		}
	}

	return false
}
