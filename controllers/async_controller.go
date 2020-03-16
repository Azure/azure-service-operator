// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	keyvaultsecretlib "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"
	multierror "github.com/hashicorp/go-multierror"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const (
	finalizerName string        = "azure.microsoft.com/finalizer"
	requeDuration time.Duration = time.Second * 20
	successMsg    string        = "successfully provisioned"
)

// AsyncReconciler is a generic reconciler for Azure resources.
// It reconciles Kubernets objects which require long running operations in Azure.
type AsyncReconciler struct {
	client.Client
	AzureClient resourcemanager.ARMClient
	Telemetry   telemetry.TelemetryClient
	Recorder    record.EventRecorder
	Scheme      *runtime.Scheme
}

// Reconcile reconciles the change request
func (r *AsyncReconciler) Reconcile(req ctrl.Request, local runtime.Object) (result ctrl.Result, err error) {
	ctx := context.Background()

	if err := r.Get(ctx, req.NamespacedName, local); err != nil {
		r.Telemetry.LogInfoByInstance("ignorable error", "error during fetch from api server", req.String())
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// get the ASOStatus struct
	status, err := r.AzureClient.GetStatus(local)
	if err != nil {
		return ctrl.Result{}, err
	}

	// record the time that this request was requested at
	if status.RequestedAt == nil {
		timeNow := metav1.NewTime(time.Now())
		status.RequestedAt = &timeNow
	}

	res, err := meta.Accessor(local)
	if err != nil {
		r.Telemetry.LogErrorByInstance("accessor fail", err, req.String())
		return ctrl.Result{}, err
	}

	var keyvaultSecretClient secrets.SecretClient

	// Determine if we need to check KeyVault for secrets
	KeyVaultName := keyvaultsecretlib.GetKeyVaultName(local)

	if len(KeyVaultName) != 0 {
		// Instantiate the KeyVault Secret Client
		keyvaultSecretClient = keyvaultsecretlib.New(KeyVaultName)

		r.Telemetry.LogInfoByInstance("status", "ensuring vault", req.String())

		if !keyvaultsecretlib.IsKeyVaultAccessible(keyvaultSecretClient) {
			r.Telemetry.LogInfoByInstance("requeuing", "awaiting vault verification", req.String())
			return ctrl.Result{RequeueAfter: requeDuration}, nil
		}
	}

	// Check to see if the skipreconcile annotation is on
	var skipReconcile bool
	annotations := res.GetAnnotations()
	if val, ok := annotations["skipreconcile"]; ok {
		if strings.ToLower(val) == "true" {
			skipReconcile = true
		}
	}

	if skipReconcile {
		// if this is a delete we should delete the finalizer to allow the kube instance to be deleted
		if !res.GetDeletionTimestamp().IsZero() {
			if HasFinalizer(res, finalizerName) {
				RemoveFinalizer(res, finalizerName)
			}
		}
		r.Recorder.Event(local, corev1.EventTypeNormal, "Skipping", "Skipping reconcile based on provided annotation")
		return ctrl.Result{}, r.Update(ctx, local)
	}

	var configOptions []resourcemanager.ConfigOption
	if res.GetDeletionTimestamp().IsZero() {
		if !HasFinalizer(res, finalizerName) {
			AddFinalizer(res, finalizerName)
			r.Recorder.Event(local, corev1.EventTypeNormal, "Added", "Object finalizer is added")
			return ctrl.Result{}, r.Update(ctx, local)
		}
	} else {
		if HasFinalizer(res, finalizerName) {
			if len(KeyVaultName) != 0 { //KeyVault was specified in Spec, so use that for secrets
				configOptions = append(configOptions, resourcemanager.WithSecretClient(keyvaultSecretClient))
			}
			found, deleteErr := r.AzureClient.Delete(ctx, local, configOptions...)
			final := multierror.Append(deleteErr)
			if err := final.ErrorOrNil(); err != nil {
				r.Telemetry.LogErrorByInstance("error deleting object", err, req.String())
				r.Recorder.Event(local, corev1.EventTypeWarning, "FailedDelete", fmt.Sprintf("Failed to delete resource: %s", err.Error()))
				return ctrl.Result{}, err
			}
			if !found {
				r.Recorder.Event(local, corev1.EventTypeNormal, "Deleted", "Successfully deleted")
				RemoveFinalizer(res, finalizerName)
				return ctrl.Result{}, r.Update(ctx, local)
			}
			r.Telemetry.LogInfoByInstance("requeuing", "deletion unfinished", req.String())
			return ctrl.Result{RequeueAfter: requeDuration}, r.Status().Update(ctx, local)
		}
		return ctrl.Result{}, nil
	}

	// loop through parents until one is successfully referenced
	parents, err := r.AzureClient.GetParents(local)
	for _, p := range parents {
		if err := r.Get(ctx, p.Key, p.Target); err == nil {
			if pAccessor, err := meta.Accessor(p.Target); err == nil {
				if err := controllerutil.SetControllerReference(pAccessor, res, r.Scheme); err == nil {
					r.Telemetry.LogInfoByInstance("status", "setting parent reference", req.String())
					err := r.Update(ctx, local)
					if err != nil {
						r.Telemetry.LogErrorByInstance("failed to reference parent", err, req.String())
					}
					break
				}
			}
		}
	}

	r.Telemetry.LogInfoByInstance("status", "reconciling object", req.String())

	if len(KeyVaultName) != 0 { //KeyVault was specified in Spec, so use that for secrets
		configOptions = append(configOptions, resourcemanager.WithSecretClient(keyvaultSecretClient))
	}

	done, ensureErr := r.AzureClient.Ensure(ctx, local, configOptions...)
	if ensureErr != nil {
		r.Telemetry.LogErrorByInstance("ensure err", ensureErr, req.String())
	}
	if !done && !status.Provisioning {
		status.RequestedAt = nil
	}
	if done && !status.Provisioned && ensureErr == nil {
		status.FailedProvisioning = true
	}

	// update the status of the resource in kubernetes
	// Implementations of Ensure() tend to set their outcomes in local.Status
	err = r.Status().Update(ctx, local)
	if err != nil {
		r.Telemetry.LogInfoByInstance("status", "failed updating status", req.String())
	}

	final := multierror.Append(ensureErr, r.Update(ctx, local))
	err = final.ErrorOrNil()
	if err != nil {
		r.Recorder.Event(local, corev1.EventTypeWarning, "FailedReconcile", fmt.Sprintf("Failed to reconcile resource: %s", err.Error()))
	} else if done {
		r.Recorder.Event(local, corev1.EventTypeNormal, "Reconciled", "Successfully reconciled")
	}

	result = ctrl.Result{}
	if !done {
		r.Telemetry.LogInfoByInstance("status", "reconciling object not finished", req.String())
		result.RequeueAfter = requeDuration
	} else {
		r.Telemetry.LogInfoByInstance("reconciling", "success", req.String())

		// record the duration of the request
		if status.CompletedAt == nil || status.CompletedAt.IsZero() {
			compTime := metav1.Now()
			status.CompletedAt = &compTime
			if status.RequestedAt == nil {
				r.Telemetry.LogErrorByInstance("Cannot find request time", fmt.Errorf("Request time was nil"), req.String())
			} else {
				durationInSecs := (*status.CompletedAt).Sub((*status.RequestedAt).Time).Seconds()
				r.Telemetry.LogDuration(durationInSecs)
			}
		}
	}

	r.Telemetry.LogInfo("status", "exiting reconciliation")

	return result, err
}
