/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/pkg/errors"
)

// PollerResponse is the response from issuing a PUT to Azure. It contains a poller (for polling the long-running
// operation URL) and a RawResponse containing the raw HTTP response.
type PollerResponse struct {
	// Poller contains an initialized poller.
	Poller *azcore.Poller

	// ID is the ID of the poller (not the ID of the resource). This is used to prevent another kind of poller from
	// being resumed with this pollers URL (which would cause deserialization issues and other problems).
	ID string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Resume rehydrates a ResourcesCreateOrUpdateByIDPollerResponse from the provided client and resume token.
func (l *PollerResponse) Resume(ctx context.Context, client *GenericClient, token string) error {
	poller, err := armruntime.NewPollerFromResumeToken(l.ID, token, client.pl)
	if err != nil {
		return err
	}
	// The linter doesn't realize that we don't need to close the resp body because it's already done by the poller.
	// Suppressing it as it is a false positive.
	// nolint:bodyclose
	resp, err := poller.Poll(ctx)
	if err != nil {
		var typedError *azcore.ResponseError
		if errors.As(err, &typedError) {
			if typedError.RawResponse != nil {
				return client.createOrUpdateByIDHandleError(typedError.RawResponse)
			}
		}
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}
