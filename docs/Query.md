

# Query

A wrapper around the SQL query statement. This is needed so that the JSON representation of ExecuteSqlQueryRequest has the following format: `{\"query\":\"select *\"}`.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**maxRows** | **String** | Sets the limit for the maximum number of rows returned after the query execution. |  [optional] |
|**query** | **String** | Required. Sql query to execute. |  [optional] |
|**queryParameters** | [**List&lt;QueryParameter&gt;**](QueryParameter.md) | In the struct, the value corresponds to the value of query parameter and date type corresponds to the date type of the query parameter. |  [optional] |
|**timeout** | **String** | Sets the number of seconds the driver will wait for a query to execute. |  [optional] |



