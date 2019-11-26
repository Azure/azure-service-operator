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
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest/to"

	model "github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"

	eventhubsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"

	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// EventhubReconciler reconciles a Eventhub object
type EventhubReconciler struct {
	client.Client
	Log             logr.Logger
	Recorder        record.EventRecorder
	Scheme          *runtime.Scheme
	EventHubManager eventhubsresourcemanager.EventHubManager
	SecretClient    secrets.SecretClient
}

// +kubebuilder:rbac:groups=azure.microsoft.com,resources=eventhubs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=eventhubs/status,verbs=get;update;patch

// Reconcile function does the main reconciliation loop of the operator
func (r *EventhubReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("eventhub", req.NamespacedName)

	// your logic here
	var instance azurev1alpha1.Eventhub

	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		log.Info("Unable to retrieve eventhub resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if instance.IsBeingDeleted() {
		err := r.handleFinalizer(&instance)
		if err != nil {
			return reconcile.Result{}, fmt.Errorf("error when handling finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	if !instance.HasFinalizer(eventhubFinalizerName) {
		err := r.addFinalizer(&instance)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error when removing finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	requeueAfter, err := strconv.Atoi(os.Getenv("REQUEUE_AFTER"))
	if err != nil {
		requeueAfter = 30
	}

	if !instance.IsSubmitted() {
		err := r.reconcileExternal(&instance)
		if err != nil {
			catch := []string{
				errhelp.ParentNotFoundErrorCode,
				errhelp.ResourceGroupNotFoundErrorCode,
				errhelp.NotFoundErrorCode,
			}
			if azerr, ok := err.(*errhelp.AzureError); ok {
				if helpers.ContainsString(catch, azerr.Type) {
					log.Info("Got ignorable error", "type", azerr.Type)
					return ctrl.Result{Requeue: true, RequeueAfter: time.Duration(requeueAfter) * time.Second}, nil
				}
			}

			return ctrl.Result{}, fmt.Errorf("error when creating resource in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	if instance.IsSubmitted() {
		err := r.reapply(&instance)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error when reapplying resource in azure: %v", err)
		}
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}

// SetupWithManager binds the reconciler to a manager instance
func (r *EventhubReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.Eventhub{}).
		Owns(&v1.Secret{}).
		Complete(r)
}

func (r *EventhubReconciler) reapply(instance *azurev1alpha1.Eventhub) error {
	ctx := context.Background()
	var err error
	eventhubName := instance.ObjectMeta.Name
	eventhubNamespace := instance.Spec.Namespace
	resourcegroup := instance.Spec.ResourceGroup
	secretName := instance.Spec.SecretName

	result, _ := r.EventHubManager.GetHub(ctx, resourcegroup, eventhubNamespace, eventhubName)
	if result.Response.StatusCode == http.StatusNotFound {
		r.reconcileExternal(instance)
		r.Recorder.Event(instance, v1.EventTypeNormal, "Updated", "Resource does not exist in azure, reapplied it")
	}

	if result.Response.StatusCode == http.StatusOK {
		// get secret for eventhub
		key := types.NamespacedName{Name: secretName, Namespace: instance.Namespace}
		_, err = r.SecretClient.Get(ctx, key)
		if err != nil {
			//check if access policy exists, create if it does not and apply secret
			_, err := r.EventHubManager.ListKeys(ctx, resourcegroup, eventhubNamespace, eventhubName, instance.Spec.AuthorizationRule.Name)
			if err != nil {
				err = r.createOrUpdateAccessPolicyEventHub(resourcegroup, eventhubNamespace, eventhubName, instance)
				if err != nil {
					r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to reapply createAccessPolicyEventHub")
					return err
				}
			}

			err = r.listAccessKeysAndCreateSecrets(resourcegroup, eventhubNamespace, eventhubName, secretName, instance.Spec.AuthorizationRule.Name, instance)
			if err != nil {
				r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to reapply listAccessKeysAndCreateSecrets")
				return err
			}
		}
	}

	return nil
}

func (r *EventhubReconciler) reconcileExternal(instance *azurev1alpha1.Eventhub) error {
	ctx := context.Background()

	var err error

	eventhubName := instance.ObjectMeta.Name
	eventhubNamespace := instance.Spec.Namespace
	resourcegroup := instance.Spec.ResourceGroup
	partitionCount := instance.Spec.Properties.PartitionCount
	messageRetentionInDays := instance.Spec.Properties.MessageRetentionInDays
	captureDescription := instance.Spec.Properties.CaptureDescription
	secretName := instance.Spec.SecretName

	if len(secretName) == 0 {
		secretName = eventhubName
		instance.Spec.SecretName = eventhubName
	}

	// write information back to instance
	instance.Status.Provisioning = true

	//get owner instance
	var ownerInstance azurev1alpha1.EventhubNamespace
	eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespace, Namespace: instance.Namespace}

	err = r.Get(ctx, eventhubNamespacedName, &ownerInstance)
	if err != nil {
		//log error and kill it, as the parent might not exist in the cluster. It could have been created elsewhere or through the portal directly
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to get owner instance of eventhubnamespace")
	} else {
		//set owner reference for eventhub if it exists
		references := []metav1.OwnerReference{
			{
				APIVersion: "v1",
				Kind:       "EventhubNamespace",
				Name:       ownerInstance.GetName(),
				UID:        ownerInstance.GetUID(),
			},
		}
		instance.ObjectMeta.SetOwnerReferences(references)
	}

	err = r.Update(ctx, instance)
	if err != nil {
		//log error and kill it
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	capturePtr := getCaptureDescriptionPtr(captureDescription)

	_, err = r.EventHubManager.CreateHub(ctx, resourcegroup, eventhubNamespace, eventhubName, messageRetentionInDays, partitionCount, capturePtr)
	if err != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Couldn't create resource in azure")
		instance.Status.Provisioning = false
		errUpdate := r.Update(ctx, instance)
		if errUpdate != nil {
			//log error and kill it
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
		}
		return errhelp.NewAzureError(err)
	}

	err = r.createOrUpdateAccessPolicyEventHub(resourcegroup, eventhubNamespace, eventhubName, instance)
	if err != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to createAccessPolicyEventHub")
		return err
	}

	err = r.listAccessKeysAndCreateSecrets(resourcegroup, eventhubNamespace, eventhubName, secretName, instance.Spec.AuthorizationRule.Name, instance)
	if err != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to listAccessKeysAndCreateSecrets")
		return err
	}

	// write information back to instance
	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	err = r.Update(ctx, instance)
	if err != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
	}
	return nil
}

