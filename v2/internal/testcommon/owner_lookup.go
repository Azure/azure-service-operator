/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func objectRanksByOwner(objs ...client.Object) [][]client.Object {
	var result [][]client.Object

	roots := findRootResources(objs)
	toExamine := roots
	for len(toExamine) != 0 {
		result = append(result, toExamine)
		var owned []client.Object
		for _, obj := range toExamine {
			owned = append(owned, findOwnedResources(objs, obj)...)
		}
		toExamine = owned
	}

	return result
}

func findRootResources(objs []client.Object) []client.Object {
	oids := make(map[types.UID]client.Object)
	for _, obj := range objs {
		oids[obj.GetUID()] = obj
	}

	// Note: This is imperfect. If you have a situation like A owns B owns C and you pass just A and C here, we won't detect it
	var roots []client.Object

	for _, obj := range objs {
		refs := obj.GetOwnerReferences()
		foundOwner := false
		for _, ref := range refs {
			if _, ok := oids[ref.UID]; ok {
				foundOwner = true
				break
			}
		}
		if !foundOwner {
			roots = append(roots, obj)
		}
	}

	return roots
}

func findOwnedResources(objs []client.Object, obj client.Object) []client.Object {
	var result []client.Object

	for _, o := range objs {
		ownerRefs := o.GetOwnerReferences()
		for _, ownerRef := range ownerRefs {
			if ownerRef.UID == obj.GetUID() {
				result = append(result, o)
				break
			}
		}
	}

	return result
}
