// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

// Container App Job
type Job_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Identity: Managed identities needed by a container app job to interact with other Azure services to not maintain any
	// secrets or credentials in code.
	Identity *ManagedServiceIdentity_STATUS `json:"identity,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Container Apps Job resource specific properties.
	Properties *Job_Properties_STATUS `json:"properties,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type Job_Properties_STATUS struct {
	// Configuration: Container Apps Job configuration properties.
	Configuration *JobConfiguration_STATUS `json:"configuration,omitempty"`

	// EnvironmentId: Resource ID of environment.
	EnvironmentId *string `json:"environmentId,omitempty"`

	// EventStreamEndpoint: The endpoint of the eventstream of the container apps job.
	EventStreamEndpoint *string `json:"eventStreamEndpoint,omitempty"`

	// OutboundIpAddresses: Outbound IP Addresses of a container apps job.
	OutboundIpAddresses []string `json:"outboundIpAddresses,omitempty"`

	// ProvisioningState: Provisioning state of the Container Apps Job.
	ProvisioningState *Job_Properties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// Template: Container Apps job definition.
	Template *JobTemplate_STATUS `json:"template,omitempty"`

	// WorkloadProfileName: Workload profile name to pin for container apps job execution.
	WorkloadProfileName *string `json:"workloadProfileName,omitempty"`
}

type Job_Properties_ProvisioningState_STATUS string

const (
	Job_Properties_ProvisioningState_STATUS_Canceled   = Job_Properties_ProvisioningState_STATUS("Canceled")
	Job_Properties_ProvisioningState_STATUS_Deleting   = Job_Properties_ProvisioningState_STATUS("Deleting")
	Job_Properties_ProvisioningState_STATUS_Failed     = Job_Properties_ProvisioningState_STATUS("Failed")
	Job_Properties_ProvisioningState_STATUS_InProgress = Job_Properties_ProvisioningState_STATUS("InProgress")
	Job_Properties_ProvisioningState_STATUS_Succeeded  = Job_Properties_ProvisioningState_STATUS("Succeeded")
)

// Mapping from string to Job_Properties_ProvisioningState_STATUS
var job_Properties_ProvisioningState_STATUS_Values = map[string]Job_Properties_ProvisioningState_STATUS{
	"canceled":   Job_Properties_ProvisioningState_STATUS_Canceled,
	"deleting":   Job_Properties_ProvisioningState_STATUS_Deleting,
	"failed":     Job_Properties_ProvisioningState_STATUS_Failed,
	"inprogress": Job_Properties_ProvisioningState_STATUS_InProgress,
	"succeeded":  Job_Properties_ProvisioningState_STATUS_Succeeded,
}

// Non versioned Container Apps Job configuration properties
type JobConfiguration_STATUS struct {
	// EventTriggerConfig: Trigger configuration of an event driven job.
	EventTriggerConfig *JobConfiguration_EventTriggerConfig_STATUS `json:"eventTriggerConfig,omitempty"`

	// ManualTriggerConfig: Manual trigger configuration for a single execution job. Properties replicaCompletionCount and
	// parallelism would be set to 1 by default
	ManualTriggerConfig *JobConfiguration_ManualTriggerConfig_STATUS `json:"manualTriggerConfig,omitempty"`

	// Registries: Collection of private container registry credentials used by a Container apps job
	Registries []RegistryCredentials_STATUS `json:"registries,omitempty"`

	// ReplicaRetryLimit: Maximum number of retries before failing the job.
	ReplicaRetryLimit *int `json:"replicaRetryLimit,omitempty"`

	// ReplicaTimeout: Maximum number of seconds a replica is allowed to run.
	ReplicaTimeout *int `json:"replicaTimeout,omitempty"`

	// ScheduleTriggerConfig: Cron formatted repeating trigger schedule ("* * * * *") for cronjobs. Properties completions and
	// parallelism would be set to 1 by default
	ScheduleTriggerConfig *JobConfiguration_ScheduleTriggerConfig_STATUS `json:"scheduleTriggerConfig,omitempty"`

	// Secrets: Collection of secrets used by a Container Apps Job
	Secrets []Secret_STATUS `json:"secrets,omitempty"`

	// TriggerType: Trigger type of the job
	TriggerType *JobConfiguration_TriggerType_STATUS `json:"triggerType,omitempty"`
}

