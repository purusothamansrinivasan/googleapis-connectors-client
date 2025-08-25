

# Action

Action message contains metadata information about a single action present in the external system.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**description** | **String** | Brief Description of action |  [optional] |
|**displayName** | **String** | Display Name of action to be shown on client side |  [optional] |
|**inputJsonSchema** | [**JsonSchema**](JsonSchema.md) |  |  [optional] |
|**inputParameters** | [**List&lt;InputParameter&gt;**](InputParameter.md) | List containing input parameter metadata. |  [optional] |
|**name** | **String** | Name of the action. |  [optional] |
|**resultJsonSchema** | [**JsonSchema**](JsonSchema.md) |  |  [optional] |
|**resultMetadata** | [**List&lt;ResultMetadata&gt;**](ResultMetadata.md) | List containing the metadata of result fields. |  [optional] |



