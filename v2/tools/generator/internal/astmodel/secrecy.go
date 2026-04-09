/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// Secrecy classifies whether a property contains a secret value.
type Secrecy string

const (
	SecrecyAlways = Secrecy("always") // Property always contains a secret
	SecrecyNever  = Secrecy("never")  // Property never contains a secret
)
