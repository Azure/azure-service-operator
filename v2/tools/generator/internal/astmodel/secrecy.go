/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// Secrecy classifies whether a property contains a secret value.
type Secrecy string

const (
	SecrecyRequired = Secrecy("required") // Property must be specified as a secret
	SecrecyNever    = Secrecy("never")    // Property never contains a secret
	SecrecyOptional = Secrecy("optional") // Property may optionally be specified as a secret
)
