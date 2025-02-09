# CreateContainerGroupNetworking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | [**ContainerNetworkingProtocol**](ContainerNetworkingProtocol.md) |  | 
**Port** | **int32** |  | 
**Auth** | **bool** |  | 
**LoadBalancer** | Pointer to **string** |  | [optional] [default to "round_robin"]
**SingleConnectionLimit** | Pointer to **bool** |  | [optional] [default to false]
**ClientRequestTimeout** | Pointer to **int32** |  | [optional] [default to 100000]
**ServerResponseTimeout** | Pointer to **int32** |  | [optional] [default to 100000]

## Methods

### NewCreateContainerGroupNetworking

`func NewCreateContainerGroupNetworking(protocol ContainerNetworkingProtocol, port int32, auth bool, ) *CreateContainerGroupNetworking`

NewCreateContainerGroupNetworking instantiates a new CreateContainerGroupNetworking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerGroupNetworkingWithDefaults

`func NewCreateContainerGroupNetworkingWithDefaults() *CreateContainerGroupNetworking`

NewCreateContainerGroupNetworkingWithDefaults instantiates a new CreateContainerGroupNetworking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *CreateContainerGroupNetworking) GetProtocol() ContainerNetworkingProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateContainerGroupNetworking) GetProtocolOk() (*ContainerNetworkingProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateContainerGroupNetworking) SetProtocol(v ContainerNetworkingProtocol)`

SetProtocol sets Protocol field to given value.


### GetPort

`func (o *CreateContainerGroupNetworking) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateContainerGroupNetworking) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateContainerGroupNetworking) SetPort(v int32)`

SetPort sets Port field to given value.


### GetAuth

`func (o *CreateContainerGroupNetworking) GetAuth() bool`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *CreateContainerGroupNetworking) GetAuthOk() (*bool, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *CreateContainerGroupNetworking) SetAuth(v bool)`

SetAuth sets Auth field to given value.


### GetLoadBalancer

`func (o *CreateContainerGroupNetworking) GetLoadBalancer() string`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *CreateContainerGroupNetworking) GetLoadBalancerOk() (*string, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *CreateContainerGroupNetworking) SetLoadBalancer(v string)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *CreateContainerGroupNetworking) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetSingleConnectionLimit

`func (o *CreateContainerGroupNetworking) GetSingleConnectionLimit() bool`

GetSingleConnectionLimit returns the SingleConnectionLimit field if non-nil, zero value otherwise.

### GetSingleConnectionLimitOk

`func (o *CreateContainerGroupNetworking) GetSingleConnectionLimitOk() (*bool, bool)`

GetSingleConnectionLimitOk returns a tuple with the SingleConnectionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleConnectionLimit

`func (o *CreateContainerGroupNetworking) SetSingleConnectionLimit(v bool)`

SetSingleConnectionLimit sets SingleConnectionLimit field to given value.

### HasSingleConnectionLimit

`func (o *CreateContainerGroupNetworking) HasSingleConnectionLimit() bool`

HasSingleConnectionLimit returns a boolean if a field has been set.

### GetClientRequestTimeout

`func (o *CreateContainerGroupNetworking) GetClientRequestTimeout() int32`

GetClientRequestTimeout returns the ClientRequestTimeout field if non-nil, zero value otherwise.

### GetClientRequestTimeoutOk

`func (o *CreateContainerGroupNetworking) GetClientRequestTimeoutOk() (*int32, bool)`

GetClientRequestTimeoutOk returns a tuple with the ClientRequestTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRequestTimeout

`func (o *CreateContainerGroupNetworking) SetClientRequestTimeout(v int32)`

SetClientRequestTimeout sets ClientRequestTimeout field to given value.

### HasClientRequestTimeout

`func (o *CreateContainerGroupNetworking) HasClientRequestTimeout() bool`

HasClientRequestTimeout returns a boolean if a field has been set.

### GetServerResponseTimeout

`func (o *CreateContainerGroupNetworking) GetServerResponseTimeout() int32`

GetServerResponseTimeout returns the ServerResponseTimeout field if non-nil, zero value otherwise.

### GetServerResponseTimeoutOk

`func (o *CreateContainerGroupNetworking) GetServerResponseTimeoutOk() (*int32, bool)`

GetServerResponseTimeoutOk returns a tuple with the ServerResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerResponseTimeout

`func (o *CreateContainerGroupNetworking) SetServerResponseTimeout(v int32)`

SetServerResponseTimeout sets ServerResponseTimeout field to given value.

### HasServerResponseTimeout

`func (o *CreateContainerGroupNetworking) HasServerResponseTimeout() bool`

HasServerResponseTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


