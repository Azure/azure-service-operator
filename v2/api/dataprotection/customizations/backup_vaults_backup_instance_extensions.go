/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"fmt"
	"strings"

	armdataprotection "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
	dataprotection "github.com/Azure/azure-service-operator/v2/api/dataprotection/v1api20231101/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	annotation "github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

var (
	_ extensions.PreReconciliationChecker  = &BackupVaultsBackupInstanceExtension{}
	_ extensions.PostReconciliationChecker = &BackupVaultsBackupInstanceExtension{}
)

var nonRetryableStates = set.Make(
	"ConfiguringProtectionFailed",
	"Invalid",
	"NotProtected",
	"ProtectionConfigured",
	"SoftDeleted",
	"ProtectionStopped",
	"BackupSchedulesSuspended",
	"RetentionSchedulesSuspended",
)

func (extension *BackupVaultsBackupInstanceExtension) PostReconcileCheck(
	ctx context.Context,
	obj genruntime.MetaObject,
	owner genruntime.MetaObject,
	_ *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
	_ extensions.PostReconcileCheckFunc,
) (extensions.PostReconcileCheckResult, error) {
	log.V(Debug).Info("########################## Starting Post-reconcilation for Backup Instance ##########################")
	backupInstance, ok := obj.(*dataprotection.BackupVaultsBackupInstance)
	if !ok {
		return extensions.PostReconcileCheckResult{},
			errors.Errorf("cannot run on unknown resource type %T, expected *dataprotection.BackupVaultsBackupInstance", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = backupInstance

	if owner != nil &&
		backupInstance != nil &&
		backupInstance.Status.Id != nil &&
		backupInstance.Status.Properties != nil &&
		backupInstance.Status.Properties.ProtectionStatus != nil &&
		backupInstance.Status.Properties.ProtectionStatus.Status != nil {

		protectionStatus := *backupInstance.Status.Properties.ProtectionStatus.Status
		log.V(Debug).Info(fmt.Sprintf("########################## Protection Status is  %q ##########################", protectionStatus))

		// return success for reconcilation if the state is non-retryable
		if nonRetryableStates.Contains(protectionStatus) {
			log.V(Debug).Info("########################## Returning PostReconcileCheckResultSuccess ##########################")
			return extensions.PostReconcileCheckResultSuccess(), nil
		}

		// call sync api only when protection status is ProtectionError and error code is usererror
		if protectionStatus == "ProtectionError" {
			var protectionStatusErrorCode string
			protectionStatusErrorCode = strings.ToLower(*backupInstance.Status.Properties.ProtectionStatus.ErrorDetails.Code)
			log.V(Debug).Info(fmt.Sprintf("########################## Protection Error code is  %q ##########################", protectionStatusErrorCode))

			if protectionStatusErrorCode != "" && strings.Contains(protectionStatusErrorCode, "usererror") {
				id, _ := genruntime.GetAndParseResourceID(backupInstance)
				subscription := id.SubscriptionID
				rg := id.ResourceGroupName
				vaultName := id.Parent.Name

				log.V(Debug).Info("########################## Starting NewBackupInstancesClient  ##########################")
				clientFactory, err := armdataprotection.NewClientFactory(subscription, armClient.Creds(), armClient.ClientOptions())
				if err != nil {
					return extensions.PostReconcileCheckResultFailure("failed to create armdataprotection client"), err
				}

				var parameters armdataprotection.SyncBackupInstanceRequest
				parameters.SyncType = to.Ptr(armdataprotection.SyncTypeDefault)

				// get the resume token from the resource
				pollerID, pollerResumeToken, hasToken := annotation.GetPollerResumeToken(backupInstance)

				log.V(Debug).Info(fmt.Sprintf("########################## pollerID is  %q ##########################", pollerID))
				log.V(Debug).Info(fmt.Sprintf("########################## pollerResumeToken is  %q ##########################", pollerResumeToken))

				// if resume token is empty then call the begin sync api and set the resume token
				if !hasToken {
					log.V(Debug).Info("########################## Starting BeginSyncBackupInstance  ##########################")
					poller, err := clientFactory.NewBackupInstancesClient().BeginSyncBackupInstance(ctx, rg, vaultName, backupInstance.AzureName(), parameters, nil)
					if err != nil {
						return extensions.PostReconcileCheckResultFailure("failed BeginSyncBackupInstance to get the result"), err
					}

					resumeToken, err := poller.ResumeToken()
					if err != nil {
						return extensions.PostReconcileCheckResultFailure("couldn't create PUT resume token for resource"), err
					}

					annotation.SetPollerResumeToken(backupInstance, "BeginSyncBackupInstance.PollerID", resumeToken)
				}

				// BeginSyncBackupInstance is in-progress - poller id and resume token is available to poll
				if pollerID != "BeginSyncBackupInstance.PollerID" {
					return extensions.PostReconcileCheckResultFailure(fmt.Sprintf("cannot Poll with pollerID=%s", pollerID)), nil
				}

				log.V(Debug).Info("########################## Polling for BeginSyncBackupInstance ##########################")
				var options *armdataprotection.BackupInstancesClientBeginSyncBackupInstanceOptions
				options.ResumeToken = pollerResumeToken
				poller, clientSyncErr := clientFactory.NewBackupInstancesClient().BeginSyncBackupInstance(ctx, rg, vaultName, backupInstance.AzureName(), parameters, options)
				if clientSyncErr != nil {
					return extensions.PostReconcileCheckResultFailure("failed BeginSyncBackupInstance to get the result"), err
				}
				if poller.Done() {
					log.V(Debug).Info("########################## ClearPollerResumeToken for BeginSyncBackupInstance ##########################")
					annotation.ClearPollerResumeToken(backupInstance)
				}
			}
		}
	}
	return extensions.PostReconcileCheckResultFailure("Backup Instance is in non terminal state"), nil
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
