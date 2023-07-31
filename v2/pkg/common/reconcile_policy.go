// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package common

import "github.com/Azure/azure-service-operator/v2/internal/reconcilers"

// =================== Reconcile Policy Constants ===================

const (
	// ReconcilePolicyManage instructs the operator to manage the resource in question.
	// This includes issuing PUTs to update it and DELETE's to delete it from Azure if deleted in Kuberentes.
	// This is the default policy when no policy is specified.
	ReconcilePolicyManage = reconcilers.ReconcilePolicy("manage")

	// ReconcilePolicySkip instructs the operator to skip all reconciliation actions. This includes creating
	// the resource.
	ReconcilePolicySkip = reconcilers.ReconcilePolicy("skip")

	// ReconcilePolicyDetachOnDelete instructs the operator to skip deletion of resources in Azure. This allows
	// deletion of the resource in Kubernetes to go through but does not delete the underlying Azure resource.
	ReconcilePolicyDetachOnDelete = reconcilers.ReconcilePolicy("detach-on-delete")
)

// ReconcilePolicyAnnotation describes the reconcile policy for the resource in question.
// A reconcile policy describes what action (if any) the operator is allowed to take when
// reconciling the resource.
// If no reconcile policy is specified, the default is "run"
const ReconcilePolicyAnnotation = "serviceoperator.azure.com/reconcile-policy"
