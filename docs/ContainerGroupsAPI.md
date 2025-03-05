# \ContainerGroupsAPI

All URIs are relative to *https://api.salad.com/api/public*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContainerGroup**](ContainerGroupsAPI.md#CreateContainerGroup) | **Post** /organizations/{organization_name}/projects/{project_name}/containers | Create Container Group
[**DeleteContainerGroup**](ContainerGroupsAPI.md#DeleteContainerGroup) | **Delete** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name} | Delete Container Group
[**GetContainerGroup**](ContainerGroupsAPI.md#GetContainerGroup) | **Get** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name} | Get Container Group
[**GetContainerGroupInstance**](ContainerGroupsAPI.md#GetContainerGroupInstance) | **Get** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id} | Get Container Group Instance
[**ListContainerGroupInstances**](ContainerGroupsAPI.md#ListContainerGroupInstances) | **Get** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances | List Container Group Instances
[**ListContainerGroups**](ContainerGroupsAPI.md#ListContainerGroups) | **Get** /organizations/{organization_name}/projects/{project_name}/containers | List Container Groups
[**ReallocateContainerGroupInstance**](ContainerGroupsAPI.md#ReallocateContainerGroupInstance) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}/reallocate | Reallocate Container Group Instance
[**RecreateContainerGroupInstance**](ContainerGroupsAPI.md#RecreateContainerGroupInstance) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}/recreate | Recreate Container Group Instance
[**RestartContainerGroupInstance**](ContainerGroupsAPI.md#RestartContainerGroupInstance) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/instances/{container_group_instance_id}/restart | Restart container Group Instance
[**StartContainerGroup**](ContainerGroupsAPI.md#StartContainerGroup) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/start | Start Container Group
[**StopContainerGroup**](ContainerGroupsAPI.md#StopContainerGroup) | **Post** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/stop | Stop Container Group
[**UpdateContainerGroup**](ContainerGroupsAPI.md#UpdateContainerGroup) | **Patch** /organizations/{organization_name}/projects/{project_name}/containers/{container_group_name} | Update Container Group



## CreateContainerGroup

> ContainerGroup CreateContainerGroup(ctx, organizationName, projectName).CreateContainerGroup(createContainerGroup).Execute()

Create Container Group



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
	createContainerGroup := *openapiclient.NewCreateContainerGroup("Name_example", *openapiclient.NewCreateContainer("Image_example", *openapiclient.NewContainerResourceRequirements(int32(123), int32(123))), false, openapiclient.ContainerRestartPolicy("always"), int32(123)) // CreateContainerGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerGroupsAPI.CreateContainerGroup(context.Background(), organizationName, projectName).CreateContainerGroup(createContainerGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.CreateContainerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContainerGroup`: ContainerGroup
	fmt.Fprintf(os.Stdout, "Response from `ContainerGroupsAPI.CreateContainerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**projectName** | **string** | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createContainerGroup** | [**CreateContainerGroup**](CreateContainerGroup.md) |  | 

### Return type

[**ContainerGroup**](ContainerGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainerGroup

> DeleteContainerGroup(ctx, organizationName, projectName, containerGroupName).Execute()

Delete Container Group



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
	r, err := apiClient.ContainerGroupsAPI.DeleteContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.DeleteContainerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

Other parameters are passed through a pointer to a apiDeleteContainerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerGroup

> ContainerGroup GetContainerGroup(ctx, organizationName, projectName, containerGroupName).Execute()

Get Container Group



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
	resp, r, err := apiClient.ContainerGroupsAPI.GetContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.GetContainerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerGroup`: ContainerGroup
	fmt.Fprintf(os.Stdout, "Response from `ContainerGroupsAPI.GetContainerGroup`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetContainerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ContainerGroup**](ContainerGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerGroupInstance

> ContainerGroupInstance GetContainerGroupInstance(ctx, organizationName, projectName, containerGroupName, containerGroupInstanceId).Execute()

Get Container Group Instance



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
	containerGroupInstanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique instance identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerGroupsAPI.GetContainerGroupInstance(context.Background(), organizationName, projectName, containerGroupName, containerGroupInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.GetContainerGroupInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerGroupInstance`: ContainerGroupInstance
	fmt.Fprintf(os.Stdout, "Response from `ContainerGroupsAPI.GetContainerGroupInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**projectName** | **string** | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API. | 
**containerGroupName** | **string** | The unique container group name | 
**containerGroupInstanceId** | **string** | The unique instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerGroupInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ContainerGroupInstance**](ContainerGroupInstance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerGroupInstances

> ContainerGroupInstances ListContainerGroupInstances(ctx, organizationName, projectName, containerGroupName).Execute()

List Container Group Instances



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
	resp, r, err := apiClient.ContainerGroupsAPI.ListContainerGroupInstances(context.Background(), organizationName, projectName, containerGroupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.ListContainerGroupInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContainerGroupInstances`: ContainerGroupInstances
	fmt.Fprintf(os.Stdout, "Response from `ContainerGroupsAPI.ListContainerGroupInstances`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiListContainerGroupInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ContainerGroupInstances**](ContainerGroupInstances.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContainerGroups

> ContainerGroupList ListContainerGroups(ctx, organizationName, projectName).Execute()

List Container Groups



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerGroupsAPI.ListContainerGroups(context.Background(), organizationName, projectName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.ListContainerGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContainerGroups`: ContainerGroupList
	fmt.Fprintf(os.Stdout, "Response from `ContainerGroupsAPI.ListContainerGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**projectName** | **string** | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContainerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContainerGroupList**](ContainerGroupList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReallocateContainerGroupInstance

> ReallocateContainerGroupInstance(ctx, organizationName, projectName, containerGroupName, containerGroupInstanceId).Execute()

Reallocate Container Group Instance



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
	containerGroupInstanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique instance identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContainerGroupsAPI.ReallocateContainerGroupInstance(context.Background(), organizationName, projectName, containerGroupName, containerGroupInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.ReallocateContainerGroupInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**projectName** | **string** | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API. | 
**containerGroupName** | **string** | The unique container group name | 
**containerGroupInstanceId** | **string** | The unique instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiReallocateContainerGroupInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecreateContainerGroupInstance

> RecreateContainerGroupInstance(ctx, organizationName, projectName, containerGroupName, containerGroupInstanceId).Execute()

Recreate Container Group Instance



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
	containerGroupInstanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique instance identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContainerGroupsAPI.RecreateContainerGroupInstance(context.Background(), organizationName, projectName, containerGroupName, containerGroupInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.RecreateContainerGroupInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**projectName** | **string** | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API. | 
**containerGroupName** | **string** | The unique container group name | 
**containerGroupInstanceId** | **string** | The unique instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecreateContainerGroupInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartContainerGroupInstance

> RestartContainerGroupInstance(ctx, organizationName, projectName, containerGroupName, containerGroupInstanceId).Execute()

Restart container Group Instance



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
	containerGroupInstanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique instance identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContainerGroupsAPI.RestartContainerGroupInstance(context.Background(), organizationName, projectName, containerGroupName, containerGroupInstanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.RestartContainerGroupInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**projectName** | **string** | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API. | 
**containerGroupName** | **string** | The unique container group name | 
**containerGroupInstanceId** | **string** | The unique instance identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartContainerGroupInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartContainerGroup

> StartContainerGroup(ctx, organizationName, projectName, containerGroupName).Execute()

Start Container Group



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
	r, err := apiClient.ContainerGroupsAPI.StartContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.StartContainerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

Other parameters are passed through a pointer to a apiStartContainerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopContainerGroup

> StopContainerGroup(ctx, organizationName, projectName, containerGroupName).Execute()

Stop Container Group



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
	r, err := apiClient.ContainerGroupsAPI.StopContainerGroup(context.Background(), organizationName, projectName, containerGroupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.StopContainerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

Other parameters are passed through a pointer to a apiStopContainerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContainerGroup

> ContainerGroup UpdateContainerGroup(ctx, organizationName, projectName, containerGroupName).UpdateContainerGroup(updateContainerGroup).Execute()

Update Container Group



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
	updateContainerGroup := *openapiclient.NewUpdateContainerGroup() // UpdateContainerGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerGroupsAPI.UpdateContainerGroup(context.Background(), organizationName, projectName, containerGroupName).UpdateContainerGroup(updateContainerGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerGroupsAPI.UpdateContainerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContainerGroup`: ContainerGroup
	fmt.Fprintf(os.Stdout, "Response from `ContainerGroupsAPI.UpdateContainerGroup`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateContainerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateContainerGroup** | [**UpdateContainerGroup**](UpdateContainerGroup.md) |  | 

### Return type

[**ContainerGroup**](ContainerGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/merge-patch+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

