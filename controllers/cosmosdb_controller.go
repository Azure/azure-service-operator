/*
MIT License

Copyright (c) Microsoft Corporation. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE
*/

package controllers

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	servicev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/client/deployment"
	"github.com/Azure/azure-service-operator/pkg/client/group"
	"github.com/Azure/azure-service-operator/pkg/config"
	cosmosdbtemplate "github.com/Azure/azure-service-operator/pkg/cosmosdb"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/to"
)

// CosmosDBReconciler reconciles a CosmosDB object
type CosmosDBReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=service.azure,resources=cosmosdbs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=service.azure,resources=cosmosdbs/status,verbs=get;update;patch

func (r *CosmosDBReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("cosmosdb", req.NamespacedName)

	// Fetch the CosmosDB instance
	instance := &servicev1alpha1.CosmosDB{}
	err := r.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		log.Error(err, "unable to fetch CosmosDB")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	log.Info("Getting CosmosDB Account", "CosmosDB.Namespace", instance.Namespace, "CosmosDB.Name", instance.Name)
	log.V(1).Info("Describing CosmosDB Account", "CosmosDB", instance)

	cosmosDBFinalizerName := "cosmosdb.finalizers.azure"
	// examine DeletionTimestamp to determine if object is under deletion
	if instance.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent
		// registering our finalizer.
		if !helpers.ContainsString(instance.ObjectMeta.Finalizers, cosmosDBFinalizerName) {
			instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, cosmosDBFinalizerName)
			if err := r.Update(ctx, instance); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else {
		// The object is being deleted
		if helpers.ContainsString(instance.ObjectMeta.Finalizers, cosmosDBFinalizerName) {
			// our finalizer is present, so lets handle any external dependency
			if err := r.deleteExternalResources(instance); err != nil {
				// if fail to delete the external dependency here, return with error
				// so that it can be retried
				return ctrl.Result{}, err
			}

			// remove our finalizer from the list and update it.
			instance.ObjectMeta.Finalizers = helpers.RemoveString(instance.ObjectMeta.Finalizers, cosmosDBFinalizerName)
			if err := r.Update(ctx, instance); err != nil {
				return ctrl.Result{}, err
			}
		}

		return ctrl.Result{}, err
	}

	resourceGroupName := helpers.AzrueResourceGroupName(config.Instance.SubscriptionID, config.Instance.ClusterName, "cosmosdb", instance.Name, instance.Namespace)
	deploymentName := instance.Status.DeploymentName
	if deploymentName != "" {
		log.Info("Checking deployment", "ResourceGroupName", resourceGroupName, "DeploymentName", deploymentName)
		de, _ := deployment.GetDeployment(ctx, resourceGroupName, deploymentName)
		provisioningState := *de.Properties.ProvisioningState
		if helpers.IsDeploymentComplete(provisioningState) {
			log.Info("Deployment is complete", "ProvisioningState", provisioningState)
			_, err = r.updateStatus(req, resourceGroupName, deploymentName, provisioningState, de.Properties.Outputs)
			if err != nil {
				return ctrl.Result{}, err
			}
			return ctrl.Result{}, nil
		} else {
			log.Info("Requeue the request", "ProvisioningState", provisioningState)
			return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
		}
	}

	log.Info("Creating a new resource group", "ResourceGroupName", resourceGroupName)
	tags := map[string]*string{
		"name":      to.StringPtr(instance.Name),
		"namespace": to.StringPtr(instance.Namespace),
		"kind":      to.StringPtr("cosmosdb"),
	}
	group.CreateGroup(ctx, resourceGroupName, instance.Spec.Location, tags)

	log.Info("Reconciling CosmosDB", "CosmosDB.Namespace", instance.Namespace, "CosmosDB.Name", instance.Name)
	template := cosmosdbtemplate.New(instance)
	deploymentName, err = template.CreateDeployment(ctx, resourceGroupName)
	if err != nil {
		log.Error(err, "Failed to reconcile CosmosDB")
		return ctrl.Result{}, err
	}

	de, _ := deployment.GetDeployment(ctx, resourceGroupName, deploymentName)
	_, err = r.updateStatus(req, resourceGroupName, deploymentName, *de.Properties.ProvisioningState, nil)
	if err != nil {
		return ctrl.Result{}, err
	}

	// CosmosDB created successfully - don't requeue
	return ctrl.Result{}, nil
}

