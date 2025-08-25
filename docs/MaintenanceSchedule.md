

# MaintenanceSchedule

Maintenance schedule which is exposed to customer and potentially end user, indicating published upcoming future maintenance schedule

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**canReschedule** | **Boolean** | This field is deprecated, and will be always set to true since reschedule can happen multiple times now. This field should not be removed until all service producers remove this for their customers. |  [optional] |
|**endTime** | **String** | The scheduled end time for the maintenance. |  [optional] |
|**rolloutManagementPolicy** | **String** | The rollout management policy this maintenance schedule is associated with. When doing reschedule update request, the reschedule should be against this given policy. |  [optional] |
|**scheduleDeadlineTime** | **String** | schedule_deadline_time is the time deadline any schedule start time cannot go beyond, including reschedule. It&#39;s normally the initial schedule start time plus maintenance window length (1 day or 1 week). Maintenance cannot be scheduled to start beyond this deadline. |  [optional] |
|**startTime** | **String** | The scheduled start time for the maintenance. |  [optional] |



