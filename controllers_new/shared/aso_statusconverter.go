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

package shared

import (
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/reconciler"
)

func ToAsoStatus(c *reconciler.Status) v1alpha1.ASOStatus {
	cs := c.State
	return v1alpha1.ASOStatus{
		Provisioning: cs != reconciler.Pending &&
			cs != reconciler.Succeeded &&
			cs != reconciler.Failed,
		Provisioned: cs == reconciler.Succeeded,
		State:       string(c.State),
		Message:     c.Message,
	}
}

func ToControllerStatus(a v1alpha1.ASOStatus) *reconciler.Status {
	return &reconciler.Status{
		State:   reconciler.ProvisionState(a.State),
		Message: a.Message,
	}
}
