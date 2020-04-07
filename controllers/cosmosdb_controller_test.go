// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all cosmos

package controllers

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestCosmosDBHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	cosmosDBAccountName := GenerateTestResourceNameWithRandom("cosmosdb", 8)
	cosmosDBNamespace := "default"

	instance := &v1alpha1.CosmosDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cosmosDBAccountName,
			Namespace: cosmosDBNamespace,
		},
		Spec: v1alpha1.CosmosDBSpec{
			Location:      "westus",
			ResourceGroup: tc.resourceGroupName,
			Kind:          v1alpha1.CosmosDBKindGlobalDocumentDB,
			Properties: v1alpha1.CosmosDBProperties{
				DatabaseAccountOfferType: v1alpha1.CosmosDBDatabaseAccountOfferTypeStandard,
			},
		},
	}

	assert := assert.New(t)
	typeOf := fmt.Sprintf("%T", instance)

	err := tc.k8sClient.Create(ctx, instance)
	assert.Equal(nil, err, fmt.Sprintf("create %s in k8s", typeOf))

	res, err := meta.Accessor(instance)
	assert.Equal(nil, err, "not a metav1 object")

	names := types.NamespacedName{Name: res.GetName(), Namespace: res.GetNamespace()}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, names, instance)
		return HasFinalizer(res, finalizerName)
	}, tc.timeoutFast, tc.retry, "error waiting for %s to have finalizer", typeOf)

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, names, instance)
		log.Printf("[COSMOSDB][REF] Message=%v Provisioned=%v", instance.Status.Message, instance.Status.Provisioned)
		statused := ConvertToStatus(instance)
		log.Printf("[COSMOSDB][CON] Message=%v Provisioned=%v", statused.Status.Message, statused.Status.Provisioned)
		return strings.Contains(statused.Status.Message, successMsg) && statused.Status.Provisioned
	}, tc.timeout, tc.retry, "wait for %s to provision", typeOf)

	EnsureDelete(ctx, t, tc, instance)

}
