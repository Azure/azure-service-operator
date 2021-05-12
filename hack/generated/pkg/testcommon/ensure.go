/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"

	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
)

type Ensure struct {
	kubeClient      client.Client
	stateAnnotation string
	errorAnnotation string
}

func NewEnsure(c client.Client, stateAnnotation string, errorAnnotation string) *Ensure {
	return &Ensure{
		kubeClient:      c,
		stateAnnotation: stateAnnotation,
		errorAnnotation: errorAnnotation,
	}
}

// Provisioned checks to ensure the provisioning state of the resource is successful.
func (e *Ensure) Provisioned(ctx context.Context, obj runtime.Object) (bool, error) {
	key, err := client.ObjectKeyFromObject(obj)
	if err != nil {
		return false, err
	}

	err = e.kubeClient.Get(ctx, key, obj)
	if err != nil {
		return false, err
	}

	// Have to cast because return of kubeClient.Get is not a metav1.Object unfortunately.
	metaObj, ok := obj.(metav1.Object)
	if !ok {
		return false, errors.Errorf("result of get was not metav1.Object, was: %T", obj)
	}

	state := metaObj.GetAnnotations()[e.stateAnnotation]
	return state == string(armclient.SucceededProvisioningState), nil
}

// Deleted ensures that the object specified has been deleted
func (e *Ensure) Deleted(ctx context.Context, obj runtime.Object) (bool, error) {
	key, err := client.ObjectKeyFromObject(obj)
	if err != nil {
		return false, err
	}

	// Note that obj won't be modified if it's already deleted, so
	// could be "stale" after this call
	err = e.kubeClient.Get(ctx, key, obj)
	if apierrors.IsNotFound(err) {
		return true, nil
	}
	if err != nil {
		return false, err
	}

	return false, nil
}

// AllDeleted ensures that all of the specified objects are deleted
func (e *Ensure) AllDeleted(ctx context.Context, objs []runtime.Object) (bool, error) {
	for _, obj := range objs {
		// It's possible that this is horribly inefficient. Should be good enough for now though
		deleted, err := e.Deleted(ctx, obj)
		if err != nil {
			return false, err
		}
		if !deleted {
			return false, nil
		}
	}

	return true, nil
}
