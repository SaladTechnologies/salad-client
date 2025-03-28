# ContainerGroupProbeGrpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int32** | The port number on which the gRPC health check service is exposed. | 
**Service** | **string** | The name of the gRPC service that implements the health check protocol. | 

## Methods

### NewContainerGroupProbeGrpc

`func NewContainerGroupProbeGrpc(port int32, service string, ) *ContainerGroupProbeGrpc`

NewContainerGroupProbeGrpc instantiates a new ContainerGroupProbeGrpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupProbeGrpcWithDefaults

`func NewContainerGroupProbeGrpcWithDefaults() *ContainerGroupProbeGrpc`

NewContainerGroupProbeGrpcWithDefaults instantiates a new ContainerGroupProbeGrpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *ContainerGroupProbeGrpc) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ContainerGroupProbeGrpc) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ContainerGroupProbeGrpc) SetPort(v int32)`

SetPort sets Port field to given value.


### GetService

`func (o *ContainerGroupProbeGrpc) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ContainerGroupProbeGrpc) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ContainerGroupProbeGrpc) SetService(v string)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


