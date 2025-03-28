# CreateContainerGroupNetworking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | **bool** | Determines whether authentication is required for network connections to the container group | 
**ClientRequestTimeout** | Pointer to **int32** | The container group networking client request timeout. | [optional] [default to 100000]
**LoadBalancer** | Pointer to [**ContainerGroupNetworkingLoadBalancer**](ContainerGroupNetworkingLoadBalancer.md) |  | [optional] [default to CONTAINERGROUPNETWORKINGLOADBALANCER_ROUND_ROBIN]
**Port** | **int32** | The container group networking port. | 
**Protocol** | [**ContainerNetworkingProtocol**](ContainerNetworkingProtocol.md) |  | 
**ServerResponseTimeout** | Pointer to **int32** | The container group networking server response timeout. | [optional] [default to 100000]
**SingleConnectionLimit** | Pointer to **bool** | The container group networking single connection limit flag. | [optional] [default to false]

## Methods

### NewCreateContainerGroupNetworking

`func NewCreateContainerGroupNetworking(auth bool, port int32, protocol ContainerNetworkingProtocol, ) *CreateContainerGroupNetworking`

NewCreateContainerGroupNetworking instantiates a new CreateContainerGroupNetworking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerGroupNetworkingWithDefaults

`func NewCreateContainerGroupNetworkingWithDefaults() *CreateContainerGroupNetworking`

NewCreateContainerGroupNetworkingWithDefaults instantiates a new CreateContainerGroupNetworking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetLoadBalancer

`func (o *CreateContainerGroupNetworking) GetLoadBalancer() ContainerGroupNetworkingLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *CreateContainerGroupNetworking) GetLoadBalancerOk() (*ContainerGroupNetworkingLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *CreateContainerGroupNetworking) SetLoadBalancer(v ContainerGroupNetworkingLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *CreateContainerGroupNetworking) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


