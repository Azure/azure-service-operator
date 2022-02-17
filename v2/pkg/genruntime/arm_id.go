/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/pkg/errors"
)

// GetAndParseResourceID gets the ARM ID from the given MetaObject and parses it into its constituent parts
func GetAndParseResourceID(obj MetaObject) (*arm.ResourceID, error) {
	resourceID, hasResourceID := GetResourceID(obj)
	if !hasResourceID {
		return nil, errors.Errorf("cannot find resource id for obj %s/%s", obj.GetNamespace(), obj.GetName())
	}

	return arm.ParseResourceID(resourceID)
}
