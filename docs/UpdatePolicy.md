

# UpdatePolicy

Maintenance policy applicable to instance updates.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**channel** | [**ChannelEnum**](#ChannelEnum) | Optional. Relative scheduling channel applied to resource. |  [optional] |
|**denyMaintenancePeriods** | [**List&lt;DenyMaintenancePeriod&gt;**](DenyMaintenancePeriod.md) | Deny Maintenance Period that is applied to resource to indicate when maintenance is forbidden. The protocol supports zero-to-many such periods, but the current SLM Rollout implementation only supports zero-to-one. |  [optional] |
|**window** | [**MaintenanceWindow**](MaintenanceWindow.md) |  |  [optional] |



## Enum: ChannelEnum

| Name | Value |
|---- | -----|
| UPDATE_CHANNEL_UNSPECIFIED | &quot;UPDATE_CHANNEL_UNSPECIFIED&quot; |
| EARLIER | &quot;EARLIER&quot; |
| LATER | &quot;LATER&quot; |
| WEEK1 | &quot;WEEK1&quot; |
| WEEK2 | &quot;WEEK2&quot; |
| WEEK5 | &quot;WEEK5&quot; |



