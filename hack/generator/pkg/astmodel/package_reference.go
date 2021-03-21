/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"sort"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

const (
	genRuntimePathPrefix  = "github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
	GenRuntimePackageName = "genruntime"
	GroupSuffix           = ".infra.azure.com"
)

var MetaV1PackageReference = MakeExternalPackageReference("k8s.io/apimachinery/pkg/apis/meta/v1")

type PackageReference interface {
	// AsLocalPackage attempts conversion to a LocalPackageReference
	AsLocalPackage() (LocalPackageReference, bool)
	// Package returns the package name of this reference
	PackageName() string
	// PackagePath returns the fully qualified package path
	PackagePath() string
	// Equals returns true if the passed package reference references the same package, false otherwise
	Equals(ref PackageReference) bool
	// String returns the string representation of the package reference
	String() string
	// IsPreview returns true if this package reference has a suffix indicating it's a preview
	// release, false otherwise
	IsPreview() bool
}

// PackageAsLocalPackage converts the given PackageReference into a LocalPackageReference if possible.
// If the provided PackageReference does not represent a local package an error is returned.
func PackageAsLocalPackage(pkg PackageReference) (LocalPackageReference, error) {
	if localPkg, ok := pkg.AsLocalPackage(); ok {
		return localPkg, nil
	}
	return LocalPackageReference{}, errors.Errorf("%q is not a local package", pkg.PackagePath())
}

func SortPackageReferencesByPathAndVersion(packages []PackageReference) {
	sort.Slice(packages, func(i, j int) bool {
		comparer := versionComparer{
			left:  []rune(packages[i].PackagePath()),
			right: []rune(packages[j].PackagePath()),
		}
		compare := comparer.Compare()
		return compare < 0
	})
}

// versionComparer captures our state while doing an alphanumeric version comparision
// We need separate indexes for each side because we're doing a numeric comparison, which will
// compare "100" and "0100" as equal (leading zeros are not significant)
type versionComparer struct {
	left       []rune
	leftIndex  int
	right      []rune
	rightIndex int
}

func (v *versionComparer) Compare() int {
	for {
		if v.leftIndex >= len(v.left) && v.rightIndex >= len(v.right) {
			// Ran out of both arrays at the same time
			break
		}

		if v.leftIndex >= len(v.left) {
			// Ran out of the left array only; if the right array has an identifier, right goes first, otherwise left
			if unicode.IsLetter(v.right[v.rightIndex]) {
				return 1
			}

			return -1
		}

		if v.rightIndex >= len(v.right) {
			// Ran out of the right array only; if the left array has an identifier, left goes first, otherwise right
			if unicode.IsLetter(v.left[v.leftIndex]) {
				return -1
			}

			return 1
		}

		leftRune := v.left[v.leftIndex]
		rightRune := v.right[v.rightIndex]

		if unicode.IsDigit(leftRune) && unicode.IsDigit(rightRune) {
			// Found the start of a number
			compare := v.compareNumeric()
			if compare != 0 {
				return compare
			}

			continue
		}

		if unicode.IsLetter(leftRune) && unicode.IsLetter(rightRune) {
			// Found the start of an identifier
			compare := v.compareIdentifier()
			if compare != 0 {
				return compare
			}

			continue
		}

		if leftRune == rightRune {
			// Both runes the same, skip to the next one
			v.leftIndex++
			v.rightIndex++
			continue
		}

		// Runes are different, make a decision
		if leftRune < rightRune {
			return -1
		} else {
			return 1
		}

	}

	return 0
}

// compareNumeric compares two digit sequences as though they represent an integer number. We don't
// convert the number to a literal int because we don't want to run the risk of overflow, but
// fortunately we can compare integer numbers by length and digit by digit.
func (v *versionComparer) compareNumeric() int {
	// Start by skipping any leading zeros as they don't change the value
	v.leftIndex = v.endOfSpan(v.left, v.leftIndex, v.IsZero)
	v.rightIndex = v.endOfSpan(v.right, v.rightIndex, v.IsZero)

	// Find the length of each digit sequence
	leftLength := v.endOfSpan(v.left, v.leftIndex, unicode.IsDigit) - v.leftIndex
	rightLength := v.endOfSpan(v.right, v.rightIndex, unicode.IsDigit) - v.rightIndex

	// A longer number is larger
	// (this is safe because we've already skipped leading zeros)
	if leftLength > rightLength {
		return 1
	} else if leftLength < rightLength {
		return -1
	}

	// Both digit sequences are the same length.
	// Need to compare digit by digit, without falling off the end of the sequence
	length := leftLength
	for length > 0 {
		if v.left[v.leftIndex] != v.right[v.rightIndex] {
			// Found different digits
			break
		}

		v.leftIndex++
		v.rightIndex++
		length--
	}

	if length == 0 {
		// Ran out of digits, they're equal
		return 0
	}

	if v.left[v.leftIndex] < v.right[v.rightIndex] {
		return -1
	}

	return 1
}

// compareIdentifier compares two letter sequences, with certain identifiers being given priority in the sequence
func (v *versionComparer) compareIdentifier() int {
	// Find the end of each letter sequence and extract the strings
	leftEnd := v.endOfSpan(v.left, v.leftIndex, unicode.IsLetter)
	rightEnd := v.endOfSpan(v.right, v.rightIndex, unicode.IsLetter)

	left := string(v.left[v.leftIndex:leftEnd])
	right := string(v.right[v.rightIndex:rightEnd])

	// Early exit if the strings are equal
	if left == right {
		v.leftIndex = leftEnd
		v.rightIndex = rightEnd
		return 0
	}

	// Check to see if we have any special identifiers
	leftRank, leftIsSpecial := v.isPreviewVersionLabel(left)
	rightRank, rightIsSpecial := v.isPreviewVersionLabel(right)

	// Check to see if both identifiers are special
	if leftIsSpecial && rightIsSpecial {
		// Don't need to check for equality because we did that above
		if leftRank < rightRank {
			return -1
		}

		return 1
	}

	// They're not both special - but we might have one
	if leftIsSpecial {
		// Left is special, right is not
		return -1
	}
	if rightIsSpecial {
		// Left is not special, right is
		return 1
	}

	// Neither is special, just compare them
	if left < right {
		return -1
	}

	return 1
}

func (v *versionComparer) endOfSpan(runes []rune, start int, predicate func(rune) bool) int {
	index := start
	for index < len(runes) {
		if !predicate(runes[index]) {
			break
		}

		index++
	}

	return index
}

func (v *versionComparer) IsZero(r rune) bool {
	return r == '0'
}

// previewVersionLabels is a sequence of specially treated identifiers
// These come before all others and are compared in the order listed here
var previewVersionLabels []string = []string{
	"alpha",
	"beta",
	"preview",
}

// isPreviewVersionLabel checks the passed identifier to see if it is one of our special set, and
// if so returns its rank and true. If the passed identifier is not special, returns -1 and false.
func (v *versionComparer) isPreviewVersionLabel(identifier string) (int, bool) {
	for rank, id := range previewVersionLabels {
		if identifier == id {
			return rank, true
		}
	}

	return -1, false
}

// containsPreviewVersionLabel checks the passed identifier to see if it contains one of our
// special set, and if so returns its true. If the passed identifier does not contain one,
// returns false.
func containsPreviewVersionLabel(identifier string) bool {
	for _, id := range previewVersionLabels {
		if strings.Contains(identifier, id) {
			return true
		}
	}

	return false
}