const storageAccountResourceFmt = "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s"

func getCaptureDescriptionPtr(captureDescription azurev1alpha1.CaptureDescription) *model.CaptureDescription {
	// add capture details
	var capturePtr *model.CaptureDescription

	storage := captureDescription.Destination.StorageAccount
	storageAccountResourceID := fmt.Sprintf(storageAccountResourceFmt, config.SubscriptionID(), storage.ResourceGroup, storage.AccountName)

	if captureDescription.Enabled {
		capturePtr = &model.CaptureDescription{
			Enabled:           to.BoolPtr(true),
			Encoding:          model.Avro,
			IntervalInSeconds: &captureDescription.IntervalInSeconds,
			SizeLimitInBytes:  &captureDescription.SizeLimitInBytes,
			Destination: &model.Destination{
				Name: &captureDescription.Destination.Name,
				DestinationProperties: &model.DestinationProperties{
					StorageAccountResourceID: &storageAccountResourceID,
					BlobContainer:            &captureDescription.Destination.BlobContainer,
					ArchiveNameFormat:        &captureDescription.Destination.ArchiveNameFormat,
				},
			},
			SkipEmptyArchives: to.BoolPtr(true),
		}
	}
	return capturePtr
}

func (r *EventhubReconciler) deleteEventhub(instance *azurev1alpha1.Eventhub) error {

	ctx := context.Background()

	eventhubName := instance.ObjectMeta.Name
	namespaceName := instance.Spec.Namespace
	resourcegroup := instance.Spec.ResourceGroup

	var err error
	_, err = r.EventHubManager.DeleteHub(ctx, resourcegroup, namespaceName, eventhubName)
	if err != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Couldn't delete resource in azure")
		return err
	}
	return nil
}

func (r *EventhubReconciler) createOrUpdateAccessPolicyEventHub(resourcegroup string, eventhubNamespace string, eventhubName string, instance *azurev1alpha1.Eventhub) error {

	var err error
	ctx := context.Background()

	authorizationRuleName := instance.Spec.AuthorizationRule.Name
	accessRights := make([]model.AccessRights, len(instance.Spec.AuthorizationRule.Rights))
	for i, v := range instance.Spec.AuthorizationRule.Rights {
		accessRights[i] = model.AccessRights(v)
	}
	//accessRights := r.toAccessRights(instance.Spec.AuthorizationRule.Rights)
	parameters := model.AuthorizationRule{
		AuthorizationRuleProperties: &model.AuthorizationRuleProperties{
			Rights: &accessRights,
		},
	}
	_, err = r.EventHubManager.CreateOrUpdateAuthorizationRule(ctx, resourcegroup, eventhubNamespace, eventhubName, authorizationRuleName, parameters)
	if err != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to createorupdateauthorizationrule")
		return err
	}
	return nil
}

func (r *EventhubReconciler) listAccessKeysAndCreateSecrets(resourcegroup string, eventhubNamespace string, eventhubName string, secretName string, authorizationRuleName string, instance *azurev1alpha1.Eventhub) error {

	var err error
	var result model.AccessKeys
	ctx := context.Background()

	result, err = r.EventHubManager.ListKeys(ctx, resourcegroup, eventhubNamespace, eventhubName, authorizationRuleName)
	if err != nil {
		//log error and kill it
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to list keys")
	} else {
		//create secrets in the k8s with the listed keys
		data := map[string][]byte{
			"primaryconnectionstring":   []byte(*result.PrimaryConnectionString),
			"secondaryconnectionstring": []byte(*result.SecondaryConnectionString),
			"primaryKey":                []byte(*result.PrimaryKey),
			"secondaryKey":              []byte(*result.SecondaryKey),
			"sharedaccesskey":           []byte(authorizationRuleName),
			"eventhubnamespace":         []byte(eventhubNamespace),
			"eventhubName":              []byte(eventhubName),
		}
		err = r.createEventhubSecrets(
			ctx,
			secretName,
			instance,
			data,
		)
		if err != nil {
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", fmt.Sprintf("unable to create secret for %s", eventhubName))
			return err
		}
	}
	return nil

}

func (r *EventhubReconciler) createEventhubSecrets(ctx context.Context, secretName string, instance *azurev1alpha1.Eventhub, data map[string][]byte) error {
	key := types.NamespacedName{
		Name:      secretName,
		Namespace: instance.Namespace,
	}

	err := r.SecretClient.Upsert(ctx, key, data)
	if err != nil {
		return err
	}

	// attempt to set owner, gracefully fail
	kSecret := &v1.Secret{}
	if err := r.Get(ctx, key, kSecret); err == nil {
		if err := controllerutil.SetControllerReference(instance, kSecret, r.Scheme); err != nil {
			return err
		}
		return r.Update(ctx, kSecret)
	}

	return nil
}