// Container Apps Job versioned application definition. Defines the desired state of an immutable revision. Any changes to
// this section Will result in a new revision being created
type JobTemplate_STATUS struct {
	// Containers: List of container definitions for the Container App.
	Containers []Container_STATUS `json:"containers,omitempty"`

	// InitContainers: List of specialized containers that run before app containers.
	InitContainers []BaseContainer_STATUS `json:"initContainers,omitempty"`

	// Volumes: List of volume definitions for the Container App.
	Volumes []Volume_STATUS `json:"volumes,omitempty"`
}

type JobConfiguration_EventTriggerConfig_STATUS struct {
	Parallelism            *int `json:"parallelism,omitempty"`
	ReplicaCompletionCount *int `json:"replicaCompletionCount,omitempty"`

	// Scale: Scaling configurations for event driven jobs.
	Scale *JobScale_STATUS `json:"scale,omitempty"`
}

type JobConfiguration_ManualTriggerConfig_STATUS struct {
	Parallelism            *int `json:"parallelism,omitempty"`
	ReplicaCompletionCount *int `json:"replicaCompletionCount,omitempty"`
}

type JobConfiguration_ScheduleTriggerConfig_STATUS struct {
	// CronExpression: Cron formatted repeating schedule ("* * * * *") of a Cron Job.
	CronExpression         *string `json:"cronExpression,omitempty"`
	Parallelism            *int    `json:"parallelism,omitempty"`
	ReplicaCompletionCount *int    `json:"replicaCompletionCount,omitempty"`
}

type JobConfiguration_TriggerType_STATUS string

const (
	JobConfiguration_TriggerType_STATUS_Event    = JobConfiguration_TriggerType_STATUS("Event")
	JobConfiguration_TriggerType_STATUS_Manual   = JobConfiguration_TriggerType_STATUS("Manual")
	JobConfiguration_TriggerType_STATUS_Schedule = JobConfiguration_TriggerType_STATUS("Schedule")
)

// Mapping from string to JobConfiguration_TriggerType_STATUS
var jobConfiguration_TriggerType_STATUS_Values = map[string]JobConfiguration_TriggerType_STATUS{
	"event":    JobConfiguration_TriggerType_STATUS_Event,
	"manual":   JobConfiguration_TriggerType_STATUS_Manual,
	"schedule": JobConfiguration_TriggerType_STATUS_Schedule,
}

// Scaling configurations for event driven jobs.
type JobScale_STATUS struct {
	// MaxExecutions: Maximum number of job executions that are created for a trigger, default 100.
	MaxExecutions *int `json:"maxExecutions,omitempty"`

	// MinExecutions: Minimum number of job executions that are created for a trigger, default 0
	MinExecutions   *int `json:"minExecutions,omitempty"`
	PollingInterval *int `json:"pollingInterval,omitempty"`

	// Rules: Scaling rules.
	Rules []JobScaleRule_STATUS `json:"rules,omitempty"`
}

// Scaling rule.
type JobScaleRule_STATUS struct {
	// Auth: Authentication secrets for the scale rule.
	Auth []ScaleRuleAuth_STATUS `json:"auth,omitempty"`

	// Metadata: Metadata properties to describe the scale rule.
	Metadata map[string]v1.JSON `json:"metadata,omitempty"`

	// Name: Scale Rule Name
	Name *string `json:"name,omitempty"`

	// Type: Type of the scale rule
	// eg: azure-servicebus, redis etc.
	Type *string `json:"type,omitempty"`
}
