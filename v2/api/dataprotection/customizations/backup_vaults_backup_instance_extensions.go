/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"fmt"
	"strings"

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

func (ext *BackupVaultsBackupInstanceExtension) PreReconcileCheck(
	ctx context.Context,
	obj genruntime.MetaObject,
	owner genruntime.MetaObject,
	_ *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
	_ extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {

	log.V(Debug).Info("########################## Starting reconcilation for Backup Instance ##########################")
	fmt.Sprintf("########################## Starting reconcilation for Backup Instance ##########################")

	backupInstance, ok := obj.(*dataprotection.BackupVaultsBackupInstance)

	if !ok {
		return extensions.PreReconcileCheckResult{},
			errors.Errorf("cannot run on unknown resource type %T, expected *dataprotection.BackupVaultsBackupInstance", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = backupInstance

	protectionStatus := strings.ToLower(*backupInstance.Status.Properties.ProtectionStatus.Status)
	protectionStatusErrorCode := ""

	if !strings.Contains(protectionStatus, "protectionerror") {
		log.V(Debug).Info("########################## BlockReconcile protectionerror for Backup Instance ##########################")
		return extensions.BlockReconcile(
				fmt.Sprintf("Backup Instance %q is not in protectionerror state. Provisioning state: %q", owner.GetName(), protectionStatus)),
			nil
	}

	if backupInstance.Status.Properties.ProtectionStatus.ErrorDetails == nil {
		log.V(Debug).Info("########################## BlockReconcile no error for Backup Instance ##########################")
		return extensions.BlockReconcile(
				fmt.Sprintf("Backup Instance %q has no errors %q", owner.GetName(), protectionStatus)),
			nil
	}

	protectionStatusErrorCode = strings.ToLower(*backupInstance.Status.Properties.ProtectionStatus.ErrorDetails.Code)

	if protectionStatusErrorCode == "" || !strings.Contains(protectionStatusErrorCode, "usererror") {
		return extensions.BlockReconcile(
				fmt.Sprintf("Backup Instance %q having Protection Status ErrorCode as %q", owner.GetName(), protectionStatusErrorCode)),
			nil
	}

	id, err := genruntime.GetAndParseResourceID(backupInstance)
	if err != nil {
		return extensions.BlockReconcile(
				fmt.Sprintf("Backup Instance Id %q is not parsed", owner.GetName())),
			nil
	}

	subscription := id.SubscriptionID
	rg := id.ResourceGroupName
	vaultName := id.Parent.Name

	log.V(Debug).Info("########################## Starting NewBackupInstancesClient for Backup Instance ##########################")
	fmt.Sprintf("########################## Starting NewBackupInstancesClient for Backup Instance ##########################")

	clientFactory, err := armdataprotection.NewClientFactory(subscription, armClient.Creds(), armClient.ClientOptions())

	if err != nil {
		log.Error(err, "failed to create armdataprotection client")
	}

	var parameters armdataprotection.SyncBackupInstanceRequest
	parameters.SyncType = to.Ptr(armdataprotection.SyncTypeDefault)

	log.V(Debug).Info("########################## Starting BeginSyncBackupInstance for Backup Instance ##########################")
	fmt.Sprintf("########################## Starting BeginSyncBackupInstance for Backup Instance ##########################")

	poller, err := clientFactory.NewBackupInstancesClient().BeginSyncBackupInstance(ctx, rg, vaultName, backupInstance.AzureName(), parameters, nil)

	_, err = poller.PollUntilDone(ctx, nil)

	if err != nil {
		log.Error(err, "failed BeginSyncBackupInstance to pull the result")
	}

	log.V(Debug).Info("########################## Ending reconcilation for Backup Instance ##########################")
	fmt.Sprintf("########################## Ending reconcilation for Backup Instance ##########################")

	return extensions.ProceedWithReconcile(), nil
}
