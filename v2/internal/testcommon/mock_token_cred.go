/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
)

// PolicyFunc is a type that implements the Policy interface.
// Use this type when implementing a stateless policy as a first-class function.
type PolicyFunc func(*policy.Request) (*http.Response, error)

// Do implements the Policy interface on PolicyFunc.
func (pf PolicyFunc) Do(req *policy.Request) (*http.Response, error) {
	return pf(req)
}

var _ policy.Policy = PolicyFunc(nil)

type MockTokenCredential struct{}

func (MockTokenCredential) NewAuthenticationPolicy() policy.Policy {
	return PolicyFunc(func(req *policy.Request) (*http.Response, error) {
		return req.Next()
	})
}

func (MockTokenCredential) GetToken(context.Context, policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{
		Token:     "abc123",
		ExpiresOn: time.Now().Add(1 * time.Hour),
	}, nil
}

var _ azcore.TokenCredential = MockTokenCredential{}
