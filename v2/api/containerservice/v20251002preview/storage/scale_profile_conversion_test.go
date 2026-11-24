// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package storage

import (
	"testing"

	. "github.com/onsi/gomega"

	"k8s.io/utils/ptr"

	v20250801s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20250801/storage"
	hub "github.com/Azure/azure-service-operator/v2/api/containerservice/v20260501/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func TestScaleProfile_AutoscaleRoundTripsThroughHub(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := ScaleProfile{
		Autoscale: &AutoScaleProfile{
			MinCount: ptr.To(1),
			MaxCount: ptr.To(5),
			Size:     ptr.To("Standard_D2s_v3"),
		},
		Manual:      []ManualScaleProfile{{Count: ptr.To(2)}},
		PropertyBag: genruntime.PropertyBag{"unrelated": "retained"},
	}

	var intermediate v20250801s.ScaleProfile
	g.Expect(original.AssignProperties_To_ScaleProfile(&intermediate)).To(Succeed())

	var newest hub.ScaleProfile
	g.Expect(intermediate.AssignProperties_To_ScaleProfile(&newest)).To(Succeed())
	g.Expect(newest.Autoscale).To(Equal([]hub.AutoScaleProfile{{
		MinCount: ptr.To(1),
		MaxCount: ptr.To(5),
		Size:     ptr.To("Standard_D2s_v3"),
	}}))

	var restoredIntermediate v20250801s.ScaleProfile
	g.Expect(restoredIntermediate.AssignProperties_From_ScaleProfile(&newest)).To(Succeed())

	var restored ScaleProfile
	g.Expect(restored.AssignProperties_From_ScaleProfile(&restoredIntermediate)).To(Succeed())
	g.Expect(restored).To(Equal(original))
}

func TestScaleProfileStatus_AutoscaleRoundTripsThroughHub(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := ScaleProfile_STATUS{
		Autoscale: &AutoScaleProfile_STATUS{
			MinCount: ptr.To(1),
			MaxCount: ptr.To(5),
			Size:     ptr.To("Standard_D2s_v3"),
		},
		Manual:      []ManualScaleProfile_STATUS{{Count: ptr.To(2)}},
		PropertyBag: genruntime.PropertyBag{"unrelated": "retained"},
	}

	var intermediate v20250801s.ScaleProfile_STATUS
	g.Expect(original.AssignProperties_To_ScaleProfile_STATUS(&intermediate)).To(Succeed())

	var newest hub.ScaleProfile_STATUS
	g.Expect(intermediate.AssignProperties_To_ScaleProfile_STATUS(&newest)).To(Succeed())
	g.Expect(newest.Autoscale).To(Equal([]hub.AutoScaleProfile_STATUS{{
		MinCount: ptr.To(1),
		MaxCount: ptr.To(5),
		Size:     ptr.To("Standard_D2s_v3"),
	}}))

	var restoredIntermediate v20250801s.ScaleProfile_STATUS
	g.Expect(restoredIntermediate.AssignProperties_From_ScaleProfile_STATUS(&newest)).To(Succeed())

	var restored ScaleProfile_STATUS
	g.Expect(restored.AssignProperties_From_ScaleProfile_STATUS(&restoredIntermediate)).To(Succeed())
	g.Expect(restored).To(Equal(original))
}

