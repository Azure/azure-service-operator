/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package helpers

import (
	"context"
	"strings"

	aadpodv1 "github.com/Azure/aad-pod-identity/pkg/apis/aadpodidentity/v1"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Keep this in sync with the value in config/manager/manager.yaml
const selector = "aso_manager_binding"

type AADIdentity struct {
	IdentityName string
	ResourceId   string
	ClientID     string
}

// +kubebuilder:rbac:groups=aadpodidentity.k8s.io,resources=azureassignedidentities,verbs=get;list;watch
// +kubebuilder:rbac:groups=aadpodidentity.k8s.io,resources=azureassignedidentities/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=aadpodidentity.k8s.io,resources=azureidentities;azureidentities/status,verbs=get;list;watch
// +kubebuilder:rbac:groups=aadpodidentity.k8s.io,resources=azureidentitybindings;azureidentitybindings/status,verbs=get;list;watch

type AADIdentityFinder struct {
	client    client.Reader
	namespace string
}

func NewAADIdentityFinder(client client.Reader, namespace string) *AADIdentityFinder {
	return &AADIdentityFinder{
		client:    client,
		namespace: namespace,
	}
}

func (a *AADIdentityFinder) FindIdentity(ctx context.Context) (*AADIdentity, error) {

	var bindings aadpodv1.AzureIdentityBindingList
	// Can't use client.MatchingFields{"spec.selector": selector} here because by default
	// fields other than namespace aren't indexed. See https://github.com/kubernetes-sigs/kubebuilder/issues/547
	if err := a.client.List(ctx, &bindings, client.InNamespace(a.namespace)); err != nil {
		return nil, err
	}

	// Client side filtering, for now ok because we're expecting there aren't many of these
	var matchingBindings []aadpodv1.AzureIdentityBinding
	for _, binding := range bindings.Items {
		if selector == binding.Spec.Selector {
			matchingBindings = append(matchingBindings, binding)
		}
	}

	// The expectation here is that there is only a single AAD Identity in our namespace.
	// At some point we'll support multiple and we'll need to fix this logic
	if len(matchingBindings) == 0 {
		return nil, errors.Errorf("couldn't find any azure identity bindings matching %s", selector)
	}
	if len(matchingBindings) > 1 {
		return nil, errors.Errorf("found %d bindings matching %s, expected 1", len(bindings.Items), selector)
	}

	binding := matchingBindings[0]

	var identity aadpodv1.AzureIdentity
	key := client.ObjectKey{
		Namespace: binding.Namespace,
		Name:      binding.Spec.AzureIdentity,
	}
	if err := a.client.Get(ctx, key, &identity); err != nil {
		return nil, errors.Wrapf(err, "failed getting identity %s", key)
	}

	// TODO: Confirm the identity we found was the right kind (MSI) rather than SP

	return &AADIdentity{
		IdentityName: extractIdentityName(identity.Spec.ResourceID),
		ResourceId:   identity.Spec.ResourceID,
		ClientID:     identity.Spec.ClientID,
	}, nil
}

func extractIdentityName(resourceId string) string {
	items := strings.Split(resourceId, "/")
	return items[len(items)-1]
}
