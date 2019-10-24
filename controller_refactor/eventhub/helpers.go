package eventhub

import (
	"fmt"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
)

func getStatus(instance runtime.Object) (*v1alpha1.ASOStatus, error) {
	x, err := convertInstance(instance)
	if err != nil {
		return nil, err
	}
	return &x.Status, nil
}

func updateStatus(instance runtime.Object, status *v1alpha1.ASOStatus) error {
	x, err := convertInstance(instance)
	if err != nil {
		return err
	}
	x.Status = *status
	return nil
}

func convertInstance(obj runtime.Object) (*v1alpha1.Eventhub, error) {
	local, ok := obj.(*v1alpha1.Eventhub)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: EventhubNamespace")
	}
	return local, nil
}
