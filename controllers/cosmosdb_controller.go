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
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/cosmosdbs"
	"k8s.io/client-go/tools/record"
)

const cosmosDBFinalizerName = "cosmosdb.finalizers.azure.com"

// CosmosDBReconciler reconciles a CosmosDB object
type CosmosDBReconciler struct {
	client.Client
	Log         logr.Logger
	Recorder    record.EventRecorder
	RequeueTime time.Duration
}

// +kubebuilder:rbac:groups=service.azure,resources=cosmosdbs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=service.azure,resources=cosmosdbs/status,verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *CosmosDBReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("cosmosdb", req.NamespacedName)

	// Fetch the CosmosDB instance
	var instance azurev1.CosmosDB

	requeueAfter, err := strconv.Atoi(os.Getenv("REQUEUE_AFTER"))
	if err != nil {
		requeueAfter = 30
	}

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Error(err, "unable to fetch CosmosDB")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	log.Info("Getting CosmosDB Account", "CosmosDB.Namespace", instance.Namespace, "CosmosDB.Name", instance.Name)
	log.V(1).Info("Describing CosmosDB Account", "CosmosDB", instance)	

	if helpers.IsBeingDeleted(&instance) {
		if helpers.HasFinalizer(&instance, cosmosDBFinalizerName) {
			if err := r.deleteExternal(&instance); err != nil {
				log.Info("Delete CosmosDB failed with ", err.Error())
				return ctrl.Result{}, err
			}

			helpers.RemoveFinalizer(&instance, cosmosDBFinalizerName)
			if err := r.Update(context.Background(), &instance); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !helpers.HasFinalizer(&instance, cosmosDBFinalizerName) {
		if err := r.addFinalizer(&instance); err != nil {
			log.Info("Adding cosmosDB finalizer failed with ", err.Error())
			return ctrl.Result{}, err
		}
	}

	if !instance.IsSubmitted() {
		if err := r.reconcileExternal(&instance); err != nil {
			if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
				log.Info("Requeuing as the async operation is not complete")
				return ctrl.Result{
					Requeue:      true,
					RequeueAfter: time.Second * time.Duration(requeueAfter),
				}, nil
			}
			return ctrl.Result{}, fmt.Errorf("error reconciling cosmosdb in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	r.Recorder.Event(&instance, "Normal", "Provisioned", "CosmosDB "+instance.ObjectMeta.Name+" provisioned ")
	return ctrl.Result{}, nil
}

func (r *CosmosDBReconciler) addFinalizer(instance *azurev1.CosmosDB) error {
	helpers.AddFinalizer(instance, cosmosDBFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update finalizer: %v", err)
	}
	r.Recorder.Event(instance, "Normal", "Updated", fmt.Sprintf("finalizer %s added", cosmosDBFinalizerName))
	return nil
}

func (r *CosmosDBReconciler) reconcileExternal(instance *azurev1.CosmosDB) error {
	ctx := context.Background()
	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	kind := instance.Spec.Kind
	dbType := instance.Spec.Properties.DatabaseAccountOfferType

	// write information back to instance
	instance.Status.Provisioning = true

	if err := r.Status().Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	_, err := cosmosdbs.CreateCosmosDB(ctx, groupName, name, location, kind, dbType, nil)
	if err != nil {
		if errhelp.IsAsynchronousOperationNotComplete(err) || errhelp.IsGroupNotFound(err) {
			r.Recorder.Event(instance, "Normal", "Provisioning", name+" provisioning")
			return err
		}
		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't create resource in azure")
		instance.Status.Provisioning = false
		errUpdate := r.Status().Update(ctx, instance)
		if errUpdate != nil {
			r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
		}
		return err
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	if err = r.Status().Update(ctx, instance); err != nil {
		r.Recorder.Event(instance, "Warning", "Failed", "Unable to update instance")
	}

	return nil
}

func (r *CosmosDBReconciler) deleteExternal(instance *azurev1.CosmosDB) error {
	ctx := context.Background()
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	_, err := cosmosdbs.DeleteCosmosDB(ctx, groupName, name)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			r.Recorder.Event(instance, "Warning", "DoesNotExist", "Resource to delete does not exist")
			return nil
		}

		r.Recorder.Event(instance, "Warning", "Failed", "Couldn't delete resouce in azure")
		return err
	}

	r.Recorder.Event(instance, "Normal", "Deleted", name+" deleted")
	return nil
}

// SetupWithManager sets up the controller functions
func (r *CosmosDBReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1.CosmosDB{}).
		Complete(r)
}

/* Below code was from prior to refactor.  
   Left here for future reference for pulling out values post deployment.
   
   
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
}*/
