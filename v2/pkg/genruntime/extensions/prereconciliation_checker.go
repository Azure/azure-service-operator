/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package extensions

import (
	"context"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/go-logr/logr"
)

// PreReconciliationChecker is implemented by resources that want to do extra checks before proceeding with
// a full ARM reconcile.
type PreReconciliationChecker interface {
	// PreReconcileCheck does a pre-reconcile check to see if the resource is in a state that can be reconciled.
	// ARM resources should implement this to avoid reconciliation attempts that cannot possibly succeed.
	// Returns ProceedWithReconcile if the reconciliation should go ahead.
	// Returns SkipReconcile and a human-readable reason if the reconciliation should be skipped.
	// ctx is the current operation context.
	// obj is the resource about to be reconciled. The resource's State will be freshly updated.
	// kubeClient allows access to the cluster for any required queries.
	// armClient allows access to ARM for any required queries.
	// log is the logger for the current operation.
	// next is the next (nested) implementation to call.
	PreReconcileCheck(
		ctx context.Context,
		obj genruntime.MetaObject,
		owner genruntime.MetaObject,
		kubeClient kubeclient.Client,
		armClient *genericarmclient.GenericClient,
		log logr.Logger,
		next PreReconcileCheckFunc,
	) (PreReconcileCheckResult, error)
}

type PreReconcileCheckFunc func(
	ctx context.Context,
	obj genruntime.MetaObject,
	owner genruntime.MetaObject,
	kubeClient kubeclient.Client,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
) (PreReconcileCheckResult, error)

type PreReconcileCheckResult struct {
	action preReconcileCheckResultType
	reason string
}

// ProceedWithReconcile returns a PreReconcileCheckResult indicating that the resource is ready for reconciliation.
func ProceedWithReconcile() PreReconcileCheckResult {
	return PreReconcileCheckResult{
		action: preReconcileCheckResultTypeProceed,
	}
}

// SkipReconcile returns a PreReconcileCheckResult indicating that the resource is not ready for reconciliation.
// An explanatory reason is included.
func SkipReconcile(reason string) PreReconcileCheckResult {
	return PreReconcileCheckResult{
		action: preReconcileCheckResultTypeSkip,
		reason: reason,
	}
}

// ShouldReconcile returns true if the provided PreReconcileCheckResult indicates that the resource should be reconciled.
func (r PreReconcileCheckResult) ShouldReconcile() bool {
	return r.action == preReconcileCheckResultTypeProceed
}

// Reason returns the reason for the PreReconcileCheckResult.
func (r PreReconcileCheckResult) Reason() string {
	return r.reason
}

// PreReconcileCheckResultType is the type of result returned by PreReconcileCheck.
type preReconcileCheckResultType string

const (
	preReconcileCheckResultTypeProceed preReconcileCheckResultType = "Proceed"
	preReconcileCheckResultTypeSkip    preReconcileCheckResultType = "Skip"
)

// CreatePreReconciliationChecker creates a checker that can be used to check if a resource is ready for reconciliation.
// If the resource in question has not implemented the PreReconciliationChecker interface, the provided default checker
// is returned directly.
// We also return a bool indicating whether the resource extension implements the PreReconciliationChecker interface.
// host is a resource extension that may implement the PreReconciliationChecker interface.
// checker is the nested checker to use if the resource extension does not implement the PreReconciliationChecker interface.
func CreatePreReconciliationChecker(
	host genruntime.ResourceExtension,
	checker PreReconcileCheckFunc,
) (PreReconcileCheckFunc, bool) {
	impl, ok := host.(PreReconciliationChecker)
	if !ok {
		return checker, false
	}

	return func(
		ctx context.Context,
		obj genruntime.MetaObject,
		owner genruntime.MetaObject,
		kubeClient kubeclient.Client,
		armClient *genericarmclient.GenericClient,
		log logr.Logger,
	) (PreReconcileCheckResult, error) {
		log.V(Status).Info("Extension pre-reconcile check running")

		result, err := impl.PreReconcileCheck(ctx, obj, owner, kubeClient, armClient, log, checker)
		if err != nil {
			log.V(Status).Info(
				"Extension pre-reconcile check failed",
				"Error", err.Error())

			// We choose to skip here so that things are definitely broken and the user will notice
			// If we defaulted to always reconciling, the user might not notice that something is wrong
			return SkipReconcile("Extension PreReconcileCheck failed"), err
		}

		log.V(Status).Info(
			"Extension pre-reconcile check succeeded",
			"Result", result)

		return result, nil
	}, true
}
