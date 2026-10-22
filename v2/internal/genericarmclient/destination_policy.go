/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient

import (
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"

	"github.com/Azure/azure-service-operator/v2/internal/httpclient"
)

// destinationPolicy is a policy that requires requests to match a specific allowed origin.
type destinationPolicy struct {
	allowed httpclient.Origin
}

var _ policy.Policy = &destinationPolicy{}

// newDestinationPolicy creates a new policy that requires requests to match the allowed origin.
// allowed is the origin that requests must match.
// Requests to any other origin will be rejected with a httpclient.DestinationError.
func newDestinationPolicy(allowed httpclient.Origin) policy.Policy {
	return &destinationPolicy{
		allowed: allowed,
	}
}

func (p *destinationPolicy) Do(req *policy.Request) (*http.Response, error) {
	if err := p.allowed.Validate(req.Raw().URL); err != nil {
		return nil, err
	}

	return req.Next()
}
