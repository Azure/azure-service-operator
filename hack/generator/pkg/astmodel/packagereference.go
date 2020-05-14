/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "path/filepath"

// PackageReference indicates which package
// a struct belongs to.
type PackageReference struct {
	groupName   string
	packageName string
}

// PackagePath is the path to the package reference
func (pr *PackageReference) PackagePath() string {
	return filepath.Join(pr.GroupName(), pr.PackageName())
}

// GroupName is the group name of the package reference
func (pr *PackageReference) GroupName() string {
	return pr.groupName
}

// PackageName is the package name of the package reference
func (pr *PackageReference) PackageName() string {
	return pr.packageName
}
