/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// ImportSecretMode classifies whether a property contains a secret value.
type ImportSecretMode string

const (
	ImportSecretModeRequired = ImportSecretMode("required") // Property must be specified as a secret
	ImportSecretModeNever    = ImportSecretMode("never")    // Property never contains a secret
	ImportSecretModeOptional = ImportSecretMode("optional") // Property may optionally be specified as a secret
)
