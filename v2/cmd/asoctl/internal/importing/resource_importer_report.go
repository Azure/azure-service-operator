/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"errors"
	"sync"

	"github.com/go-logr/logr"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
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

var resourceImportReportStatusOrder = map[resourceImportReportStatus]int{
	Imported: 0,
	Skipped:  1,
	Failed:   2,
}

// newResourceImportReport creates a new resourceImportReport
func newResourceImportReport() *resourceImportReport {
	return &resourceImportReport{
		// We don't know how many resources will be processed, but initial tests show that there are
		// often a considerable number of extension resources floating around (especially RoleAssignment)
		content: make(map[resourceImportReportKey]int, 100),
	}
}

// AddSuccessfulImport adds a successful import to the report
func (r *resourceImportReport) AddSuccessfulImport(importer ImportableResource) {
	key := resourceImportReportKey{
		group:  importer.GroupKind().Group,
		kind:   importer.GroupKind().Kind,
		status: Imported,
	}

	r.add(key)
}

// AddSkippedImport adds a skipped import to the report
func (r *resourceImportReport) AddSkippedImport(importer ImportableResource, reason string) {
	key := resourceImportReportKey{
		group:  importer.GroupKind().Group,
		kind:   importer.GroupKind().Kind,
		status: Skipped,
		reason: reason,
	}

	r.add(key)
}

// AddFailedImport adds a failed import to the report
func (r *resourceImportReport) AddFailedImport(importer ImportableResource, reason string) {
	key := resourceImportReportKey{
		group:  importer.GroupKind().Group,
		kind:   importer.GroupKind().Kind,
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
		func(left resourceImportReportKey, right resourceImportReportKey) bool {
			return left.lessThan(right)
		})

	for _, key := range keys {
		count := r.content[key]
		key.WriteToLog(log, count)
	}
}

func (k *resourceImportReportKey) lessThan(other resourceImportReportKey) bool {
	if k.group != other.group {
		return k.group < other.group
	}

	if k.kind != other.kind {
		return k.kind < other.kind
	}

	if k.status != other.status {
		return resourceImportReportStatusOrder[k.status] < resourceImportReportStatusOrder[other.status]
	}

	return k.reason < other.reason
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
