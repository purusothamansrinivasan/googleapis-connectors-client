

# Instance

Instance represents the interface for SLM services to actuate the state of control plane resources. Example Instance in JSON, where consumer-project-number=123456, producer-project-id=cloud-sql: ```json Instance: { \"name\": \"projects/123456/locations/us-east1/instances/prod-instance\", \"create_time\": { \"seconds\": 1526406431, }, \"labels\": { \"env\": \"prod\", \"foo\": \"bar\" }, \"state\": READY, \"software_versions\": { \"software_update\": \"cloud-sql-09-28-2018\", }, \"maintenance_policy_names\": { \"UpdatePolicy\": \"projects/123456/locations/us-east1/maintenancePolicies/prod-update-policy\", } \"tenant_project_id\": \"cloud-sql-test-tenant\", \"producer_metadata\": { \"cloud-sql-tier\": \"basic\", \"cloud-sql-instance-size\": \"1G\", }, \"provisioned_resources\": [ { \"resource-type\": \"compute-instance\", \"resource-url\": \"https://www.googleapis.com/compute/v1/projects/cloud-sql/zones/us-east1-b/instances/vm-1\", } ], \"maintenance_schedules\": { \"csa_rollout\": { \"start_time\": { \"seconds\": 1526406431, }, \"end_time\": { \"seconds\": 1535406431, }, }, \"ncsa_rollout\": { \"start_time\": { \"seconds\": 1526406431, }, \"end_time\": { \"seconds\": 1535406431, }, } }, \"consumer_defined_name\": \"my-sql-instance1\", } ``` LINT.IfChange

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**consumerDefinedName** | **String** | consumer_defined_name is the name of the instance set by the service consumers. Generally this is different from the &#x60;name&#x60; field which reperesents the system-assigned id of the instance which the service consumers do not recognize. This is a required field for tenants onboarding to Maintenance Window notifications (go/slm-rollout-maintenance-policies#prerequisites). |  [optional] |
|**createTime** | **String** | Output only. Timestamp when the resource was created. |  [optional] [readonly] |
|**instanceType** | **String** | Optional. The instance_type of this instance of format: projects/{project_number}/locations/{location_id}/instanceTypes/{instance_type_id}. Instance Type represents a high-level tier or SKU of the service that this instance belong to. When enabled(eg: Maintenance Rollout), Rollout uses &#39;instance_type&#39; along with &#39;software_versions&#39; to determine whether instance needs an update or not. |  [optional] |
|**labels** | **Map&lt;String, String&gt;** | Optional. Resource labels to represent user provided metadata. Each label is a key-value pair, where both the key and the value are arbitrary strings provided by the user. |  [optional] |
|**maintenancePolicyNames** | **Map&lt;String, String&gt;** | Optional. The MaintenancePolicies that have been attached to the instance. The key must be of the type name of the oneof policy name defined in MaintenancePolicy, and the referenced policy must define the same policy type. For details, please refer to go/mr-user-guide. Should not be set if maintenance_settings.maintenance_policies is set. |  [optional] |
|**maintenanceSchedules** | [**Map&lt;String, MaintenanceSchedule&gt;**](MaintenanceSchedule.md) | The MaintenanceSchedule contains the scheduling information of published maintenance schedule with same key as software_versions. |  [optional] |
|**maintenanceSettings** | [**MaintenanceSettings**](MaintenanceSettings.md) |  |  [optional] |
|**name** | **String** | Unique name of the resource. It uses the form: &#x60;projects/{project_number}/locations/{location_id}/instances/{instance_id}&#x60; Note: This name is passed, stored and logged across the rollout system. So use of consumer project_id or any other consumer PII in the name is strongly discouraged for wipeout (go/wipeout) compliance. See go/elysium/project_ids#storage-guidance for more details. |  [optional] |
|**notificationParameters** | [**Map&lt;String, NotificationParameter&gt;**](NotificationParameter.md) | Optional. notification_parameter are information that service producers may like to include that is not relevant to Rollout. This parameter will only be passed to Gamma and Cloud Logging for notification/logging purpose. |  [optional] |
|**producerMetadata** | **Map&lt;String, String&gt;** | Output only. Custom string attributes used primarily to expose producer-specific information in monitoring dashboards. See go/get-instance-metadata. |  [optional] [readonly] |
|**provisionedResources** | [**List&lt;ProvisionedResource&gt;**](ProvisionedResource.md) | Output only. The list of data plane resources provisioned for this instance, e.g. compute VMs. See go/get-instance-metadata. |  [optional] [readonly] |
|**slmInstanceTemplate** | **String** | Link to the SLM instance template. Only populated when updating SLM instances via SSA&#39;s Actuation service adaptor. Service producers with custom control plane (e.g. Cloud SQL) doesn&#39;t need to populate this field. Instead they should use software_versions. |  [optional] |
|**sloMetadata** | [**SloMetadata**](SloMetadata.md) |  |  [optional] |
|**softwareVersions** | **Map&lt;String, String&gt;** | Software versions that are used to deploy this instance. This can be mutated by rollout services. |  [optional] |
|**state** | [**StateEnum**](#StateEnum) | Output only. Current lifecycle state of the resource (e.g. if it&#39;s being created or ready to use). |  [optional] [readonly] |
|**tenantProjectId** | **String** | Output only. ID of the associated GCP tenant project. See go/get-instance-metadata. |  [optional] [readonly] |
|**updateTime** | **String** | Output only. Timestamp when the resource was last modified. |  [optional] [readonly] |



## Enum: StateEnum

| Name | Value |
|---- | -----|
| STATE_UNSPECIFIED | &quot;STATE_UNSPECIFIED&quot; |
| CREATING | &quot;CREATING&quot; |
| READY | &quot;READY&quot; |
| UPDATING | &quot;UPDATING&quot; |
| REPAIRING | &quot;REPAIRING&quot; |
| DELETING | &quot;DELETING&quot; |
| ERROR | &quot;ERROR&quot; |



