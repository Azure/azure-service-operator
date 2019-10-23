package controller_refactor

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type instanceFetcher struct {
	kubeClient client.Client
}

func (fetcher *instanceFetcher) fetch(ctx context.Context, namespacedName types.NamespacedName, object runtime.Object) error {
	return fetcher.kubeClient.Get(ctx, namespacedName, object)
}
