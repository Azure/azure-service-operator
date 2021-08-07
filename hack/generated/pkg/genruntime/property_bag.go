/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// PropertyBag is an unordered set of stashed information that used for properties not directly supported by storage
// resources, allowing for full fidelity round trip conversions
type PropertyBag map[string]string

// PropertyBag returns a new property bag
// originals is a (potentially empty) sequence of existing property bags who's content will be copied into the new
// property bag. In the case of key overlaps, the bag later in the parameter list overwrites the earlier value.
func NewPropertyBag(originals ...PropertyBag) PropertyBag {
	result := make(PropertyBag)

	for _, orig := range originals {
		for k, v := range orig {
			result[k] = v
		}
	}

	return result
}

// Contains returns true if the specified name is present in the bag; false otherwise
func (bag PropertyBag) Contains(name string) bool {
	_, found := bag[name]
	return found
}

// Add is used to add a value into the bag; exact formatting depends on the type.
// Any existing value will be overwritten.
// property is the name of the item to put into the bag
// value is the instance to be stashed away for later
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
// property is the name of the item to remove and return
// destination should be a pointer to where the value is to be placed
// If the item is present and successfully deserialized, returns no error (nil); otherwise returns an error.
// If an error happens deserializing an item from the bag, it is still removed from the bag.
func (bag PropertyBag) Pull(property string, destination interface{}) error {
	value, found := bag[property]
	if !found {
		// Property not found in the bag
		return errors.Errorf("property bag does not contain %q", property)
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
			return errors.Wrapf(err, "pulling %q from PropertyBag", property)
		}
	}

	return nil
}
