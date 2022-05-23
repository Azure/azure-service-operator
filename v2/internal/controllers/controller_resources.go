/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	ctrlconversion "sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	networkstorage "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
	resourcesalpha "github.com/Azure/azure-service-operator/v2/api/resources/v1alpha1api20200601"
	resourcesbeta "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/registration"
)

func GetKnownStorageTypes(
	mgr ctrl.Manager,
	armClientFactory arm.ARMClientFactory,
	kubeClient kubeclient.Client,
	positiveConditions *conditions.PositiveConditionBuilder,
	options Options) ([]*registration.StorageType, error) {

	resourceResolver := resolver.NewResolver(kubeClient)
	knownStorageTypes := getKnownStorageTypes()

	knownStorageTypes = append(
		knownStorageTypes,
		registration.NewStorageType(&resourcesbeta.ResourceGroup{}))

	// Verify we're using the hub version of VirtualNetworksSubnet in the loop below
	var _ ctrlconversion.Hub = &networkstorage.VirtualNetworksSubnet{}

	// TODO: Modifying this list would be easier if it were a map
	for _, t := range knownStorageTypes {
		if _, ok := t.Obj.(*networkstorage.VirtualNetworksSubnet); ok {
			t.Indexes = append(t.Indexes, registration.Index{
				Key:  ".metadata.ownerReferences[0]",
				Func: indexOwner,
			})
		}
		if _, ok := t.Obj.(*networkstorage.RouteTablesRoute); ok {
			t.Indexes = append(t.Indexes, registration.Index{
				Key:  ".metadata.ownerReferences[0]",
				Func: indexOwner,
			})
		}
	}

	err := resourceResolver.IndexStorageTypes(mgr.GetScheme(), knownStorageTypes)
	if err != nil {
		return nil, errors.Wrap(err, "failed add storage types to resource resolver")
	}

	var extensions map[schema.GroupVersionKind]genruntime.ResourceExtension
	extensions, err = GetResourceExtensions(mgr.GetScheme())
	if err != nil {
		return nil, errors.Wrap(err, "failed getting extensions")
	}

	for _, t := range knownStorageTypes {
		// Use the provided GVK to construct a new runtime object of the desired concrete type.
		var gvk schema.GroupVersionKind
		gvk, err = apiutil.GVKForObject(t.Obj, mgr.GetScheme())
		if err != nil {
			return nil, errors.Wrapf(err, "creating GVK for obj %T", t.Obj)
		}
		extension := extensions[gvk]

		err = augmentWithARMReconciler(
			armClientFactory,
			kubeClient,
			resourceResolver,
			positiveConditions,
			options,
			extension,
			t)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't create reconciler")
		}
	}

	for _, t := range knownStorageTypes {
		err = augmentWithControllerName(t)
		if err != nil {
			return nil, err
		}
	}

	return knownStorageTypes, nil
}

func augmentWithARMReconciler(
	armClientFactory arm.ARMClientFactory,
	kubeClient kubeclient.Client,
	resourceResolver *resolver.Resolver,
	positiveConditions *conditions.PositiveConditionBuilder,
	options Options,
	extension genruntime.ResourceExtension,
	t *registration.StorageType) error {
	t.Reconciler = arm.NewAzureDeploymentReconciler(
		armClientFactory,
		kubeClient,
		resourceResolver,
		positiveConditions,
		options.Config,
		extension)

	return nil
}

func augmentWithControllerName(t *registration.StorageType) error {
	controllerName, err := getControllerName(t.Obj)
	if err != nil {
		return errors.Wrapf(err, "failed to get controller name for obj %T", t.Obj)
	}

	t.Name = controllerName

	return nil
}

func getControllerName(obj client.Object) (string, error) {
	v, err := conversion.EnforcePtr(obj)
	if err != nil {
		return "", errors.Wrap(err, "t.Obj was expected to be ptr but was not")
	}

	typ := v.Type()
	return fmt.Sprintf("%sController", typ.Name()), nil
}

func indexOwner(rawObj client.Object) []string {
	owners := rawObj.GetOwnerReferences()
	if len(owners) == 0 {
		return nil
	}

	// Only works for 1 owner now but that's fine
	return []string{string(owners[0].UID)}
}