func TestScaleProfile_AutoscaleListIsReplacedByPreviewEdits(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := hub.ScaleProfile{
		Autoscale: []hub.AutoScaleProfile{
			{MinCount: ptr.To(1), Size: ptr.To("first")},
			{MinCount: ptr.To(2), Size: ptr.To("second"), PropertyBag: genruntime.PropertyBag{"future": "retained"}},
			{MinCount: ptr.To(3), Size: ptr.To("third")},
		},
		PropertyBag: genruntime.PropertyBag{"unrelated": "retained"},
	}
	snapshot := original.DeepCopy()

	var intermediate v20250801s.ScaleProfile
	g.Expect(intermediate.AssignProperties_From_ScaleProfile(&original)).To(Succeed())
	bagSnapshot := genruntime.NewPropertyBag(intermediate.PropertyBag)

	var preview ScaleProfile
	g.Expect(preview.AssignProperties_From_ScaleProfile(&intermediate)).To(Succeed())
	g.Expect(preview.Autoscale).To(Equal(&AutoScaleProfile{MinCount: ptr.To(1), Size: ptr.To("first")}))
	g.Expect(intermediate.PropertyBag).To(Equal(bagSnapshot))
	g.Expect(preview.PropertyBag).To(Equal(bagSnapshot))

	var restoredIntermediate v20250801s.ScaleProfile
	g.Expect(preview.AssignProperties_To_ScaleProfile(&restoredIntermediate)).To(Succeed())
	var restored hub.ScaleProfile
	g.Expect(restoredIntermediate.AssignProperties_To_ScaleProfile(&restored)).To(Succeed())
	g.Expect(restored).To(Equal(*snapshot))
	g.Expect(preview.PropertyBag).To(Equal(bagSnapshot))

	preview.Autoscale.MinCount = ptr.To(5)
	g.Expect(preview.AssignProperties_To_ScaleProfile(&restoredIntermediate)).To(Succeed())
	g.Expect(restoredIntermediate.AssignProperties_To_ScaleProfile(&restored)).To(Succeed())

	want := snapshot.DeepCopy()
	want.Autoscale = want.Autoscale[:1]
	want.Autoscale[0].MinCount = ptr.To(5)
	g.Expect(restored).To(Equal(*want))
	g.Expect(original).To(Equal(*snapshot))
	g.Expect(preview.PropertyBag).To(Equal(bagSnapshot))

	preview.Autoscale = nil
	g.Expect(preview.AssignProperties_To_ScaleProfile(&restoredIntermediate)).To(Succeed())
	g.Expect(restoredIntermediate.AssignProperties_To_ScaleProfile(&restored)).To(Succeed())
	g.Expect(restored.Autoscale).To(BeNil())
	g.Expect(restored.PropertyBag).To(Equal(genruntime.PropertyBag{"unrelated": "retained"}))
}

func TestScaleProfileStatus_AutoscaleListIsReplacedByPreviewEdits(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	original := hub.ScaleProfile_STATUS{
		Autoscale: []hub.AutoScaleProfile_STATUS{
			{MinCount: ptr.To(1), Size: ptr.To("first")},
			{MinCount: ptr.To(2), Size: ptr.To("second"), PropertyBag: genruntime.PropertyBag{"future": "retained"}},
			{MinCount: ptr.To(3), Size: ptr.To("third")},
		},
		PropertyBag: genruntime.PropertyBag{"unrelated": "retained"},
	}
	snapshot := original.DeepCopy()

	var intermediate v20250801s.ScaleProfile_STATUS
	g.Expect(intermediate.AssignProperties_From_ScaleProfile_STATUS(&original)).To(Succeed())
	bagSnapshot := genruntime.NewPropertyBag(intermediate.PropertyBag)

	var preview ScaleProfile_STATUS
	g.Expect(preview.AssignProperties_From_ScaleProfile_STATUS(&intermediate)).To(Succeed())
	g.Expect(preview.Autoscale).To(Equal(&AutoScaleProfile_STATUS{MinCount: ptr.To(1), Size: ptr.To("first")}))
	g.Expect(intermediate.PropertyBag).To(Equal(bagSnapshot))
	g.Expect(preview.PropertyBag).To(Equal(bagSnapshot))

	var restoredIntermediate v20250801s.ScaleProfile_STATUS
	g.Expect(preview.AssignProperties_To_ScaleProfile_STATUS(&restoredIntermediate)).To(Succeed())
	var restored hub.ScaleProfile_STATUS
	g.Expect(restoredIntermediate.AssignProperties_To_ScaleProfile_STATUS(&restored)).To(Succeed())
	g.Expect(restored).To(Equal(*snapshot))
	g.Expect(preview.PropertyBag).To(Equal(bagSnapshot))

	preview.Autoscale.MinCount = ptr.To(5)
	g.Expect(preview.AssignProperties_To_ScaleProfile_STATUS(&restoredIntermediate)).To(Succeed())
	g.Expect(restoredIntermediate.AssignProperties_To_ScaleProfile_STATUS(&restored)).To(Succeed())

	want := snapshot.DeepCopy()
	want.Autoscale = want.Autoscale[:1]
	want.Autoscale[0].MinCount = ptr.To(5)
	g.Expect(restored).To(Equal(*want))
	g.Expect(original).To(Equal(*snapshot))
	g.Expect(preview.PropertyBag).To(Equal(bagSnapshot))

	preview.Autoscale = nil
	g.Expect(preview.AssignProperties_To_ScaleProfile_STATUS(&restoredIntermediate)).To(Succeed())
	g.Expect(restoredIntermediate.AssignProperties_To_ScaleProfile_STATUS(&restored)).To(Succeed())
	g.Expect(restored.Autoscale).To(BeNil())
	g.Expect(restored.PropertyBag).To(Equal(genruntime.PropertyBag{"unrelated": "retained"}))
}

