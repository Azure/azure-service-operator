/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources

import (
	"errors"
	"sync"

	"github.com/go-logr/logr"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// resourceImportReport is used to generate a report of the resources that were imported (or not).
type resourceImportReport struct {
	content map[resourceImportReportKey]int // Track the number of occurrences for each resource type and result
	lock    sync.Mutex                      // Lock to protect the above map
}

// resourceImportReportKey is a key used to accumulate the results of resource imports
type resourceImportReportKey struct {
	group  string
	kind   string
	status resourceImportReportStatus
	reason string
}

// resourceImportReportStatus is the status of a resource import key
type resourceImportReportStatus string

const (
	Imported resourceImportReportStatus = "Imported"
	Skipped  resourceImportReportStatus = "Skipped"
	Failed   resourceImportReportStatus = "Failed"
)

// newResourceImportReport creates a new resourceImportReport
func newResourceImportReport() *resourceImportReport {
	return &resourceImportReport{
		// We don't know how many resources will be processed, but initial tests show that there are
		// often a considerable number of extension resources floating around (especially RoleAssignment)
		content: make(map[resourceImportReportKey]int, 100),
	}
}

// AddSuccessfulImport adds a successful import to the report
func (r *resourceImportReport) AddSuccessfulImport(
	gk schema.GroupKind,
) {
	key := resourceImportReportKey{
		group:  gk.Group,
		kind:   gk.Kind,
		status: Imported,
	}

	r.add(key)
}

// AddSkippedImport adds a skipped import to the report
func (r *resourceImportReport) AddSkippedImport(
	gk schema.GroupKind,
	reason string,
) {
	key := resourceImportReportKey{
		group:  gk.Group,
		kind:   gk.Kind,
		status: Skipped,
		reason: reason,
	}

	r.add(key)
}

// AddFailedImport adds a failed import to the report
func (r *resourceImportReport) AddFailedImport(
	gk schema.GroupKind,
	reason string,
) {
	key := resourceImportReportKey{
		group:  gk.Group,
		kind:   gk.Kind,
		status: Failed,
		reason: reason,
	}

	r.add(key)
}

func (r *resourceImportReport) add(key resourceImportReportKey) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.content[key]++
}

func (r *resourceImportReport) WriteToLog(log logr.Logger) {
	keys := maps.Keys(r.content)
	slices.SortFunc(
		keys,
		func(left resourceImportReportKey, right resourceImportReportKey) int {
			return left.compareTo(right)
		})

	for _, key := range keys {
		count := r.content[key]
		key.WriteToLog(log, count)
	}
}

func (k *resourceImportReportKey) compareTo(other resourceImportReportKey) int {
	if k.group < other.group {
		return -1
	} else if k.group > other.group {
		return 1
	}

	if k.kind < other.kind {
		return -1
	} else if k.kind > other.kind {
		return 1
	}

	if k.status < other.status {
		return -1
	} else if k.status > other.status {
		return 1
	}

	if k.reason < other.reason {
		return -1
	} else if k.reason > other.reason {
		return 1
	}

	return 0
}

func (k *resourceImportReportKey) WriteToLog(log logr.Logger, count int) {
	switch k.status {
	case Imported:
		log.Info(
			"Successful imports",
			"Group", k.group,
			"Kind", k.kind,
			"Count", count)
	case Skipped:
		log.V(1).Info(
			"Skipped imports",
			"Group", k.group,
			"Kind", k.kind,
			"Count", count,
			"Reason", k.reason)
	case Failed:
		log.Error(
			errors.New(k.reason),
			"Failed imports",
			"Group", k.group,
			"Kind", k.kind,
			"Count", count)
	}
}
