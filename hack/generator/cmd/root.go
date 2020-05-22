/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/klog/v2"

	"github.com/Azure/k8s-infra/hack/generator/cmd/gen"
)

// Execute kicks off the command line
func Execute() {
	cmd, err := newRootCommand()
	if err != nil {
		klog.Fatalf("fatal error: commands failed to build! %v\n", err)
	}

	if err := cmd.Execute(); err != nil {
		klog.Fatalln(err)
	}
}

func newRootCommand() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:              "k8sinfra",
		Short:            "k8sinfra provides a cmdline interface for generating k8s-infra types from Azure deployment template schema",
		TraverseChildren: true,
	}

	rootCmd.Flags().SortFlags = false

	var cfgFile string
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (defaults are ./config.yaml, $HOME/.k8sinfra/config.yaml, /etc/k8sinfra/config.yaml)")
	cobra.OnInitialize(func() {
		initConfig(&cfgFile)
	})

	cmdFuncs := []func() (*cobra.Command, error){
		gen.NewGenCommand,
	}

	for _, f := range cmdFuncs {
		cmd, err := f()
		if err != nil {
			return rootCmd, err
		}
		rootCmd.AddCommand(cmd)
	}

	return rootCmd, nil
}

func initConfig(cfgFilePtr *string) {
	cfgFile := *cfgFilePtr
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("/etc/k8sinfra/")
		viper.AddConfigPath("$HOME/.k8sinfra")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; set sane defaults and carry on.
			fmt.Println("Configuration file not found.")
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("Fatal error reading config file: %s", err))
		}
	}
	fmt.Println("Found configuration file.")
}
