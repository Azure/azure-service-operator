/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/hbollon/go-edlib"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/set"
)

// TypoAdvisor is a utility that helps augment errors with guidance when mistakes are made in configuration.
type TypoAdvisor struct {
	lock  sync.RWMutex
	terms set.Set[string] // set of terms we know to exist
}

func NewTypoAdvisor() *TypoAdvisor {
	return &TypoAdvisor{
		terms: set.Make[string](),
	}
}

// AddTerm records that we saw the specified item
func (advisor *TypoAdvisor) AddTerm(item string) {
	advisor.lock.Lock()
	defer advisor.lock.Unlock()
	advisor.terms.Add(item)
}

// HasTerms returns true if this advisor has any terms available
func (advisor *TypoAdvisor) HasTerms() bool {
	return len(advisor.terms) > 0
}

// Errorf creates a new error with advice, or a simple error if no advice possible
func (advisor *TypoAdvisor) Errorf(typo string, format string, args ...interface{}) error {
	advisor.lock.RLock()
	defer advisor.lock.RUnlock()

	if !advisor.HasTerms() || advisor.terms.Contains(typo) {
		// Can't make any suggestions,
		return errors.Errorf(format, args...)
	}

	msg := fmt.Sprintf(format, args...)

	suggestion, err := edlib.FuzzySearch(
		strings.ToLower(typo),
		set.AsSortedSlice(advisor.terms),
		edlib.Levenshtein)
	if err != nil {
		// Can't offer a suggestion
		return errors.Errorf(
			"%s (unable to provide suggestion: %s)",
			msg,
			err)
	}

	return errors.Errorf(
		"%s (did you mean %s?)",
		msg,
		suggestion)
}

// Wrapf adds any guidance to the provided error, if possible
func (advisor *TypoAdvisor) Wrapf(originalError error, typo string, format string, args ...interface{}) error {
	advisor.lock.RLock()
	defer advisor.lock.RUnlock()

	if !advisor.HasTerms() || advisor.terms.Contains(typo) || originalError == nil {
		// Can't make any suggestions, or don't need one
		return originalError
	}

	msg := fmt.Sprintf(format, args...)

	suggestion, err := edlib.FuzzySearch(
		strings.ToLower(typo),
		set.AsSortedSlice(advisor.terms),
		edlib.Levenshtein)
	if err != nil {
		// Can't offer a suggestion
		return errors.Wrapf(
			originalError,
			"%s (unable to provide suggestion: %s)",
			msg,
			err)
	}

	return errors.Wrapf(
		originalError,
		"%s (did you mean %s?)",
		msg,
		suggestion)
}
