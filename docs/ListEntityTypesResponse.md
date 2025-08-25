

# ListEntityTypesResponse

Response message for EntityService.ListEntityTypes

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**nextPageToken** | **String** | Next page token if more entity types available. |  [optional] |
|**types** | [**List&lt;EntityType&gt;**](EntityType.md) | List of metadata related to all entity types. |  [optional] |
|**unsupportedTypeNames** | **List&lt;String&gt;** | List of entity type names which contain unsupported Datatypes. Check datatype.proto for more information. |  [optional] |



