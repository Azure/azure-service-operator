/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	goflag "flag"
	"os"
	"runtime/debug"

	flag "github.com/spf13/pflag"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/tools/asoctl/cmd"
)

func main() {
	flagSet := goflag.NewFlagSet(os.Args[0], goflag.ExitOnError)
	klog.InitFlags(flagSet)
	flagSet.Parse(os.Args[1:]) //nolint:errcheck // error will never be returned due to ExitOnError
	flag.CommandLine.AddGoFlagSet(flagSet)

	defer func() {
		// If a panic occurs, log details as debug
		if r := recover(); r != nil {
			klog.ErrorDepth(1, r)
			klog.V(1).InfoDepth(1, string(debug.Stack()))
			os.Exit(1)
		}
	}()

	cmd.Execute()
}