func (r *CosmosDBReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&servicev1alpha1.CosmosDB{}).
		Complete(r)
}

func (r *CosmosDBReconciler) updateStatus(req ctrl.Request, resourceGroupName, deploymentName, provisioningState string, outputs interface{}) (*servicev1alpha1.CosmosDB, error) {
	ctx := context.Background()
	log := r.Log.WithValues("cosmosdb", req.NamespacedName)

	resource := &servicev1alpha1.CosmosDB{}
	r.Get(ctx, req.NamespacedName, resource)
	log.Info("Getting CosmosDB Account", "CosmosDB.Namespace", resource.Namespace, "CosmosDB.Name", resource.Name)

	resourceCopy := resource.DeepCopy()
	resourceCopy.Status.DeploymentName = deploymentName
	resourceCopy.Status.ProvisioningState = provisioningState

	err := r.Status().Update(ctx, resourceCopy)
	if err != nil {
		log.Error(err, "unable to update CosmosDB status")
		return nil, err
	}
	log.V(1).Info("Updated Status", "CosmosDB.Namespace", resourceCopy.Namespace, "CosmosDB.Name", resourceCopy.Name, "CosmosDB.Status", resourceCopy.Status)

	if helpers.IsDeploymentComplete(provisioningState) {
		if outputs != nil {
			resourceCopy.Output.CosmosDBName = helpers.GetOutput(outputs, "cosmosDBName")
			resourceCopy.Output.PrimaryMasterKey = helpers.GetOutput(outputs, "primaryMasterKey")
		}

		err := r.syncAdditionalResourcesAndOutput(req, resourceCopy)
		if err != nil {
			log.Error(err, "error syncing resources")
			return nil, err
		}
		log.V(1).Info("Updated additional resources", "CosmosDB.Namespace", resourceCopy.Namespace, "CosmosDB.Name", resourceCopy.Name, "CosmosDB.AdditionalResources", resourceCopy.AdditionalResources, "CosmosDB.Output", resourceCopy.Output)
	}

	return resourceCopy, nil
}

func (r *CosmosDBReconciler) deleteExternalResources(instance *servicev1alpha1.CosmosDB) error {
	//
	// delete any external resources associated with the cosmosdb
	//
	// Ensure that delete implementation is idempotent and safe to invoke
	// multiple types for same object.
	ctx := context.Background()
	log := r.Log.WithValues("CosmosDB.Namespace", instance.Namespace, "CosmosDB.Name", instance.Name)

	resourceGroupName := helpers.AzrueResourceGroupName(config.Instance.SubscriptionID, config.Instance.ClusterName, "cosmosdb", instance.Name, instance.Namespace)
	log.Info("Deleting CosmosDB Account", "ResourceGroupName", resourceGroupName)
	_, err := group.DeleteGroup(ctx, resourceGroupName)
	if err != nil && helpers.IgnoreAzureResourceNotFound(err) != nil {
		return err
	}

	err = helpers.DeleteSecret(instance.Name, instance.Namespace)
	if err != nil && helpers.IgnoreKubernetesResourceNotFound(err) != nil {
		return err
	}

	return nil
}

func (r *CosmosDBReconciler) syncAdditionalResourcesAndOutput(req ctrl.Request, s *servicev1alpha1.CosmosDB) (err error) {
	ctx := context.Background()
	log := r.Log.WithValues("cosmosdb", req.NamespacedName)

	secrets := []string{}
	secretData := map[string]string{
		"cosmosDBName":     "{{.Obj.Output.CosmosDBName}}",
		"primaryMasterKey": "{{.Obj.Output.PrimaryMasterKey}}",
	}
	secret := helpers.CreateSecret(s, s.Name, s.Namespace, secretData)
	secrets = append(secrets, secret)

	resourceCopy := s.DeepCopy()
	resourceCopy.AdditionalResources.Secrets = secrets

	err = r.Update(ctx, resourceCopy)
	if err != nil {
		log.Error(err, "unable to update CosmosDB status")
		return err
	}

	return nil
}
