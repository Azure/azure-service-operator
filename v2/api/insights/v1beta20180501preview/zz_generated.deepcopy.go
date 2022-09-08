//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta20180501preview

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderField) DeepCopyInto(out *HeaderField) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderField.
func (in *HeaderField) DeepCopy() *HeaderField {
	if in == nil {
		return nil
	}
	out := new(HeaderField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderFieldARM) DeepCopyInto(out *HeaderFieldARM) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderFieldARM.
func (in *HeaderFieldARM) DeepCopy() *HeaderFieldARM {
	if in == nil {
		return nil
	}
	out := new(HeaderFieldARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderField_STATUS) DeepCopyInto(out *HeaderField_STATUS) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderField_STATUS.
func (in *HeaderField_STATUS) DeepCopy() *HeaderField_STATUS {
	if in == nil {
		return nil
	}
	out := new(HeaderField_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderField_STATUSARM) DeepCopyInto(out *HeaderField_STATUSARM) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderField_STATUSARM.
func (in *HeaderField_STATUSARM) DeepCopy() *HeaderField_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(HeaderField_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestGeolocation) DeepCopyInto(out *WebTestGeolocation) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestGeolocation.
func (in *WebTestGeolocation) DeepCopy() *WebTestGeolocation {
	if in == nil {
		return nil
	}
	out := new(WebTestGeolocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestGeolocationARM) DeepCopyInto(out *WebTestGeolocationARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestGeolocationARM.
func (in *WebTestGeolocationARM) DeepCopy() *WebTestGeolocationARM {
	if in == nil {
		return nil
	}
	out := new(WebTestGeolocationARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestGeolocation_STATUS) DeepCopyInto(out *WebTestGeolocation_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestGeolocation_STATUS.
func (in *WebTestGeolocation_STATUS) DeepCopy() *WebTestGeolocation_STATUS {
	if in == nil {
		return nil
	}
	out := new(WebTestGeolocation_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestGeolocation_STATUSARM) DeepCopyInto(out *WebTestGeolocation_STATUSARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestGeolocation_STATUSARM.
func (in *WebTestGeolocation_STATUSARM) DeepCopy() *WebTestGeolocation_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(WebTestGeolocation_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestPropertiesARM) DeepCopyInto(out *WebTestPropertiesARM) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(WebTestPropertiesConfigurationARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(int)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(WebTestProperties_Kind)
		**out = **in
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]WebTestGeolocationARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(WebTestPropertiesRequestARM)
		(*in).DeepCopyInto(*out)
	}
	if in.RetryEnabled != nil {
		in, out := &in.RetryEnabled, &out.RetryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SyntheticMonitorId != nil {
		in, out := &in.SyntheticMonitorId, &out.SyntheticMonitorId
		*out = new(string)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
	if in.ValidationRules != nil {
		in, out := &in.ValidationRules, &out.ValidationRules
		*out = new(WebTestPropertiesValidationRulesARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestPropertiesARM.
func (in *WebTestPropertiesARM) DeepCopy() *WebTestPropertiesARM {
	if in == nil {
		return nil
	}
	out := new(WebTestPropertiesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestPropertiesConfiguration) DeepCopyInto(out *WebTestPropertiesConfiguration) {
	*out = *in
	if in.WebTest != nil {
		in, out := &in.WebTest, &out.WebTest
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestPropertiesConfiguration.
func (in *WebTestPropertiesConfiguration) DeepCopy() *WebTestPropertiesConfiguration {
	if in == nil {
		return nil
	}
	out := new(WebTestPropertiesConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestPropertiesConfigurationARM) DeepCopyInto(out *WebTestPropertiesConfigurationARM) {
	*out = *in
	if in.WebTest != nil {
		in, out := &in.WebTest, &out.WebTest
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestPropertiesConfigurationARM.
func (in *WebTestPropertiesConfigurationARM) DeepCopy() *WebTestPropertiesConfigurationARM {
	if in == nil {
		return nil
	}
	out := new(WebTestPropertiesConfigurationARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestPropertiesRequest) DeepCopyInto(out *WebTestPropertiesRequest) {
	*out = *in
	if in.FollowRedirects != nil {
		in, out := &in.FollowRedirects, &out.FollowRedirects
		*out = new(bool)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HeaderField, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HttpVerb != nil {
		in, out := &in.HttpVerb, &out.HttpVerb
		*out = new(string)
		**out = **in
	}
	if in.ParseDependentRequests != nil {
		in, out := &in.ParseDependentRequests, &out.ParseDependentRequests
		*out = new(bool)
		**out = **in
	}
	if in.RequestBody != nil {
		in, out := &in.RequestBody, &out.RequestBody
		*out = new(string)
		**out = **in
	}
	if in.RequestUrl != nil {
		in, out := &in.RequestUrl, &out.RequestUrl
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestPropertiesRequest.
func (in *WebTestPropertiesRequest) DeepCopy() *WebTestPropertiesRequest {
	if in == nil {
		return nil
	}
	out := new(WebTestPropertiesRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestPropertiesRequestARM) DeepCopyInto(out *WebTestPropertiesRequestARM) {
	*out = *in
	if in.FollowRedirects != nil {
		in, out := &in.FollowRedirects, &out.FollowRedirects
		*out = new(bool)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HeaderFieldARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HttpVerb != nil {
		in, out := &in.HttpVerb, &out.HttpVerb
		*out = new(string)
		**out = **in
	}
	if in.ParseDependentRequests != nil {
		in, out := &in.ParseDependentRequests, &out.ParseDependentRequests
		*out = new(bool)
		**out = **in
	}
	if in.RequestBody != nil {
		in, out := &in.RequestBody, &out.RequestBody
		*out = new(string)
		**out = **in
	}
	if in.RequestUrl != nil {
		in, out := &in.RequestUrl, &out.RequestUrl
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestPropertiesRequestARM.
func (in *WebTestPropertiesRequestARM) DeepCopy() *WebTestPropertiesRequestARM {
	if in == nil {
		return nil
	}
	out := new(WebTestPropertiesRequestARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestPropertiesValidationRules) DeepCopyInto(out *WebTestPropertiesValidationRules) {
	*out = *in
	if in.ContentValidation != nil {
		in, out := &in.ContentValidation, &out.ContentValidation
		*out = new(WebTestPropertiesValidationRulesContentValidation)
		(*in).DeepCopyInto(*out)
	}
	if in.ExpectedHttpStatusCode != nil {
		in, out := &in.ExpectedHttpStatusCode, &out.ExpectedHttpStatusCode
		*out = new(int)
		**out = **in
	}
	if in.IgnoreHttpsStatusCode != nil {
		in, out := &in.IgnoreHttpsStatusCode, &out.IgnoreHttpsStatusCode
		*out = new(bool)
		**out = **in
	}
	if in.SSLCertRemainingLifetimeCheck != nil {
		in, out := &in.SSLCertRemainingLifetimeCheck, &out.SSLCertRemainingLifetimeCheck
		*out = new(int)
		**out = **in
	}
	if in.SSLCheck != nil {
		in, out := &in.SSLCheck, &out.SSLCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestPropertiesValidationRules.
func (in *WebTestPropertiesValidationRules) DeepCopy() *WebTestPropertiesValidationRules {
	if in == nil {
		return nil
	}
	out := new(WebTestPropertiesValidationRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestPropertiesValidationRulesARM) DeepCopyInto(out *WebTestPropertiesValidationRulesARM) {
	*out = *in
	if in.ContentValidation != nil {
		in, out := &in.ContentValidation, &out.ContentValidation
		*out = new(WebTestPropertiesValidationRulesContentValidationARM)
		(*in).DeepCopyInto(*out)
	}
	if in.ExpectedHttpStatusCode != nil {
		in, out := &in.ExpectedHttpStatusCode, &out.ExpectedHttpStatusCode
		*out = new(int)
		**out = **in
	}
	if in.IgnoreHttpsStatusCode != nil {
		in, out := &in.IgnoreHttpsStatusCode, &out.IgnoreHttpsStatusCode
		*out = new(bool)
		**out = **in
	}
	if in.SSLCertRemainingLifetimeCheck != nil {
		in, out := &in.SSLCertRemainingLifetimeCheck, &out.SSLCertRemainingLifetimeCheck
		*out = new(int)
		**out = **in
	}
	if in.SSLCheck != nil {
		in, out := &in.SSLCheck, &out.SSLCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestPropertiesValidationRulesARM.
func (in *WebTestPropertiesValidationRulesARM) DeepCopy() *WebTestPropertiesValidationRulesARM {
	if in == nil {
		return nil
	}
	out := new(WebTestPropertiesValidationRulesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestPropertiesValidationRulesContentValidation) DeepCopyInto(out *WebTestPropertiesValidationRulesContentValidation) {
	*out = *in
	if in.ContentMatch != nil {
		in, out := &in.ContentMatch, &out.ContentMatch
		*out = new(string)
		**out = **in
	}
	if in.IgnoreCase != nil {
		in, out := &in.IgnoreCase, &out.IgnoreCase
		*out = new(bool)
		**out = **in
	}
	if in.PassIfTextFound != nil {
		in, out := &in.PassIfTextFound, &out.PassIfTextFound
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestPropertiesValidationRulesContentValidation.
func (in *WebTestPropertiesValidationRulesContentValidation) DeepCopy() *WebTestPropertiesValidationRulesContentValidation {
	if in == nil {
		return nil
	}
	out := new(WebTestPropertiesValidationRulesContentValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestPropertiesValidationRulesContentValidationARM) DeepCopyInto(out *WebTestPropertiesValidationRulesContentValidationARM) {
	*out = *in
	if in.ContentMatch != nil {
		in, out := &in.ContentMatch, &out.ContentMatch
		*out = new(string)
		**out = **in
	}
	if in.IgnoreCase != nil {
		in, out := &in.IgnoreCase, &out.IgnoreCase
		*out = new(bool)
		**out = **in
	}
	if in.PassIfTextFound != nil {
		in, out := &in.PassIfTextFound, &out.PassIfTextFound
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestPropertiesValidationRulesContentValidationARM.
func (in *WebTestPropertiesValidationRulesContentValidationARM) DeepCopy() *WebTestPropertiesValidationRulesContentValidationARM {
	if in == nil {
		return nil
	}
	out := new(WebTestPropertiesValidationRulesContentValidationARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_STATUSARM) DeepCopyInto(out *WebTestProperties_STATUSARM) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(WebTestProperties_STATUS_ConfigurationARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(int)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(WebTestProperties_STATUS_Kind)
		**out = **in
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]WebTestGeolocation_STATUSARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(WebTestProperties_STATUS_RequestARM)
		(*in).DeepCopyInto(*out)
	}
	if in.RetryEnabled != nil {
		in, out := &in.RetryEnabled, &out.RetryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SyntheticMonitorId != nil {
		in, out := &in.SyntheticMonitorId, &out.SyntheticMonitorId
		*out = new(string)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
	if in.ValidationRules != nil {
		in, out := &in.ValidationRules, &out.ValidationRules
		*out = new(WebTestProperties_STATUS_ValidationRulesARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_STATUSARM.
func (in *WebTestProperties_STATUSARM) DeepCopy() *WebTestProperties_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_STATUS_Configuration) DeepCopyInto(out *WebTestProperties_STATUS_Configuration) {
	*out = *in
	if in.WebTest != nil {
		in, out := &in.WebTest, &out.WebTest
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_STATUS_Configuration.
func (in *WebTestProperties_STATUS_Configuration) DeepCopy() *WebTestProperties_STATUS_Configuration {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_STATUS_Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_STATUS_ConfigurationARM) DeepCopyInto(out *WebTestProperties_STATUS_ConfigurationARM) {
	*out = *in
	if in.WebTest != nil {
		in, out := &in.WebTest, &out.WebTest
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_STATUS_ConfigurationARM.
func (in *WebTestProperties_STATUS_ConfigurationARM) DeepCopy() *WebTestProperties_STATUS_ConfigurationARM {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_STATUS_ConfigurationARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_STATUS_Request) DeepCopyInto(out *WebTestProperties_STATUS_Request) {
	*out = *in
	if in.FollowRedirects != nil {
		in, out := &in.FollowRedirects, &out.FollowRedirects
		*out = new(bool)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HeaderField_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HttpVerb != nil {
		in, out := &in.HttpVerb, &out.HttpVerb
		*out = new(string)
		**out = **in
	}
	if in.ParseDependentRequests != nil {
		in, out := &in.ParseDependentRequests, &out.ParseDependentRequests
		*out = new(bool)
		**out = **in
	}
	if in.RequestBody != nil {
		in, out := &in.RequestBody, &out.RequestBody
		*out = new(string)
		**out = **in
	}
	if in.RequestUrl != nil {
		in, out := &in.RequestUrl, &out.RequestUrl
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_STATUS_Request.
func (in *WebTestProperties_STATUS_Request) DeepCopy() *WebTestProperties_STATUS_Request {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_STATUS_Request)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_STATUS_RequestARM) DeepCopyInto(out *WebTestProperties_STATUS_RequestARM) {
	*out = *in
	if in.FollowRedirects != nil {
		in, out := &in.FollowRedirects, &out.FollowRedirects
		*out = new(bool)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HeaderField_STATUSARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HttpVerb != nil {
		in, out := &in.HttpVerb, &out.HttpVerb
		*out = new(string)
		**out = **in
	}
	if in.ParseDependentRequests != nil {
		in, out := &in.ParseDependentRequests, &out.ParseDependentRequests
		*out = new(bool)
		**out = **in
	}
	if in.RequestBody != nil {
		in, out := &in.RequestBody, &out.RequestBody
		*out = new(string)
		**out = **in
	}
	if in.RequestUrl != nil {
		in, out := &in.RequestUrl, &out.RequestUrl
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_STATUS_RequestARM.
func (in *WebTestProperties_STATUS_RequestARM) DeepCopy() *WebTestProperties_STATUS_RequestARM {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_STATUS_RequestARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_STATUS_ValidationRules) DeepCopyInto(out *WebTestProperties_STATUS_ValidationRules) {
	*out = *in
	if in.ContentValidation != nil {
		in, out := &in.ContentValidation, &out.ContentValidation
		*out = new(WebTestProperties_STATUS_ValidationRules_ContentValidation)
		(*in).DeepCopyInto(*out)
	}
	if in.ExpectedHttpStatusCode != nil {
		in, out := &in.ExpectedHttpStatusCode, &out.ExpectedHttpStatusCode
		*out = new(int)
		**out = **in
	}
	if in.IgnoreHttpsStatusCode != nil {
		in, out := &in.IgnoreHttpsStatusCode, &out.IgnoreHttpsStatusCode
		*out = new(bool)
		**out = **in
	}
	if in.SSLCertRemainingLifetimeCheck != nil {
		in, out := &in.SSLCertRemainingLifetimeCheck, &out.SSLCertRemainingLifetimeCheck
		*out = new(int)
		**out = **in
	}
	if in.SSLCheck != nil {
		in, out := &in.SSLCheck, &out.SSLCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_STATUS_ValidationRules.
func (in *WebTestProperties_STATUS_ValidationRules) DeepCopy() *WebTestProperties_STATUS_ValidationRules {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_STATUS_ValidationRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_STATUS_ValidationRulesARM) DeepCopyInto(out *WebTestProperties_STATUS_ValidationRulesARM) {
	*out = *in
	if in.ContentValidation != nil {
		in, out := &in.ContentValidation, &out.ContentValidation
		*out = new(WebTestProperties_STATUS_ValidationRules_ContentValidationARM)
		(*in).DeepCopyInto(*out)
	}
	if in.ExpectedHttpStatusCode != nil {
		in, out := &in.ExpectedHttpStatusCode, &out.ExpectedHttpStatusCode
		*out = new(int)
		**out = **in
	}
	if in.IgnoreHttpsStatusCode != nil {
		in, out := &in.IgnoreHttpsStatusCode, &out.IgnoreHttpsStatusCode
		*out = new(bool)
		**out = **in
	}
	if in.SSLCertRemainingLifetimeCheck != nil {
		in, out := &in.SSLCertRemainingLifetimeCheck, &out.SSLCertRemainingLifetimeCheck
		*out = new(int)
		**out = **in
	}
	if in.SSLCheck != nil {
		in, out := &in.SSLCheck, &out.SSLCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_STATUS_ValidationRulesARM.
func (in *WebTestProperties_STATUS_ValidationRulesARM) DeepCopy() *WebTestProperties_STATUS_ValidationRulesARM {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_STATUS_ValidationRulesARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_STATUS_ValidationRules_ContentValidation) DeepCopyInto(out *WebTestProperties_STATUS_ValidationRules_ContentValidation) {
	*out = *in
	if in.ContentMatch != nil {
		in, out := &in.ContentMatch, &out.ContentMatch
		*out = new(string)
		**out = **in
	}
	if in.IgnoreCase != nil {
		in, out := &in.IgnoreCase, &out.IgnoreCase
		*out = new(bool)
		**out = **in
	}
	if in.PassIfTextFound != nil {
		in, out := &in.PassIfTextFound, &out.PassIfTextFound
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_STATUS_ValidationRules_ContentValidation.
func (in *WebTestProperties_STATUS_ValidationRules_ContentValidation) DeepCopy() *WebTestProperties_STATUS_ValidationRules_ContentValidation {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_STATUS_ValidationRules_ContentValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_STATUS_ValidationRules_ContentValidationARM) DeepCopyInto(out *WebTestProperties_STATUS_ValidationRules_ContentValidationARM) {
	*out = *in
	if in.ContentMatch != nil {
		in, out := &in.ContentMatch, &out.ContentMatch
		*out = new(string)
		**out = **in
	}
	if in.IgnoreCase != nil {
		in, out := &in.IgnoreCase, &out.IgnoreCase
		*out = new(bool)
		**out = **in
	}
	if in.PassIfTextFound != nil {
		in, out := &in.PassIfTextFound, &out.PassIfTextFound
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_STATUS_ValidationRules_ContentValidationARM.
func (in *WebTestProperties_STATUS_ValidationRules_ContentValidationARM) DeepCopy() *WebTestProperties_STATUS_ValidationRules_ContentValidationARM {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_STATUS_ValidationRules_ContentValidationARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTest_STATUS) DeepCopyInto(out *WebTest_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(WebTestProperties_STATUS_Configuration)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(int)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(WebTestProperties_STATUS_Kind)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]WebTestGeolocation_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PropertiesName != nil {
		in, out := &in.PropertiesName, &out.PropertiesName
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(WebTestProperties_STATUS_Request)
		(*in).DeepCopyInto(*out)
	}
	if in.RetryEnabled != nil {
		in, out := &in.RetryEnabled, &out.RetryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SyntheticMonitorId != nil {
		in, out := &in.SyntheticMonitorId, &out.SyntheticMonitorId
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.ValidationRules != nil {
		in, out := &in.ValidationRules, &out.ValidationRules
		*out = new(WebTestProperties_STATUS_ValidationRules)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTest_STATUS.
func (in *WebTest_STATUS) DeepCopy() *WebTest_STATUS {
	if in == nil {
		return nil
	}
	out := new(WebTest_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTest_STATUSARM) DeepCopyInto(out *WebTest_STATUSARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(WebTestProperties_STATUSARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTest_STATUSARM.
func (in *WebTest_STATUSARM) DeepCopy() *WebTest_STATUSARM {
	if in == nil {
		return nil
	}
	out := new(WebTest_STATUSARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webtest) DeepCopyInto(out *Webtest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webtest.
func (in *Webtest) DeepCopy() *Webtest {
	if in == nil {
		return nil
	}
	out := new(Webtest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Webtest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebtestList) DeepCopyInto(out *WebtestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Webtest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebtestList.
func (in *WebtestList) DeepCopy() *WebtestList {
	if in == nil {
		return nil
	}
	out := new(WebtestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebtestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webtest_Spec) DeepCopyInto(out *Webtest_Spec) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(WebTestPropertiesConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(int)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(WebTestProperties_Kind)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]WebTestGeolocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(WebTestPropertiesRequest)
		(*in).DeepCopyInto(*out)
	}
	if in.RetryEnabled != nil {
		in, out := &in.RetryEnabled, &out.RetryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SyntheticMonitorId != nil {
		in, out := &in.SyntheticMonitorId, &out.SyntheticMonitorId
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
	if in.ValidationRules != nil {
		in, out := &in.ValidationRules, &out.ValidationRules
		*out = new(WebTestPropertiesValidationRules)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webtest_Spec.
func (in *Webtest_Spec) DeepCopy() *Webtest_Spec {
	if in == nil {
		return nil
	}
	out := new(Webtest_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webtest_SpecARM) DeepCopyInto(out *Webtest_SpecARM) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(WebTestPropertiesARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webtest_SpecARM.
func (in *Webtest_SpecARM) DeepCopy() *Webtest_SpecARM {
	if in == nil {
		return nil
	}
	out := new(Webtest_SpecARM)
	in.DeepCopyInto(out)
	return out
}
