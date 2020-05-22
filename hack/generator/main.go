/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	goflag "flag"
	flag "github.com/spf13/pflag"
	"k8s.io/klog/v2"

	"github.com/Azure/k8s-infra/hack/generator/cmd"
)

func main() {
	klog.InitFlags(nil)
	goflag.Parse()
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	cmd.Execute()
}
