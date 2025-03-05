# \WorkloadErrorsAPI

All URIs are relative to *https://api.salad.com/api/public*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorkloadErrors**](WorkloadErrorsAPI.md#GetWorkloadErrors) | **Get** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/errors | Get Workload Errors



## GetWorkloadErrors

> WorkloadErrorList GetWorkloadErrors(ctx, organizationName, projectName, containerGroupName).Execute()

Get Workload Errors



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/SaladTechnologies/salad-client"
)

func main() {
	organizationName := "organizationName_example" // string | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization.
	projectName := "projectName_example" // string | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.
	containerGroupName := "containerGroupName_example" // string | The unique container group name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadErrorsAPI.GetWorkloadErrors(context.Background(), organizationName, projectName, containerGroupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadErrorsAPI.GetWorkloadErrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkloadErrors`: WorkloadErrorList
	fmt.Fprintf(os.Stdout, "Response from `WorkloadErrorsAPI.GetWorkloadErrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**projectName** | **string** | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API. | 
**containerGroupName** | **string** | The unique container group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkloadErrorList**](WorkloadErrorList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

