/*
 *Copyright (c) Microsoft Corporation.
 *Licensed under the MIT license.
 */

package arm

import (
	"context"

	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// CommitObject persists the contents of obj to etcd by using the Kubernetes client.
// Note that after this method has been called, obj contains the result of the update
// from APIServer (including an updated resourceVersion).
func CommitObject(ctx context.Context, kubeClient *kubeclient.Client, obj genruntime.MetaObject) error {
	// Order of updates (spec first or status first) matters here.
	// If the status is updated first: clients that are waiting on status
	// Condition Ready == true might see that quickly enough, and make a spec
	// update fast enough, to conflict with the second write (that of the spec).
	// This will trigger extra requests to Azure and fail our recording tests but is
	// otherwise harmless in an actual deployment.
	// We update the spec first to avoid the above problem.

	// We must clone here because the result of this update could contain
	// fields such as status.location that may not be set but are not omitempty.
	// This will cause the contents we have in Status.Location to be overwritten.
	clone := obj.DeepCopyObject().(client.Object)

	err := kubeClient.Client.Update(ctx, clone)
	if err != nil {
		return errors.Wrapf(err, "updating %s/%s resource", obj.GetNamespace(), obj.GetName())
	}

	obj.SetResourceVersion(clone.GetResourceVersion())

	// Note that subsequent calls to GET can (if using a cached client) can miss the updates we've just done.
	// See: https://github.com/kubernetes-sigs/controller-runtime/issues/1464.
	err = kubeClient.Client.Status().Update(ctx, obj)
	if err != nil {
		return errors.Wrapf(err, "updating %s/%s resource status", obj.GetNamespace(), obj.GetName())
	}

	return nil
}

func IgnoreNotFoundAndConflict(err error) error {
	if apierrors.IsConflict(err) {
		return nil
	}

	return client.IgnoreNotFound(err)
}
