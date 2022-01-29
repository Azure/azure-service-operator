/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

func Test_ExportFilterWithRenameTo_ErrorWhenNotIncludeTransitive(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	exportFilter := config.ExportFilter{
		Action: config.ExportFilterInclude,
		TypeMatcher: config.TypeMatcher{
			Group:   "foo",
			Version: "bar",
			Name:    "baz",
		},
		RenameTo: "Test",
	}
	err := exportFilter.Initialize()
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err.Error()).To(ContainSubstring("RenameTo can only be specified on ExportFilters with Action \"include-transitive\""))
}
