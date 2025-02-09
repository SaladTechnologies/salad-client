# \SystemLogsAPI

All URIs are relative to *https://api.salad.com/api/public*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSystemLogs**](SystemLogsAPI.md#GetSystemLogs) | **Get** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/system-logs | Get System Logs



## GetSystemLogs

> SystemLogList GetSystemLogs(ctx, organizationName, projectName, containerGroupName).Execute()

Get System Logs



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
	organizationName := "acme-corp" // string | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization.
	projectName := "default" // string | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.
	containerGroupName := "containerGroupName_example" // string | The unique container group name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SystemLogsAPI.GetSystemLogs(context.Background(), organizationName, projectName, containerGroupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SystemLogsAPI.GetSystemLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemLogs`: SystemLogList
	fmt.Fprintf(os.Stdout, "Response from `SystemLogsAPI.GetSystemLogs`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSystemLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SystemLogList**](SystemLogList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

