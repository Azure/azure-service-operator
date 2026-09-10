/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"github.com/rotisserie/eris"

	v20250801s "github.com/Azure/azure-service-operator/v2/api/containerservice/v20250801/storage"
	v20260301s "github.com/Azure/azure-service-operator/v2/api/containerservice/v20260301/storage"
)

var (
	_ augmentConversionForManagedCluster_Spec                      = &ManagedCluster_Spec{}
	_ augmentConversionForManagedCluster_STATUS                    = &ManagedCluster_STATUS{}
	_ augmentConversionForManagedClusterIdentity_STATUS            = &ManagedClusterIdentity_STATUS{}
	_ augmentConversionForManagedClusterLoadBalancerProfile        = &ManagedClusterLoadBalancerProfile{}
	_ augmentConversionForManagedClusterLoadBalancerProfile_STATUS = &ManagedClusterLoadBalancerProfile_STATUS{}
	_ augmentConversionForManagedClusterPodIdentity_STATUS         = &ManagedClusterPodIdentity_STATUS{}
)

func (cluster *ManagedCluster_Spec) AssignPropertiesFrom(source *v20250801s.ManagedCluster_Spec) error {
	if source.AutoScalerProfile != nil {
		var intermediate v20260301s.ManagedClusterPropertiesAutoScalerProfile
		if err := source.AutoScalerProfile.AssignProperties_To_ManagedClusterPropertiesAutoScalerProfile(&intermediate); err != nil {
			return eris.Wrap(err, "converting AutoScalerProfile to intermediate version")
		}

		var profile ManagedClusterPropertiesAutoScalerProfile
		if err := profile.AssignProperties_From_ManagedClusterPropertiesAutoScalerProfile(&intermediate); err != nil {
			return eris.Wrap(err, "converting AutoScalerProfile from intermediate version")
		}

		cluster.AutoScalerProfile = &profile
	} else {
		cluster.AutoScalerProfile = nil
	}

	return nil
}

func (cluster *ManagedCluster_Spec) AssignPropertiesTo(destination *v20250801s.ManagedCluster_Spec) error {
	if cluster.AutoScalerProfile != nil {
		var intermediate v20260301s.ManagedClusterPropertiesAutoScalerProfile
		if err := cluster.AutoScalerProfile.AssignProperties_To_ManagedClusterPropertiesAutoScalerProfile(&intermediate); err != nil {
			return eris.Wrap(err, "converting AutoScalerProfile to intermediate version")
		}

		var profile v20250801s.ManagedClusterProperties_AutoScalerProfile
		if err := profile.AssignProperties_From_ManagedClusterPropertiesAutoScalerProfile(&intermediate); err != nil {
			return eris.Wrap(err, "converting AutoScalerProfile from intermediate version")
		}

		destination.AutoScalerProfile = &profile
	} else {
		destination.AutoScalerProfile = nil
	}

	return nil
}

func (cluster *ManagedCluster_STATUS) AssignPropertiesFrom(source *v20250801s.ManagedCluster_STATUS) error {
	if source.AutoScalerProfile != nil {
		var intermediate v20260301s.ManagedClusterPropertiesAutoScalerProfile_STATUS
		if err := source.AutoScalerProfile.AssignProperties_To_ManagedClusterPropertiesAutoScalerProfile_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting AutoScalerProfile to intermediate version")
		}

		var profile ManagedClusterPropertiesAutoScalerProfile_STATUS
		if err := profile.AssignProperties_From_ManagedClusterPropertiesAutoScalerProfile_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting AutoScalerProfile from intermediate version")
		}

		cluster.AutoScalerProfile = &profile
	} else {
		cluster.AutoScalerProfile = nil
	}

	return nil
}

func (cluster *ManagedCluster_STATUS) AssignPropertiesTo(destination *v20250801s.ManagedCluster_STATUS) error {
	if cluster.AutoScalerProfile != nil {
		var intermediate v20260301s.ManagedClusterPropertiesAutoScalerProfile_STATUS
		if err := cluster.AutoScalerProfile.AssignProperties_To_ManagedClusterPropertiesAutoScalerProfile_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting AutoScalerProfile to intermediate version")
		}

		var profile v20250801s.ManagedClusterProperties_AutoScalerProfile_STATUS
		if err := profile.AssignProperties_From_ManagedClusterPropertiesAutoScalerProfile_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting AutoScalerProfile from intermediate version")
		}

		destination.AutoScalerProfile = &profile
	} else {
		destination.AutoScalerProfile = nil
	}

	return nil
}

