

# ExecuteSqlQueryResponse

A response returned by the connection after executing the sql query.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**results** | **List&lt;Map&lt;String, Object&gt;&gt;** | In the case of successful execution of the query the response contains results returned by the external system. For example, the result rows of the query are contained in the &#39;results&#39; Struct list - \&quot;results\&quot;: [ { \&quot;field1\&quot;: \&quot;val1\&quot;, \&quot;field2\&quot;: \&quot;val2\&quot;,.. },.. ] Each Struct row can contain fields any type of like nested Structs or lists. |  [optional] |



