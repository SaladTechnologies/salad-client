# ContainerGroupProbeTcp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int32** | The TCP port number that the probe should connect to. Must be a valid port number between 0 and 65535. | 

## Methods

### NewContainerGroupProbeTcp

`func NewContainerGroupProbeTcp(port int32, ) *ContainerGroupProbeTcp`

NewContainerGroupProbeTcp instantiates a new ContainerGroupProbeTcp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupProbeTcpWithDefaults

`func NewContainerGroupProbeTcpWithDefaults() *ContainerGroupProbeTcp`

NewContainerGroupProbeTcpWithDefaults instantiates a new ContainerGroupProbeTcp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *ContainerGroupProbeTcp) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ContainerGroupProbeTcp) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ContainerGroupProbeTcp) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