func (identity *ManagedClusterIdentity_STATUS) AssignPropertiesFrom(source *v20250801s.ManagedClusterIdentity_STATUS) error {
	if source.UserAssignedIdentities != nil {
		identities := make(map[string]ManagedServiceIdentityUserAssignedIdentitiesValue_STATUS, len(source.UserAssignedIdentities))
		for key, value := range source.UserAssignedIdentities {
			var intermediate v20260301s.ManagedServiceIdentityUserAssignedIdentitiesValue_STATUS
			if err := value.AssignProperties_To_ManagedServiceIdentityUserAssignedIdentitiesValue_STATUS(&intermediate); err != nil {
				return eris.Wrapf(err, "converting UserAssignedIdentities[%q] to intermediate version", key)
			}

			var converted ManagedServiceIdentityUserAssignedIdentitiesValue_STATUS
			if err := converted.AssignProperties_From_ManagedServiceIdentityUserAssignedIdentitiesValue_STATUS(&intermediate); err != nil {
				return eris.Wrapf(err, "converting UserAssignedIdentities[%q] from intermediate version", key)
			}

			identities[key] = converted
		}

		identity.UserAssignedIdentities = identities
	} else {
		identity.UserAssignedIdentities = nil
	}

	return nil
}

func (identity *ManagedClusterIdentity_STATUS) AssignPropertiesTo(destination *v20250801s.ManagedClusterIdentity_STATUS) error {
	if identity.UserAssignedIdentities != nil {
		identities := make(map[string]v20250801s.ManagedClusterIdentity_UserAssignedIdentities_STATUS, len(identity.UserAssignedIdentities))
		for key, value := range identity.UserAssignedIdentities {
			var intermediate v20260301s.ManagedServiceIdentityUserAssignedIdentitiesValue_STATUS
			if err := value.AssignProperties_To_ManagedServiceIdentityUserAssignedIdentitiesValue_STATUS(&intermediate); err != nil {
				return eris.Wrapf(err, "converting UserAssignedIdentities[%q] to intermediate version", key)
			}

			var converted v20250801s.ManagedClusterIdentity_UserAssignedIdentities_STATUS
			if err := converted.AssignProperties_From_ManagedServiceIdentityUserAssignedIdentitiesValue_STATUS(&intermediate); err != nil {
				return eris.Wrapf(err, "converting UserAssignedIdentities[%q] from intermediate version", key)
			}

			identities[key] = converted
		}

		destination.UserAssignedIdentities = identities
	} else {
		destination.UserAssignedIdentities = nil
	}

	return nil
}

func (profile *ManagedClusterLoadBalancerProfile) AssignPropertiesFrom(source *v20250801s.ManagedClusterLoadBalancerProfile) error {
	if source.ManagedOutboundIPs != nil {
		var intermediate v20260301s.ManagedClusterLoadBalancerProfileManagedOutboundIPs
		if err := source.ManagedOutboundIPs.AssignProperties_To_ManagedClusterLoadBalancerProfileManagedOutboundIPs(&intermediate); err != nil {
			return eris.Wrap(err, "converting ManagedOutboundIPs to intermediate version")
		}

		var converted ManagedClusterLoadBalancerProfileManagedOutboundIPs
		if err := converted.AssignProperties_From_ManagedClusterLoadBalancerProfileManagedOutboundIPs(&intermediate); err != nil {
			return eris.Wrap(err, "converting ManagedOutboundIPs from intermediate version")
		}

		profile.ManagedOutboundIPs = &converted
	} else {
		profile.ManagedOutboundIPs = nil
	}

	if source.OutboundIPPrefixes != nil {
		var intermediate v20260301s.ManagedClusterLoadBalancerProfileOutboundIPPrefixes
		if err := source.OutboundIPPrefixes.AssignProperties_To_ManagedClusterLoadBalancerProfileOutboundIPPrefixes(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPPrefixes to intermediate version")
		}

		var converted ManagedClusterLoadBalancerProfileOutboundIPPrefixes
		if err := converted.AssignProperties_From_ManagedClusterLoadBalancerProfileOutboundIPPrefixes(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPPrefixes from intermediate version")
		}

		profile.OutboundIPPrefixes = &converted
	} else {
		profile.OutboundIPPrefixes = nil
	}

	if source.OutboundIPs != nil {
		var intermediate v20260301s.ManagedClusterLoadBalancerProfileOutboundIPs
		if err := source.OutboundIPs.AssignProperties_To_ManagedClusterLoadBalancerProfileOutboundIPs(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPs to intermediate version")
		}

		var converted ManagedClusterLoadBalancerProfileOutboundIPs
		if err := converted.AssignProperties_From_ManagedClusterLoadBalancerProfileOutboundIPs(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPs from intermediate version")
		}

		profile.OutboundIPs = &converted
	} else {
		profile.OutboundIPs = nil
	}

	return nil
}

