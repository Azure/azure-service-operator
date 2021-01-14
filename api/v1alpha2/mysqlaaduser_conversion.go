// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha2

import "sigs.k8s.io/controller-runtime/pkg/conversion"

var _ conversion.Hub = &MySQLAADUser{}

func (*MySQLAADUser) Hub() {}
