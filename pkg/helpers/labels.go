// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package helpers

import "strings"

// LabelsToTags converts labels from a kube resource to the data structure expected by Azure for tags
func LabelsToTags(in map[string]string) (map[string]*string, [][]string) {
	out := map[string]*string{}
	issues := [][]string{}
	for k, v := range in {
		if strings.ContainsAny(k, "<>%/?\\") {
			issues = append(issues, []string{k, "contains a character not allowed in Azure tags"})
		} else {
			value := v
			out[k] = &value
		}
	}

	return out, issues
}
