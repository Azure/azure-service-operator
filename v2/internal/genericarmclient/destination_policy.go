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

type destinationPolicy struct{ allowed httpclient.Origin }

func newDestinationPolicy(allowed httpclient.Origin) policy.Policy {
	return &destinationPolicy{allowed: allowed}
}

func (p *destinationPolicy) Do(req *policy.Request) (*http.Response, error) {
	if err := p.allowed.Validate(req.Raw().URL); err != nil {
		return nil, err
	}
	return req.Next()
}
