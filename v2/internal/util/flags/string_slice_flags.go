/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package flags

import (
	"flag"
	"strings"
)

type SliceFlags []string

var _ flag.Value = &SliceFlags{}

func (i *SliceFlags) String() string {
	return strings.Join(*i, ",")
}

func (i *SliceFlags) Set(value string) error {
	if value == "" {
		return nil
	}

	if strings.Contains(value, ",") {
		// split
		split := strings.Split(value, ",")
		for _, s := range split {
			*i = append(*i, s)
		}
		return nil
	}

	*i = append(*i, value)
	return nil
}