func TestScaleProfile_AutoscaleAbsenceClearsReusedDestination(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var source v20250801s.ScaleProfile
	preview := ScaleProfile{Autoscale: &AutoScaleProfile{MinCount: ptr.To(1)}}
	g.Expect(preview.AssignProperties_From_ScaleProfile(&source)).To(Succeed())
	g.Expect(preview.Autoscale).To(BeNil())
	g.Expect(preview.PropertyBag).To(BeNil())

	var statusSource v20250801s.ScaleProfile_STATUS
	previewStatus := ScaleProfile_STATUS{Autoscale: &AutoScaleProfile_STATUS{MinCount: ptr.To(1)}}
	g.Expect(previewStatus.AssignProperties_From_ScaleProfile_STATUS(&statusSource)).To(Succeed())
	g.Expect(previewStatus.Autoscale).To(BeNil())
	g.Expect(previewStatus.PropertyBag).To(BeNil())
}

func TestScaleProfile_InvalidAutoscaleIsReportedWithoutMutatingSource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	source := v20250801s.ScaleProfile{PropertyBag: genruntime.PropertyBag{"Autoscale": `{"unexpected":1}`}}
	snapshot := genruntime.NewPropertyBag(source.PropertyBag)
	var preview ScaleProfile
	g.Expect(preview.AssignProperties_From_ScaleProfile(&source)).NotTo(Succeed())
	g.Expect(source.PropertyBag).To(Equal(snapshot))

	var newest hub.ScaleProfile
	g.Expect(source.AssignProperties_To_ScaleProfile(&newest)).NotTo(Succeed())
	g.Expect(source.PropertyBag).To(Equal(snapshot))

	statusSource := v20250801s.ScaleProfile_STATUS{PropertyBag: genruntime.PropertyBag{"Autoscale": `{"unexpected":1}`}}
	var previewStatus ScaleProfile_STATUS
	g.Expect(previewStatus.AssignProperties_From_ScaleProfile_STATUS(&statusSource)).NotTo(Succeed())
	g.Expect(statusSource.PropertyBag).To(Equal(snapshot))

	var newestStatus hub.ScaleProfile_STATUS
	g.Expect(statusSource.AssignProperties_To_ScaleProfile_STATUS(&newestStatus)).NotTo(Succeed())
	g.Expect(statusSource.PropertyBag).To(Equal(snapshot))
}

func TestScaleProfile_LegacySingletonPropertyBagCanBeRead(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	bag := genruntime.PropertyBag{"Autoscale": `{"minCount":1,"size":"legacy"}`, "unrelated": "retained"}
	legacy := v20250801s.ScaleProfile{PropertyBag: bag}
	var newest hub.ScaleProfile
	g.Expect(legacy.AssignProperties_To_ScaleProfile(&newest)).To(Succeed())
	g.Expect(newest.Autoscale).To(Equal([]hub.AutoScaleProfile{{MinCount: ptr.To(1), Size: ptr.To("legacy")}}))
	g.Expect(newest.PropertyBag).To(Equal(genruntime.PropertyBag{"unrelated": "retained"}))

	var preview ScaleProfile
	g.Expect(preview.AssignProperties_From_ScaleProfile(&legacy)).To(Succeed())
	g.Expect(preview.Autoscale).To(Equal(&AutoScaleProfile{MinCount: ptr.To(1), Size: ptr.To("legacy")}))

	legacyStatus := v20250801s.ScaleProfile_STATUS{PropertyBag: bag}
	var newestStatus hub.ScaleProfile_STATUS
	g.Expect(legacyStatus.AssignProperties_To_ScaleProfile_STATUS(&newestStatus)).To(Succeed())
	g.Expect(newestStatus.Autoscale).To(Equal([]hub.AutoScaleProfile_STATUS{{MinCount: ptr.To(1), Size: ptr.To("legacy")}}))
	g.Expect(newestStatus.PropertyBag).To(Equal(genruntime.PropertyBag{"unrelated": "retained"}))

	var previewStatus ScaleProfile_STATUS
	g.Expect(previewStatus.AssignProperties_From_ScaleProfile_STATUS(&legacyStatus)).To(Succeed())
	g.Expect(previewStatus.Autoscale).To(Equal(&AutoScaleProfile_STATUS{MinCount: ptr.To(1), Size: ptr.To("legacy")}))
	g.Expect(bag).To(Equal(genruntime.PropertyBag{"Autoscale": `{"minCount":1,"size":"legacy"}`, "unrelated": "retained"}))
}
