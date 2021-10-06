/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

type StringSet map[string]struct{}

func MakeStringSet() StringSet {
	return make(StringSet)
}

func (set StringSet) Contains(s string) bool {
	_, ok := set[s]
	return ok
}

func (set StringSet) Add(s string) {
	set[s] = struct{}{}
}

func (set StringSet) Remove(s string) {
	delete(set, s)
}
