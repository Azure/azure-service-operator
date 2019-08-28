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

package v1

import (
	"time"

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
)

type TestConfig struct {
	Key           string
	Namespace     string
	Location      string
	ResourceGroup string
	Timeout       time.Duration
	Poll          time.Duration
}

func NewTestConfig() *TestConfig {
	key := helpers.RandomString(10)
	return &TestConfig{
		Key:           key,
		Namespace:     "default",
		Location:      "westus",
		ResourceGroup: "aso-dev-rg-" + key,
		Timeout:       time.Second * 240,
		Poll:          time.Second * 10,
	}
}
