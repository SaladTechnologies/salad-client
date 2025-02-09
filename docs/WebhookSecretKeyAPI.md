# \WebhookSecretKeyAPI

All URIs are relative to *https://api.salad.com/api/public*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebhookSecretKey**](WebhookSecretKeyAPI.md#GetWebhookSecretKey) | **Get** /organizations/{organization_name}/webhook-secret-key | Gets the webhook secret key
[**UpdateWebhookSecretKey**](WebhookSecretKeyAPI.md#UpdateWebhookSecretKey) | **Post** /organizations/{organization_name}/webhook-secret-key | Updates the webhook secret key



## GetWebhookSecretKey

> WebhookSecretKey GetWebhookSecretKey(ctx, organizationName).Execute()

Gets the webhook secret key



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSecretKeyAPI.GetWebhookSecretKey(context.Background(), organizationName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSecretKeyAPI.GetWebhookSecretKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookSecretKey`: WebhookSecretKey
	fmt.Fprintf(os.Stdout, "Response from `WebhookSecretKeyAPI.GetWebhookSecretKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookSecretKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookSecretKey**](WebhookSecretKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookSecretKey

> WebhookSecretKey UpdateWebhookSecretKey(ctx, organizationName).Execute()

Updates the webhook secret key



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookSecretKeyAPI.UpdateWebhookSecretKey(context.Background(), organizationName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookSecretKeyAPI.UpdateWebhookSecretKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhookSecretKey`: WebhookSecretKey
	fmt.Fprintf(os.Stdout, "Response from `WebhookSecretKeyAPI.UpdateWebhookSecretKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationName** | **string** | Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookSecretKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookSecretKey**](WebhookSecretKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

