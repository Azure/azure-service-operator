/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"

	"github.com/Azure/azure-service-operator/v2/internal/version"
)

// NOTE: All of these methods (and types) were adapted from
// https://github.com/Azure/azure-sdk-for-go/blob/sdk/resources/armresources/v0.3.0/sdk/resources/armresources/zz_generated_resources_client.go

type GenericClient struct {
	endpoint       string
	pl             runtime.Pipeline
	subscriptionID string
}

// TODO: Need to do retryAfter detection in each call?

// NewGenericClient creates a new instance of GenericClient
func NewGenericClient(endpoint arm.Endpoint, creds azcore.TokenCredential, subscriptionID string) *GenericClient {
	return NewGenericClientFromHTTPClient(endpoint, creds, nil, subscriptionID)
}

// NewGenericClientFromHTTPClient creates a new instance of GenericClient from the provided connection.
func NewGenericClientFromHTTPClient(endpoint arm.Endpoint, creds azcore.TokenCredential, httpClient *http.Client, subscriptionID string) *GenericClient {
	opts := &arm.ClientOptions{
		ClientOptions: policy.ClientOptions{
			Retry: policy.RetryOptions{
				MaxRetries: 0,
			},
			PerCallPolicies: []policy.Policy{NewUserAgentPolicy(userAgent)},
		},
		DisableRPRegistration: true,
	}

	// We assign this HTTPClient like this because if we actually set it to nil, due to the way
	// go interfaces wrap values, the subsequent if httpClient == nil check returns false (even though
	// the value IN the interface IS nil).
	if httpClient != nil {
		opts.Transport = httpClient
	}

	pipeline := armruntime.NewPipeline(
		"generic",
		version.BuildVersion,
		creds,
		opts)

	return &GenericClient{endpoint: string(endpoint), pl: pipeline, subscriptionID: subscriptionID}
}

// SubscriptionID returns the subscription the client is configured for
func (client *GenericClient) SubscriptionID() string {
	return client.subscriptionID
}

func (client *GenericClient) BeginCreateOrUpdateByID(ctx context.Context, resourceID string, apiVersion string, resource interface{}) (*PollerResponse, error) {
	// The linter doesn't realize that the response is closed in the course of
	// the autorest.NewPoller call below. Suppressing it as it is a false positive.
	// nolint:bodyclose
	resp, err := client.createOrUpdateByID(ctx, resourceID, apiVersion, resource)
	if err != nil {
		return nil, err
	}
	result := PollerResponse{
		RawResponse: resp,
		ID:          "GenericClient.CreateOrUpdateByID",
	}
	pt, err := armruntime.NewPoller("GenericClient.CreateOrUpdateByID", "", resp, client.pl, client.createOrUpdateByIDHandleError)
	if err != nil {
		return nil, err
	}
	result.Poller = pt
	return &result, nil
}

func (client *GenericClient) createOrUpdateByID(
	ctx context.Context,
	resourceID string,
	apiVersion string,
	resource interface{}) (*http.Response, error) {

	req, err := client.createOrUpdateByIDCreateRequest(ctx, resourceID, apiVersion, resource)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateByIDHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateByIDCreateRequest creates the CreateOrUpdateByID request.
func (client *GenericClient) createOrUpdateByIDCreateRequest(
	ctx context.Context,
	resourceID string,
	apiVersion string,
	resource interface{}) (*policy.Request, error) {

	if resourceID == "" {
		return nil, errors.New("parameter resourceID cannot be empty")
	}

	urlPath := resourceID
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, resource)
}

// createOrUpdateByIDHandleError handles the CreateOrUpdateByID error response.
func (client *GenericClient) createOrUpdateByIDHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(errors.Wrapf(err, "\n%s", string(body)), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetByID - Gets a resource by ID.
// If the operation fails it returns the *CloudError error type.
func (client *GenericClient) GetByID(ctx context.Context, resourceID string, apiVersion string, resource interface{}) (time.Duration, error) {
	req, err := client.getByIDCreateRequest(ctx, resourceID, apiVersion)
	if err != nil {
		return zeroDuration, err
	}
	resp, err := client.pl.Do(req)
	retryAfter := GetRetryAfter(resp)
	if err != nil {
		return retryAfter, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return retryAfter, client.getByIDHandleError(resp)
	}
	return zeroDuration, client.getByIDHandleResponse(resp, resource)
}

// getByIDCreateRequest creates the GetByID request.
func (client *GenericClient) getByIDCreateRequest(ctx context.Context, resourceID string, apiVersion string) (*policy.Request, error) {
	urlPath := "/{resourceId}"
	if resourceID == "" {
		return nil, errors.New("parameter resourceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", resourceID)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *GenericClient) getByIDHandleResponse(resp *http.Response, resource interface{}) error {
	if err := runtime.UnmarshalAsJSON(resp, resource); err != nil {
		return err
	}
	return nil
}

// getByIDHandleError handles the GetByID error response.
func (client *GenericClient) getByIDHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(errors.Wrapf(err, "\n%s", string(body)), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// DeleteByID - Deletes a resource by ID.
// If the operation fails it returns the *CloudError error type.
func (client *GenericClient) DeleteByID(ctx context.Context, resourceID string, apiVersion string) (time.Duration, error) {
	resp, err := client.deleteByID(ctx, resourceID, apiVersion)
	retryAfter := GetRetryAfter(resp)
	if err != nil {
		return retryAfter, err
	}

	return retryAfter, nil
}

// DeleteByID - Deletes a resource by ID.
// If the operation fails it returns the *CloudError error type.
func (client *GenericClient) deleteByID(ctx context.Context, resourceID string, apiVersion string) (*http.Response, error) {
	req, err := client.deleteByIDCreateRequest(ctx, resourceID, apiVersion)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent, http.StatusNotFound) {
		return nil, client.deleteByIDHandleError(resp)
	}
	return resp, nil
}

// deleteByIDCreateRequest creates the DeleteByID request.
func (client *GenericClient) deleteByIDCreateRequest(ctx context.Context, resourceID string, apiVersion string) (*policy.Request, error) {
	urlPath := "/{resourceId}"
	if resourceID == "" {
		return nil, errors.New("parameter resourceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", resourceID)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", apiVersion)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteByIDHandleError handles the DeleteByID error response.
func (client *GenericClient) deleteByIDHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}

	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(errors.Wrapf(err, "\n%s", string(body)), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

func (client *GenericClient) HeadByID(ctx context.Context, resourceID string, apiVersion string) (bool, time.Duration, error) {
	if resourceID == "" {
		return false, zeroDuration, errors.New("parameter resourceID cannot be empty")
	}

	ignored := struct{}{}
	retryAfter, err := client.GetByID(ctx, resourceID, apiVersion, &ignored)

	switch {
	case IsNotFoundError(err):
		return false, retryAfter, nil
	case err != nil:
		return false, retryAfter, err
	default:
		return true, retryAfter, nil
	}
}

func IsNotFoundError(err error) bool {
	var typedError azcore.HTTPResponse
	if errors.As(err, &typedError) {
		// The linter doesn't realize that we don't need to close RawResponse() in this context.
		// Suppressing it as it is a false positive.
		// nolint:bodyclose
		if typedError.RawResponse() != nil && typedError.RawResponse().StatusCode == 404 {
			return true
		}
	}

	return false
}
