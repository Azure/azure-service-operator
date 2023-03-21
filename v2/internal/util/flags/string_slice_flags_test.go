/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package flags

import (
	"flag"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ParsesTwoParameters(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	f := SliceFlags{}

	flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
	flagSet.Var(&f, "list", "A list")
	err := flagSet.Parse([]string{"--list", "abc", "--list", "def"})
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(f).To(HaveLen(2))
	g.Expect(f[0]).To(Equal("abc"))
	g.Expect(f[1]).To(Equal("def"))
}

func Test_ParsesCommaSeparatedParameter(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	f := SliceFlags{}

	flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
	flagSet.Var(&f, "list", "A list")
	err := flagSet.Parse([]string{"--list", "abc,def"})
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(f).To(HaveLen(2))
	g.Expect(f[0]).To(Equal("abc"))
	g.Expect(f[1]).To(Equal("def"))
}
