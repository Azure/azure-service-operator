/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package httpclient

import (
	"errors"
	"net/http"
)

// Client preserves HTTP behaviour while checking redirected destinations.
type Client struct {
	client *http.Client
}

// NewRedirectGuardedClient copies base without changing its shared transport.
func NewRedirectGuardedClient(base *http.Client, allowed Origin) *Client {
	if base == nil {
		base = http.DefaultClient
	}
	client := *base
	previous := base.CheckRedirect
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if err := allowed.Validate(req.URL); err != nil {
			return err
		}
		if previous != nil {
			if err := previous(req, via); err != nil {
				return err
			}
		} else if len(via) >= 10 {
			return errors.New("stopped after 10 redirects")
		}
		return allowed.Validate(req.URL)
	}
	return &Client{
		client: &client,
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	var rejected *DestinationError
	if errors.As(err, &rejected) {
		// http.Client wraps redirect errors in a url.Error containing the original URL.
		return resp, rejected
	}
	return resp, err
}
