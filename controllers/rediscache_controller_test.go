// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/stretchr/testify/assert"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestRedisCacheControllerHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)

	var rgLocation string
	var rgName string
	var redisCacheName string
	var redisCacheSecret string
	var err error

	rgName = tc.resourceGroupName
	rgLocation = tc.resourceGroupLocation
	redisCacheName = "t-rediscache-" + helpers.RandomString(10)
	redisCacheSecret = "RedisCacheSecret"

	// Create the RedisCache object and expect the Reconcile to be created
	redisCacheInstance := &azurev1alpha1.RedisCache{
		ObjectMeta: metav1.ObjectMeta{
			Name:      redisCacheName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.RedisCacheSpec{
			Location:          rgLocation,
			ResourceGroupName: rgName,
			Properties: azurev1alpha1.RedisCacheProperties{
				Sku: azurev1alpha1.RedisCacheSku{
					Name:     "B_Gen5_2",
					Family:   "Gen5",
					Capacity: 2,
				},
				EnableNonSslPort: true,
			},
			SecretName: redisCacheSecret,
		},
	}

	// create rc
	EnsureInstance(ctx, t, tc, redisCacheInstance)

	//verify secret exists in k8s for rc
	secret := &v1.Secret{}
	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, types.NamespacedName{Name: redisCacheSecret, Namespace: redisCacheInstance.Namespace}, secret)

		if err == nil {
			if (secret.ObjectMeta.Name == redisCacheSecret) && (secret.ObjectMeta.Namespace == redisCacheInstance.Namespace) {
				return true
			}
		}
		return false
	}, tc.timeoutFast, tc.retry, "wait for rc to have secret")

	// delete rc
	EnsureDelete(ctx, t, tc, redisCacheInstance)
}
