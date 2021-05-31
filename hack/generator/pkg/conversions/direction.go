/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

// Direction specifies the direction of conversion we're implementing with this function
type Direction int

const (
	// ConvertFrom indicates the conversion is from the passed 'other', populating the receiver with properties from the other
	ConvertFrom = Direction(1)
	// ConvertTo indicates the conversion is to the passed 'other', populating the other with properties from the receiver
	ConvertTo = Direction(2)
)
