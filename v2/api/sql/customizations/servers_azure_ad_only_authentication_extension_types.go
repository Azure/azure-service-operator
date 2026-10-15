/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"strings"

	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/retry"
)

var _ extensions.ErrorClassifier = &ServersAzureADOnlyAuthenticationExtension{}

func (extension *ServersAzureADOnlyAuthenticationExtension) ClassifyError(
	cloudError *genericarmclient.CloudError,
	apiVersion string,
	log logr.Logger,
	next extensions.ErrorClassifierFunc,
) (core.CloudErrorDetails, error) {
	details, err := next(cloudError)
	if err != nil {
		return core.CloudErrorDetails{}, err
	}

	// A separately managed administrator may still be reconciling when authentication is enabled.
	if cloudError != nil && strings.EqualFold(cloudError.Code(), "InvalidServerAADOnlyAuthNoAADAdminPropertyName") {
		details.Classification = core.ErrorRetryable
		details.Retry = retry.Slow
	}

	return details, nil
}
