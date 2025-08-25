

# TimeOfDay

Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**hours** | **Integer** | Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value \&quot;24:00:00\&quot; for scenarios like business closing time. |  [optional] |
|**minutes** | **Integer** | Minutes of hour of day. Must be from 0 to 59. |  [optional] |
|**nanos** | **Integer** | Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999. |  [optional] |
|**seconds** | **Integer** | Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds. |  [optional] |



