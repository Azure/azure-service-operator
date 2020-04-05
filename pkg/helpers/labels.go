// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package helpers

import (
	"strings"
)

// LabelsToTags converts labels from a kube resource to the data structure expected by Azure for tags
// this function will translate characters that are not allows by Azure to "."s
func LabelsToTags(in map[string]string) map[string]*string {
	out := map[string]*string{}
	for k, v := range in {
		newK := k
		value := v
		if strings.ContainsAny(k, "<>%/?\\") {
			newK = ReplaceAny(k, []string{"<", ">", "%", "/", "\\\\", "\\?"})
		}
		out[newK] = &value
	}

	return out
}
