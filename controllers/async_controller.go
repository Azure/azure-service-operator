// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"
	"fmt"
	"strings"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	keyvaultsecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"
	kubesecrets "github.com/Azure/azure-service-operator/pkg/secrets/kube"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"
)

const (
	finalizerName    string        = "azure.microsoft.com/finalizer"
	requeueDuration  time.Duration = time.Second * 20
	successMsg       string        = "successfully provisioned"
	reconcileTimeout time.Duration = time.Minute * 5
)

// AsyncReconciler is a generic reconciler for Azure resources.
// It reconciles Kubernetes objects which require long running operations in Azure.
type AsyncReconciler struct {
	client.Client
	ARMFactory resourcemanager.ClientFactory
	Telemetry  telemetry.TelemetryClient
	Recorder   record.EventRecorder
	Scheme     *runtime.Scheme
}

// Reconcile reconciles the change request
func (r *AsyncReconciler) Reconcile(req ctrl.Request, obj runtime.Object) (result ctrl.Result, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), reconcileTimeout)
	defer cancel()

	if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
		r.Telemetry.LogInfoByInstance("ignorable error", "error during fetch from api server", req.String())
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	creds := config.GlobalCredentials()
	keyvaultName := creds.OperatorKeyvault()

	var secretClient secrets.SecretClient
	if keyvaultName == "" {
		r.Telemetry.LogInfoByInstance("status", "keyvault name is empty", req.String())
		secretClient = kubesecrets.New(r.Client, config.SecretNamingVersion())
	} else {
		r.Telemetry.LogInfoByInstance("status", "Instantiating secrets client for keyvault "+keyvaultName, req.String())
		secretClient = keyvaultsecrets.New(keyvaultName, creds, config.SecretNamingVersion())
	}

	armClient := r.ARMFactory(creds, secretClient, r.Scheme)

	// get the ASOStatus struct
	status, err := armClient.GetStatus(obj)
	if err != nil {
		r.Telemetry.LogErrorByInstance("unable to fetch status", err, req.String())
		return ctrl.Result{}, err
	}

	// record the time that this request was requested at
	if status.RequestedAt == nil {
		timeNow := metav1.NewTime(time.Now())
		status.RequestedAt = &timeNow
	}

	res, err := meta.Accessor(obj)
	if err != nil {
		r.Telemetry.LogErrorByInstance("accessor fail", err, req.String())
		return ctrl.Result{}, err
	}

	// TODO(creds-refactor): combine this keyvault lookup with the
	// secret client code above (now that it's done per request). At
	// the moment this is ridiculously fiddly - it's really just a
	// fallback of:
	// 1. keyvault from the object if present
	// 2. keyvault from the credentials if present
	// 3. secret
	// The complication is that if the object keyvault is inaccessible
	// we want to report that on the instance status, but we can't get
	// that until we have created the client to use
	// GetStatus. Untangle this - GetStatus is just a cast and getting
	// the status field from the specific struct. Do this with
	// reflection maybe?
	var keyvaultSecretClient secrets.SecretClient

	// Determine if we need to check KeyVault for secrets
	keyVaultName := keyvaultsecrets.GetKeyVaultName(obj)

	if len(keyVaultName) != 0 {
		// Instantiate the KeyVault Secret Client
		keyvaultSecretClient = keyvaultsecrets.New(keyVaultName, config.GlobalCredentials(), config.SecretNamingVersion())

		r.Telemetry.LogInfoByInstance("status", "ensuring vault", req.String())

		if !keyvaultsecrets.IsKeyVaultAccessible(keyvaultSecretClient) {
			r.Telemetry.LogInfoByInstance("requeuing", "awaiting vault verification", req.String())

			// update the status of the resource in kubernetes
			status.Message = "Waiting for secretclient keyvault to be available"
			return ctrl.Result{RequeueAfter: requeueDuration}, r.Status().Update(ctx, obj)
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
		r.Recorder.Event(obj, corev1.EventTypeNormal, "Skipping", "Skipping reconcile based on provided annotation")
		return ctrl.Result{}, r.Update(ctx, obj)
	}

	var configOptions []resourcemanager.ConfigOption
	if res.GetDeletionTimestamp().IsZero() {
		if !HasFinalizer(res, finalizerName) {
			AddFinalizer(res, finalizerName)
			r.Recorder.Event(obj, corev1.EventTypeNormal, "Added", "Object finalizer is added")
			return ctrl.Result{}, r.Update(ctx, obj)
		}
	} else {
		if HasFinalizer(res, finalizerName) {
			if len(keyVaultName) != 0 { // keyVault was specified in Spec, so use that for secrets
				configOptions = append(configOptions, resourcemanager.WithSecretClient(keyvaultSecretClient))
			}
			found, deleteErr := armClient.Delete(ctx, obj, configOptions...)
			final := multierror.Append(deleteErr)
			if err := final.ErrorOrNil(); err != nil {
				r.Telemetry.LogErrorByInstance("error deleting object", err, req.String())
				r.Recorder.Event(obj, corev1.EventTypeWarning, "FailedDelete", fmt.Sprintf("Failed to delete resource: %s", err.Error()))
				return ctrl.Result{}, err
			}
			if !found {
				r.Recorder.Event(obj, corev1.EventTypeNormal, "Deleted", "Successfully deleted")
				RemoveFinalizer(res, finalizerName)
				return ctrl.Result{}, r.Update(ctx, obj)
			}
			r.Telemetry.LogInfoByInstance("requeuing", "deletion unfinished", req.String())
			return ctrl.Result{RequeueAfter: requeueDuration}, r.Status().Update(ctx, obj)
		}
		return ctrl.Result{}, nil
	}

	// loop through parents until one is successfully referenced
	parents, err := armClient.GetParents(obj)
	for _, p := range parents {
		if err := r.Get(ctx, p.Key, p.Target); err == nil {
			if pAccessor, err := meta.Accessor(p.Target); err == nil {
				if err := controllerutil.SetControllerReference(pAccessor, res, r.Scheme); err == nil {
					r.Telemetry.LogInfoByInstance("status", "setting parent reference", req.String())
					err := r.Update(ctx, obj)
					if err != nil {
						r.Telemetry.LogErrorByInstance("failed to reference parent", err, req.String())
					}
					break
				}
			}
		}
	}

	r.Telemetry.LogInfoByInstance("status", "reconciling object", req.String())

	if len(keyVaultName) != 0 { //KeyVault was specified in Spec, so use that for secrets
		configOptions = append(configOptions, resourcemanager.WithSecretClient(keyvaultSecretClient))
	}

	done, ensureErr := armClient.Ensure(ctx, obj, configOptions...)
	if ensureErr != nil {
		r.Telemetry.LogErrorByInstance("ensure err", ensureErr, req.String())
	}
	if !done && !status.Provisioning {
		status.RequestedAt = nil
	}
	if done && !status.Provisioned && ensureErr == nil {
		status.SetFailedProvisioning(status.Message) // Keep the same message
	}

	// update the status of the resource in kubernetes
	// Implementations of Ensure() tend to set their outcomes in obj.Status
	err = r.Status().Update(ctx, obj)
	if err != nil {
		r.Telemetry.LogInfoByInstance("status", "failed updating status", req.String())
	}

	final := multierror.Append(ensureErr, r.Update(ctx, obj))
	err = final.ErrorOrNil()
	if err != nil {
		r.Recorder.Event(obj, corev1.EventTypeWarning, "FailedReconcile", fmt.Sprintf("Failed to reconcile resource: %s", err.Error()))
	} else if done {
		r.Recorder.Event(obj, corev1.EventTypeNormal, "Reconciled", "Successfully reconciled")
	}

	result = ctrl.Result{}
	if !done {
		r.Telemetry.LogInfoByInstance("status", "reconciling object not finished", req.String())
		result.RequeueAfter = requeueDuration
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

	r.Telemetry.LogInfoByInstance("status", "exiting reconciliation", req.String())

	return result, err
}
