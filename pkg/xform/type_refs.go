/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package xform

import (
	"fmt"
	"reflect"
	"strings"

	azcorev1 "github.com/Azure/k8s-infra/apis/core/v1"
)

type (
	TypeReferenceLocation struct {
		JSONFieldName     string
		TemplateFieldName string
		Path              []string
		Group             string
		Kind              string
		IsOwned           bool
		IsSlice           bool
	}
)

func GetTypeReferenceData(obj azcorev1.MetaObject) ([]TypeReferenceLocation, error) {
	// expect the obj.Spec.Properties
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	specField, found := t.FieldByName("Spec")
	if !found {
		return nil, fmt.Errorf("GetTypeReferenceData could not find obj.Spec field")
	}

	propField, found := specField.Type.FieldByName("Properties")
	if !found {
		return []TypeReferenceLocation{}, nil
	}

	refs, err := getResourceReferences(propField.Type)
	if err != nil {
		return refs, err
	}

	for i := range refs {
		refs[i].Path = append([]string{"spec", "properties"}, refs[i].Path...)
	}

	return refs, nil
}

func getResourceReferences(t reflect.Type) ([]TypeReferenceLocation, error) {
	var refs []TypeReferenceLocation
	var err error

	switch t.Kind() {
	case reflect.Ptr:
		refs, err = gatherPtr(t)
	case reflect.Struct:
		refs, err = gatherStruct(t)
	}

	return refs, err
}

func gatherPtr(t reflect.Type) ([]TypeReferenceLocation, error) {
	return getResourceReferences(t.Elem())
}

func gatherStruct(t reflect.Type) ([]TypeReferenceLocation, error) {
	var refs []TypeReferenceLocation

	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i)

		if structField.Anonymous {
			embeddedRefs, err := getResourceReferences(structField.Type)
			if err != nil {
				return refs, err
			}

			refs = append(refs, embeddedRefs...)
			continue
		}

		jsonTag, ok := structField.Tag.Lookup("json")
		if !ok {
			continue
		}

		groupTag, groupOk := structField.Tag.Lookup("group")
		kindTag, kindOk := structField.Tag.Lookup("kind")
		isSlice := structField.Type.Kind() == reflect.Slice
		isOwned := false
		if ownedTag, ownedOk := structField.Tag.Lookup("owned"); ownedOk {
			isOwned = ownedTag == "true"
		}

		jsonFieldName := strings.Split(jsonTag, ",")[0]
		var templateFieldName string
		if isSlice {
			templateFieldName = strings.TrimSuffix(jsonFieldName, "Refs") + "s"
		} else {
			templateFieldName = strings.TrimSuffix(jsonFieldName, "Ref")
		}

		switch {
		case groupOk && kindOk:
			refs = append(refs, TypeReferenceLocation{
				TemplateFieldName: templateFieldName,
				JSONFieldName:     jsonFieldName,
				Group:             groupTag,
				Kind:              kindTag,
				IsSlice:           structField.Type.Kind() == reflect.Slice,
				IsOwned:           isOwned,
			})
		default:
			references, err := getResourceReferences(structField.Type)
			if err != nil {
				return refs, err
			}

			for i := range references {
				references[i].Path = append(references[i].Path, strings.Split(jsonTag, ",")[0])
			}

			refs = append(refs, references...)
		}
	}
	return refs, nil
}

func (trl *TypeReferenceLocation) JSONFields() []string {
	return append(trl.Path, trl.JSONFieldName)
}

func (trl *TypeReferenceLocation) TemplateFields() []string {
	return append(trl.Path, trl.TemplateFieldName)
}
