/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package gen

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/Azure/k8s-infra/hack/generator/pkg/jsonast"
	"github.com/Azure/k8s-infra/hack/generator/pkg/xcobra"
)

// NewGenCommand creates a new cobra Command when invoked from the command line
func NewGenCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "gen",
		Short: "generate K8s infrastructure resources from Azure deployment template schema",
		Run: xcobra.RunWithCtx(func(ctx context.Context, cmd *cobra.Command, args []string) error {

			//TODO extract into a new type (maybe 'Generator') so that this file can concentrate on command line processing
			configuration, err := loadConfiguration("azure-cloud.yaml")
			if err != nil {
				log.Printf("Error loading configuration: %v\n", err)
				writeSampleConfig("public-cloud-sample.yaml")
				return err
			}

			err = configuration.Validate()
			if err != nil {
				log.Printf("Configuration invalid: %v\n", err)
				return err
			}

			log.Printf("Loading schema %s", configuration.SchemaURL)
			schema, err := loadSchema(configuration.SchemaURL)
			if err != nil {
				log.Printf("Failed to load schema: %v\n", err)
				return err
			}

			root := schema.Root()

			idfactory := astmodel.NewIdentifierFactory()
			scanner := jsonast.NewSchemaScanner(idfactory)
			scanner.AddFilters(viper.GetStringSlice("resources"))

			_, err = scanner.ToNodes(ctx, root)
			if err != nil {
				log.Printf("Error: %v\n", err)
				return err
			}

			err = os.RemoveAll("resources")
			if err != nil {
				log.Printf("Error: %v\n", err)
				return err
			}

			err = os.Mkdir("resources", 0700)
			if err != nil {
				log.Printf("Error %v\n", err)
				return err
			}

			log.Printf("INF Checkpoint\n")

			for _, st := range scanner.Structs {
				shouldExport, reason := configuration.ShouldExport(st)
				var motivation string
				if reason != "" {
					motivation = "because " + reason
				}
				switch shouldExport {
				case jsonast.Skip:
					log.Printf("Skipping struct %s/%s %s", st.Version(), st.Name(), motivation)

				case jsonast.Export:
					log.Printf("Exporting struct %s/%s %s", st.Version(), st.Name(), motivation)
					exportType(st)
				}
			}

			log.Printf("Completed writing %v resources\n", len(scanner.Structs))

			return nil
		}),
	}

	cmd.Flags().StringArrayP("resources", "r", nil, "list of resource type / versions to generate")
	if err := viper.BindPFlag("resources", cmd.Flags().Lookup("resources")); err != nil {
		return cmd, err
	}

	return cmd, nil
}

func loadSchema(source string) (*gojsonschema.Schema, error) {
	sl := gojsonschema.NewSchemaLoader()
	schema, err := sl.Compile(gojsonschema.NewReferenceLoader(source))
	if err != nil {
		return nil, err
	}

	return schema, nil
}

func createNamespace(base string, version string) string {
	var builder []rune

	for _, r := range base {
		builder = append(builder, rune(r))
	}

	for _, r := range version {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			builder = append(builder, rune(r))
		}
	}

	return string(builder)
}

func loadConfiguration(configurationFile string) (*jsonast.ExportConfiguration, error) {
	data, err := ioutil.ReadFile(configurationFile)
	if err != nil {
		return nil, err
	}

	result := jsonast.ExportConfiguration{}

	err = yaml.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func writeSampleConfig(configFile string) error {
	sample := jsonast.ExportConfiguration{
		TypeFilters: []*jsonast.TypeFilter{
			{
				Because: "all 2020 API versions are included",
				Action:  jsonast.IncludeType,
				Version: "2020-*",
			},
			{
				Because: "preview SDK versions are excluded by default",
				Action:  jsonast.ExcludeType,
				Version: "*preview",
			},
		},
	}

	data, err := yaml.Marshal(sample)
	if err != nil {
		return err
	}

	ioutil.WriteFile(configFile, data, os.FileMode(0644))

	return nil
}

func exportType(st *astmodel.StructDefinition) {
	ns := createNamespace("api", st.Version())
	dirName := fmt.Sprintf("resources/%v", ns)
	fileName := fmt.Sprintf("%v/%v.go", dirName, st.Name())

	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		log.Printf("Creating folder '%s'\n", dirName)
		os.Mkdir(dirName, 0700)
	}

	genFile := astmodel.NewFileDefinition(ns, st)
	genFile.SaveTo(fileName)
}
