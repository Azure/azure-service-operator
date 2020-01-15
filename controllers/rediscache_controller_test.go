package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/stretchr/testify/assert"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
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
	var redisCacheNamespacedName types.NamespacedName
	var err error

	rgName = tc.resourceGroupName
	rgLocation = tc.resourceGroupLocation
	redisCacheName = "t-rediscache-" + helpers.RandomString(10)

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
		},
	}

	// create redis
	err = tc.k8sClient.Create(ctx, redisCacheInstance)
	assert.Equal(nil, err, "create redisCacheInstance in k8s")

	redisCacheNamespacedName = types.NamespacedName{Name: redisCacheName, Namespace: "default"}

	// make sure redis has a finalizer
	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, redisCacheNamespacedName, redisCacheInstance)
		return helpers.HasFinalizer(redisCacheInstance, finalizerName)
	}, tc.timeout, tc.retry, "wait for redisCacheInstance to have finalizer")

	// make sure redis provisions
	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, redisCacheNamespacedName, redisCacheInstance)
		return redisCacheInstance.Status.Provisioned
	}, tc.timeout, tc.retry, "wait for redisCacheInstance to be provisioned")

	// delete redis
	err = tc.k8sClient.Delete(ctx, redisCacheInstance)
	assert.Equal(nil, err, "delete redisCacheInstance in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, redisCacheNamespacedName, redisCacheInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for redisCache to be gone from k8s")
}
