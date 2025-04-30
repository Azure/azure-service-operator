/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"

	api "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20240801"
	hub "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20240801/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

var _ extensions.PreReconciliationChecker = &FlexibleServersDatabaseExtension{}

func (extension *FlexibleServersDatabaseExtension) PreReconcileCheck(
	ctx context.Context,
	obj genruntime.MetaObject,
	owner genruntime.MetaObject,
	resourceResolver *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
	next extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
	// Check to see if our owning server is ready for the database to be reconciled
	// Owner nil can happen if the server owner of the database is referenced by armID
	if owner != nil {
		if server, ok := owner.(*hub.FlexibleServer); ok {
			serverState := server.Status.State
			if serverState != nil && flexibleServerStateBlocksReconciliation(*serverState) {
				return extensions.BlockReconcile(
					fmt.Sprintf(
						"Owning FlexibleServer is in provisioning state %q",
						*serverState)), nil
			}
		}
	}

	return next(ctx, obj, owner, resourceResolver, armClient, log)
}

var _ extensions.Importer = &FlexibleServersDatabaseExtension{}

// Import skips databases that can't be managed by ARM
func (extension *FlexibleServersDatabaseExtension) Import(
	ctx context.Context,
	rsrc genruntime.ImportableResource,
	owner *genruntime.ResourceReference,
	next extensions.ImporterFunc,
) (extensions.ImportResult, error) {
	// If this cast doesn't compile, update the `api` import to reference the now latest
	// stable version of the dbforpostgresql group (this will happen when we import a new
	// API version in the generator.)
	if server, ok := rsrc.(*api.FlexibleServersDatabase); ok {
		if server.Spec.AzureName == "azure_maintenance" {
			return extensions.ImportSkipped("azure_maintenance database is not accessible by users"), nil
		}

		if server.Spec.AzureName == "azure_sys" {
			return extensions.ImportSkipped("built in databases cannot be managed by ARM"), nil
		}

		if server.Spec.AzureName == "postgres" {
			return extensions.ImportSkipped("built in databases cannot be managed by ARM"), nil
		}
	}

	return next(ctx, rsrc, owner)
}
