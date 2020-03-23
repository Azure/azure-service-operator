/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package zips

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/devigned/tab"
)

type (
	// Client is the HTTP client for the Cloud Partner Portal
	Client struct {
		HTTPClient *http.Client
		Authorizer autorest.Authorizer
		Host       string
		mwStack    []MiddlewareFunc
	}

	// ClientOption is a variadic optional configuration func
	ClientOption func(c *Client) error

	// MiddlewareFunc allows a consumer of the Client to inject handlers within the request / response pipeline
	//
	// The example below adds the atom xml content type to the request, calls the next middleware and returns the
	// result.
	//
	// addAtomXMLContentType MiddlewareFunc = func(next RestHandler) RestHandler {
	//		return func(ctx context.Context, req *http.Request) (res *http.Response, e error) {
	//			if req.Method != http.MethodGet && req.Method != http.MethodHead {
	//				req.Header.Add("content-Type", "application/atom+xml;type=entry;charset=utf-8")
	//			}
	//			return next(ctx, req)
	//		}
	//	}
	MiddlewareFunc func(next RestHandler) RestHandler

	// RestHandler is used to transform a request and response within the http pipeline
	RestHandler func(ctx context.Context, req *http.Request) (*http.Response, error)

	// SimpleTokenProvider makes it easy to authorize with a string bearer token
	SimpleTokenProvider struct{}

	HttpError struct {
		StatusCode int
		Body       string
		Response   *http.Response
	}

	NotFoundError struct {
		Response *http.Response
	}
)

var (
	httpLogger MiddlewareFunc = func(next RestHandler) RestHandler {
		return func(ctx context.Context, req *http.Request) (*http.Response, error) {
			printErrLine := func(format string, args ...interface{}) {
				_, _ = fmt.Fprintf(os.Stderr, format, args...)
			}

			requestDump, err := httputil.DumpRequest(req, true)
			if err != nil {
				printErrLine("+%v\n", err)
			}
			printErrLine(string(requestDump))

			res, err := next(ctx, req)
			if err != nil {
				return res, err
			}

			resDump, err := httputil.DumpResponse(res, true)
			if err != nil {
				printErrLine("+%v\n", err)
			}
			printErrLine(string(resDump))

			return res, err
		}
	}
)

func NewClient(authorizer autorest.Authorizer, opts ...ClientOption) (*Client, error) {
	c := &Client{
		Authorizer: authorizer,
		Host:       azure.PublicCloud.ResourceManagerEndpoint,
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) PutDeployment(ctx context.Context, deployment *Deployment, mw ...MiddlewareFunc) (*Deployment, error) {
	entityPath, err := deployment.GetEntityPath()
	if err != nil {
		return nil, err
	}

	bits, err := json.Marshal(deployment)
	if err != nil {
		return nil, err
	}

	res, err := c.Put(ctx, entityPath, bytes.NewReader(bits), mw...)
	defer closeResponse(ctx, res)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, NewHttpError(res, string(body))
	}

	if err := json.Unmarshal(body, deployment); err != nil {
		return deployment, err
	}

	return deployment, nil
}

