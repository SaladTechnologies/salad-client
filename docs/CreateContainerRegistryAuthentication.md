# CreateContainerRegistryAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basic** | Pointer to [**NullableCreateContainerRegistryAuthenticationBasic**](CreateContainerRegistryAuthenticationBasic.md) |  | [optional] 
**GcpGcr** | Pointer to [**NullableCreateContainerRegistryAuthenticationGcpGcr**](CreateContainerRegistryAuthenticationGcpGcr.md) |  | [optional] 
**AwsEcr** | Pointer to [**NullableCreateContainerRegistryAuthenticationAwsEcr**](CreateContainerRegistryAuthenticationAwsEcr.md) |  | [optional] 
**DockerHub** | Pointer to [**NullableCreateContainerRegistryAuthenticationDockerHub**](CreateContainerRegistryAuthenticationDockerHub.md) |  | [optional] 
**GcpGar** | Pointer to [**NullableCreateContainerRegistryAuthenticationGcpGcr**](CreateContainerRegistryAuthenticationGcpGcr.md) |  | [optional] 

## Methods

### NewCreateContainerRegistryAuthentication

`func NewCreateContainerRegistryAuthentication() *CreateContainerRegistryAuthentication`

NewCreateContainerRegistryAuthentication instantiates a new CreateContainerRegistryAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerRegistryAuthenticationWithDefaults

`func NewCreateContainerRegistryAuthenticationWithDefaults() *CreateContainerRegistryAuthentication`

NewCreateContainerRegistryAuthenticationWithDefaults instantiates a new CreateContainerRegistryAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasic

`func (o *CreateContainerRegistryAuthentication) GetBasic() CreateContainerRegistryAuthenticationBasic`

GetBasic returns the Basic field if non-nil, zero value otherwise.

### GetBasicOk

`func (o *CreateContainerRegistryAuthentication) GetBasicOk() (*CreateContainerRegistryAuthenticationBasic, bool)`

GetBasicOk returns a tuple with the Basic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasic

`func (o *CreateContainerRegistryAuthentication) SetBasic(v CreateContainerRegistryAuthenticationBasic)`

SetBasic sets Basic field to given value.

### HasBasic

`func (o *CreateContainerRegistryAuthentication) HasBasic() bool`

HasBasic returns a boolean if a field has been set.

### SetBasicNil

`func (o *CreateContainerRegistryAuthentication) SetBasicNil(b bool)`

 SetBasicNil sets the value for Basic to be an explicit nil

### UnsetBasic
`func (o *CreateContainerRegistryAuthentication) UnsetBasic()`

UnsetBasic ensures that no value is present for Basic, not even an explicit nil
### GetGcpGcr

`func (o *CreateContainerRegistryAuthentication) GetGcpGcr() CreateContainerRegistryAuthenticationGcpGcr`

GetGcpGcr returns the GcpGcr field if non-nil, zero value otherwise.

### GetGcpGcrOk

`func (o *CreateContainerRegistryAuthentication) GetGcpGcrOk() (*CreateContainerRegistryAuthenticationGcpGcr, bool)`

GetGcpGcrOk returns a tuple with the GcpGcr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpGcr

`func (o *CreateContainerRegistryAuthentication) SetGcpGcr(v CreateContainerRegistryAuthenticationGcpGcr)`

SetGcpGcr sets GcpGcr field to given value.

### HasGcpGcr

`func (o *CreateContainerRegistryAuthentication) HasGcpGcr() bool`

HasGcpGcr returns a boolean if a field has been set.

### SetGcpGcrNil

`func (o *CreateContainerRegistryAuthentication) SetGcpGcrNil(b bool)`

 SetGcpGcrNil sets the value for GcpGcr to be an explicit nil

### UnsetGcpGcr
`func (o *CreateContainerRegistryAuthentication) UnsetGcpGcr()`

UnsetGcpGcr ensures that no value is present for GcpGcr, not even an explicit nil
### GetAwsEcr

`func (o *CreateContainerRegistryAuthentication) GetAwsEcr() CreateContainerRegistryAuthenticationAwsEcr`

GetAwsEcr returns the AwsEcr field if non-nil, zero value otherwise.

### GetAwsEcrOk

`func (o *CreateContainerRegistryAuthentication) GetAwsEcrOk() (*CreateContainerRegistryAuthenticationAwsEcr, bool)`

GetAwsEcrOk returns a tuple with the AwsEcr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEcr

`func (o *CreateContainerRegistryAuthentication) SetAwsEcr(v CreateContainerRegistryAuthenticationAwsEcr)`

SetAwsEcr sets AwsEcr field to given value.

### HasAwsEcr

`func (o *CreateContainerRegistryAuthentication) HasAwsEcr() bool`

HasAwsEcr returns a boolean if a field has been set.

### SetAwsEcrNil

`func (o *CreateContainerRegistryAuthentication) SetAwsEcrNil(b bool)`

 SetAwsEcrNil sets the value for AwsEcr to be an explicit nil

### UnsetAwsEcr
`func (o *CreateContainerRegistryAuthentication) UnsetAwsEcr()`

UnsetAwsEcr ensures that no value is present for AwsEcr, not even an explicit nil
### GetDockerHub

`func (o *CreateContainerRegistryAuthentication) GetDockerHub() CreateContainerRegistryAuthenticationDockerHub`

GetDockerHub returns the DockerHub field if non-nil, zero value otherwise.

### GetDockerHubOk

`func (o *CreateContainerRegistryAuthentication) GetDockerHubOk() (*CreateContainerRegistryAuthenticationDockerHub, bool)`

GetDockerHubOk returns a tuple with the DockerHub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerHub

`func (o *CreateContainerRegistryAuthentication) SetDockerHub(v CreateContainerRegistryAuthenticationDockerHub)`

SetDockerHub sets DockerHub field to given value.

### HasDockerHub

`func (o *CreateContainerRegistryAuthentication) HasDockerHub() bool`

HasDockerHub returns a boolean if a field has been set.

### SetDockerHubNil

`func (o *CreateContainerRegistryAuthentication) SetDockerHubNil(b bool)`

 SetDockerHubNil sets the value for DockerHub to be an explicit nil

### UnsetDockerHub
`func (o *CreateContainerRegistryAuthentication) UnsetDockerHub()`

UnsetDockerHub ensures that no value is present for DockerHub, not even an explicit nil
### GetGcpGar

`func (o *CreateContainerRegistryAuthentication) GetGcpGar() CreateContainerRegistryAuthenticationGcpGcr`

GetGcpGar returns the GcpGar field if non-nil, zero value otherwise.

### GetGcpGarOk

`func (o *CreateContainerRegistryAuthentication) GetGcpGarOk() (*CreateContainerRegistryAuthenticationGcpGcr, bool)`

GetGcpGarOk returns a tuple with the GcpGar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpGar

`func (o *CreateContainerRegistryAuthentication) SetGcpGar(v CreateContainerRegistryAuthenticationGcpGcr)`

SetGcpGar sets GcpGar field to given value.

### HasGcpGar

`func (o *CreateContainerRegistryAuthentication) HasGcpGar() bool`

HasGcpGar returns a boolean if a field has been set.

### SetGcpGarNil

`func (o *CreateContainerRegistryAuthentication) SetGcpGarNil(b bool)`

 SetGcpGarNil sets the value for GcpGar to be an explicit nil

### UnsetGcpGar
`func (o *CreateContainerRegistryAuthentication) UnsetGcpGar()`

UnsetGcpGar ensures that no value is present for GcpGar, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


