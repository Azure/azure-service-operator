/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package mysql

import (
	"context"

	"github.com/pkg/errors"
	ctrlconversion "sigs.k8s.io/controller-runtime/pkg/conversion"

	asomysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1"
	dbformysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20210501storage"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

type Connector interface {
	CreateOrUpdate(ctx context.Context) error
	Delete(ctx context.Context) error
	Exists(ctx context.Context) (bool, error)
}

func getServerFQDN(ctx context.Context, resourceResolver *resolver.Resolver, user *asomysql.User) (string, error) {
	// Get the owner - at this point it must exist
	owner, err := resourceResolver.ResolveOwner(ctx, user)
	if err != nil {
		return "", err
	}

	flexibleServer, ok := owner.(*dbformysql.FlexibleServer)
	if !ok {
		return "", errors.Errorf("owner was not type FlexibleServer, instead: %T", owner)
	}

	// Magical assertion to ensure that this is still the storage type
	var _ ctrlconversion.Hub = &dbformysql.FlexibleServer{}

	if flexibleServer.Status.FullyQualifiedDomainName == nil {
		// This possibly means that the server hasn't finished deploying yet
		err = errors.Errorf("owning Flexibleserver %q '.status.fullyQualifiedDomainName' not set. Has the server been provisioned successfully?", flexibleServer.Name)
		return "", conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonWaitingForOwner)
	}
	serverFQDN := *flexibleServer.Status.FullyQualifiedDomainName

	return serverFQDN, nil
}
