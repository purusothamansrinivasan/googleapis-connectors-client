package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// UpdatePolicy represents the UpdatePolicy schema from the OpenAPI specification
type UpdatePolicy struct {
	Channel string `json:"channel,omitempty"` // Optional. Relative scheduling channel applied to resource.
	Denymaintenanceperiods []DenyMaintenancePeriod `json:"denyMaintenancePeriods,omitempty"` // Deny Maintenance Period that is applied to resource to indicate when maintenance is forbidden. The protocol supports zero-to-many such periods, but the current SLM Rollout implementation only supports zero-to-one.
	Window MaintenanceWindow `json:"window,omitempty"` // MaintenanceWindow definition.
}

// InputParameter represents the InputParameter schema from the OpenAPI specification
type InputParameter struct {
	Additionaldetails map[string]interface{} `json:"additionalDetails,omitempty"` // The following map contains fields that are not explicitly mentioned above,this give connectors the flexibility to add new metadata fields.
	Datatype string `json:"dataType,omitempty"` // The data type of the Parameter
	Defaultvalue interface{} `json:"defaultValue,omitempty"` // The following field specifies the default value of the Parameter provided by the external system if a value is not provided.
	Description string `json:"description,omitempty"` // A brief description of the Parameter.
	Jsonschema JsonSchema `json:"jsonSchema,omitempty"` // JsonSchema representation of schema metadata
	Name string `json:"name,omitempty"` // Name of the Parameter.
	Nullable bool `json:"nullable,omitempty"` // Specifies whether a null value is allowed.
}

// DenyMaintenancePeriod represents the DenyMaintenancePeriod schema from the OpenAPI specification
type DenyMaintenancePeriod struct {
	Time TimeOfDay `json:"time,omitempty"` // Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.
	Enddate Date `json:"endDate,omitempty"` // Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following: * A full date, with non-zero year, month, and day values. * A month and day, with a zero year (for example, an anniversary). * A year on its own, with a zero month and a zero day. * A year and month, with a zero day (for example, a credit card expiration date). Related types: * google.type.TimeOfDay * google.type.DateTime * google.protobuf.Timestamp
	Startdate Date `json:"startDate,omitempty"` // Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following: * A full date, with non-zero year, month, and day values. * A month and day, with a zero year (for example, an anniversary). * A year on its own, with a zero month and a zero day. * A year and month, with a zero day (for example, a credit card expiration date). Related types: * google.type.TimeOfDay * google.type.DateTime * google.protobuf.Timestamp
}

// MaintenanceSettings represents the MaintenanceSettings schema from the OpenAPI specification
type MaintenanceSettings struct {
	Exclude bool `json:"exclude,omitempty"` // Optional. Exclude instance from maintenance. When true, rollout service will not attempt maintenance on the instance. Rollout service will include the instance in reported rollout progress as not attempted.
	Isrollback bool `json:"isRollback,omitempty"` // Optional. If the update call is triggered from rollback, set the value as true.
	Maintenancepolicies map[string]interface{} `json:"maintenancePolicies,omitempty"` // Optional. The MaintenancePolicies that have been attached to the instance. The key must be of the type name of the oneof policy name defined in MaintenancePolicy, and the embedded policy must define the same policy type. For details, please refer to go/mr-user-guide. Should not be set if maintenance_policy_names is set. If only the name is needed, then only populate MaintenancePolicy.name.
}

// ProvisionedResource represents the ProvisionedResource schema from the OpenAPI specification
type ProvisionedResource struct {
	Resourcetype string `json:"resourceType,omitempty"` // Type of the resource. This can be either a GCP resource or a custom one (e.g. another cloud provider's VM). For GCP compute resources use singular form of the names listed in GCP compute API documentation (https://cloud.google.com/compute/docs/reference/rest/v1/), prefixed with 'compute-', for example: 'compute-instance', 'compute-disk', 'compute-autoscaler'.
	Resourceurl string `json:"resourceUrl,omitempty"` // URL identifying the resource, e.g. "https://www.googleapis.com/compute/v1/projects/...)".
}

