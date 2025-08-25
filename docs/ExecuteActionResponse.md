

# ExecuteActionResponse

Response message for ActionService.ExecuteAction

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**results** | **List&lt;Map&lt;String, Object&gt;&gt;** | In the case of successful invocation of the specified action, the results Struct contains values based on the response of the action invoked. 1. If the action execution produces any entities as a result, they are returned as an array of Structs with the &#39;key&#39; being the field name and the &#39;value&#39; being the value of that field in each result row. { &#39;results&#39;: [{&#39;key&#39;: &#39;value&#39;}, ...] } |  [optional] |



