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

var _ extensions.ErrorClassifier = &AfdCustomDomainExtension{}

func (extension *AfdCustomDomainExtension) ClassifyError(
	cloudError *genericarmclient.CloudError,
	apiVersion string,
	log logr.Logger,
	next extensions.ErrorClassifierFunc,
) (core.CloudErrorDetails, error) {
	details, err := next(cloudError)
	if err != nil {
		return core.CloudErrorDetails{}, err
	}

	if isAfdCustomDomainRouteAssociationError(cloudError) {
		details.Retry = retry.Slow
	}

	return details, nil
}

func isAfdCustomDomainRouteAssociationError(err *genericarmclient.CloudError) bool {
	if err == nil {
		return false
	}

	return err.Code() == "BadRequest" &&
		strings.Contains(err.Message(), "This resource is still associated with a route")
}