// RefreshAccessTokenResponse represents the RefreshAccessTokenResponse schema from the OpenAPI specification
type RefreshAccessTokenResponse struct {
	Accesscredentials AccessCredentials `json:"accessCredentials,omitempty"` // AccessCredentials includes the OAuth access token, and the other fields returned along with it.
}

// Action represents the Action schema from the OpenAPI specification
type Action struct {
	Description string `json:"description,omitempty"` // Brief Description of action
	Displayname string `json:"displayName,omitempty"` // Display Name of action to be shown on client side
	Inputjsonschema JsonSchema `json:"inputJsonSchema,omitempty"` // JsonSchema representation of schema metadata
	Inputparameters []InputParameter `json:"inputParameters,omitempty"` // List containing input parameter metadata.
	Name string `json:"name,omitempty"` // Name of the action.
	Resultjsonschema JsonSchema `json:"resultJsonSchema,omitempty"` // JsonSchema representation of schema metadata
	Resultmetadata []ResultMetadata `json:"resultMetadata,omitempty"` // List containing the metadata of result fields.
}

// ExecuteSqlQueryResponse represents the ExecuteSqlQueryResponse schema from the OpenAPI specification
type ExecuteSqlQueryResponse struct {
	Results []map[string]interface{} `json:"results,omitempty"` // In the case of successful execution of the query the response contains results returned by the external system. For example, the result rows of the query are contained in the 'results' Struct list - "results": [ { "field1": "val1", "field2": "val2",.. },.. ] Each Struct row can contain fields any type of like nested Structs or lists.
}

// MaintenanceSchedule represents the MaintenanceSchedule schema from the OpenAPI specification
type MaintenanceSchedule struct {
	Canreschedule bool `json:"canReschedule,omitempty"` // This field is deprecated, and will be always set to true since reschedule can happen multiple times now. This field should not be removed until all service producers remove this for their customers.
	Endtime string `json:"endTime,omitempty"` // The scheduled end time for the maintenance.
	Rolloutmanagementpolicy string `json:"rolloutManagementPolicy,omitempty"` // The rollout management policy this maintenance schedule is associated with. When doing reschedule update request, the reschedule should be against this given policy.
	Scheduledeadlinetime string `json:"scheduleDeadlineTime,omitempty"` // schedule_deadline_time is the time deadline any schedule start time cannot go beyond, including reschedule. It's normally the initial schedule start time plus maintenance window length (1 day or 1 week). Maintenance cannot be scheduled to start beyond this deadline.
	Starttime string `json:"startTime,omitempty"` // The scheduled start time for the maintenance.
}

// Date represents the Date schema from the OpenAPI specification
type Date struct {
	Day int `json:"day,omitempty"` // Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	Month int `json:"month,omitempty"` // Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
	Year int `json:"year,omitempty"` // Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
}

// UpdateEntitiesWithConditionsResponse represents the UpdateEntitiesWithConditionsResponse schema from the OpenAPI specification
type UpdateEntitiesWithConditionsResponse struct {
	Response map[string]interface{} `json:"response,omitempty"` // Response returned by the external system.
}

// MaintenanceWindow represents the MaintenanceWindow schema from the OpenAPI specification
type MaintenanceWindow struct {
	Dailycycle DailyCycle `json:"dailyCycle,omitempty"` // Time window specified for daily operations.
	Weeklycycle WeeklyCycle `json:"weeklyCycle,omitempty"` // Time window specified for weekly operations.
}

// ExecuteSqlQueryRequest represents the ExecuteSqlQueryRequest schema from the OpenAPI specification
type ExecuteSqlQueryRequest struct {
	Query Query `json:"query,omitempty"` // A wrapper around the SQL query statement. This is needed so that the JSON representation of ExecuteSqlQueryRequest has the following format: `{"query":"select *"}`.
}

// Entity represents the Entity schema from the OpenAPI specification
type Entity struct {
	Name string `json:"name,omitempty"` // Output only. Resource name of the Entity. Format: projects/{project}/locations/{location}/connections/{connection}/entityTypes/{type}/entities/{id}
	Fields map[string]interface{} `json:"fields,omitempty"` // Fields of the entity. The key is name of the field and the value contains the applicable `google.protobuf.Value` entry for this field.
}

