/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

/*
This file contains manual implementations to reduce code bloat in generated code.
*/

type ResourceExtensions interface {
	GetExtendedResources() *[]KubernetesResource
}
