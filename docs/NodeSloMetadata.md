

# NodeSloMetadata

Node information for custom per-node SLO implementations. SSA does not support per-node SLO, but producers can populate per-node information in SloMetadata for custom precomputations. SSA Eligibility Exporter will emit per-node metric based on this information.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**location** | **String** | The location of the node, if different from instance location. |  [optional] |
|**nodeId** | **String** | The id of the node. This should be equal to SaasInstanceNode.node_id. |  [optional] |
|**perSliEligibility** | [**PerSliSloEligibility**](PerSliSloEligibility.md) |  |  [optional] |



