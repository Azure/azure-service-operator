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
	Telemetry   telemetry.PrometheusTelemetry
	Recorder    record.EventRecorder
	Scheme      *runtime.Scheme
}

func (r *AsyncReconciler) Reconcile(req ctrl.Request, local runtime.Object) (result ctrl.Result, err error) {
	ctx := context.Background()

	// // log operator start
	// r.Telemetry.LogStart()

	// // log failure / success
	// defer func() {
	// 	if err != nil {
	// 		r.Telemetry.LogError(
	// 			"Failure occured during reconcilliation",
	// 			err)
	// 		r.Telemetry.LogFailure()
	// 	} else if result.Requeue {
	// 		r.Telemetry.LogFailure()
	// 	} else {
	// 		r.Telemetry.LogSuccess()
	// 	}
	// }()

	if err := r.Get(ctx, req.NamespacedName, local); err != nil {
		r.Telemetry.LogInfo("ignorable error", "error during fetch from api server")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	res, err := meta.Accessor(local)
	if err != nil {
		r.Telemetry.LogError("accessor fail", err)
		return ctrl.Result{}, err
	}

	// Instantiate the KeyVault Secret Client if KeyVault specified in Spec
	r.Telemetry.LogInfo("status", "retrieving keyvault for secrets if specified")
	var keyvaultSecretClient secrets.SecretClient
	KeyVaultName := keyvaultsecretlib.GetKeyVaultName(local)
	if len(KeyVaultName) != 0 {
		keyvaultSecretClient = keyvaultsecretlib.New(KeyVaultName)
		if !keyvaultsecretlib.IsKeyVaultAccessible(keyvaultSecretClient) {
			r.Telemetry.LogInfo("requeuing", "Waiting for Keyvault to store secrets to be available")
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
			var found bool
			var deleteErr error
			configOptions = append(configOptions, resourcemanager.WithKubeClient(r.Client))
			if len(KeyVaultName) != 0 { //KeyVault was specified in Spec, so use that for secrets
				configOptions = append(configOptions, resourcemanager.WithSecretClient(keyvaultSecretClient))
			}

			found, deleteErr = r.AzureClient.Delete(ctx, local, configOptions...)
			final := multierror.Append(deleteErr)
			if err := final.ErrorOrNil(); err != nil {
				r.Telemetry.LogError("error deleting object", err)
				r.Recorder.Event(local, corev1.EventTypeWarning, "FailedDelete", fmt.Sprintf("Failed to delete resource: %s", err.Error()))
				return ctrl.Result{}, err
			}
			if !found {
				r.Recorder.Event(local, corev1.EventTypeNormal, "Deleted", "Successfully deleted")
				RemoveFinalizer(res, finalizerName)
				return ctrl.Result{}, r.Update(ctx, local)
			}
			r.Telemetry.LogInfo("requeuing", "deletion unfinished")
			return ctrl.Result{RequeueAfter: requeDuration}, nil
		}
		return ctrl.Result{}, nil
	}

	// loop through parents until one is successfully referenced
	parents, err := r.AzureClient.GetParents(local)
	for _, p := range parents {
		//r.Telemetry.LogInfo("status", "handling parent "+p.Key.Name)

		if err := r.Get(ctx, p.Key, p.Target); err == nil {
			//r.Telemetry.LogInfo("status", "handling parent get for "+reflect.TypeOf(p.Target).String())

			if pAccessor, err := meta.Accessor(p.Target); err == nil {
				if err := controllerutil.SetControllerReference(pAccessor, res, r.Scheme); err == nil {
					r.Telemetry.LogInfo("status", "setting parent reference to object: "+pAccessor.GetName())
					err := r.Update(ctx, local)
					if err != nil {
						r.Telemetry.LogInfo("warning", "failed to update instance: "+err.Error())
					}
					break
				}
			}
		}
	}

	r.Telemetry.LogInfo("status", "reconciling object")

	configOptions = append(configOptions, resourcemanager.WithKubeClient(r.Client))
	if len(KeyVaultName) != 0 { //KeyVault was specified in Spec, so use that for secrets
		configOptions = append(configOptions, resourcemanager.WithSecretClient(keyvaultSecretClient))
	}

	done, ensureErr := r.AzureClient.Ensure(ctx, local, configOptions...)
	if ensureErr != nil {
		r.Telemetry.LogError("ensure err", ensureErr)
	}

	// update the status of the resource in kubernetes
	// Implementations of Ensure() tend to set their outcomes in local.Status
	err = r.Status().Update(ctx, local)
	if err != nil {
		r.Telemetry.LogInfo("status", "failed updating status")
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
		r.Telemetry.LogInfo("status", "reconciling object not finished")
		result.RequeueAfter = requeDuration
	}

	r.Telemetry.LogInfo("status", "exiting reconciliation")

	return result, err
}
