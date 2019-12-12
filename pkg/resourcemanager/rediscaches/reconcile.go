package rediscaches

import (
	"context"
	"fmt"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches"

	"k8s.io/apimachinery/pkg/runtime"
)

func (rc *AzureRedisCacheManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := rc.convert(obj)
	if err != nil {
		return false, err
	}

	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	sku := instance.Spec.Properties.Sku
	enableNonSSLPort := instance.Spec.Properties.EnableNonSslPort

	instance.Status.Provisioning = true

	_, err = rc.CreateRedisCache(ctx, groupName, name, location, sku, enableNonSSLPort, nil)
	if err != nil {
		instance.Status.Provisioning = false
		return false, fmt.Errorf("Redis Cache create error %v", err)
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	return true, nil
}

func (rc *AzureRedisCacheManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := rc.convert(obj)
	if err != nil {
		return false, err
	}

	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName

	_, err = rc.DeleteRedisCache(ctx, groupName, name)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			fmt.Errorf("RedisCache delete error %v", err)
			return nil
		}
	}
	return true, nil
}

func (rc *AzureRedisCacheManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	return nil, nil
}

func (rc *AzureRedisCacheManager) convert(obj runtime.Object) (*azurev1alpha1.RedisCache, error) {
	local, ok := obj.(*azurev1alpha1.RedisCache)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
