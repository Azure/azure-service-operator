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
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

var _ extensions.PostReconciliationChecker = &BackupVaultsBackupInstanceExtension{}

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

const (
	BackupInstancePollerResumeTokenAnnotation = "serviceoperator.azure.com/bi-poller-resume-token"
)

func GetPollerResumeToken(obj genruntime.MetaObject, log logr.Logger) (string, bool) {
	log.V(Debug).Info("GetPollerResumeToken")
	token, hasResumeToken := obj.GetAnnotations()[BackupInstancePollerResumeTokenAnnotation]
	return token, hasResumeToken
}

func SetPollerResumeToken(obj genruntime.MetaObject, token string, log logr.Logger) {
	log.V(Debug).Info("SetPollerResumeToken")
	genruntime.AddAnnotation(obj, BackupInstancePollerResumeTokenAnnotation, token)
}

// ClearPollerResumeToken clears the poller resume token and ID annotations
func ClearPollerResumeToken(obj genruntime.MetaObject, log logr.Logger) {
	log.V(Debug).Info("ClearPollerResumeToken")
	genruntime.RemoveAnnotation(obj, BackupInstancePollerResumeTokenAnnotation)
}

func (extension *BackupVaultsBackupInstanceExtension) PostReconcileCheck(
	ctx context.Context,
	obj genruntime.MetaObject,
	owner genruntime.MetaObject,
	resolver *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
	next extensions.PostReconcileCheckFunc,
) (extensions.PostReconcileCheckResult, error) {
	log.V(Debug).Info("Starting Post-reconcilation for Backup Instance")
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
		log.V(Debug).Info(fmt.Sprintf("Protection Status is  %q", protectionStatus))

		// return success for reconcilation if the state is non-retryable
		if nonRetryableStates.Contains(protectionStatus) {
			log.V(Debug).Info("Returning PostReconcileCheckResultSuccess")
			return next(ctx, obj, owner, resolver, armClient, log)
		}

		// call sync api only when protection status is ProtectionError and error code is usererror
		if protectionStatus == "ProtectionError" {
			var protectionStatusErrorCode string
			protectionStatusErrorCode = strings.ToLower(*backupInstance.Status.Properties.ProtectionStatus.ErrorDetails.Code)
			log.V(Debug).Info(fmt.Sprintf("Protection Error code is  %q", protectionStatusErrorCode))

			if protectionStatusErrorCode != "" && strings.Contains(protectionStatusErrorCode, "usererror") {
				id, _ := genruntime.GetAndParseResourceID(backupInstance)
				subscription := id.SubscriptionID
				rg := id.ResourceGroupName
				vaultName := id.Parent.Name

				clientFactory, err := armdataprotection.NewClientFactory(subscription, armClient.Creds(), armClient.ClientOptions())
				if err != nil {
					return extensions.PostReconcileCheckResultFailure("failed to create armdataprotection client"), err
				}

				var parameters armdataprotection.SyncBackupInstanceRequest
				parameters.SyncType = to.Ptr(armdataprotection.SyncTypeDefault)

				// get the resume token from the resource
				pollerResumeToken, _ := GetPollerResumeToken(obj, log)

				// BeginSyncBackupInstance is in-progress - poller resume token is available

				log.V(Debug).Info("Starting BeginSyncBackupInstance")

				poller, err := clientFactory.NewBackupInstancesClient().BeginSyncBackupInstance(ctx, rg, vaultName, backupInstance.AzureName(), parameters, &armdataprotection.BackupInstancesClientBeginSyncBackupInstanceOptions{
					ResumeToken: pollerResumeToken,
				})
				if err != nil {
					return extensions.PostReconcileCheckResultFailure("Failed Polling for BeginSyncBackupInstance to get the result"), err
				}

				if pollerResumeToken == "" {
					resumeToken, resumeTokenErr := poller.ResumeToken()
					if resumeTokenErr != nil {
						return extensions.PostReconcileCheckResultFailure("couldn't create PUT resume token for resource"), resumeTokenErr
					} else {
						SetPollerResumeToken(obj, resumeToken, log)
					}
				}

				_, pollErr := poller.Poll(ctx)
				if pollErr != nil {
					return extensions.PostReconcileCheckResultFailure("couldn't create PUT resume token for resource"), pollErr
				}

				if poller.Done() {
					log.V(Debug).Info("Polling is completed")
					ClearPollerResumeToken(obj, log)
					_, err := poller.Result(ctx)
					if err != nil {
						return extensions.PostReconcileCheckResultFailure("couldn't create PUT resume token for resource"), err
					}
				}
				log.V(Debug).Info("Polling is in-progress")
			}
		}
	}
	return extensions.PostReconcileCheckResultFailure("Backup Instance is in non terminal state"), nil
}
