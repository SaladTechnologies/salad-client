# UpdateContainerGroupNetworking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **NullableInt32** | The port number to expose on the container group | [optional] 

## Methods

### NewUpdateContainerGroupNetworking

`func NewUpdateContainerGroupNetworking() *UpdateContainerGroupNetworking`

NewUpdateContainerGroupNetworking instantiates a new UpdateContainerGroupNetworking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContainerGroupNetworkingWithDefaults

`func NewUpdateContainerGroupNetworkingWithDefaults() *UpdateContainerGroupNetworking`

NewUpdateContainerGroupNetworkingWithDefaults instantiates a new UpdateContainerGroupNetworking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *UpdateContainerGroupNetworking) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateContainerGroupNetworking) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateContainerGroupNetworking) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateContainerGroupNetworking) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *UpdateContainerGroupNetworking) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *UpdateContainerGroupNetworking) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


