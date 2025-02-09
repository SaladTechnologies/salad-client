# ContainerLoggingHttp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**Port** | **int32** |  | 
**User** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Format** | **string** |  | 
**Headers** | Pointer to [**[]HttpHeadersInner**](HttpHeadersInner.md) |  | [optional] 
**Compression** | **string** |  | 

## Methods

### NewContainerLoggingHttp

`func NewContainerLoggingHttp(host string, port int32, format string, compression string, ) *ContainerLoggingHttp`

NewContainerLoggingHttp instantiates a new ContainerLoggingHttp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerLoggingHttpWithDefaults

`func NewContainerLoggingHttpWithDefaults() *ContainerLoggingHttp`

NewContainerLoggingHttpWithDefaults instantiates a new ContainerLoggingHttp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ContainerLoggingHttp) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ContainerLoggingHttp) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ContainerLoggingHttp) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *ContainerLoggingHttp) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ContainerLoggingHttp) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ContainerLoggingHttp) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUser

`func (o *ContainerLoggingHttp) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ContainerLoggingHttp) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ContainerLoggingHttp) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ContainerLoggingHttp) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *ContainerLoggingHttp) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *ContainerLoggingHttp) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPassword

`func (o *ContainerLoggingHttp) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ContainerLoggingHttp) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ContainerLoggingHttp) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ContainerLoggingHttp) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ContainerLoggingHttp) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ContainerLoggingHttp) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPath

`func (o *ContainerLoggingHttp) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ContainerLoggingHttp) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ContainerLoggingHttp) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ContainerLoggingHttp) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ContainerLoggingHttp) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ContainerLoggingHttp) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetFormat

`func (o *ContainerLoggingHttp) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ContainerLoggingHttp) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ContainerLoggingHttp) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetHeaders

`func (o *ContainerLoggingHttp) GetHeaders() []HttpHeadersInner`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ContainerLoggingHttp) GetHeadersOk() (*[]HttpHeadersInner, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ContainerLoggingHttp) SetHeaders(v []HttpHeadersInner)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ContainerLoggingHttp) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *ContainerLoggingHttp) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *ContainerLoggingHttp) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetCompression

`func (o *ContainerLoggingHttp) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *ContainerLoggingHttp) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *ContainerLoggingHttp) SetCompression(v string)`

SetCompression sets Compression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


