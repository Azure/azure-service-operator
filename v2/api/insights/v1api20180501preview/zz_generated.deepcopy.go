//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20180501preview

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
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
func (in *WebTestProperties_Configuration) DeepCopyInto(out *WebTestProperties_Configuration) {
	*out = *in
	if in.WebTest != nil {
		in, out := &in.WebTest, &out.WebTest
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Configuration.
func (in *WebTestProperties_Configuration) DeepCopy() *WebTestProperties_Configuration {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Configuration_STATUS) DeepCopyInto(out *WebTestProperties_Configuration_STATUS) {
	*out = *in
	if in.WebTest != nil {
		in, out := &in.WebTest, &out.WebTest
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Configuration_STATUS.
func (in *WebTestProperties_Configuration_STATUS) DeepCopy() *WebTestProperties_Configuration_STATUS {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Configuration_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Request) DeepCopyInto(out *WebTestProperties_Request) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Request.
func (in *WebTestProperties_Request) DeepCopy() *WebTestProperties_Request {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Request)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Request_STATUS) DeepCopyInto(out *WebTestProperties_Request_STATUS) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Request_STATUS.
func (in *WebTestProperties_Request_STATUS) DeepCopy() *WebTestProperties_Request_STATUS {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Request_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_ValidationRules) DeepCopyInto(out *WebTestProperties_ValidationRules) {
	*out = *in
	if in.ContentValidation != nil {
		in, out := &in.ContentValidation, &out.ContentValidation
		*out = new(WebTestProperties_ValidationRules_ContentValidation)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_ValidationRules.
func (in *WebTestProperties_ValidationRules) DeepCopy() *WebTestProperties_ValidationRules {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_ValidationRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_ValidationRules_ContentValidation) DeepCopyInto(out *WebTestProperties_ValidationRules_ContentValidation) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_ValidationRules_ContentValidation.
func (in *WebTestProperties_ValidationRules_ContentValidation) DeepCopy() *WebTestProperties_ValidationRules_ContentValidation {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_ValidationRules_ContentValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_ValidationRules_ContentValidation_STATUS) DeepCopyInto(out *WebTestProperties_ValidationRules_ContentValidation_STATUS) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_ValidationRules_ContentValidation_STATUS.
func (in *WebTestProperties_ValidationRules_ContentValidation_STATUS) DeepCopy() *WebTestProperties_ValidationRules_ContentValidation_STATUS {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_ValidationRules_ContentValidation_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_ValidationRules_STATUS) DeepCopyInto(out *WebTestProperties_ValidationRules_STATUS) {
	*out = *in
	if in.ContentValidation != nil {
		in, out := &in.ContentValidation, &out.ContentValidation
		*out = new(WebTestProperties_ValidationRules_ContentValidation_STATUS)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_ValidationRules_STATUS.
func (in *WebTestProperties_ValidationRules_STATUS) DeepCopy() *WebTestProperties_ValidationRules_STATUS {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_ValidationRules_STATUS)
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
func (in *Webtest_STATUS) DeepCopyInto(out *Webtest_STATUS) {
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
		*out = new(WebTestProperties_Configuration_STATUS)
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
		*out = new(WebTestProperties_Kind_STATUS)
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
		*out = new(WebTestProperties_Request_STATUS)
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
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.ValidationRules != nil {
		in, out := &in.ValidationRules, &out.ValidationRules
		*out = new(WebTestProperties_ValidationRules_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webtest_STATUS.
func (in *Webtest_STATUS) DeepCopy() *Webtest_STATUS {
	if in == nil {
		return nil
	}
	out := new(Webtest_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webtest_Spec) DeepCopyInto(out *Webtest_Spec) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(WebTestProperties_Configuration)
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
		*out = new(WebTestProperties_Request)
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
		*out = new(WebTestProperties_ValidationRules)
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
