// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package telemetry

import (
	log "github.com/go-logr/logr"
)

// BaseClient stores information for all implementations of telemetry
type BaseClient struct {
	Logger log.Logger
}
