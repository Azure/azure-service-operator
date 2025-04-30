/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"

	postgresql "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

var _ extensions.PreReconciliationChecker = &FlexibleServersFirewallRuleExtension{}

func (ext *FlexibleServersFirewallRuleExtension) PreReconcileCheck(
	ctx context.Context,
	obj genruntime.MetaObject,
	owner genruntime.MetaObject,
	resourceResolver *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
	next extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
	// Check to see if our owning server is ready for the database to be reconciled
	// Owner nil can happen if the server/owner of the firewall rule is referenced by armID
	if owner != nil {
		if server, ok := owner.(*postgresql.FlexibleServer); ok {
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
