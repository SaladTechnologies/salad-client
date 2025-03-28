# CreateContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **[]string** | Pass a command (and optional arguments) to override the ENTRYPOINT and CMD of a container image. Each element in the array represents a command segment or argument. | [optional] 
**EnvironmentVariables** | Pointer to **map[string]string** | Key-value pairs of environment variables to set within the container. These variables will be available to processes running inside the container. | [optional] 
**Image** | **string** | The container image. | 
**ImageCaching** | Pointer to **bool** | The container image caching. | [optional] 
**Logging** | Pointer to [**CreateContainerLogging**](CreateContainerLogging.md) |  | [optional] 
**Priority** | Pointer to [**NullableContainerGroupPriority**](ContainerGroupPriority.md) |  | [optional] 
**RegistryAuthentication** | Pointer to [**ContainerRegistryAuthentication**](ContainerRegistryAuthentication.md) |  | [optional] 
**Resources** | [**ContainerResourceRequirements**](ContainerResourceRequirements.md) |  | 

## Methods

### NewCreateContainer

`func NewCreateContainer(image string, resources ContainerResourceRequirements, ) *CreateContainer`

NewCreateContainer instantiates a new CreateContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerWithDefaults

`func NewCreateContainerWithDefaults() *CreateContainer`

NewCreateContainerWithDefaults instantiates a new CreateContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *CreateContainer) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CreateContainer) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CreateContainer) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CreateContainer) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *CreateContainer) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *CreateContainer) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetEnvironmentVariables

`func (o *CreateContainer) GetEnvironmentVariables() map[string]string`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *CreateContainer) GetEnvironmentVariablesOk() (*map[string]string, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *CreateContainer) SetEnvironmentVariables(v map[string]string)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *CreateContainer) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetImage

`func (o *CreateContainer) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CreateContainer) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CreateContainer) SetImage(v string)`

SetImage sets Image field to given value.


### GetImageCaching

`func (o *CreateContainer) GetImageCaching() bool`

GetImageCaching returns the ImageCaching field if non-nil, zero value otherwise.

### GetImageCachingOk

`func (o *CreateContainer) GetImageCachingOk() (*bool, bool)`

GetImageCachingOk returns a tuple with the ImageCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageCaching

`func (o *CreateContainer) SetImageCaching(v bool)`

SetImageCaching sets ImageCaching field to given value.

### HasImageCaching

`func (o *CreateContainer) HasImageCaching() bool`

HasImageCaching returns a boolean if a field has been set.

### GetLogging

`func (o *CreateContainer) GetLogging() CreateContainerLogging`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *CreateContainer) GetLoggingOk() (*CreateContainerLogging, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *CreateContainer) SetLogging(v CreateContainerLogging)`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *CreateContainer) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### GetPriority

`func (o *CreateContainer) GetPriority() ContainerGroupPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateContainer) GetPriorityOk() (*ContainerGroupPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateContainer) SetPriority(v ContainerGroupPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateContainer) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *CreateContainer) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *CreateContainer) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetRegistryAuthentication

`func (o *CreateContainer) GetRegistryAuthentication() ContainerRegistryAuthentication`

GetRegistryAuthentication returns the RegistryAuthentication field if non-nil, zero value otherwise.

### GetRegistryAuthenticationOk

`func (o *CreateContainer) GetRegistryAuthenticationOk() (*ContainerRegistryAuthentication, bool)`

GetRegistryAuthenticationOk returns a tuple with the RegistryAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryAuthentication

`func (o *CreateContainer) SetRegistryAuthentication(v ContainerRegistryAuthentication)`

SetRegistryAuthentication sets RegistryAuthentication field to given value.

### HasRegistryAuthentication

`func (o *CreateContainer) HasRegistryAuthentication() bool`

HasRegistryAuthentication returns a boolean if a field has been set.

### GetResources

`func (o *CreateContainer) GetResources() ContainerResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CreateContainer) GetResourcesOk() (*ContainerResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CreateContainer) SetResources(v ContainerResourceRequirements)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


