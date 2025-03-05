# \QueuesAPI

All URIs are relative to *https://api.salad.com/api/public*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQueue**](QueuesAPI.md#CreateQueue) | **Post** /organizations/{organization_name}/projects/{project_name}/queues | Create Queue
[**CreateQueueJob**](QueuesAPI.md#CreateQueueJob) | **Post** /organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs | Create Job
[**DeleteQueue**](QueuesAPI.md#DeleteQueue) | **Delete** /organizations/{organization_name}/projects/{project_name}/queues/{queue_name} | Delete Queue
[**DeleteQueueJob**](QueuesAPI.md#DeleteQueueJob) | **Delete** /organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs/{queue_job_id} | Delete Job
[**GetQueue**](QueuesAPI.md#GetQueue) | **Get** /organizations/{organization_name}/projects/{project_name}/queues/{queue_name} | Get Queue
[**GetQueueJob**](QueuesAPI.md#GetQueueJob) | **Get** /organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs/{queue_job_id} | Get Job
[**ListQueueJobs**](QueuesAPI.md#ListQueueJobs) | **Get** /organizations/{organization_name}/projects/{project_name}/queues/{queue_name}/jobs | List Jobs
[**ListQueues**](QueuesAPI.md#ListQueues) | **Get** /organizations/{organization_name}/projects/{project_name}/queues | List Queues
[**UpdateQueue**](QueuesAPI.md#UpdateQueue) | **Patch** /organizations/{organization_name}/projects/{project_name}/queues/{queue_name} | Update Queue



## CreateQueue

> Queue CreateQueue(ctx, organizationName, projectName).CreateQueue(createQueue).Execute()

Create Queue



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
	createQueue := *openapiclient.NewCreateQueue("Name_example") // CreateQueue | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueuesAPI.CreateQueue(context.Background(), organizationName, projectName).CreateQueue(createQueue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueuesAPI.CreateQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQueue`: Queue
	fmt.Fprintf(os.Stdout, "Response from `QueuesAPI.CreateQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**projectName** | **string** | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createQueue** | [**CreateQueue**](CreateQueue.md) |  | 

### Return type

[**Queue**](Queue.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQueueJob

> QueueJob CreateQueueJob(ctx, organizationName, projectName, queueName).CreateQueueJob(createQueueJob).Execute()

Create Job



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
	queueName := "queueName_example" // string | The queue name.
	createQueueJob := *openapiclient.NewCreateQueueJob(interface{}(123)) // CreateQueueJob | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueuesAPI.CreateQueueJob(context.Background(), organizationName, projectName, queueName).CreateQueueJob(createQueueJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueuesAPI.CreateQueueJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateQueueJob`: QueueJob
	fmt.Fprintf(os.Stdout, "Response from `QueuesAPI.CreateQueueJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**projectName** | **string** | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API. | 
**queueName** | **string** | The queue name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateQueueJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createQueueJob** | [**CreateQueueJob**](CreateQueueJob.md) |  | 

### Return type

[**QueueJob**](QueueJob.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQueue

> DeleteQueue(ctx, organizationName, projectName, queueName).Execute()

Delete Queue



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
	queueName := "queueName_example" // string | The queue name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QueuesAPI.DeleteQueue(context.Background(), organizationName, projectName, queueName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueuesAPI.DeleteQueue``: %v\n", err)
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
**queueName** | **string** | The queue name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQueueRequest struct via the builder pattern


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


## DeleteQueueJob

> DeleteQueueJob(ctx, organizationName, projectName, queueName, queueJobId).Execute()

Delete Job



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
	queueName := "queueName_example" // string | The queue name.
	queueJobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The job identifier. This is automatically generated and assigned when the job is created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QueuesAPI.DeleteQueueJob(context.Background(), organizationName, projectName, queueName, queueJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueuesAPI.DeleteQueueJob``: %v\n", err)
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
**queueName** | **string** | The queue name. | 
**queueJobId** | **string** | The job identifier. This is automatically generated and assigned when the job is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteQueueJobRequest struct via the builder pattern


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


## GetQueue

> Queue GetQueue(ctx, organizationName, projectName, queueName).Execute()

Get Queue



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
	queueName := "queueName_example" // string | The queue name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueuesAPI.GetQueue(context.Background(), organizationName, projectName, queueName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueuesAPI.GetQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueue`: Queue
	fmt.Fprintf(os.Stdout, "Response from `QueuesAPI.GetQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**projectName** | **string** | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API. | 
**queueName** | **string** | The queue name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Queue**](Queue.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueueJob

> QueueJob GetQueueJob(ctx, organizationName, projectName, queueName, queueJobId).Execute()

Get Job



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
	queueName := "queueName_example" // string | The queue name.
	queueJobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The job identifier. This is automatically generated and assigned when the job is created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueuesAPI.GetQueueJob(context.Background(), organizationName, projectName, queueName, queueJobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueuesAPI.GetQueueJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueueJob`: QueueJob
	fmt.Fprintf(os.Stdout, "Response from `QueuesAPI.GetQueueJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**projectName** | **string** | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API. | 
**queueName** | **string** | The queue name. | 
**queueJobId** | **string** | The job identifier. This is automatically generated and assigned when the job is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**QueueJob**](QueueJob.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQueueJobs

> QueueJobList ListQueueJobs(ctx, organizationName, projectName, queueName).Page(page).PageSize(pageSize).Execute()

List Jobs



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
	queueName := "queueName_example" // string | The queue name.
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueuesAPI.ListQueueJobs(context.Background(), organizationName, projectName, queueName).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueuesAPI.ListQueueJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQueueJobs`: QueueJobList
	fmt.Fprintf(os.Stdout, "Response from `QueuesAPI.ListQueueJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**projectName** | **string** | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API. | 
**queueName** | **string** | The queue name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQueueJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**QueueJobList**](QueueJobList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQueues

> QueueList ListQueues(ctx, organizationName, projectName).Execute()

List Queues



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
	resp, r, err := apiClient.QueuesAPI.ListQueues(context.Background(), organizationName, projectName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueuesAPI.ListQueues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListQueues`: QueueList
	fmt.Fprintf(os.Stdout, "Response from `QueuesAPI.ListQueues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**projectName** | **string** | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListQueuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**QueueList**](QueueList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQueue

> Queue UpdateQueue(ctx, organizationName, projectName, queueName).UpdateQueue(updateQueue).Execute()

Update Queue



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
	queueName := "queueName_example" // string | The queue name.
	updateQueue := *openapiclient.NewUpdateQueue() // UpdateQueue | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QueuesAPI.UpdateQueue(context.Background(), organizationName, projectName, queueName).UpdateQueue(updateQueue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QueuesAPI.UpdateQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateQueue`: Queue
	fmt.Fprintf(os.Stdout, "Response from `QueuesAPI.UpdateQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 
**projectName** | **string** | Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API. | 
**queueName** | **string** | The queue name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateQueue** | [**UpdateQueue**](UpdateQueue.md) |  | 

### Return type

[**Queue**](Queue.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

