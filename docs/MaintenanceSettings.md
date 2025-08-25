

# MaintenanceSettings

Maintenance settings associated with instance. Allows service producers and end users to assign settings that controls maintenance on this instance.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**exclude** | **Boolean** | Optional. Exclude instance from maintenance. When true, rollout service will not attempt maintenance on the instance. Rollout service will include the instance in reported rollout progress as not attempted. |  [optional] |
|**isRollback** | **Boolean** | Optional. If the update call is triggered from rollback, set the value as true. |  [optional] |
|**maintenancePolicies** | [**Map&lt;String, MaintenancePolicy&gt;**](MaintenancePolicy.md) | Optional. The MaintenancePolicies that have been attached to the instance. The key must be of the type name of the oneof policy name defined in MaintenancePolicy, and the embedded policy must define the same policy type. For details, please refer to go/mr-user-guide. Should not be set if maintenance_policy_names is set. If only the name is needed, then only populate MaintenancePolicy.name. |  [optional] |



