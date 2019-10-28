/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package eventhub

import (
	"fmt"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	converter "github.com/Azure/azure-service-operator/controllers_new/shared"
	"github.com/Azure/azure-service-operator/pkg/reconciler"
	"k8s.io/apimachinery/pkg/runtime"
)

func GetStatus(instance runtime.Object) (*reconciler.Status, error) {
	x, err := convertInstance(instance)
	if err != nil {
		return nil, err
	}
	return converter.ToControllerStatus(x.Status), nil
}

func updateStatus(instance runtime.Object, status *reconciler.Status) error {
	x, err := convertInstance(instance)
	if err != nil {
		return err
	}
	x.Status = converter.ToAsoStatus(status)
	return nil
}

func convertInstance(obj runtime.Object) (*v1alpha1.Eventhub, error) {
	local, ok := obj.(*v1alpha1.Eventhub)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: EventhubNamespace")
	}
	return local, nil
}
