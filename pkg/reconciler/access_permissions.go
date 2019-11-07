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

package reconciler

import "strings"

const (
	AccessPermissionAnnotation = "/access-permissions"
)

type AccessPermissions string

func (p AccessPermissions) create() bool { return containsOrEmpty(p, "c") }
func (p AccessPermissions) update() bool { return containsOrEmpty(p, "u") }
func (p AccessPermissions) delete() bool { return containsOrEmpty(p, "d") }

// Note that all permissions are enabled by default
func containsOrEmpty(p AccessPermissions, c string) bool {
	return strings.Contains(strings.ToLower(string(p)), c) || p == ""
}
