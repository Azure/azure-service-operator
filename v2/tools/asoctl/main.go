/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	goflag "flag"
	"os"

	flag "github.com/spf13/pflag"
	"k8s.io/klog"

	"github.com/Azure/azure-service-operator/v2/tools/asoctl/cmd"
)

func main() {
	flagSet := goflag.NewFlagSet(os.Args[0], goflag.ExitOnError)
	klog.InitFlags(flagSet)
	flagSet.Parse(os.Args[1:]) //nolint:errcheck // error will never be returned due to ExitOnError
	flag.CommandLine.AddGoFlagSet(flagSet)
	cmd.Execute()
}
