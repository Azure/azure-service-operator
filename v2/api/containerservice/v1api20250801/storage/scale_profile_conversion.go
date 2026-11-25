// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package storage

import (
	"github.com/Azure/azure-service-operator/v2/api/containerservice/internal/conversion"
	"github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20250801/storage/compat"
	hub "github.com/Azure/azure-service-operator/v2/api/containerservice/v20260501/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

var (
	_ augmentConversionForScaleProfile        = &ScaleProfile{}
	_ augmentConversionForScaleProfile_STATUS = &ScaleProfile_STATUS{}
)

func (profile *ScaleProfile) AssignPropertiesFrom(src *hub.ScaleProfile) error {
	bag := genruntime.NewPropertyBag(profile.PropertyBag)
	if len(src.Autoscale) > 0 {
		profiles := make([]compat.AutoScaleProfile, len(src.Autoscale))
		for i := range src.Autoscale {
			if err := profiles[i].AssignProperties_From_AutoScaleProfile(&src.Autoscale[i]); err != nil {
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
	profile.PropertyBag = bag
	return nil
}

func (profile *ScaleProfile) AssignPropertiesTo(dst *hub.ScaleProfile) error {
	bag := genruntime.NewPropertyBag(dst.PropertyBag)
	profiles, err := conversion.PullAutoscaleProfiles[compat.AutoScaleProfile](bag)
	if err != nil {
		return err
	}

	dst.Autoscale = nil
	if profiles != nil {
		dst.Autoscale = make([]hub.AutoScaleProfile, len(profiles))
		for i := range profiles {
			if err := profiles[i].AssignProperties_To_AutoScaleProfile(&dst.Autoscale[i]); err != nil {
				return err
			}
		}
	}

	if len(bag) == 0 {
		bag = nil
	}
	dst.PropertyBag = bag
	return nil
}

func (profile *ScaleProfile_STATUS) AssignPropertiesFrom(src *hub.ScaleProfile_STATUS) error {
	bag := genruntime.NewPropertyBag(profile.PropertyBag)
	if len(src.Autoscale) > 0 {
		profiles := make([]compat.AutoScaleProfile_STATUS, len(src.Autoscale))
		for i := range src.Autoscale {
			if err := profiles[i].AssignProperties_From_AutoScaleProfile_STATUS(&src.Autoscale[i]); err != nil {
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
	profile.PropertyBag = bag
	return nil
}

func (profile *ScaleProfile_STATUS) AssignPropertiesTo(dst *hub.ScaleProfile_STATUS) error {
	bag := genruntime.NewPropertyBag(dst.PropertyBag)
	profiles, err := conversion.PullAutoscaleProfiles[compat.AutoScaleProfile_STATUS](bag)
	if err != nil {
		return err
	}

	dst.Autoscale = nil
	if profiles != nil {
		dst.Autoscale = make([]hub.AutoScaleProfile_STATUS, len(profiles))
		for i := range profiles {
			if err := profiles[i].AssignProperties_To_AutoScaleProfile_STATUS(&dst.Autoscale[i]); err != nil {
				return err
			}
		}
	}

	if len(bag) == 0 {
		bag = nil
	}
	dst.PropertyBag = bag
	return nil
}
