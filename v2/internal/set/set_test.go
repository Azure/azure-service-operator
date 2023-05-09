/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package set

import (
	"testing"
)

func TestSet(t *testing.T) {
	// Test creating a set with multiple values
	set := Make(1, 2, 3)
	if len(set) != 3 {
		t.Errorf("Expected set length 3, but got %v", len(set))
	}

	// Test adding a value to the set
	set.Add(4)
	if !set.Contains(4) {
		t.Errorf("Expected set to contain 4, but it didn't")
	}

	// Test adding all values from another set
	otherSet := Make(3, 4, 5)
	set.AddAll(otherSet)
	if len(set) != 5 {
		t.Errorf("Expected set length 4, but got %v", len(set))
	}

	// Test removing a value from the set
	set.Remove(4)
	if set.Contains(4) {
		t.Errorf("Expected set not to contain 4, but it did")
	}

	// Test copying the set
	copySet := set.Copy()
	if !AreEqual(set, copySet) {
		t.Errorf("Expected copied set to be equal to original set, but they were different")
	}

	// Test clearing the set
	set.Clear()
	if len(set) != 0 {
		t.Errorf("Expected set length 0 after clear, but got %v", len(set))
	}

	// Test checking if two sets are equal
	set1 := Make(1, 2, 3)
	set2 := Make(1, 2, 3)
	if !set1.Equals(set2) {
		t.Errorf("Expected sets to be equal, but they were different")
	}

	// Test getting the values from the set
	set3 := Make("foo", "bar", "baz")
	values := set3.Values()
	if len(values) != 3 {
		t.Errorf("Expected values length 3, but got %v", len(values))
	}

	// Test getting sorted slice of values
	set4 := Make(3, 1, 2)
	sortedValues := AsSortedSlice(set4)
	if len(sortedValues) != 3 {
		t.Errorf("Expected sorted values length 3, but got %v", len(sortedValues))
	}
	if sortedValues[0] != 1 || sortedValues[1] != 2 || sortedValues[2] != 3 {
		t.Errorf("Expected sorted values to be [1 2 3], but got %v", sortedValues)
	}
}
