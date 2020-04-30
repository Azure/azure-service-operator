/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/devigned/tab/opencensus"

	"contrib.go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"

	"github.com/Azure/k8s-infra/hack/generator/cmd"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if os.Getenv("TRACING") == "true" {
		closer, err := initOpenCensus()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer closer()
	}

	cmd.Execute()
}

func initOpenCensus() (func(), error) {
	exporter, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint:     "localhost:6831",
		CollectorEndpoint: "http://localhost:14268/api/traces",
		Process: jaeger.Process{
			ServiceName: "generator",
		},
	})

	if err != nil {
		return nil, err
	}

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)
	return exporter.Flush, nil
}