// ExecuteActionRequest represents the ExecuteActionRequest schema from the OpenAPI specification
type ExecuteActionRequest struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"` // Parameters for executing the action. The parameters can be key/value pairs or nested structs.
}

// ResultMetadata represents the ResultMetadata schema from the OpenAPI specification
type ResultMetadata struct {
	Description string `json:"description,omitempty"` // A brief description of the metadata field.
	Jsonschema JsonSchema `json:"jsonSchema,omitempty"` // JsonSchema representation of schema metadata
	Name string `json:"name,omitempty"` // Name of the metadata field.
	Datatype string `json:"dataType,omitempty"` // The data type of the metadata field
}

// ExecuteActionResponse represents the ExecuteActionResponse schema from the OpenAPI specification
type ExecuteActionResponse struct {
	Results []map[string]interface{} `json:"results,omitempty"` // In the case of successful invocation of the specified action, the results Struct contains values based on the response of the action invoked. 1. If the action execution produces any entities as a result, they are returned as an array of Structs with the 'key' being the field name and the 'value' being the value of that field in each result row. { 'results': [{'key': 'value'}, ...] }
}

// ExchangeAuthCodeResponse represents the ExchangeAuthCodeResponse schema from the OpenAPI specification
type ExchangeAuthCodeResponse struct {
	Accesscredentials AccessCredentials `json:"accessCredentials,omitempty"` // AccessCredentials includes the OAuth access token, and the other fields returned along with it.
}

// MaintenancePolicy represents the MaintenancePolicy schema from the OpenAPI specification
type MaintenancePolicy struct {
	Updatepolicy UpdatePolicy `json:"updatePolicy,omitempty"` // Maintenance policy applicable to instance updates.
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time when the resource was updated.
	Createtime string `json:"createTime,omitempty"` // Output only. The time when the resource was created.
	Description string `json:"description,omitempty"` // Optional. Description of what this policy is for. Create/Update methods return INVALID_ARGUMENT if the length is greater than 512.
	Labels map[string]interface{} `json:"labels,omitempty"` // Optional. Resource labels to represent user provided metadata. Each label is a key-value pair, where both the key and the value are arbitrary strings provided by the user.
	Name string `json:"name,omitempty"` // Required. MaintenancePolicy name using the form: `projects/{project_id}/locations/{location_id}/maintenancePolicies/{maintenance_policy_id}` where {project_id} refers to a GCP consumer project ID, {location_id} refers to a GCP region/zone, {maintenance_policy_id} must be 1-63 characters long and match the regular expression `[a-z0-9]([-a-z0-9]*[a-z0-9])?`.
	State string `json:"state,omitempty"` // Optional. The state of the policy.
}

// SloMetadata represents the SloMetadata schema from the OpenAPI specification
type SloMetadata struct {
	Perslieligibility PerSliSloEligibility `json:"perSliEligibility,omitempty"` // PerSliSloEligibility is a mapping from an SLI name to eligibility.
	Tier string `json:"tier,omitempty"` // Name of the SLO tier the Instance belongs to. This name will be expected to match the tiers specified in the service SLO configuration. Field is mandatory and must not be empty.
	Nodes []NodeSloMetadata `json:"nodes,omitempty"` // Optional. List of nodes. Some producers need to use per-node metadata to calculate SLO. This field allows such producers to publish per-node SLO meta data, which will be consumed by SSA Eligibility Exporter and published in the form of per node metric to Monarch.
}

