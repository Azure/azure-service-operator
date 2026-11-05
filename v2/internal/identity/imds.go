/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package identity

import (
	"context"
	"net/http"
	"time"
)

// IsIMDSAvailable reports whether the Azure Instance Metadata Service answers within one second.
// Any HTTP response counts as available. A timeout or connection error does not.
func IsIMDSAvailable(ctx context.Context) bool {
	probeCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		probeCtx,
		http.MethodGet,
		"http://169.254.169.254/metadata/identity/oauth2/token",
		nil,
	)
	if err != nil {
		return false
	}
	req.Header.Set("Metadata", "true")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
		return false
	}
	defer resp.Body.Close()
	return true
}
