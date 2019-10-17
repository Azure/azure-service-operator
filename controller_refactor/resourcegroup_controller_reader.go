package controller_refactor

import (
	"context"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ResourceGroupDefinitionFetcher struct{}

func (fetcher *ResourceGroupDefinitionFetcher) GetThis(ctx context.Context, kubeClient client.Client, req controllerruntime.Request) (*ThisResourceDefinitions, error) {
	var instance v1alpha1.ResourceGroup
	err := kubeClient.Get(ctx, req.NamespacedName, &instance)
	crdInfo := fetcher.getDefinition(&instance)
	return &ThisResourceDefinitions{
		CRDInfo:    crdInfo,
		CRDUpdater: fetcher.getUpdater(&instance, crdInfo),
	}, err
}

func (_ *ResourceGroupDefinitionFetcher) GetDependencies(ctx context.Context, kubeClient client.Client, req controllerruntime.Request) (*DependencyDefinitions, error) {
	return &DependencyDefinitions{
		Dependencies: []*CRDInfo{},
		Owner:        nil,
	}, nil
}

func (_ *ResourceGroupDefinitionFetcher) getDefinition(instance *v1alpha1.ResourceGroup) *CRDInfo {
	return &CRDInfo{
		ProvisionState:  instance.Status.ProvisionState,
		Name:            instance.Name,
		Parameters:      instance.Spec.Parameters,
		CRDInstance:     instance,
		GetBaseInstance: &instance.ResourceBaseDefinition,
		IsBeingDeleted:  !instance.ObjectMeta.DeletionTimestamp.IsZero(),
	}
}

func (_ *ResourceGroupDefinitionFetcher) getUpdater(instance *v1alpha1.ResourceGroup, crdInfo *CRDInfo) *CRDUpdater {
	return &CRDUpdater{
		UpdateInstance: func(state *v1alpha1.ResourceBaseDefinition) {
			instance.ResourceBaseDefinition = *state
		},
		CRDInfo: crdInfo,
	}
}
