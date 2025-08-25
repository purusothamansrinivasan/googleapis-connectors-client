

# JsonSchema

JsonSchema representation of schema metadata

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**additionalDetails** | **Map&lt;String, Object&gt;** | Additional details apart from standard json schema fields, this gives flexibility to store metadata about the schema |  [optional] |
|**description** | **String** | A description of this schema. |  [optional] |
|**_enum** | **List&lt;Object&gt;** | Possible values for an enumeration. This works in conjunction with &#x60;type&#x60; to represent types with a fixed set of legal values |  [optional] |
|**format** | **String** | Format of the value as per https://json-schema.org/understanding-json-schema/reference/string.html#format |  [optional] |
|**items** | [**JsonSchema**](JsonSchema.md) |  |  [optional] |
|**jdbcType** | [**JdbcTypeEnum**](#JdbcTypeEnum) | JDBC datatype of the field. |  [optional] |
|**properties** | [**Map&lt;String, JsonSchema&gt;**](JsonSchema.md) | The child schemas, applicable only if this is of type &#x60;object&#x60;. The key is the name of the property and the value is the json schema that describes that property |  [optional] |
|**required** | **List&lt;String&gt;** | Whether this property is required. |  [optional] |
|**type** | **List&lt;String&gt;** | JSON Schema Validation: A Vocabulary for Structural Validation of JSON |  [optional] |



## Enum: JdbcTypeEnum

| Name | Value |
|---- | -----|
| DATA_TYPE_UNSPECIFIED | &quot;DATA_TYPE_UNSPECIFIED&quot; |
| INT | &quot;INT&quot; |
| SMALLINT | &quot;SMALLINT&quot; |
| DOUBLE | &quot;DOUBLE&quot; |
| DATE | &quot;DATE&quot; |
| DATETIME | &quot;DATETIME&quot; |
| TIME | &quot;TIME&quot; |
| STRING | &quot;STRING&quot; |
| LONG | &quot;LONG&quot; |
| BOOLEAN | &quot;BOOLEAN&quot; |
| DECIMAL | &quot;DECIMAL&quot; |
| UUID | &quot;UUID&quot; |
| BLOB | &quot;BLOB&quot; |
| BIT | &quot;BIT&quot; |
| TINYINT | &quot;TINYINT&quot; |
| INTEGER | &quot;INTEGER&quot; |
| BIGINT | &quot;BIGINT&quot; |
| FLOAT | &quot;FLOAT&quot; |
| REAL | &quot;REAL&quot; |
| NUMERIC | &quot;NUMERIC&quot; |
| CHAR | &quot;CHAR&quot; |
| VARCHAR | &quot;VARCHAR&quot; |
| LONGVARCHAR | &quot;LONGVARCHAR&quot; |
| TIMESTAMP | &quot;TIMESTAMP&quot; |
| NCHAR | &quot;NCHAR&quot; |
| NVARCHAR | &quot;NVARCHAR&quot; |
| LONGNVARCHAR | &quot;LONGNVARCHAR&quot; |
| NULL | &quot;NULL&quot; |
| OTHER | &quot;OTHER&quot; |
| JAVA_OBJECT | &quot;JAVA_OBJECT&quot; |
| DISTINCT | &quot;DISTINCT&quot; |
| STRUCT | &quot;STRUCT&quot; |
| ARRAY | &quot;ARRAY&quot; |
| CLOB | &quot;CLOB&quot; |
| REF | &quot;REF&quot; |
| DATALINK | &quot;DATALINK&quot; |
| ROWID | &quot;ROWID&quot; |
| BINARY | &quot;BINARY&quot; |
| VARBINARY | &quot;VARBINARY&quot; |
| LONGVARBINARY | &quot;LONGVARBINARY&quot; |
| NCLOB | &quot;NCLOB&quot; |
| SQLXML | &quot;SQLXML&quot; |
| REF_CURSOR | &quot;REF_CURSOR&quot; |
| TIME_WITH_TIMEZONE | &quot;TIME_WITH_TIMEZONE&quot; |
| TIMESTAMP_WITH_TIMEZONE | &quot;TIMESTAMP_WITH_TIMEZONE&quot; |



