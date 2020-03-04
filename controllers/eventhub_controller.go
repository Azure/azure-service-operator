// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"github.com/Azure/azure-service-operator/pkg/secrets"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	eventhubsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"

	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// EventhubReconciler reconciles a Eventhub object
type EventhubReconciler struct {
	client.Client
	Log             logr.Logger
	Recorder        record.EventRecorder
	Scheme          *runtime.Scheme
	EventHubManager eventhubsresourcemanager.EventHubManager
	SecretClient    secrets.SecretClient
	Reconciler      *AsyncReconciler
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=eventhubs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=eventhubs/status,verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *EventhubReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(req, &azurev1alpha1.Eventhub{})
}

// SetupWithManager binds the reconciler to a manager instance
func (r *EventhubReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.Eventhub{}).
		Owns(&v1.Secret{}).
		Complete(r)
}
