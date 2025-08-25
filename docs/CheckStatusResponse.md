

# CheckStatusResponse

The status of the connector.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**description** | **String** | When the connector is not in ACTIVE state, the description must be populated to specify the reason why it&#39;s not in ACTIVE state. |  [optional] |
|**state** | [**StateEnum**](#StateEnum) | State of the connector. |  [optional] |



## Enum: StateEnum

| Name | Value |
|---- | -----|
| STATE_UNSPECIFIED | &quot;STATE_UNSPECIFIED&quot; |
| ACTIVE | &quot;ACTIVE&quot; |
| ERROR | &quot;ERROR&quot; |
| AUTH_ERROR | &quot;AUTH_ERROR&quot; |



