/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	. "github.com/onsi/gomega"
	"testing"
)

/*
 * TypeAsObjectType() tests
 */

func TestTypeAsObjectType_GivenWrappedType_UnwrapsExpectedObject(t *testing.T) {

	objectType := NewObjectType()
	flagged := ArmFlag.ApplyTo(objectType)
	optional := NewOptionalType(objectType)

	cases := []struct {
		name    string
		subject Type
		object  *ObjectType
	}{
		{"Object is", objectType, objectType},
		{"Flagged type unwraps", flagged, objectType},
		{"Optional type unwraps", optional, objectType},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			ot, err := TypeAsObjectType(c.subject)
			g.Expect(err).To(BeNil())
			g.Expect(ot).To(Equal(c.object))
			g.Expect(ot.Equals(c.object)).To(BeTrue())
		})
	}
}
