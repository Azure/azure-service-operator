// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	"context"
	"strings"
	"time"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"

	armstorage "github.com/Azure/ARO-HCP/test/sdk/resourcemanager/redhatopenshifthcp/armredhatopenshifthcp"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/Azure/azure-service-operator/v2/api/redhatopenshift/v1api20240610preview/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

var _ extensions.PreReconciliationChecker = &HcpOpenShiftClusterExtension{}

// PreReconcileCheck does a pre-reconcile check to see if the resource is in a state that can be reconciled.
// ARM resources should implement this to avoid reconciliation attempts that cannot possibly succeed.
// Returns ProceedWithReconcile if the reconciliation should go ahead.
// Returns BlockReconcile and a human-readable reason if the reconciliation should be skipped.
// ctx is the current operation context.
// obj is the resource about to be reconciled. The resource's State will be freshly updated.
// owner is the owner of the resource being reconciled.
// resourceResolver allows resolving references to other resources.
// armClient allows access to ARM for any required queries.
// log is the logger for the current operation.
// next is the next (nested) implementation to call.
func (ext *HcpOpenShiftClusterExtension) PreReconcileCheck(ctx context.Context,
	obj genruntime.MetaObject,
	owner genruntime.MetaObject,
	resourceResolver *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
	next extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
	// This has to be the current hub storage version of the hcpOpenShiftCluster.
	// It will need to be updated if the hub storage version changes.
	hcpOpenShiftCluster, ok := obj.(*storage.HcpOpenShiftCluster)
	if !ok {
		return extensions.PreReconcileCheckResult{}, eris.Errorf("cannot run on unknown resource type %T, expected *storage.HcpOpenShiftCluster", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = hcpOpenShiftCluster

	// If the hcpOpenShiftCluster is already deleting, we have to wait for that to finish
	// before trying anything else
	if hcpOpenShiftCluster.Status.Properties != nil &&
		hcpOpenShiftCluster.Status.Properties.ProvisioningState != nil &&
		strings.EqualFold(*hcpOpenShiftCluster.Status.Properties.ProvisioningState, "Deleting") {
		return extensions.BlockReconcile("reconcile blocked while hcpOpenShiftCluster is at status deleting"), nil
	}

	return next(ctx, obj, owner, resourceResolver, armClient, log)
}

var _ genruntime.KubernetesSecretExporter = &HcpOpenShiftClusterExtension{}

func (ext *HcpOpenShiftClusterExtension) ExportKubernetesSecrets(
	ctx context.Context,
	obj genruntime.MetaObject,
	additionalSecrets set.Set[string],
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
) (*genruntime.KubernetesSecretExportResult, error) {
	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	typedObj, ok := obj.(*storage.HcpOpenShiftCluster)
	if !ok {
		return nil, eris.Errorf("cannot run on unknown resource type %T, expected *storage.HcpOpenShiftCluster", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = typedObj

	primarySecrets := secretsSpecifiedHcp(typedObj)
	requestedSecrets := set.Union(primarySecrets, additionalSecrets)

	if len(requestedSecrets) == 0 {
		log.V(Debug).Info("No secrets retrieval to perform as operatorSpec is empty")
		return nil, nil
	}

	id, err := genruntime.GetAndParseResourceID(typedObj)
	if err != nil {
		return nil, err
	}

	subscription := id.SubscriptionID
	// Using armClient.ClientOptions() here ensures we share the same HTTP connection, so this is not opening a new
	// connection each time through
	var clusterClient *armstorage.HcpOpenShiftClustersClient
	clusterClient, err = armstorage.NewHcpOpenShiftClustersClient(subscription, armClient.Creds(), armClient.ClientOptions())
	if err != nil {
		return nil, eris.Wrapf(err, "failed to create new NewOpenShiftClustersClient")
	}

	var adminCredentials string
	if requestedSecrets.Contains(adminCredentialsKey) {
		log.V(Debug).Info("Starting BeginRequestAdminCredential")
		var poller *runtime.Poller[armstorage.HcpOpenShiftClustersClientRequestAdminCredentialResponse]
		poller, err = clusterClient.BeginRequestAdminCredential(ctx, id.ResourceGroupName, typedObj.AzureName(), nil)
		if err != nil {
			return nil, eris.Wrapf(err, "failed creating admin credentials")
		}

		log.V(Debug).Info("Waiting for admin credential request to complete")
		pollCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
		defer cancel()
		resp, pollErr := poller.PollUntilDone(pollCtx, &runtime.PollUntilDoneOptions{
			Frequency: 15 * time.Second,
		})
		if pollErr != nil {
			if ctx.Err() != nil {
				return nil, eris.Wrapf(pollErr, "parent context cancelled while waiting for admin credentials")
			}
			if pollCtx.Err() == context.DeadlineExceeded {
				return nil, eris.Wrapf(pollErr, "timed out after 5 minutes waiting for admin credentials to be ready")
			}
			return nil, eris.Wrapf(pollErr, "failed waiting for admin credentials to be ready")
		}

		log.V(Debug).Info("Admin credential request completed")
		adminCredentials = to.Value(resp.HcpOpenShiftClusterAdminCredential.Kubeconfig)
		if adminCredentials == "" {
			return nil, eris.Errorf(
				"admin credential response for cluster %s in resource group %s contained an empty kubeconfig",
				typedObj.AzureName(), id.ResourceGroupName)
		}
	}

	secretSlice, err := secretsToWriteHcp(typedObj, adminCredentials)
	if err != nil {
		return nil, err
	}

	resolvedSecrets := map[string]string{}
	if adminCredentials != "" {
		resolvedSecrets[adminCredentialsKey] = adminCredentials
	}
	return &genruntime.KubernetesSecretExportResult{
		Objs:       secrets.SliceToClientObjectSlice(secretSlice),
		RawSecrets: secrets.SelectSecrets(additionalSecrets, resolvedSecrets),
	}, nil
}

func secretsSpecifiedHcp(obj *storage.HcpOpenShiftCluster) set.Set[string] {
	if obj.Spec.OperatorSpec == nil || obj.Spec.OperatorSpec.Secrets == nil {
		return nil
	}

	operatorSecrets := obj.Spec.OperatorSpec.Secrets
	result := set.Set[string]{}
	if operatorSecrets.AdminCredentials != nil {
		result.Add(adminCredentialsKey)
	}

	return result
}

func secretsToWriteHcp(obj *storage.HcpOpenShiftCluster, adminCredentials string) ([]*v1.Secret, error) {
	operatorSpecSecrets := obj.Spec.OperatorSpec.Secrets
	if operatorSpecSecrets == nil {
		return nil, nil
	}

	collector := secrets.NewCollector(obj.Namespace)
	collector.AddValue(operatorSpecSecrets.AdminCredentials, adminCredentials)

	return collector.Values()
}
