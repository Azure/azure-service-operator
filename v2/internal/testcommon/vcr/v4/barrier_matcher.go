/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v4

import (
	"net/url"
)

// urlPath extracts the path portion of a URL string, stripping query parameters.
// Used by both the barrier matcher and the replay round tripper to match resources by path.
func urlPath(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		// Should never happen with URLs from HTTP requests
		return rawURL
	}

	return u.Path
}
