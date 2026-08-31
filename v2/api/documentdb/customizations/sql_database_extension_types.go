/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"context"
	"strings"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	documentdb "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

const (
	migrateToAutoscale        = "migrateToAutoscale"
	migrateToManualThroughput = "migrateToManualThroughput"
	throughputSettingsSuffix  = "/throughputSettings/default"
)

var (
	_ extensions.ErrorClassifier          = &SqlDatabaseExtension{}
	_ extensions.PreReconciliationChecker = &SqlDatabaseExtension{}
)

type sqlDatabaseThroughputSettings struct {
	Properties *sqlDatabaseThroughputProperties `json:"properties,omitempty"`
}

type sqlDatabaseThroughputProperties struct {
	Resource *documentdb.ThroughputSettingsGetProperties_Resource_STATUS `json:"resource,omitempty"`
}

// PreReconcileCheck migrates existing dedicated throughput when the desired throughput mode differs.
func (extension *SqlDatabaseExtension) PreReconcileCheck(
	ctx context.Context,
	obj genruntime.MetaObject,
	resourceResolver *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
	next extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
	database, ok := obj.(*documentdb.SqlDatabase)
	if !ok {
		return extensions.PreReconcileCheckResult{}, eris.Errorf(
			"cannot run on unknown resource type %T, expected *documentdb.SqlDatabase",
			obj)
	}

	var _ conversion.Hub = database

	if database.Status.Id == nil || *database.Status.Id == "" || database.Spec.Options == nil {
		return next(ctx, obj, resourceResolver, armClient, log)
	}

	settingsID := strings.TrimSuffix(*database.Status.Id, "/") + throughputSettingsSuffix
	var settings sqlDatabaseThroughputSettings
	_, err := armClient.GetByID(ctx, settingsID, database.GetAPIVersion(), &settings)
	if genericarmclient.IsNotFoundError(err) {
		return next(ctx, obj, resourceResolver, armClient, log)
	}
	if err != nil {
		return extensions.PreReconcileCheckResult{}, eris.Wrap(err, "getting SQL database throughput settings")
	}

	if settings.Properties != nil &&
		settings.Properties.Resource != nil &&
		settings.Properties.Resource.OfferReplacePending != nil &&
		strings.EqualFold(*settings.Properties.Resource.OfferReplacePending, "true") {
		return extensions.BlockReconcile("SQL database throughput update is in progress"), nil
	}

	action := sqlDatabaseMigrationAction(database.Spec.Options, settings.Properties)
	if action == "" {
		return next(ctx, obj, resourceResolver, armClient, log)
	}

	log.Info("Migrating SQL database throughput", "action", action)
	_, err = armClient.BeginPostActionByID(ctx, settingsID, action, database.GetAPIVersion())
	if err != nil {
		return extensions.PreReconcileCheckResult{}, eris.Wrap(err, "starting SQL database throughput migration")
	}

	return extensions.BlockReconcile("SQL database throughput migration is in progress"), nil
}

func sqlDatabaseMigrationAction(
	desired *documentdb.CreateUpdateOptions,
	observed *sqlDatabaseThroughputProperties,
) string {
	if desired == nil || observed == nil || observed.Resource == nil {
		return ""
	}

	if desired.AutoscaleSettings != nil &&
		desired.Throughput == nil &&
		observed.Resource.AutoscaleSettings == nil &&
		observed.Resource.Throughput != nil {
		return migrateToAutoscale
	}

	if desired.Throughput != nil && desired.AutoscaleSettings == nil && observed.Resource.AutoscaleSettings != nil {
		return migrateToManualThroughput
	}

	return ""
}

// ClassifyError evaluates the provided error, returning whether it is fatal or can be retried.
// A BadRequest (400) is normally fatal, but CosmosDB Databases may return 400 if database creation is attempted while
// the parent account is still being created, so we make BadRequest retryable for this case.
// cloudError is the error returned from ARM.
// next is the next implementation to call.
func (extension *SqlDatabaseExtension) ClassifyError(
	cloudError *genericarmclient.CloudError,
	_ string,
	_ logr.Logger,
	next extensions.ErrorClassifierFunc,
) (core.CloudErrorDetails, error) {
	details, err := next(cloudError)
	if err != nil {
		return core.CloudErrorDetails{}, err
	}

	// Override is to treat BadRequest as retryable for SqlDatabases
	if isRetryableBadRequest(cloudError) {
		details.Classification = core.ErrorRetryable
	}

	return details, nil
}

// isRetryableBadRequest checks the passed error to see if it is a retryable conflict, returning true if it is.
func isRetryableBadRequest(err *genericarmclient.CloudError) bool {
	if err == nil {
		return false
	}

	return err.Code() == "BadRequest"
}
