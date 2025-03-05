/*
SaladCloud API

The SaladCloud REST API. Please refer to the [SaladCloud API Documentation](https://docs.salad.com/api-reference) for more details.

API version: 0.9.0-alpha.7
Contact: cloud@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package saladclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// SystemLogsAPIService SystemLogsAPI service
type SystemLogsAPIService service

type ApiGetSystemLogsRequest struct {
	ctx context.Context
	ApiService *SystemLogsAPIService
	organizationName string
	projectName string
	containerGroupName string
}

func (r ApiGetSystemLogsRequest) Execute() (*SystemLogList, *http.Response, error) {
	return r.ApiService.GetSystemLogsExecute(r)
}

/*
GetSystemLogs Get System Logs

Gets the System Logs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationName Your organization name. This identifies the billing context for the API operation and represents a security boundary for SaladCloud resources. The organization must be created before using the API, and you must be a member of the organization.
 @param projectName Your project name. This represents a collection of related SaladCloud resources. The project must be created before using the API.
 @param containerGroupName The unique container group name
 @return ApiGetSystemLogsRequest
*/
func (a *SystemLogsAPIService) GetSystemLogs(ctx context.Context, organizationName string, projectName string, containerGroupName string) ApiGetSystemLogsRequest {
	return ApiGetSystemLogsRequest{
		ApiService: a,
		ctx: ctx,
		organizationName: organizationName,
		projectName: projectName,
		containerGroupName: containerGroupName,
	}
}

// Execute executes the request
//  @return SystemLogList
func (a *SystemLogsAPIService) GetSystemLogsExecute(r ApiGetSystemLogsRequest) (*SystemLogList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SystemLogList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemLogsAPIService.GetSystemLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organization_name}/projects/{project_name}/containers/{container_group_name}/system-logs"
	localVarPath = strings.Replace(localVarPath, "{"+"organization_name"+"}", url.PathEscape(parameterValueToString(r.organizationName, "organizationName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project_name"+"}", url.PathEscape(parameterValueToString(r.projectName, "projectName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"container_group_name"+"}", url.PathEscape(parameterValueToString(r.containerGroupName, "containerGroupName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.organizationName) < 2 {
		return localVarReturnValue, nil, reportError("organizationName must have at least 2 elements")
	}
	if strlen(r.organizationName) > 63 {
		return localVarReturnValue, nil, reportError("organizationName must have less than 63 elements")
	}
	if strlen(r.projectName) < 2 {
		return localVarReturnValue, nil, reportError("projectName must have at least 2 elements")
	}
	if strlen(r.projectName) > 63 {
		return localVarReturnValue, nil, reportError("projectName must have less than 63 elements")
	}
	if strlen(r.containerGroupName) < 2 {
		return localVarReturnValue, nil, reportError("containerGroupName must have at least 2 elements")
	}
	if strlen(r.containerGroupName) > 63 {
		return localVarReturnValue, nil, reportError("containerGroupName must have less than 63 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Salad-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
