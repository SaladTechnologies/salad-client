# ContainerRegistryAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsEcr** | Pointer to [**ContainerRegistryAuthenticationAwsEcr**](ContainerRegistryAuthenticationAwsEcr.md) |  | [optional] 
**Basic** | Pointer to [**ContainerRegistryAuthenticationBasic**](ContainerRegistryAuthenticationBasic.md) |  | [optional] 
**DockerHub** | Pointer to [**ContainerRegistryAuthenticationDockerHub**](ContainerRegistryAuthenticationDockerHub.md) |  | [optional] 
**GcpGar** | Pointer to [**ContainerRegistryAuthenticationGcpGar**](ContainerRegistryAuthenticationGcpGar.md) |  | [optional] 
**GcpGcr** | Pointer to [**ContainerRegistryAuthenticationGcpGcr**](ContainerRegistryAuthenticationGcpGcr.md) |  | [optional] 

## Methods

### NewContainerRegistryAuthentication

`func NewContainerRegistryAuthentication() *ContainerRegistryAuthentication`

NewContainerRegistryAuthentication instantiates a new ContainerRegistryAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRegistryAuthenticationWithDefaults

`func NewContainerRegistryAuthenticationWithDefaults() *ContainerRegistryAuthentication`

NewContainerRegistryAuthenticationWithDefaults instantiates a new ContainerRegistryAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsEcr

`func (o *ContainerRegistryAuthentication) GetAwsEcr() ContainerRegistryAuthenticationAwsEcr`

GetAwsEcr returns the AwsEcr field if non-nil, zero value otherwise.

### GetAwsEcrOk

`func (o *ContainerRegistryAuthentication) GetAwsEcrOk() (*ContainerRegistryAuthenticationAwsEcr, bool)`

GetAwsEcrOk returns a tuple with the AwsEcr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEcr

`func (o *ContainerRegistryAuthentication) SetAwsEcr(v ContainerRegistryAuthenticationAwsEcr)`

SetAwsEcr sets AwsEcr field to given value.

### HasAwsEcr

`func (o *ContainerRegistryAuthentication) HasAwsEcr() bool`

HasAwsEcr returns a boolean if a field has been set.

### GetBasic

`func (o *ContainerRegistryAuthentication) GetBasic() ContainerRegistryAuthenticationBasic`

GetBasic returns the Basic field if non-nil, zero value otherwise.

### GetBasicOk

`func (o *ContainerRegistryAuthentication) GetBasicOk() (*ContainerRegistryAuthenticationBasic, bool)`

GetBasicOk returns a tuple with the Basic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasic

`func (o *ContainerRegistryAuthentication) SetBasic(v ContainerRegistryAuthenticationBasic)`

SetBasic sets Basic field to given value.

### HasBasic

`func (o *ContainerRegistryAuthentication) HasBasic() bool`

HasBasic returns a boolean if a field has been set.

### GetDockerHub

`func (o *ContainerRegistryAuthentication) GetDockerHub() ContainerRegistryAuthenticationDockerHub`

GetDockerHub returns the DockerHub field if non-nil, zero value otherwise.

### GetDockerHubOk

`func (o *ContainerRegistryAuthentication) GetDockerHubOk() (*ContainerRegistryAuthenticationDockerHub, bool)`

GetDockerHubOk returns a tuple with the DockerHub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerHub

`func (o *ContainerRegistryAuthentication) SetDockerHub(v ContainerRegistryAuthenticationDockerHub)`

SetDockerHub sets DockerHub field to given value.

### HasDockerHub

`func (o *ContainerRegistryAuthentication) HasDockerHub() bool`

HasDockerHub returns a boolean if a field has been set.

### GetGcpGar

`func (o *ContainerRegistryAuthentication) GetGcpGar() ContainerRegistryAuthenticationGcpGar`

GetGcpGar returns the GcpGar field if non-nil, zero value otherwise.

### GetGcpGarOk

`func (o *ContainerRegistryAuthentication) GetGcpGarOk() (*ContainerRegistryAuthenticationGcpGar, bool)`

GetGcpGarOk returns a tuple with the GcpGar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpGar

`func (o *ContainerRegistryAuthentication) SetGcpGar(v ContainerRegistryAuthenticationGcpGar)`

SetGcpGar sets GcpGar field to given value.

### HasGcpGar

`func (o *ContainerRegistryAuthentication) HasGcpGar() bool`

HasGcpGar returns a boolean if a field has been set.

### GetGcpGcr

`func (o *ContainerRegistryAuthentication) GetGcpGcr() ContainerRegistryAuthenticationGcpGcr`

GetGcpGcr returns the GcpGcr field if non-nil, zero value otherwise.

### GetGcpGcrOk

`func (o *ContainerRegistryAuthentication) GetGcpGcrOk() (*ContainerRegistryAuthenticationGcpGcr, bool)`

GetGcpGcrOk returns a tuple with the GcpGcr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpGcr

`func (o *ContainerRegistryAuthentication) SetGcpGcr(v ContainerRegistryAuthenticationGcpGcr)`

SetGcpGcr sets GcpGcr field to given value.

### HasGcpGcr

`func (o *ContainerRegistryAuthentication) HasGcpGcr() bool`

HasGcpGcr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


