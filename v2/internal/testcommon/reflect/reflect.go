/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reflect

import (
	"fmt"
	"reflect"
)

// PopulateStruct is far from complete/perfect and shouldn't be used in production code.
// Test usage is OK.
// deprecated: Don't use it in production code.
func PopulateStruct(s any) {
	val := reflect.ValueOf(s).Elem()
	valType := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := valType.Field(i)

		switch field.Kind() {
		case reflect.Ptr:
			newVal := reflect.New(fieldType.Type.Elem())
			field.Set(newVal)
			PopulateStruct(field.Interface())
		case reflect.Struct:
			PopulateStruct(field.Addr().Interface())
		case reflect.String:
			field.Set(reflect.ValueOf("exampleValue"))
		case reflect.Map:
			// skip
		default:
			panic(fmt.Sprintf("unsupported field type: %s", field.Kind()))
		}
	}
}
