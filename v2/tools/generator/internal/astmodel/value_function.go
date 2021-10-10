/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// ValueFunction defines a function that returns a single value suitable for storing in a property
type ValueFunction interface {
	Function

	// ReturnType specifies the return type of the function
	ReturnType() Type
}
