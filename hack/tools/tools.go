// +build tools

package tools

import (
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"     //nolint
	_ "github.com/golangci/golangci-lint" //nolint
)
