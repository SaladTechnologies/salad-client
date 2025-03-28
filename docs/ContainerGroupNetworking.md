# ContainerGroupNetworking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | **bool** | Whether authentication is required for network access to the container group | 
**ClientRequestTimeout** | Pointer to **int32** | The container group networking client request timeout. | [optional] [default to 100000]
**Dns** | **string** | Domain name or URL endpoint for the container group&#39;s network interface | 
**LoadBalancer** | [**ContainerGroupNetworkingLoadBalancer**](ContainerGroupNetworkingLoadBalancer.md) |  | [default to CONTAINERGROUPNETWORKINGLOADBALANCER_ROUND_ROBIN]
**Port** | **int32** | The container group networking port. | 
**Protocol** | [**ContainerNetworkingProtocol**](ContainerNetworkingProtocol.md) |  | 
**ServerResponseTimeout** | Pointer to **int32** | The container group networking server response timeout. | [optional] [default to 100000]
**SingleConnectionLimit** | Pointer to **bool** | The container group networking single connection limit flag. | [optional] [default to false]

## Methods

### NewContainerGroupNetworking

`func NewContainerGroupNetworking(auth bool, dns string, loadBalancer ContainerGroupNetworkingLoadBalancer, port int32, protocol ContainerNetworkingProtocol, ) *ContainerGroupNetworking`

NewContainerGroupNetworking instantiates a new ContainerGroupNetworking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupNetworkingWithDefaults

`func NewContainerGroupNetworkingWithDefaults() *ContainerGroupNetworking`

NewContainerGroupNetworkingWithDefaults instantiates a new ContainerGroupNetworking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *ContainerGroupNetworking) GetAuth() bool`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *ContainerGroupNetworking) GetAuthOk() (*bool, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *ContainerGroupNetworking) SetAuth(v bool)`

SetAuth sets Auth field to given value.


### GetClientRequestTimeout

`func (o *ContainerGroupNetworking) GetClientRequestTimeout() int32`

GetClientRequestTimeout returns the ClientRequestTimeout field if non-nil, zero value otherwise.

### GetClientRequestTimeoutOk

`func (o *ContainerGroupNetworking) GetClientRequestTimeoutOk() (*int32, bool)`

GetClientRequestTimeoutOk returns a tuple with the ClientRequestTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRequestTimeout

`func (o *ContainerGroupNetworking) SetClientRequestTimeout(v int32)`

SetClientRequestTimeout sets ClientRequestTimeout field to given value.

### HasClientRequestTimeout

`func (o *ContainerGroupNetworking) HasClientRequestTimeout() bool`

HasClientRequestTimeout returns a boolean if a field has been set.

### GetDns

`func (o *ContainerGroupNetworking) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *ContainerGroupNetworking) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *ContainerGroupNetworking) SetDns(v string)`

SetDns sets Dns field to given value.


### GetLoadBalancer

`func (o *ContainerGroupNetworking) GetLoadBalancer() ContainerGroupNetworkingLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *ContainerGroupNetworking) GetLoadBalancerOk() (*ContainerGroupNetworkingLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *ContainerGroupNetworking) SetLoadBalancer(v ContainerGroupNetworkingLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.


### GetPort

`func (o *ContainerGroupNetworking) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ContainerGroupNetworking) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ContainerGroupNetworking) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *ContainerGroupNetworking) GetProtocol() ContainerNetworkingProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ContainerGroupNetworking) GetProtocolOk() (*ContainerNetworkingProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ContainerGroupNetworking) SetProtocol(v ContainerNetworkingProtocol)`

SetProtocol sets Protocol field to given value.


### GetServerResponseTimeout

`func (o *ContainerGroupNetworking) GetServerResponseTimeout() int32`

GetServerResponseTimeout returns the ServerResponseTimeout field if non-nil, zero value otherwise.

### GetServerResponseTimeoutOk

`func (o *ContainerGroupNetworking) GetServerResponseTimeoutOk() (*int32, bool)`

GetServerResponseTimeoutOk returns a tuple with the ServerResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerResponseTimeout

`func (o *ContainerGroupNetworking) SetServerResponseTimeout(v int32)`

SetServerResponseTimeout sets ServerResponseTimeout field to given value.

### HasServerResponseTimeout

`func (o *ContainerGroupNetworking) HasServerResponseTimeout() bool`

HasServerResponseTimeout returns a boolean if a field has been set.

### GetSingleConnectionLimit

`func (o *ContainerGroupNetworking) GetSingleConnectionLimit() bool`

GetSingleConnectionLimit returns the SingleConnectionLimit field if non-nil, zero value otherwise.

### GetSingleConnectionLimitOk

`func (o *ContainerGroupNetworking) GetSingleConnectionLimitOk() (*bool, bool)`

GetSingleConnectionLimitOk returns a tuple with the SingleConnectionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleConnectionLimit

`func (o *ContainerGroupNetworking) SetSingleConnectionLimit(v bool)`

SetSingleConnectionLimit sets SingleConnectionLimit field to given value.

### HasSingleConnectionLimit

`func (o *ContainerGroupNetworking) HasSingleConnectionLimit() bool`

HasSingleConnectionLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


