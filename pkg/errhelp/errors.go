// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package errhelp

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"k8s.io/apimachinery/pkg/api/errors"
)

const (
	AccountNameInvalid                             = "AccountNameInvalid"
	AlreadyExists                                  = "AlreadyExists"
	AsyncOpIncompleteError                         = "AsyncOpIncomplete"
	BadRequest                                     = "BadRequest"
	CannotParseError                               = "CannotParseError"
	ConflictingServerOperation                     = "ConflictingServerOperation"
	ContainerOperationFailure                      = "ContainerOperationFailure"
	ConsumerGroupNotFound                          = "ConsumerGroupNotFound"
	CreationPending                                = "CreationPending"
	FailoverGroupBusy                              = "FailoverGroupBusy"
	Forbidden                                      = "Forbidden"
	InvalidAccessPolicy                            = "InvalidAccessPolicy"
	InvalidCIDRNotation                            = "InvalidCIDRNotation"
	InvalidFailoverGroupRegion                     = "InvalidFailoverGroupRegion"
	InvalidParameters                              = "InvalidParameters"
	InvalidRequestFormat                           = "InvalidRequestFormat"
	InvalidResourceLocation                        = "InvalidResourceLocation"
	InvalidServerName                              = "InvalidServerName"
	InvalidResourceReference                       = "InvalidResourceReference"
	KeyNotFound                                    = "KeyNotFound"
	LocationNotAvailableForResourceType            = "LocationNotAvailableForResourceType"
	ProvisioningDisabled                           = "ProvisioningDisabled"
	NetcfgInvalidIPAddressPrefix                   = "NetcfgInvalidIPAddressPrefix"
	NetcfgInvalidSubnet                            = "NetcfgInvalidSubnet"
	NetcfgInvalidVirtualNetworkSite                = "NetcfgInvalidVirtualNetworkSite"
	NotFoundErrorCode                              = "NotFound"
	NoSuchHost                                     = "no such host"
	ParentNotFoundErrorCode                        = "ParentResourceNotFound"
	StorageAccountIsNotProvisioned                 = "StorageAccountIsNotProvisioned"
	PreconditionFailed                             = "PreconditionFailed"
	QuotaExceeded                                  = "QuotaExceeded"
	ResourceGroupNotFoundErrorCode                 = "ResourceGroupNotFound"
	RegionDoesNotAllowProvisioning                 = "RegionDoesNotAllowProvisioning"
	ResourceNotFound                               = "ResourceNotFound"
	RequestConflictError                           = "Conflict"
	ValidationError                                = "ValidationError"
	SubscriptionDoesNotHaveServer                  = "SubscriptionDoesNotHaveServer"
	NotSupported                                   = "NotSupported"
	SecretNotFound                                 = "SecretNotFound"
	RequestDisallowedByPolicy                      = "RequestDisallowedByPolicy"
	ServiceBusy                                    = "ServiceBusy"
	NameNotAvailable                               = "NameNotAvailable"
	PublicIPIdleTimeoutIsOutOfRange                = "PublicIPIdleTimeoutIsOutOfRange"
	InvalidRequestContent                          = "InvalidRequestContent"
	InvalidMaxSizeTierCombination                  = "InvalidMaxSizeTierCombination"
	InternalServerError                            = "InternalServerError"
	NetworkAclsValidationFailure                   = "NetworkAclsValidationFailure"
	SubnetHasServiceEndpointWithInvalidServiceName = "SubnetHasServiceEndpointWithInvalidServiceName"
	InvalidAddressPrefixFormat                     = "InvalidAddressPrefixFormat"
	FeatureNotSupportedForEdition                  = "FeatureNotSupportedForEdition"
	VirtualNetworkRuleBadRequest                   = "VirtualNetworkRuleBadRequest"
	LongTermRetentionPolicyInvalid                 = "LongTermRetentionPolicyInvalid"
	BackupRetentionPolicyInvalid                   = "InvalidBackupRetentionPeriod"
	OperationIdNotFound                            = "OperationIdNotFound"
	ObjectIsBeingDeleted                           = "ObjectIsBeingDeleted"
	ObjectIsDeletedButRecoverable                  = "ObjectIsDeletedButRecoverable"
)

func NewAzureError(err error) *AzureError {
	var kind, reason string
	if err == nil {
		return nil
	}
	ae := AzureError{
		Original: err,
	}

	switch v := err.(type) {
	case autorest.DetailedError:
		ae.Code = v.StatusCode.(int)

		switch e := v.Original.(type) {
		case *azure.RequestError:
			if e.ServiceError != nil {
				kind = e.ServiceError.Code
				reason = e.ServiceError.Message
			} else {
				kind = CannotParseError
				reason = CannotParseError
			}
		case azure.RequestError:
			if e.ServiceError != nil {
				kind = e.ServiceError.Code
				reason = e.ServiceError.Message
			} else {
				kind = CannotParseError
				reason = CannotParseError
			}
		case *azure.ServiceError:
			kind = e.Code
			reason = e.Message
			if e.Code == "Failed" && len(e.AdditionalInfo) == 1 {
				if v, ok := e.AdditionalInfo[0]["code"]; ok {
					kind = v.(string)
				}
			}
		case *errors.StatusError:
			kind = "StatusError"
			reason = "StatusError"
		case *json.UnmarshalTypeError:
			kind = NotFoundErrorCode
			reason = NotFoundErrorCode
		}

	case azure.AsyncOpIncompleteError:
		kind = "AsyncOpIncomplete"
		reason = "AsyncOpIncomplete"
	case validation.Error:
		kind = "ValidationError"
		reason = v.Message
	default:
		switch e := err.Error(); e {
		case InvalidServerName, AlreadyExists, AccountNameInvalid:
			kind = e
			reason = e
		default:
			contains := []string{
				InvalidAccessPolicy,
				LocationNotAvailableForResourceType,
			}
			for _, val := range contains {
				if strings.Contains(e, val) {
					kind = val
					reason = val
					break
				}
			}
		}

	}

	ae.Reason = reason
	ae.Type = kind

	return &ae
}

type AzureError struct {
	Type     string
	Reason   string
	Code     int
	Original error
}

func (e AzureError) Error() string {
	return e.Original.Error()
}

type AdminSecretNotFound struct {
	Name string
}

func (e AdminSecretNotFound) Error() string {
	return fmt.Sprintf("admin secret '%s' not found", e.Name)
}

func NewAdminSecretNotFoundError(name string) *AdminSecretNotFound {
	return &AdminSecretNotFound{name}
}
