// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build rediscache

package controllers

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/stretchr/testify/assert"

	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const longRunningTimeout = 25 * time.Minute

func TestRedisCacheControllerHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	var rgLocation string
	var rgName string
	var redisCacheName string
	var err error

	rgName = tc.resourceGroupName
	rgLocation = tc.resourceGroupLocation
	redisCacheName = GenerateTestResourceNameWithRandom("rediscache", 10)

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
					Name:     "Basic",
					Family:   "C",
					Capacity: 0,
				},
				EnableNonSslPort: true,
			},
		},
	}

	// create rc
	//EnsureInstance(ctx, t, tc, redisCacheInstance)

	err = tc.k8sClient.Create(ctx, redisCacheInstance)
	assert.Equal(nil, err, "create redis cache in k8s")

	names := types.NamespacedName{Name: redisCacheInstance.GetName(), Namespace: redisCacheInstance.GetNamespace()}

	// Wait for first sql server to resolve
	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, names, redisCacheInstance)
		return HasFinalizer(redisCacheInstance, finalizerName)
	}, tc.timeoutFast, tc.retry, fmt.Sprintf("wait for %s to have finalizer", "rediscache"))

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, names, redisCacheInstance)
		return strings.Contains(redisCacheInstance.Status.Message, successMsg) && redisCacheInstance.Status.Provisioned == true
	}, longRunningTimeout, tc.retry, fmt.Sprintf("wait for %s to provision", "rediscache"))

	//verify secret exists in k8s for rc
	secret := &v1.Secret{}
	assert.Eventually(func() bool {
		log.Println("get secret")
		err = tc.k8sClient.Get(ctx, types.NamespacedName{Name: redisCacheInstance.Name, Namespace: redisCacheInstance.Namespace}, secret)
		if err == nil {
			return true
		}
		return false
	}, 45*time.Second, tc.retry, "wait for rc to have secret")

	// delete rc
	err = tc.k8sClient.Delete(ctx, redisCacheInstance)
	assert.Equal(nil, err, fmt.Sprintf("delete %s in k8s", "rediscache"))

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, names, redisCacheInstance)
		return apierrors.IsNotFound(err)
	}, longRunningTimeout, tc.retry, fmt.Sprintf("wait for %s to be gone from k8s", "rediscache"))

}
