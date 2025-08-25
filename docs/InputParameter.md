

# InputParameter

Input Parameter message contains metadata about the parameters required for executing an Action.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**additionalDetails** | **Map&lt;String, Object&gt;** | The following map contains fields that are not explicitly mentioned above,this give connectors the flexibility to add new metadata fields. |  [optional] |
|**dataType** | [**DataTypeEnum**](#DataTypeEnum) | The data type of the Parameter |  [optional] |
|**defaultValue** | **Object** | The following field specifies the default value of the Parameter provided by the external system if a value is not provided. |  [optional] |
|**description** | **String** | A brief description of the Parameter. |  [optional] |
|**jsonSchema** | [**JsonSchema**](JsonSchema.md) |  |  [optional] |
|**name** | **String** | Name of the Parameter. |  [optional] |
|**nullable** | **Boolean** | Specifies whether a null value is allowed. |  [optional] |



## Enum: DataTypeEnum

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



