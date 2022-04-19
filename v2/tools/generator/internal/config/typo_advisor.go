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
	terms set.StringSet // set of terms we know to exist
}

func NewTypoAdvisor() *TypoAdvisor {
	return &TypoAdvisor{
		terms: set.MakeStringSet(),
	}
}

// AddTerm records that we saw the specified item
func (advisor *TypoAdvisor) AddTerm(item string) {
	advisor.lock.Lock()
	defer advisor.lock.Unlock()
	advisor.terms.Add(item)
}

// Wrapf adds any guidance to the provided error, if possible
func (advisor *TypoAdvisor) Wrapf(originalError error, typo string, format string, args ...interface{}) error {
	advisor.lock.RLock()
	defer advisor.lock.RUnlock()

	if len(advisor.terms) == 0 {
		// Can't make any suggestions
		return originalError
	}

	if advisor.terms.Contains(typo) {
		// No suggestion needed
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
