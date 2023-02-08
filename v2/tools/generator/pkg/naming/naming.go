/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package naming

import "github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"

var idFactory = astmodel.NewIdentifierFactory()

// Singularize transforms the passed name into its singular form.
func Singularize(name string) string {
	return astmodel.Singularize(name, idFactory)
}