func (c *Client) GetResource(ctx context.Context, resourceID string, resource interface{}, mw ...MiddlewareFunc) error {
	res, err := c.Get(ctx, resourceID, mw...)
	defer closeResponse(ctx, res)

	if err != nil {
		return err
	}

	if res.StatusCode == 404 {
		return &NotFoundError{
			Response: res,
		}
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode > 299 {
		return NewHttpError(res, string(body))
	}

	if resource != nil {
		return json.Unmarshal(body, resource)
	}

	return nil
}

// DeleteResource will make an HTTP DELETE call to the resourceID and attempt to fill the resource with the response.
// If the body of the response is empty, the resource will be nil.
func (c *Client) DeleteResource(ctx context.Context, resourceID string, resource interface{}, mw ...MiddlewareFunc) error {
	res, err := c.Delete(ctx, resourceID, mw...)
	defer closeResponse(ctx, res)
	if err != nil {
		return err
	}

	if res.StatusCode == 404 {
		// you asked it to be gone, well, it is.
		return nil
	}

	if res.StatusCode > 299 {
		// do our best to read the body, but if we can't, just carry on
		body, _ := ioutil.ReadAll(res.Body)
		return NewHttpError(res, string(body))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if len(body) == 0 || resource == nil {
		return nil
	}

	return json.Unmarshal(body, resource)
}

// Get will execute a HTTP GET request
func (c *Client) Get(ctx context.Context, entityPath string, mw ...MiddlewareFunc) (*http.Response, error) {
	return c.execute(ctx, http.MethodGet, entityPath, nil, mw...)
}

// Put will execute a HTTP PUT request
func (c *Client) Put(ctx context.Context, entityPath string, body io.Reader, mw ...MiddlewareFunc) (*http.Response, error) {
	return c.execute(ctx, http.MethodPut, entityPath, body, mw...)
}

// Post will execute a HTTP POST request
func (c *Client) Post(ctx context.Context, entityPath string, body io.Reader, mw ...MiddlewareFunc) (*http.Response, error) {
	return c.execute(ctx, http.MethodPost, entityPath, body, mw...)
}

// Delete will execute a HTTP DELETE request
func (c *Client) Delete(ctx context.Context, entityPath string, mw ...MiddlewareFunc) (*http.Response, error) {
	return c.execute(ctx, http.MethodDelete, entityPath, nil, mw...)
}

// Head will execute a HTTP HEAD request
func (c *Client) Head(ctx context.Context, entityPath string, mw ...MiddlewareFunc) (*http.Response, error) {
	return c.execute(ctx, http.MethodHead, entityPath, nil, mw...)
}

func (c *Client) execute(ctx context.Context, method string, entityPath string, body io.Reader, mw ...MiddlewareFunc) (*http.Response, error) {
	req, err := http.NewRequest(method, c.Host+strings.TrimPrefix(entityPath, "/"), body)
	if err != nil {
		tab.For(ctx).Error(err)
		return nil, err
	}

	final := func(_ RestHandler) RestHandler {
		return func(reqCtx context.Context, request *http.Request) (*http.Response, error) {
			client := c.getHTTPClient()
			request = request.WithContext(reqCtx)
			request.Header.Set("Content-Type", "application/json")
			request, err := autorest.CreatePreparer(c.Authorizer.WithAuthorization()).Prepare(request)
			if err != nil {
				return nil, err
			}

			return client.Do(request)
		}
	}

	mwStack := []MiddlewareFunc{final}
	if os.Getenv("DEBUG") == "true" {
		mwStack = append(mwStack, httpLogger)
	}

	sl := len(c.mwStack) - 1
	for i := sl; i >= 0; i-- {
		mwStack = append(mwStack, c.mwStack[i])
	}

	for i := len(mw) - 1; i >= 0; i-- {
		if mw[i] != nil {
			mwStack = append(mwStack, mw[i])
		}
	}

	var h RestHandler
	for _, mw := range mwStack {
		h = mw(h)
	}

	return h(ctx, req)
}

func (c *Client) getHTTPClient() *http.Client {
	if c.HTTPClient != nil {
		return c.HTTPClient
	}

	return &http.Client{
		Timeout: 60 * time.Second,
	}
}

func closeResponse(ctx context.Context, res *http.Response) {
	if res == nil {
		return
	}

	if err := res.Body.Close(); err != nil {
		tab.For(ctx).Error(err)
	}
}

func NewHttpError(response *http.Response, body string) *HttpError {
	var statusCode int
	if response != nil {
		statusCode = response.StatusCode
	}
	return &HttpError{
		StatusCode: statusCode,
		Body:       body,
		Response:   response,
	}
}

func (e HttpError) Error() string {
	var u *url.URL
	if e.Response != nil && e.Response.Request != nil {
		u = e.Response.Request.URL
	}
	return fmt.Sprintf("uri: %s, status: %d, body: %s", u, e.StatusCode, e.Body)
}

func (e NotFoundError) Error() string {
	var u *url.URL
	var statusCode int
	if e.Response != nil {
		statusCode = e.Response.StatusCode
		if e.Response.Request != nil {
			u = e.Response.Request.URL
		}
	}
	return fmt.Sprintf("not found uri: %s, status: %d", u, statusCode)
}

func IsNotFound(err error) bool {
	_, ok := err.(*NotFoundError)
	return ok
}

// WithAuthorization will inject the AZURE_TOKEN env var as the bearer token for API auth
//
// This is useful if you want to use a token from az cli.
// `AZURE_TOKEN=$(az account get-access-token --resource https://cloudpartner.azure.com --query "accessToken" -o tsv) pub publishers list`
func (s SimpleTokenProvider) WithAuthorization() autorest.PrepareDecorator {
	return func(p autorest.Preparer) autorest.Preparer {
		return autorest.PreparerFunc(func(r *http.Request) (*http.Request, error) {
			r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("AZURE_TOKEN")))
			return r, nil
		})
	}
}
