/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package entra

type relationshipDelta struct {
	ToAdd    []string
	ToRemove []string
}

func planRelationshipDelta(current []string, desired []string) relationshipDelta {
	currentSet := make(map[string]struct{}, len(current))
	desiredSet := make(map[string]struct{}, len(desired))

	for _, value := range current {
		currentSet[value] = struct{}{}
	}

	for _, value := range desired {
		desiredSet[value] = struct{}{}
	}

	toAdd := make([]string, 0)
	for _, value := range desired {
		if _, ok := currentSet[value]; !ok {
			toAdd = append(toAdd, value)
		}
	}

	toRemove := make([]string, 0)
	for _, value := range current {
		if _, ok := desiredSet[value]; !ok {
			toRemove = append(toRemove, value)
		}
	}

	return relationshipDelta{
		ToAdd:    orderedUnique(toAdd),
		ToRemove: orderedUnique(toRemove),
	}
}

func orderedUnique(values []string) []string {
	seen := make(map[string]struct{}, len(values))
	result := make([]string, 0, len(values))

	for _, value := range values {
		if _, ok := seen[value]; ok {
			continue
		}

		seen[value] = struct{}{}
		result = append(result, value)
	}

	return result
}
