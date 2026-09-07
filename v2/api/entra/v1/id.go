// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package v1

import "net/url"

func DirectoryObjectRefURI(objectID string) string {
	return "https://graph.microsoft.com/v1.0/directoryObjects/" + url.PathEscape(objectID)
}
