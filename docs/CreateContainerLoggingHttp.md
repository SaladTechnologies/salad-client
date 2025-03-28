# CreateContainerLoggingHttp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | The hostname or IP address of the HTTP logging endpoint | 
**Port** | **int32** | The port number of the HTTP logging endpoint (1-65535) | 
**User** | Pointer to **NullableString** | Optional username for HTTP authentication | [optional] 
**Password** | Pointer to **NullableString** | Optional password for HTTP authentication | [optional] 
**Path** | Pointer to **NullableString** | Optional URL path for the HTTP endpoint | [optional] 
**Format** | [**ContainerLoggingHttpFormat**](ContainerLoggingHttpFormat.md) |  | 
**Headers** | Pointer to [**[]ContainerLoggingHttpHeader**](ContainerLoggingHttpHeader.md) | Optional HTTP headers to include in log transmission requests | [optional] 
**Compression** | [**ContainerLoggingHttpCompression**](ContainerLoggingHttpCompression.md) |  | 

## Methods

### NewCreateContainerLoggingHttp

`func NewCreateContainerLoggingHttp(host string, port int32, format ContainerLoggingHttpFormat, compression ContainerLoggingHttpCompression, ) *CreateContainerLoggingHttp`

NewCreateContainerLoggingHttp instantiates a new CreateContainerLoggingHttp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerLoggingHttpWithDefaults

`func NewCreateContainerLoggingHttpWithDefaults() *CreateContainerLoggingHttp`

NewCreateContainerLoggingHttpWithDefaults instantiates a new CreateContainerLoggingHttp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *CreateContainerLoggingHttp) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateContainerLoggingHttp) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateContainerLoggingHttp) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *CreateContainerLoggingHttp) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateContainerLoggingHttp) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateContainerLoggingHttp) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUser

`func (o *CreateContainerLoggingHttp) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateContainerLoggingHttp) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateContainerLoggingHttp) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CreateContainerLoggingHttp) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *CreateContainerLoggingHttp) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *CreateContainerLoggingHttp) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPassword

`func (o *CreateContainerLoggingHttp) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateContainerLoggingHttp) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateContainerLoggingHttp) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateContainerLoggingHttp) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CreateContainerLoggingHttp) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateContainerLoggingHttp) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPath

`func (o *CreateContainerLoggingHttp) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateContainerLoggingHttp) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateContainerLoggingHttp) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CreateContainerLoggingHttp) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *CreateContainerLoggingHttp) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *CreateContainerLoggingHttp) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetFormat

`func (o *CreateContainerLoggingHttp) GetFormat() ContainerLoggingHttpFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateContainerLoggingHttp) GetFormatOk() (*ContainerLoggingHttpFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateContainerLoggingHttp) SetFormat(v ContainerLoggingHttpFormat)`

SetFormat sets Format field to given value.


### GetHeaders

`func (o *CreateContainerLoggingHttp) GetHeaders() []ContainerLoggingHttpHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CreateContainerLoggingHttp) GetHeadersOk() (*[]ContainerLoggingHttpHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CreateContainerLoggingHttp) SetHeaders(v []ContainerLoggingHttpHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CreateContainerLoggingHttp) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetCompression

`func (o *CreateContainerLoggingHttp) GetCompression() ContainerLoggingHttpCompression`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *CreateContainerLoggingHttp) GetCompressionOk() (*ContainerLoggingHttpCompression, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *CreateContainerLoggingHttp) SetCompression(v ContainerLoggingHttpCompression)`

SetCompression sets Compression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


