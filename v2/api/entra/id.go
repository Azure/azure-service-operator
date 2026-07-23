// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package entra

func DirectoryObjectRefURI(objectID string) string {
	return "https://graph.microsoft.com/v1.0/directoryObjects/" + objectID
}
