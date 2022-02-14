/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"regexp"
	"strings"
)

type ResourceURI struct {
	Path       string
	Parameters map[string]Type
}

// SwaggerGroupRegex matches a “group” (Swagger ‘namespace’)
// based on: https://github.com/Azure/autorest/blob/85de19623bdce3ccc5000bae5afbf22a49bc4665/core/lib/pipeline/metadata-generation.ts#L25
var SwaggerGroupRegex = regexp.MustCompile(`[Mm]icrosoft\.[^/\\]+`)

func (uri *ResourceURI) GetARMType() string {
	parts := strings.Split(uri.Path, "/")

	var result strings.Builder

	reading := false
	for _, part := range parts {
		if !reading {
			if SwaggerGroupRegex.MatchString(part) {
				reading = true
				result.WriteString(part)
			}
		} else {
			if part[0] != '{' && part != "default" {
				result.WriteRune('/')
				result.WriteString(part)
			}
		}
	}

	return result.String()
}
