/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

// A TypeLoaderRename is used to resolve a naming conflict that happens during loading of types
// right at the start of the code generator.
type TypeLoaderRename struct {
	TypeMatcher `yaml:",inline"`
	Scope       *string `yaml:"scope,omitempty"`
	RenameTo    *string `yaml:"renameTo,omitempty"`
}
