/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	armdataprotection "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
	dataprotection "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20231101/storage"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

var _ extensions.PreReconciliationChecker = &BackupVaultsBackupInstanceExtension{}
var _ extensions.PostReconciliationChecker = &BackupVaultsBackupInstanceExtension{}

var nonRetryableBackupInstanceStates = set.Make(
	"protectionstopped",
	"softdeleted",
	"configuringprotectionfailed",
)

func (extension *BackupVaultsBackupInstanceExtension) PostReconcileCheck(
	ctx context.Context,
	obj genruntime.MetaObject,
	owner genruntime.MetaObject,
	_ *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
	_ extensions.PostReconcileCheckFunc) (extensions.PostReconcileCheckResult, error) {

	log.V(Debug).Info("########################## Starting Post-reconcilation for Backup Instance ##########################")
	backupInstance, ok := obj.(*dataprotection.BackupVaultsBackupInstance)
	if !ok {
		return extensions.PostReconcileCheckResult{},
			errors.Errorf("cannot run on unknown resource type %T, expected *dataprotection.BackupVaultsBackupInstance", obj)
	}

	var _ conversion.Hub = backupInstance

	if owner != nil && backupInstance != nil && backupInstance.Status.Id != nil &&
		backupInstance.Status.Properties != nil && backupInstance.Status.Properties.ProtectionStatus != nil &&
		backupInstance.Status.Properties.ProtectionStatus.Status != nil {

		protectionStatus := *backupInstance.Status.Properties.ProtectionStatus.Status
		log.V(Debug).Info("########################## Protection Status ##########################" + protectionStatus)

		if !nonRetryableBackupInstanceStates.Contains(protectionStatus) {
			if (protectionStatus == "ProtectionError") {
				log.V(Debug).Info("########################## Protection Status is Protection Error ##########################")
				protectionStatusErrorCode := ""
				if backupInstance.Status.Properties.ProtectionStatus.ErrorDetails == nil {
					log.V(Debug).Info("########################## BlockReconcile No error for Backup Instance ##########################")
				}

				protectionStatusErrorCode = strings.ToLower(*backupInstance.Status.Properties.ProtectionStatus.ErrorDetails.Code)

				log.V(Debug).Info("########################## Protection Error code is ##########################" + protectionStatusErrorCode)

				if protectionStatusErrorCode != "" &&
					strings.Contains(protectionStatusErrorCode, "usererror") {

					id, _ := genruntime.GetAndParseResourceID(backupInstance)

					subscription := id.SubscriptionID
					rg := id.ResourceGroupName
					vaultName := id.Parent.Name

					log.V(Debug).Info("########################## Starting NewBackupInstancesClient for Backup Instance ##########################")

					clientFactory, err := armdataprotection.NewClientFactory(subscription, armClient.Creds(), armClient.ClientOptions())
					if err != nil {
						log.Error(err, "failed to create armdataprotection client")
					}

					var parameters armdataprotection.SyncBackupInstanceRequest
					parameters.SyncType = to.Ptr(armdataprotection.SyncTypeDefault)

					log.V(Debug).Info("########################## Starting BeginSyncBackupInstance for Backup Instance ##########################")
					poller, err := clientFactory.NewBackupInstancesClient().BeginSyncBackupInstance(ctx, rg, vaultName, backupInstance.AzureName(), parameters, nil)
					if err != nil {
						log.Error(err, "failed BeginSyncBackupInstance to get the result")
					}

					_, err = poller.PollUntilDone(ctx, nil)
					if err != nil {
						log.Error(err, "failed BeginSyncBackupInstance to poll the result")
					}
				}
			}
			return extensions.PostReconcileCheckResultFailure(
				fmt.Sprintf("Backup Instance is in State: %q", protectionStatus)), nil
		}
	}

	log.V(Debug).Info("########################## Ending Post-reconcilation for Backup Instance ##########################")
	return extensions.PostReconcileCheckResultSuccess(), nil
}

func (ext *BackupVaultsBackupInstanceExtension) PreReconcileCheck(
	ctx context.Context,
	obj genruntime.MetaObject,
	owner genruntime.MetaObject,
	_ *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
	_ extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {

	log.V(Debug).Info("########################## Starting Pre-reconcilation for Backup Instance ##########################")
	backupInstance, ok := obj.(*dataprotection.BackupVaultsBackupInstance)
	if !ok {
		return extensions.PreReconcileCheckResult{},
			errors.Errorf("cannot run on unknown resource type %T, expected *dataprotection.BackupVaultsBackupInstance", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = backupInstance

	log.V(Debug).Info("########################## Ending Pre-reconcilation for Backup Instance ##########################")

	return extensions.ProceedWithReconcile(), nil
}
