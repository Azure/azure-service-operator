/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	goflag "flag"
	"os"

	flag "github.com/spf13/pflag"
)

func main() {
	flagSet := goflag.NewFlagSet(os.Args[0], goflag.ExitOnError)
	flagSet.Parse(os.Args[1:]) //nolint:errcheck // error will never be returned due to ExitOnError
	flag.CommandLine.AddGoFlagSet(flagSet)
	Execute()
}
