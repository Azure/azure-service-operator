// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package entra

import "net/url"

func DirectoryObjectRefURI(objectID string) string {
	return "https://graph.microsoft.com/v1.0/directoryObjects/" + url.PathEscape(objectID)
}
