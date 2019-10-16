package controller_refactor

import (
	"context"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ResourceGroupDefinitionFetcher struct{}

func (fetcher *ResourceGroupDefinitionFetcher) GetThis(ctx context.Context, kubeClient client.Client, req controllerruntime.Request) (CRDInfo, CRDUpdater, error) {
	var instance v1alpha1.ResourceGroup
	err := kubeClient.Get(ctx, req.NamespacedName, &instance)
	return fetcher.getDefinition(&instance), fetcher.getUpdater(&instance), err
}

func (_ *ResourceGroupDefinitionFetcher) GetDependencies(ctx context.Context, kubeClient client.Client, req controllerruntime.Request) ([]CRDInfo, error) {
	return []CRDInfo{}, nil
}

func (_ *ResourceGroupDefinitionFetcher) getDefinition(instance *v1alpha1.ResourceGroup) CRDInfo {
	return CRDInfo{
		ProvisionState: instance.Status.ProvisionState,
		Name:           instance.Name,
		CRDInstance:    instance,
		IsBeingDeleted: !instance.ObjectMeta.DeletionTimestamp.IsZero(),
	}
}

func (_ *ResourceGroupDefinitionFetcher) getUpdater(instance *v1alpha1.ResourceGroup) CRDUpdater {
	return CRDUpdater{
		UpdateInstance: func(state *v1alpha1.ResourceBaseState) {
			instance.ResourceBaseState = *state
		},
		GetBaseInstance: &(*instance).ResourceBaseState,
	}
}
