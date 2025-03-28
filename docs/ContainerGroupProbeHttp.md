# ContainerGroupProbeHttp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | [**[]ContainerGroupProbeHttpHeader**](ContainerGroupProbeHttpHeader.md) | A collection of HTTP header name-value pairs used for configuring requests and responses in container group endpoints. Each header consists of a name and its corresponding value. | 
**Path** | **string** | The HTTP path that will be probed to check container health. | 
**Port** | **int32** | The TCP port number to which the HTTP request will be sent. | 
**Scheme** | [**NullableContainerProbeHttpScheme**](ContainerProbeHttpScheme.md) |  | 

## Methods

### NewContainerGroupProbeHttp

`func NewContainerGroupProbeHttp(headers []ContainerGroupProbeHttpHeader, path string, port int32, scheme NullableContainerProbeHttpScheme, ) *ContainerGroupProbeHttp`

NewContainerGroupProbeHttp instantiates a new ContainerGroupProbeHttp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupProbeHttpWithDefaults

`func NewContainerGroupProbeHttpWithDefaults() *ContainerGroupProbeHttp`

NewContainerGroupProbeHttpWithDefaults instantiates a new ContainerGroupProbeHttp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *ContainerGroupProbeHttp) GetHeaders() []ContainerGroupProbeHttpHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ContainerGroupProbeHttp) GetHeadersOk() (*[]ContainerGroupProbeHttpHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ContainerGroupProbeHttp) SetHeaders(v []ContainerGroupProbeHttpHeader)`

SetHeaders sets Headers field to given value.


### GetPath

`func (o *ContainerGroupProbeHttp) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ContainerGroupProbeHttp) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ContainerGroupProbeHttp) SetPath(v string)`

SetPath sets Path field to given value.


### GetPort

`func (o *ContainerGroupProbeHttp) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ContainerGroupProbeHttp) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ContainerGroupProbeHttp) SetPort(v int32)`

SetPort sets Port field to given value.


### GetScheme

`func (o *ContainerGroupProbeHttp) GetScheme() ContainerProbeHttpScheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ContainerGroupProbeHttp) GetSchemeOk() (*ContainerProbeHttpScheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ContainerGroupProbeHttp) SetScheme(v ContainerProbeHttpScheme)`

SetScheme sets Scheme field to given value.


### SetSchemeNil

`func (o *ContainerGroupProbeHttp) SetSchemeNil(b bool)`

 SetSchemeNil sets the value for Scheme to be an explicit nil

### UnsetScheme
`func (o *ContainerGroupProbeHttp) UnsetScheme()`

UnsetScheme ensures that no value is present for Scheme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


