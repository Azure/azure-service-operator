/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	creatorv1 "Telstra.Dx.AzureOperator/api/v1"

	resources "Telstra.Dx.AzureOperator/aztestcreator"
	"Telstra.Dx.AzureOperator/aztestcreator/config"
	eventhubs "Telstra.Dx.AzureOperator/aztestcreator/eventhubs"

	"github.com/go-logr/logr"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// EventhubReconciler reconciles a Eventhub object
type EventhubReconciler struct {
	client.Client
	Log logr.Logger
}

func ignoreNotFound(err error) error {

	if apierrs.IsNotFound(err) {
		return nil
	}
	return err
}

// +kubebuilder:rbac:groups=creator.microsoft.k8.io,resources=eventhubs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=creator.microsoft.k8.io,resources=eventhubs/status,verbs=get;update;patch

//Reconcile blah blah
func (r *EventhubReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("eventhub", req.NamespacedName)

	// your logic here
	var instance creatorv1.Eventhub
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Error(err, "unable to fetch Eventhub")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, ignoreNotFound(err)
	}

	r.createeventhub(&instance)
	return ctrl.Result{}, nil
}

func (r *EventhubReconciler) createeventhub(instance *creatorv1.Eventhub) {
	log := r.Log.WithValues("eventhub", instance)
	ctx := context.Background()
	// create group
	// var err error
	var groupName = instance.ObjectMeta.Name + instance.Spec.ResourceGroupSpec.Name
	var err error
	// err = config.ParseEnvironment()
	err = config.LoadSettings()
	if err != nil {
		log.Error(err, "unable to parse")
	}

	grouplocation := instance.Spec.ResourceGroupSpec.Location

	_, err = resources.CreateGroup(ctx, groupName, grouplocation)
	if err != nil {
		log.Error(err, "unable to create group")
	}

	hublocation := instance.Spec.EventHubResourceSpec.Location

	nsName := instance.ObjectMeta.Name + instance.Spec.EventHubResourceSpec.NSName
	// create Event Hubs namespace
	_, err = eventhubs.CreateNamespace(ctx, groupName, nsName, hublocation)
	if err != nil {
		log.Error(err, "ERROR")
	}

	hubName := instance.ObjectMeta.Name + instance.Spec.EventHubResourceSpec.HubName
	// create Event Hubs hub
	_, err = eventhubs.CreateHub(ctx, groupName, nsName, hubName)
	if err != nil {
		log.Error(err, "ERROR")
	}

}

// SetupWithManager blah
func (r *EventhubReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&creatorv1.Eventhub{}).
		Complete(r)
}