func (profile *ManagedClusterLoadBalancerProfile) AssignPropertiesTo(destination *v20250801s.ManagedClusterLoadBalancerProfile) error {
	if profile.ManagedOutboundIPs != nil {
		var intermediate v20260301s.ManagedClusterLoadBalancerProfileManagedOutboundIPs
		if err := profile.ManagedOutboundIPs.AssignProperties_To_ManagedClusterLoadBalancerProfileManagedOutboundIPs(&intermediate); err != nil {
			return eris.Wrap(err, "converting ManagedOutboundIPs to intermediate version")
		}

		var converted v20250801s.ManagedClusterLoadBalancerProfile_ManagedOutboundIPs
		if err := converted.AssignProperties_From_ManagedClusterLoadBalancerProfileManagedOutboundIPs(&intermediate); err != nil {
			return eris.Wrap(err, "converting ManagedOutboundIPs from intermediate version")
		}

		destination.ManagedOutboundIPs = &converted
	} else {
		destination.ManagedOutboundIPs = nil
	}

	if profile.OutboundIPPrefixes != nil {
		var intermediate v20260301s.ManagedClusterLoadBalancerProfileOutboundIPPrefixes
		if err := profile.OutboundIPPrefixes.AssignProperties_To_ManagedClusterLoadBalancerProfileOutboundIPPrefixes(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPPrefixes to intermediate version")
		}

		var converted v20250801s.ManagedClusterLoadBalancerProfile_OutboundIPPrefixes
		if err := converted.AssignProperties_From_ManagedClusterLoadBalancerProfileOutboundIPPrefixes(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPPrefixes from intermediate version")
		}

		destination.OutboundIPPrefixes = &converted
	} else {
		destination.OutboundIPPrefixes = nil
	}

	if profile.OutboundIPs != nil {
		var intermediate v20260301s.ManagedClusterLoadBalancerProfileOutboundIPs
		if err := profile.OutboundIPs.AssignProperties_To_ManagedClusterLoadBalancerProfileOutboundIPs(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPs to intermediate version")
		}

		var converted v20250801s.ManagedClusterLoadBalancerProfile_OutboundIPs
		if err := converted.AssignProperties_From_ManagedClusterLoadBalancerProfileOutboundIPs(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPs from intermediate version")
		}

		destination.OutboundIPs = &converted
	} else {
		destination.OutboundIPs = nil
	}

	return nil
}

func (profile *ManagedClusterLoadBalancerProfile_STATUS) AssignPropertiesFrom(source *v20250801s.ManagedClusterLoadBalancerProfile_STATUS) error {
	if source.ManagedOutboundIPs != nil {
		var intermediate v20260301s.ManagedClusterLoadBalancerProfileManagedOutboundIPs_STATUS
		if err := source.ManagedOutboundIPs.AssignProperties_To_ManagedClusterLoadBalancerProfileManagedOutboundIPs_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting ManagedOutboundIPs to intermediate version")
		}

		var converted ManagedClusterLoadBalancerProfileManagedOutboundIPs_STATUS
		if err := converted.AssignProperties_From_ManagedClusterLoadBalancerProfileManagedOutboundIPs_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting ManagedOutboundIPs from intermediate version")
		}

		profile.ManagedOutboundIPs = &converted
	} else {
		profile.ManagedOutboundIPs = nil
	}

	if source.OutboundIPPrefixes != nil {
		var intermediate v20260301s.ManagedClusterLoadBalancerProfileOutboundIPPrefixes_STATUS
		if err := source.OutboundIPPrefixes.AssignProperties_To_ManagedClusterLoadBalancerProfileOutboundIPPrefixes_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPPrefixes to intermediate version")
		}

		var converted ManagedClusterLoadBalancerProfileOutboundIPPrefixes_STATUS
		if err := converted.AssignProperties_From_ManagedClusterLoadBalancerProfileOutboundIPPrefixes_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPPrefixes from intermediate version")
		}

		profile.OutboundIPPrefixes = &converted
	} else {
		profile.OutboundIPPrefixes = nil
	}

	if source.OutboundIPs != nil {
		var intermediate v20260301s.ManagedClusterLoadBalancerProfileOutboundIPs_STATUS
		if err := source.OutboundIPs.AssignProperties_To_ManagedClusterLoadBalancerProfileOutboundIPs_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPs to intermediate version")
		}

		var converted ManagedClusterLoadBalancerProfileOutboundIPs_STATUS
		if err := converted.AssignProperties_From_ManagedClusterLoadBalancerProfileOutboundIPs_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPs from intermediate version")
		}

		profile.OutboundIPs = &converted
	} else {
		profile.OutboundIPs = nil
	}

	return nil
}

