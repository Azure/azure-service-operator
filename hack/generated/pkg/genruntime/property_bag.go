/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"encoding/json"
	"github.com/pkg/errors"
)

// PropertyBag is used to stash additional information that's not directly supported by storage
// resources, allowing for full fidelity round trip conversions
type PropertyBag map[string]string

// Clone creates an independent copy of the property bag
func (bag PropertyBag) Clone() PropertyBag {
	result := make(PropertyBag)
	for k, v := range bag {
		result[k] = v
	}

	return result
}

// Add is used to add a value into the bag; exact formatting depends on the type.
// Any existing value will be overwritten.
func (bag PropertyBag) Add(property string, value interface{}) error {
	switch v := value.(type) {
	case string:
		bag[property] = v
	default:
		// Default to treating as a JSON blob
		j, err := json.Marshal(v)
		if err != nil {
			return errors.Wrapf(err, "adding %s as JSON", property)
		}
		bag[property] = string(j)
	}

	return nil
}

// Pull removes a value from the bag, using it to populate the destination
// destination should be a pointer to where the value is to be placed
// Returns (false, nil) if the value isn't found, (true, nil) if the value is successfully stored,
// and (true, error) if something goes wrong during type conversion of the value.
// If an error occurs, the property value is still removed from the bag
func (bag PropertyBag) Pull(property string, destination interface{}) (bool, error) {
	value, found := bag[property]
	if !found {
		// Property not found in the bag
		return false, nil
	}

	// Property found, remove the value
	delete(bag, property)

	switch d := destination.(type) {
	case *string:
		*d = value
	default:
		data := []byte(value)
		err := json.Unmarshal(data, destination)
		if err != nil {
			return true, errors.Wrapf(err, "pulling from PropertyBag into %q", property)
		}
	}

	return true, nil
}
