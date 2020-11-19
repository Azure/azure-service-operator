/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	"github.com/Azure/k8s-infra/hack/generated/pkg/reflecthelpers"
)

const (
	// TODO: Delete these later in favor of something in status?
	DeploymentIdAnnotation   = "deployment-id.infra.azure.com"
	DeploymentNameAnnotation = "deployment-name.infra.azure.com"
	ResourceStateAnnotation  = "resource-state.infra.azure.com"
	ResourceIdAnnotation     = "resource-id.infra.azure.com"
	ResourceErrorAnnotation  = "resource-error.infra.azure.com"
	ResourceSigAnnotationKey = "resource-sig.infra.azure.com"
	// PreserveDeploymentAnnotation is the key which tells the applier to keep or delete the deployment
	PreserveDeploymentAnnotation = "x-preserve-deployment"
)

// TODO: Anybody have a better name?
type ReconcileMetadata struct {
	log     logr.Logger
	metaObj genruntime.MetaObject
}

func NewReconcileMetadata(metaObj genruntime.MetaObject, log logr.Logger) *ReconcileMetadata {
	return &ReconcileMetadata{
		metaObj: metaObj,
		log:     log,
	}
}

func (r *ReconcileMetadata) IsTerminalProvisioningState() bool {
	state := r.GetResourceProvisioningState()
	return armclient.IsTerminalProvisioningState(state)
}

func (r *ReconcileMetadata) GetResourceProvisioningState() armclient.ProvisioningState {
	return armclient.ProvisioningState(r.metaObj.GetAnnotations()[ResourceStateAnnotation])
}

func (r *ReconcileMetadata) GetDeploymentId() (string, bool) {
	id, ok := r.metaObj.GetAnnotations()[DeploymentIdAnnotation]
	return id, ok
}

func (r *ReconcileMetadata) GetDeploymentIdOrDefault() string {
	id, _ := r.GetDeploymentId()
	return id
}

func (r *ReconcileMetadata) SetDeploymentId(id string) {
	r.addAnnotation(DeploymentIdAnnotation, id)
}

func (r *ReconcileMetadata) GetDeploymentName() (string, bool) {
	id, ok := r.metaObj.GetAnnotations()[DeploymentNameAnnotation]
	return id, ok
}

func (r *ReconcileMetadata) GetDeploymentNameOrDefault() string {
	id, _ := r.GetDeploymentName()
	return id
}

func (r *ReconcileMetadata) SetDeploymentName(name string) {
	r.addAnnotation(DeploymentNameAnnotation, name)
}

func (r *ReconcileMetadata) GetShouldPreserveDeployment() bool {
	preserveDeploymentString, ok := r.metaObj.GetAnnotations()[PreserveDeploymentAnnotation]
	if !ok {
		return false
	}

	preserveDeployment, err := strconv.ParseBool(preserveDeploymentString)
	// Anything other than an error is assumed to be false...
	// TODO: Would we rather have any usage of this key imply true (regardless of value?)
	if err != nil {
		// TODO: Log here
		return false
	}

	return preserveDeployment
}

func (r *ReconcileMetadata) GetResourceIdOrDefault() string {
	return r.metaObj.GetAnnotations()[ResourceIdAnnotation]
}

func (r *ReconcileMetadata) SetResourceId(id string) {
	r.addAnnotation(ResourceIdAnnotation, id)
}

func (r *ReconcileMetadata) SetResourceProvisioningState(state armclient.ProvisioningState) {
	// TODO: It's almost certainly not safe to use this as our serialized format as it's not guaranteed backwards compatible?
	r.addAnnotation(ResourceStateAnnotation, string(state))
}

func (r *ReconcileMetadata) SetResourceError(error string) {
	r.addAnnotation(ResourceErrorAnnotation, error)
}

func (r *ReconcileMetadata) SetResourceSignature(sig string) {
	r.addAnnotation(ResourceSigAnnotationKey, sig)
}

func (r *ReconcileMetadata) HasResourceSpecHashChanged() (bool, error) {
	oldSig, exists := r.metaObj.GetAnnotations()[ResourceSigAnnotationKey]
	if !exists {
		// signature does not exist, so yes, it has changed
		return true, nil
	}

	newSig, err := r.SpecSignature()
	if err != nil {
		return false, err
	}
	// check if the last signature matches the new signature
	return oldSig != newSig, nil
}

func (r *ReconcileMetadata) addAnnotation(k string, v string) {
	annotations := r.metaObj.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}
	// I think this is the behavior we want...
	if v == "" {
		delete(annotations, k)
	} else {
		annotations[k] = v
	}
	r.metaObj.SetAnnotations(annotations)
}

// SpecSignature calculates the hash of a spec. This can be used to compare specs and determine
// if there has been a change
func (r *ReconcileMetadata) SpecSignature() (string, error) {
	// Convert the resource to unstructured for easier comparison later.
	unObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(r.metaObj)
	if err != nil {
		return "", err
	}

	spec, ok, err := unstructured.NestedMap(unObj, "spec")
	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("unable to find spec within unstructured MetaObject")
	}

	bits, err := json.Marshal(spec)
	if err != nil {
		return "", errors.Wrap(err, "unable to marshal spec of unstructured MetaObject")
	}

	hash := sha256.Sum256(bits)
	return hex.EncodeToString(hash[:]), nil
}

func (r *ReconcileMetadata) Update(
	deployment *armclient.Deployment,
	status genruntime.FromArmConverter) error {

	controllerutil.AddFinalizer(r.metaObj, GenericControllerFinalizer)

	sig, err := r.SpecSignature()
	if err != nil {
		return errors.Wrap(err, "failed to compute resource spec hash")
	}

	r.SetDeploymentId(deployment.Id)
	r.SetDeploymentName(deployment.Name)
	// TODO: Do we want to just use Azure's annotations here? I bet we don't? We probably want to map
	// TODO: them onto something more robust? For now just use Azure's though.
	r.SetResourceProvisioningState(deployment.Properties.ProvisioningState)
	r.SetResourceSignature(sig)
	if deployment.IsTerminalProvisioningState() {
		if deployment.Properties.ProvisioningState == armclient.FailedProvisioningState {
			r.SetResourceError(deployment.Properties.Error.String())
		} else if len(deployment.Properties.OutputResources) > 0 {
			resourceId := deployment.Properties.OutputResources[0].ID
			r.SetResourceId(resourceId)

			if status != nil {
				err = reflecthelpers.SetStatus(r.metaObj, status)
				if err != nil {
					return err
				}
			}
		} else {
			return errors.New("template deployment didn't have any output resources")
		}
	}

	return nil
}