func GetKnownTypes() []client.Object {
	knownTypes := getKnownTypes()

	knownTypes = append(
		knownTypes,
		&resourcesalpha.ResourceGroup{},
		&resourcesbeta.ResourceGroup{})

	return knownTypes
}

func CreateScheme() *runtime.Scheme {
	scheme := createScheme()
	_ = resourcesalpha.AddToScheme(scheme)
	_ = resourcesbeta.AddToScheme(scheme)

	return scheme
}

// GetResourceExtensions returns a map between resource and resource extension
func GetResourceExtensions(scheme *runtime.Scheme) (map[schema.GroupVersionKind]genruntime.ResourceExtension, error) {

	extensionMapping := make(map[schema.GroupVersionKind]genruntime.ResourceExtension)

	for _, extension := range getResourceExtensions() {
		for _, resource := range extension.GetExtendedResources() {

			// Make sure the type casting goes well, and we can extract the GVK successfully.
			resourceObj, ok := resource.(runtime.Object)
			if !ok {
				err := errors.Errorf("unexpected resource type for resource '%s', found '%T'", resource.AzureName(), resource)
				return nil, err
			}

			gvk, err := apiutil.GVKForObject(resourceObj, scheme)
			if err != nil {
				return nil, err
			}
			extensionMapping[gvk] = extension
		}
	}

	return extensionMapping, nil
}

// watchSecretsFactory is used to register an EventHandlerFactory for watching secret mutations.
func watchSecretsFactory(keys []string, objList client.ObjectList) registration.EventHandlerFactory {
	return func(client client.Client, log logr.Logger) handler.EventHandler {
		return handler.EnqueueRequestsFromMapFunc(watchSecrets(client, log, keys, objList))
	}
}

// TODO: It may be possible where we construct the ctrl.Manager to limit what secrets we watch with
// TODO: this feature: https://github.com/kubernetes-sigs/controller-runtime/blob/master/designs/use-selectors-at-cache.md,
// TODO: likely scoped by a label selector?
func watchSecrets(c client.Client, log logr.Logger, keys []string, objList client.ObjectList) handler.MapFunc {
	listType := reflect.TypeOf(objList).Elem()

	return func(o client.Object) []reconcile.Request {
		// Safety check that we're looking at a secret
		if _, ok := o.(*corev1.Secret); !ok {
			log.V(Status).Info("Unexpected non-secret", "namespace", o.GetNamespace(), "name", o.GetName(), "actual", fmt.Sprintf("%T", o))
			return nil
		}

		// This should be fast since the list of items it's going through should always be cached
		// locally with the shared informer.
		// Unfortunately we don't have a ctx we can use here, see https://github.com/kubernetes-sigs/controller-runtime/issues/1628
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var allMatches []client.Object

		for _, key := range keys {
			matchingResources, castOk := reflect.New(listType).Interface().(client.ObjectList)
			if !castOk {
				log.V(Status).Info("Somehow reflect.New() returned non client.ObjectList", "actual", fmt.Sprintf("%T", o))
				continue
			}

			err := c.List(ctx, matchingResources, client.MatchingFields{key: o.GetName()})
			if err != nil {
				// According to https://github.com/kubernetes/kubernetes/issues/51046, APIServer itself
				// doesn't actually support this. In our envtest tests, since we use a real client that
				// goes directly to APIServer, we'll never actually detect these changes
				log.Error(err, "couldn't list resources using secret", "kind", "TODO", "field", key, "value", o.GetName())
				continue
			}

			items, err := reflecthelpers.GetObjectListItems(matchingResources)
			if err != nil {
				log.Error(err, "failed to get list items")
				continue
			}

			allMatches = append(allMatches, items...)
		}

		var result []reconcile.Request

		for _, matchingResource := range allMatches {
			result = append(result, reconcile.Request{
				NamespacedName: types.NamespacedName{
					Namespace: matchingResource.GetNamespace(),
					Name:      matchingResource.GetName(),
				},
			})
		}

		return result
	}
}
