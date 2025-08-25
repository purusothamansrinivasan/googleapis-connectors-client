

# ProvisionedResource

Describes provisioned dataplane resources.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**resourceType** | **String** | Type of the resource. This can be either a GCP resource or a custom one (e.g. another cloud provider&#39;s VM). For GCP compute resources use singular form of the names listed in GCP compute API documentation (https://cloud.google.com/compute/docs/reference/rest/v1/), prefixed with &#39;compute-&#39;, for example: &#39;compute-instance&#39;, &#39;compute-disk&#39;, &#39;compute-autoscaler&#39;. |  [optional] |
|**resourceUrl** | **String** | URL identifying the resource, e.g. \&quot;https://www.googleapis.com/compute/v1/projects/...)\&quot;. |  [optional] |



