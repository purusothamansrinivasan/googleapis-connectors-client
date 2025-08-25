

# MaintenancePolicy

LINT.IfChange Defines policies to service maintenance events.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**createTime** | **String** | Output only. The time when the resource was created. |  [optional] |
|**description** | **String** | Optional. Description of what this policy is for. Create/Update methods return INVALID_ARGUMENT if the length is greater than 512. |  [optional] |
|**labels** | **Map&lt;String, String&gt;** | Optional. Resource labels to represent user provided metadata. Each label is a key-value pair, where both the key and the value are arbitrary strings provided by the user. |  [optional] |
|**name** | **String** | Required. MaintenancePolicy name using the form: &#x60;projects/{project_id}/locations/{location_id}/maintenancePolicies/{maintenance_policy_id}&#x60; where {project_id} refers to a GCP consumer project ID, {location_id} refers to a GCP region/zone, {maintenance_policy_id} must be 1-63 characters long and match the regular expression &#x60;[a-z0-9]([-a-z0-9]*[a-z0-9])?&#x60;. |  [optional] |
|**state** | [**StateEnum**](#StateEnum) | Optional. The state of the policy. |  [optional] |
|**updatePolicy** | [**UpdatePolicy**](UpdatePolicy.md) |  |  [optional] |
|**updateTime** | **String** | Output only. The time when the resource was updated. |  [optional] |



## Enum: StateEnum

| Name | Value |
|---- | -----|
| STATE_UNSPECIFIED | &quot;STATE_UNSPECIFIED&quot; |
| READY | &quot;READY&quot; |
| DELETING | &quot;DELETING&quot; |



