/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_CanMakeUnixStylePathIntoURL(t *testing.T) {
	g := NewGomegaWithT(t)

	url := absDirectoryPathToURL("/somewhere/over/the/rainbow")
	g.Expect(url.String()).To(Equal("file:///somewhere/over/the/rainbow/"))
}

func Test_CanMakeWindowsPathIntoURL(t *testing.T) {
	g := NewGomegaWithT(t)

	url := absDirectoryPathToURL("D:\\yellow\\brick\\road")
	g.Expect(url.String()).To(Equal("file:///D:/yellow/brick/road/"))
}
