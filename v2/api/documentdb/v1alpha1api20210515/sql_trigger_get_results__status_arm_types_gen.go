// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

//Deprecated version of SqlTriggerGetResults_Status. Use v1beta20210515.SqlTriggerGetResults_Status instead
type SqlTriggerGetResults_StatusARM struct {
	Id         *string                            `json:"id,omitempty"`
	Location   *string                            `json:"location,omitempty"`
	Name       *string                            `json:"name,omitempty"`
	Properties *SqlTriggerGetProperties_StatusARM `json:"properties,omitempty"`
	Tags       map[string]string                  `json:"tags,omitempty"`
	Type       *string                            `json:"type,omitempty"`
}

//Deprecated version of SqlTriggerGetProperties_Status. Use v1beta20210515.SqlTriggerGetProperties_Status instead
type SqlTriggerGetProperties_StatusARM struct {
	Resource *SqlTriggerGetProperties_Status_ResourceARM `json:"resource,omitempty"`
}

//Deprecated version of SqlTriggerGetProperties_Status_Resource. Use v1beta20210515.SqlTriggerGetProperties_Status_Resource instead
type SqlTriggerGetProperties_Status_ResourceARM struct {
	Body             *string                                                `json:"body,omitempty"`
	Etag             *string                                                `json:"_etag,omitempty"`
	Id               *string                                                `json:"id,omitempty"`
	Rid              *string                                                `json:"_rid,omitempty"`
	TriggerOperation *SqlTriggerGetPropertiesStatusResourceTriggerOperation `json:"triggerOperation,omitempty"`
	TriggerType      *SqlTriggerGetPropertiesStatusResourceTriggerType      `json:"triggerType,omitempty"`
	Ts               *float64                                               `json:"_ts,omitempty"`
}