// Instance represents the Instance schema from the OpenAPI specification
type Instance struct {
	State string `json:"state,omitempty"` // Output only. Current lifecycle state of the resource (e.g. if it's being created or ready to use).
	Instancetype string `json:"instanceType,omitempty"` // Optional. The instance_type of this instance of format: projects/{project_number}/locations/{location_id}/instanceTypes/{instance_type_id}. Instance Type represents a high-level tier or SKU of the service that this instance belong to. When enabled(eg: Maintenance Rollout), Rollout uses 'instance_type' along with 'software_versions' to determine whether instance needs an update or not.
	Notificationparameters map[string]interface{} `json:"notificationParameters,omitempty"` // Optional. notification_parameter are information that service producers may like to include that is not relevant to Rollout. This parameter will only be passed to Gamma and Cloud Logging for notification/logging purpose.
	Producermetadata map[string]interface{} `json:"producerMetadata,omitempty"` // Output only. Custom string attributes used primarily to expose producer-specific information in monitoring dashboards. See go/get-instance-metadata.
	Provisionedresources []ProvisionedResource `json:"provisionedResources,omitempty"` // Output only. The list of data plane resources provisioned for this instance, e.g. compute VMs. See go/get-instance-metadata.
	Slometadata SloMetadata `json:"sloMetadata,omitempty"` // SloMetadata contains resources required for proper SLO classification of the instance.
	Updatetime string `json:"updateTime,omitempty"` // Output only. Timestamp when the resource was last modified.
	Slminstancetemplate string `json:"slmInstanceTemplate,omitempty"` // Link to the SLM instance template. Only populated when updating SLM instances via SSA's Actuation service adaptor. Service producers with custom control plane (e.g. Cloud SQL) doesn't need to populate this field. Instead they should use software_versions.
	Consumerdefinedname string `json:"consumerDefinedName,omitempty"` // consumer_defined_name is the name of the instance set by the service consumers. Generally this is different from the `name` field which reperesents the system-assigned id of the instance which the service consumers do not recognize. This is a required field for tenants onboarding to Maintenance Window notifications (go/slm-rollout-maintenance-policies#prerequisites).
	Maintenanceschedules map[string]interface{} `json:"maintenanceSchedules,omitempty"` // The MaintenanceSchedule contains the scheduling information of published maintenance schedule with same key as software_versions.
	Name string `json:"name,omitempty"` // Unique name of the resource. It uses the form: `projects/{project_number}/locations/{location_id}/instances/{instance_id}` Note: This name is passed, stored and logged across the rollout system. So use of consumer project_id or any other consumer PII in the name is strongly discouraged for wipeout (go/wipeout) compliance. See go/elysium/project_ids#storage-guidance for more details.
	Labels map[string]interface{} `json:"labels,omitempty"` // Optional. Resource labels to represent user provided metadata. Each label is a key-value pair, where both the key and the value are arbitrary strings provided by the user.
	Softwareversions map[string]interface{} `json:"softwareVersions,omitempty"` // Software versions that are used to deploy this instance. This can be mutated by rollout services.
	Createtime string `json:"createTime,omitempty"` // Output only. Timestamp when the resource was created.
	Maintenancesettings MaintenanceSettings `json:"maintenanceSettings,omitempty"` // Maintenance settings associated with instance. Allows service producers and end users to assign settings that controls maintenance on this instance.
	Maintenancepolicynames map[string]interface{} `json:"maintenancePolicyNames,omitempty"` // Optional. The MaintenancePolicies that have been attached to the instance. The key must be of the type name of the oneof policy name defined in MaintenancePolicy, and the referenced policy must define the same policy type. For details, please refer to go/mr-user-guide. Should not be set if maintenance_settings.maintenance_policies is set.
	Tenantprojectid string `json:"tenantProjectId,omitempty"` // Output only. ID of the associated GCP tenant project. See go/get-instance-metadata.
}

// CheckStatusResponse represents the CheckStatusResponse schema from the OpenAPI specification
type CheckStatusResponse struct {
	Description string `json:"description,omitempty"` // When the connector is not in ACTIVE state, the description must be populated to specify the reason why it's not in ACTIVE state.
	State string `json:"state,omitempty"` // State of the connector.
}

