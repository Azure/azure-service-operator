/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package gen

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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
				err2 := writeSampleConfig("public-cloud-sample.yaml")
				if err2 != nil {
					log.Printf("Failed to write sample configuration file, err: %v\n", err2)
				}
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

			rootOutputDir := "apis" // TODO: command-line argument

			idfactory := astmodel.NewIdentifierFactory()
			scanner := jsonast.NewSchemaScanner(idfactory)
			scanner.AddFilters(viper.GetStringSlice("resources"))

			_, err = scanner.ToNodes(ctx, root)
			if err != nil {
				log.Printf("Error: %v\n", err)
				return err
			}

			err = os.RemoveAll(rootOutputDir)
			if err != nil {
				log.Printf("Error: %v\n", err)
				return err
			}

			err = os.Mkdir(rootOutputDir, 0700)
			if err != nil {
				log.Printf("Error %v\n", err)
				return err
			}

			log.Printf("INF Checkpoint\n")

			// group definitions by package
			packages := make(map[astmodel.PackageReference]*astmodel.PackageDefinition)
			for _, def := range scanner.Definitions {

				shouldExport, reason := configuration.ShouldExport(def)
				var motivation string
				if reason != "" {
					motivation = "because " + reason
				}

				defRef := def.Reference()

				switch shouldExport {
				case jsonast.Skip:
					log.Printf("Skipping %s/%s %s", defRef.PackagePath(), defRef.Name(), motivation)

				case jsonast.Export:
					log.Printf("Will export %s/%s %s", defRef.PackagePath(), defRef.Name(), motivation)

					pkgRef := defRef.PackageReference
					groupName, pkgName, err := pkgRef.GroupAndPackage()
					if err != nil {
						return err
					}
					if pkg, ok := packages[pkgRef]; ok {
						pkg.AddDefinition(def)
					} else {
						pkg = astmodel.NewPackageDefinition(groupName, pkgName)
						pkg.AddDefinition(def)
						packages[pkgRef] = pkg
					}
				}
			}

			// emit each package
			for _, pkg := range packages {

				// create directory if not already there
				outputDir := filepath.Join(rootOutputDir, pkg.GroupName, pkg.PackageName)
				if _, err := os.Stat(outputDir); os.IsNotExist(err) {
					//log.Printf("Creating directory '%s'\n", outputDir)
					err = os.MkdirAll(outputDir, 0700)
					if err != nil {
						log.Fatalf("Unable to create directory '%s'", outputDir)
					}
				}

				err = pkg.EmitDefinitions(outputDir)
				if err != nil {
					log.Fatalf("Unable to emit definitions '%v'", err)
				}
			}

			log.Printf("Completed writing %v resources\n", len(scanner.Definitions))

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

	err = ioutil.WriteFile(configFile, data, os.FileMode(0644))
	if err != nil {
		return err
	}

	return nil
}
