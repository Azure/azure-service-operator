/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"os"
	"runtime/debug"

	flag "github.com/spf13/pflag"

	"github.com/Azure/azure-service-operator/v2/tools/asoctl/cmd"
)

func main() {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	cmd.Configure(flagSet)

	flagSet.Parse(os.Args[1:]) //nolint:errcheck // error will never be returned due to ExitOnError
	flag.CommandLine.AddFlagSet(flagSet)

	log := cmd.Logger()

	defer func() {
		// If a panic occurs, logging details as debug
		if r := recover(); r != nil {
			log.Info("Panic occurred", "panic", r)
			log.V(1).Info(string(debug.Stack()))
			os.Exit(1)
		}
	}()

	cmd.Execute()
}