// Field represents the Field schema from the OpenAPI specification
type Field struct {
	Defaultvalue interface{} `json:"defaultValue,omitempty"` // The following field specifies the default value of the Field provided by the external system if a value is not provided.
	Jsonschema JsonSchema `json:"jsonSchema,omitempty"` // JsonSchema representation of schema metadata
	Additionaldetails map[string]interface{} `json:"additionalDetails,omitempty"` // The following map contains fields that are not explicitly mentioned above,this give connectors the flexibility to add new metadata fields.
	Description string `json:"description,omitempty"` // A brief description of the Field.
	Name string `json:"name,omitempty"` // Name of the Field.
	Nullable bool `json:"nullable,omitempty"` // Specifies whether a null value is allowed.
	Reference Reference `json:"reference,omitempty"`
	Datatype string `json:"dataType,omitempty"` // The data type of the Field.
	Key bool `json:"key,omitempty"` // The following boolean field specifies if the current Field acts as a primary key or id if the parent is of type entity.
}

// DailyCycle represents the DailyCycle schema from the OpenAPI specification
type DailyCycle struct {
	Duration string `json:"duration,omitempty"` // Output only. Duration of the time window, set by service producer.
	Starttime TimeOfDay `json:"startTime,omitempty"` // Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.
}

// NodeSloMetadata represents the NodeSloMetadata schema from the OpenAPI specification
type NodeSloMetadata struct {
	Perslieligibility PerSliSloEligibility `json:"perSliEligibility,omitempty"` // PerSliSloEligibility is a mapping from an SLI name to eligibility.
	Location string `json:"location,omitempty"` // The location of the node, if different from instance location.
	Nodeid string `json:"nodeId,omitempty"` // The id of the node. This should be equal to SaasInstanceNode.node_id.
}