func (profile *ManagedClusterLoadBalancerProfile_STATUS) AssignPropertiesTo(destination *v20250801s.ManagedClusterLoadBalancerProfile_STATUS) error {
	if profile.ManagedOutboundIPs != nil {
		var intermediate v20260301s.ManagedClusterLoadBalancerProfileManagedOutboundIPs_STATUS
		if err := profile.ManagedOutboundIPs.AssignProperties_To_ManagedClusterLoadBalancerProfileManagedOutboundIPs_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting ManagedOutboundIPs to intermediate version")
		}

		var converted v20250801s.ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS
		if err := converted.AssignProperties_From_ManagedClusterLoadBalancerProfileManagedOutboundIPs_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting ManagedOutboundIPs from intermediate version")
		}

		destination.ManagedOutboundIPs = &converted
	} else {
		destination.ManagedOutboundIPs = nil
	}

	if profile.OutboundIPPrefixes != nil {
		var intermediate v20260301s.ManagedClusterLoadBalancerProfileOutboundIPPrefixes_STATUS
		if err := profile.OutboundIPPrefixes.AssignProperties_To_ManagedClusterLoadBalancerProfileOutboundIPPrefixes_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPPrefixes to intermediate version")
		}

		var converted v20250801s.ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS
		if err := converted.AssignProperties_From_ManagedClusterLoadBalancerProfileOutboundIPPrefixes_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPPrefixes from intermediate version")
		}

		destination.OutboundIPPrefixes = &converted
	} else {
		destination.OutboundIPPrefixes = nil
	}

	if profile.OutboundIPs != nil {
		var intermediate v20260301s.ManagedClusterLoadBalancerProfileOutboundIPs_STATUS
		if err := profile.OutboundIPs.AssignProperties_To_ManagedClusterLoadBalancerProfileOutboundIPs_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPs to intermediate version")
		}

		var converted v20250801s.ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS
		if err := converted.AssignProperties_From_ManagedClusterLoadBalancerProfileOutboundIPs_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting OutboundIPs from intermediate version")
		}

		destination.OutboundIPs = &converted
	} else {
		destination.OutboundIPs = nil
	}

	return nil
}

func (identity *ManagedClusterPodIdentity_STATUS) AssignPropertiesFrom(source *v20250801s.ManagedClusterPodIdentity_STATUS) error {
	if source.ProvisioningInfo != nil {
		var intermediate v20260301s.ManagedClusterPodIdentityProvisioningInfo_STATUS
		if err := source.ProvisioningInfo.AssignProperties_To_ManagedClusterPodIdentityProvisioningInfo_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting ProvisioningInfo to intermediate version")
		}

		var info ManagedClusterPodIdentityProvisioningInfo_STATUS
		if err := info.AssignProperties_From_ManagedClusterPodIdentityProvisioningInfo_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting ProvisioningInfo from intermediate version")
		}

		identity.ProvisioningInfo = &info
	} else {
		identity.ProvisioningInfo = nil
	}

	return nil
}

func (identity *ManagedClusterPodIdentity_STATUS) AssignPropertiesTo(destination *v20250801s.ManagedClusterPodIdentity_STATUS) error {
	if identity.ProvisioningInfo != nil {
		var intermediate v20260301s.ManagedClusterPodIdentityProvisioningInfo_STATUS
		if err := identity.ProvisioningInfo.AssignProperties_To_ManagedClusterPodIdentityProvisioningInfo_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting ProvisioningInfo to intermediate version")
		}

		var info v20250801s.ManagedClusterPodIdentity_ProvisioningInfo_STATUS
		if err := info.AssignProperties_From_ManagedClusterPodIdentityProvisioningInfo_STATUS(&intermediate); err != nil {
			return eris.Wrap(err, "converting ProvisioningInfo from intermediate version")
		}

		destination.ProvisioningInfo = &info
	} else {
		destination.ProvisioningInfo = nil
	}

	return nil
}
