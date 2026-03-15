/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package podmetrics

import (
	"encoding/csv"
	"fmt"
	"io"
	"time"
)

// Sample is a single point-in-time measurement of a container's resource usage.
type Sample struct {
	// Timestamp is when this sample was taken (wall clock).
	Timestamp time.Time

	// Elapsed is the duration since metrics collection started.
	Elapsed time.Duration

	// PodName is the name of the pod.
	PodName string

	// ContainerName is the name of the container within the pod.
	ContainerName string

	// CPUMillicores is the CPU usage in millicores.
	CPUMillicores float64

	// MemoryBytes is the memory usage in bytes.
	MemoryBytes int64
}

// csvHeader returns the CSV column headers for metrics samples.
func csvHeader() []string {
	return []string{
		"timestamp",
		"elapsed_seconds",
		"pod",
		"container",
		"cpu_millicores",
		"memory_mb",
	}
}

// CSVRow converts a Sample to a CSV row of strings.
func (s *Sample) CSVRow() []string {
	return []string{
		s.Timestamp.Format(time.RFC3339),
		fmt.Sprintf("%.1f", s.Elapsed.Seconds()),
		s.PodName,
		s.ContainerName,
		fmt.Sprintf("%.1f", s.CPUMillicores),
		fmt.Sprintf("%.2f", float64(s.MemoryBytes)/(1000*1000)),
	}
}

// WriteCSV writes the given samples to a writer in CSV format.
func WriteCSV(w io.Writer, samples []Sample) error {
	cw := csv.NewWriter(w)
	defer cw.Flush()

	if err := cw.Write(csvHeader()); err != nil {
		return fmt.Errorf("writing CSV header: %w", err)
	}

	for i := range samples {
		if err := cw.Write(samples[i].CSVRow()); err != nil {
			return fmt.Errorf("writing CSV row %d: %w", i, err)
		}
	}

	return nil
}