// QueryParameter represents the QueryParameter schema from the OpenAPI specification
type QueryParameter struct {
	Datatype string `json:"dataType,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// ExchangeAuthCodeRequest represents the ExchangeAuthCodeRequest schema from the OpenAPI specification
type ExchangeAuthCodeRequest struct {
}

// ListEntitiesResponse represents the ListEntitiesResponse schema from the OpenAPI specification
type ListEntitiesResponse struct {
	Entities []Entity `json:"entities,omitempty"` // List containing entity rows.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Next page token if more records are available.
}

// NotificationParameter represents the NotificationParameter schema from the OpenAPI specification
type NotificationParameter struct {
	Values []string `json:"values,omitempty"` // Optional. Array of string values. e.g. instance's replica information.
}

// RefreshAccessTokenRequest represents the RefreshAccessTokenRequest schema from the OpenAPI specification
type RefreshAccessTokenRequest struct {
}

// AccessCredentials represents the AccessCredentials schema from the OpenAPI specification
type AccessCredentials struct {
	Refreshtoken string `json:"refreshToken,omitempty"` // OAuth refresh token.
	Accesstoken string `json:"accessToken,omitempty"` // OAuth access token.
	Expiresin string `json:"expiresIn,omitempty"` // Duration till the access token expires.
}

// SloEligibility represents the SloEligibility schema from the OpenAPI specification
type SloEligibility struct {
	Eligible bool `json:"eligible,omitempty"` // Whether an instance is eligible or ineligible.
	Reason string `json:"reason,omitempty"` // User-defined reason for the current value of instance eligibility. Usually, this can be directly mapped to the internal state. An empty reason is allowed.
}

// TimeOfDay represents the TimeOfDay schema from the OpenAPI specification
type TimeOfDay struct {
	Minutes int `json:"minutes,omitempty"` // Minutes of hour of day. Must be from 0 to 59.
	Nanos int `json:"nanos,omitempty"` // Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	Seconds int `json:"seconds,omitempty"` // Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.
	Hours int `json:"hours,omitempty"` // Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time.
}

// Query represents the Query schema from the OpenAPI specification
type Query struct {
	Query string `json:"query,omitempty"` // Required. Sql query to execute.
	Queryparameters []QueryParameter `json:"queryParameters,omitempty"` // In the struct, the value corresponds to the value of query parameter and date type corresponds to the date type of the query parameter.
	Timeout string `json:"timeout,omitempty"` // Sets the number of seconds the driver will wait for a query to execute.
	Maxrows string `json:"maxRows,omitempty"` // Sets the limit for the maximum number of rows returned after the query execution.
}

// WeeklyCycle represents the WeeklyCycle schema from the OpenAPI specification
type WeeklyCycle struct {
	Schedule []Schedule `json:"schedule,omitempty"` // User can specify multiple windows in a week. Minimum of 1 window.
}

// JsonSchema represents the JsonSchema schema from the OpenAPI specification
type JsonSchema struct {
	Items JsonSchema `json:"items,omitempty"` // JsonSchema representation of schema metadata
	TypeField []string `json:"type,omitempty"` // JSON Schema Validation: A Vocabulary for Structural Validation of JSON
	Description string `json:"description,omitempty"` // A description of this schema.
	Properties map[string]interface{} `json:"properties,omitempty"` // The child schemas, applicable only if this is of type `object`. The key is the name of the property and the value is the json schema that describes that property
	Required []string `json:"required,omitempty"` // Whether this property is required.
	Additionaldetails map[string]interface{} `json:"additionalDetails,omitempty"` // Additional details apart from standard json schema fields, this gives flexibility to store metadata about the schema
	Enum []interface{} `json:"enum,omitempty"` // Possible values for an enumeration. This works in conjunction with `type` to represent types with a fixed set of legal values
	Format string `json:"format,omitempty"` // Format of the value as per https://json-schema.org/understanding-json-schema/reference/string.html#format
	Jdbctype string `json:"jdbcType,omitempty"` // JDBC datatype of the field.
}

// ListActionsResponse represents the ListActionsResponse schema from the OpenAPI specification
type ListActionsResponse struct {
	Unsupportedactionnames []string `json:"unsupportedActionNames,omitempty"` // List of actions which contain unsupported Datatypes. Check datatype.proto for more information.
	Actions []Action `json:"actions,omitempty"` // List of action metadata.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Next page token if more actions available.
}

// Reference represents the Reference schema from the OpenAPI specification
type Reference struct {
	Name string `json:"name,omitempty"` // Name of the reference field.
	TypeField string `json:"type,omitempty"` // Name of reference entity type.
}

// PerSliSloEligibility represents the PerSliSloEligibility schema from the OpenAPI specification
type PerSliSloEligibility struct {
	Eligibilities map[string]interface{} `json:"eligibilities,omitempty"` // An entry in the eligibilities map specifies an eligibility for a particular SLI for the given instance. The SLI key in the name must be a valid SLI name specified in the Eligibility Exporter binary flags otherwise an error will be emitted by Eligibility Exporter and the oncaller will be alerted. If an SLI has been defined in the binary flags but the eligibilities map does not contain it, the corresponding SLI time series will not be emitted by the Eligibility Exporter. This ensures a smooth rollout and compatibility between the data produced by different versions of the Eligibility Exporters. If eligibilities map contains a key for an SLI which has not been declared in the binary flags, there will be an error message emitted in the Eligibility Exporter log and the metric for the SLI in question will not be emitted.
}

// Schedule represents the Schedule schema from the OpenAPI specification
type Schedule struct {
	Duration string `json:"duration,omitempty"` // Output only. Duration of the time window, set by service producer.
	Starttime TimeOfDay `json:"startTime,omitempty"` // Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.
	Day string `json:"day,omitempty"` // Allows to define schedule that runs specified day of the week.
}

// ListEntityTypesResponse represents the ListEntityTypesResponse schema from the OpenAPI specification
type ListEntityTypesResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Next page token if more entity types available.
	Types []EntityType `json:"types,omitempty"` // List of metadata related to all entity types.
	Unsupportedtypenames []string `json:"unsupportedTypeNames,omitempty"` // List of entity type names which contain unsupported Datatypes. Check datatype.proto for more information.
}

// CheckReadinessResponse represents the CheckReadinessResponse schema from the OpenAPI specification
type CheckReadinessResponse struct {
	Status string `json:"status,omitempty"`
}

// Empty represents the Empty schema from the OpenAPI specification
type Empty struct {
}

// EntityType represents the EntityType schema from the OpenAPI specification
type EntityType struct {
	Fields []Field `json:"fields,omitempty"` // List containing metadata information about each field of the entity type.
	Jsonschema JsonSchema `json:"jsonSchema,omitempty"` // JsonSchema representation of schema metadata
	Name string `json:"name,omitempty"` // The name of the entity type.
}
