// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package conversion

import (
	"testing"

	. "github.com/onsi/gomega"

	"k8s.io/utils/ptr"

	"github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20250801/storage/compat"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func TestPullAutoscaleProfiles(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name    string
		value   string
		want    []compat.AutoScaleProfile
		wantErr bool
	}{
		{name: "singleton", value: `{"minCount":1}`, want: []compat.AutoScaleProfile{{MinCount: ptr.To(1)}}},
		{name: "whitespace", value: " \n {\"minCount\":1}", want: []compat.AutoScaleProfile{{MinCount: ptr.To(1)}}},
		{name: "list", value: `[{"minCount":1},{"minCount":2}]`, want: []compat.AutoScaleProfile{{MinCount: ptr.To(1)}, {MinCount: ptr.To(2)}}},
		{name: "emptyList", value: `[]`, want: []compat.AutoScaleProfile{}},
		{name: "null", value: `null`},
		{name: "malformedSingleton", value: `{`, wantErr: true},
		{name: "malformedList", value: `[`, wantErr: true},
		{name: "unknownSingletonField", value: `{"unexpected":1}`, wantErr: true},
		{name: "unknownListField", value: `[{"unexpected":1}]`, wantErr: true},
		{name: "wrongValueType", value: `{"minCount":"invalid"}`, wantErr: true},
		{name: "wrongKind", value: `true`, wantErr: true},
		{name: "emptyValue", wantErr: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			bag := genruntime.PropertyBag{"Autoscale": c.value, "other": "retained"}

			profiles, err := PullAutoscaleProfiles[compat.AutoScaleProfile](bag)
			if c.wantErr {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(profiles).To(Equal(c.want))
			}
			g.Expect(bag).To(Equal(genruntime.PropertyBag{"other": "retained"}))
		})
	}

	t.Run("absent", func(t *testing.T) {
		t.Parallel()
		g := NewGomegaWithT(t)
		profiles, err := PullAutoscaleProfiles[compat.AutoScaleProfile](nil)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(profiles).To(BeNil())
	})
}
