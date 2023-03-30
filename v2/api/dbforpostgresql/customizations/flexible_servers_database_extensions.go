/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"

	"github.com/go-logr/logr"

	postgresql "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

var _ extensions.PreReconciliationChecker = &FlexibleServersDatabaseExtension{}

func (ext *FlexibleServersDatabaseExtension) PreReconcileCheck(
	_ context.Context,
	_ genruntime.MetaObject,
	owner genruntime.MetaObject,
	_ kubeclient.Client,
	_ *genericarmclient.GenericClient,
	_ logr.Logger,
	_ extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
	// Check to see if our owning server is ready for the database to be reconciled
	if server, ok := owner.(*postgresql.FlexibleServer); ok {
		serverState := server.Status.State
		if serverState != nil && flexibleServerStateBlocksReconciliation(*serverState) {
			return extensions.BlockReconcile(
				fmt.Sprintf(
					"Owning FlexibleServer is in provisioning state %q",
					*serverState)), nil
		}
	}

	return extensions.ProceedWithReconcile(), nil
}
