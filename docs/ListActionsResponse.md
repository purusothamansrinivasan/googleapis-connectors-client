

# ListActionsResponse

Response message for ActionService.ListActions

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**actions** | [**List&lt;Action&gt;**](Action.md) | List of action metadata. |  [optional] |
|**nextPageToken** | **String** | Next page token if more actions available. |  [optional] |
|**unsupportedActionNames** | **List&lt;String&gt;** | List of actions which contain unsupported Datatypes. Check datatype.proto for more information. |  [optional] |



