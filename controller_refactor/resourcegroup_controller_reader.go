package controller_refactor

import (
	"context"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ResourceGroupDefinitionFetcher struct{}

func (fetcher *ResourceGroupDefinitionFetcher) GetDefinition(ctx context.Context, kubeClient client.Client, req controllerruntime.Request) (Definition, CRDUpdater, error) {
	var instance v1alpha1.ResourceGroup
	err := kubeClient.Get(ctx, req.NamespacedName, &instance)
	return fetcher.getDefinition(instance), fetcher.getUpdater(instance), err
}

func (_ *ResourceGroupDefinitionFetcher) GetDependencies(ctx context.Context, kubeClient client.Client, req controllerruntime.Request) ([]Definition, error) {
	return []Definition{}, nil
}

func (_ *ResourceGroupDefinitionFetcher) getDefinition(instance v1alpha1.ResourceGroup) Definition {
	return Definition{
		ProvisionState: ProvisionState(instance.Status.ProvisionState),
		Name:           instance.Name,
		CRDInstance:    &instance,
		IsBeingDeleted: !instance.ObjectMeta.DeletionTimestamp.IsZero(),
	}
}

func (_ *ResourceGroupDefinitionFetcher) getUpdater(instance v1alpha1.ResourceGroup) CRDUpdater {
	return CRDUpdater{
		AddFinalizer:    instance.AddFinalizer,
		RemoveFinalizer: instance.RemoveFinalizer,
		HasFinalizer:    instance.HasFinalizer,
		SetState:        func (state ProvisionState) {
			instance.Status.ProvisionState = v1alpha1.ProvisionState(state)
		},
	}
}

