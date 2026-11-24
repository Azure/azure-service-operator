// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package storage

import (
	"reflect"

	"github.com/Azure/azure-service-operator/v2/api/containerservice/internal/conversion"
	v20250801s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20250801/storage"
	"github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20250801/storage/compat"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

var (
	_ augmentConversionForScaleProfile        = &ScaleProfile{}
	_ augmentConversionForScaleProfile_STATUS = &ScaleProfile_STATUS{}
)

func (profile *ScaleProfile) AssignPropertiesFrom(src *v20250801s.ScaleProfile) error {
	bag := genruntime.NewPropertyBag(profile.PropertyBag)
	// Retain multi-entry lists so conversion back can distinguish edits from an unchanged singleton.
	profiles, err := conversion.PullAutoscaleProfiles[compat.AutoScaleProfile](genruntime.NewPropertyBag(bag))
	if err != nil {
		return err
	}

	profile.Autoscale = nil
	if len(profiles) > 0 {
		profile.Autoscale = &AutoScaleProfile{}
		if err := profile.Autoscale.AssignProperties_From_AutoScaleProfile(&profiles[0]); err != nil {
			return err
		}
	}
	if len(profiles) <= 1 {
		bag.Remove("Autoscale")
	}

	if len(bag) == 0 {
		bag = nil
	}
	profile.PropertyBag = bag
	return nil
}

func (profile *ScaleProfile) AssignPropertiesTo(dst *v20250801s.ScaleProfile) error {
	bag := genruntime.NewPropertyBag(dst.PropertyBag)
	profiles, err := conversion.PullAutoscaleProfiles[compat.AutoScaleProfile](bag)
	if err != nil {
		return err
	}

	if profile.Autoscale != nil {
		var original *AutoScaleProfile
		if len(profiles) > 0 {
			original = &AutoScaleProfile{}
			if err := original.AssignProperties_From_AutoScaleProfile(&profiles[0]); err != nil {
				return err
			}
		}

		// Editing the singleton replaces the list; an unchanged singleton preserves it.
		if !reflect.DeepEqual(profile.Autoscale, original) {
			profiles = make([]compat.AutoScaleProfile, 1)
			if err := profile.Autoscale.AssignProperties_To_AutoScaleProfile(&profiles[0]); err != nil {
				return err
			}
		}
		if err := bag.Add("Autoscale", profiles); err != nil {
			return err
		}
	} else {
		bag.Remove("Autoscale")
	}

	if len(bag) == 0 {
		bag = nil
	}
	dst.PropertyBag = bag
	return nil
}

func (profile *ScaleProfile_STATUS) AssignPropertiesFrom(src *v20250801s.ScaleProfile_STATUS) error {
	bag := genruntime.NewPropertyBag(profile.PropertyBag)
	profiles, err := conversion.PullAutoscaleProfiles[compat.AutoScaleProfile_STATUS](genruntime.NewPropertyBag(bag))
	if err != nil {
		return err
	}

	profile.Autoscale = nil
	if len(profiles) > 0 {
		profile.Autoscale = &AutoScaleProfile_STATUS{}
		if err := profile.Autoscale.AssignProperties_From_AutoScaleProfile_STATUS(&profiles[0]); err != nil {
			return err
		}
	}
	if len(profiles) <= 1 {
		bag.Remove("Autoscale")
	}

	if len(bag) == 0 {
		bag = nil
	}
	profile.PropertyBag = bag
	return nil
}

func (profile *ScaleProfile_STATUS) AssignPropertiesTo(dst *v20250801s.ScaleProfile_STATUS) error {
	bag := genruntime.NewPropertyBag(dst.PropertyBag)
	profiles, err := conversion.PullAutoscaleProfiles[compat.AutoScaleProfile_STATUS](bag)
	if err != nil {
		return err
	}

	if profile.Autoscale != nil {
		var original *AutoScaleProfile_STATUS
		if len(profiles) > 0 {
			original = &AutoScaleProfile_STATUS{}
			if err := original.AssignProperties_From_AutoScaleProfile_STATUS(&profiles[0]); err != nil {
				return err
			}
		}

		if !reflect.DeepEqual(profile.Autoscale, original) {
			profiles = make([]compat.AutoScaleProfile_STATUS, 1)
			if err := profile.Autoscale.AssignProperties_To_AutoScaleProfile_STATUS(&profiles[0]); err != nil {
				return err
			}
		}
		if err := bag.Add("Autoscale", profiles); err != nil {
			return err
		}
	} else {
		bag.Remove("Autoscale")
	}

	if len(bag) == 0 {
		bag = nil
	}
	dst.PropertyBag = bag
	return nil
}
