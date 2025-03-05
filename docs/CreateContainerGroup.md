# CreateContainerGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Container** | [**CreateContainer**](CreateContainer.md) |  | 
**AutostartPolicy** | **bool** |  | 
**RestartPolicy** | [**ContainerRestartPolicy**](ContainerRestartPolicy.md) |  | 
**Replicas** | **int32** |  | 
**CountryCodes** | Pointer to [**[]CountryCode**](CountryCode.md) | List of countries nodes must be located in. Remove this field to permit nodes from any country. | [optional] 
**Networking** | Pointer to [**CreateContainerGroupNetworking**](CreateContainerGroupNetworking.md) |  | [optional] 
**LivenessProbe** | Pointer to [**ContainerGroupLivenessProbe**](ContainerGroupLivenessProbe.md) |  | [optional] 
**ReadinessProbe** | Pointer to [**ContainerGroupReadinessProbe**](ContainerGroupReadinessProbe.md) |  | [optional] 
**StartupProbe** | Pointer to [**ContainerGroupStartupProbe**](ContainerGroupStartupProbe.md) |  | [optional] 
**QueueConnection** | Pointer to [**ContainerGroupQueueConnection**](ContainerGroupQueueConnection.md) |  | [optional] 
**QueueAutoscaler** | Pointer to [**QueueAutoscaler**](QueueAutoscaler.md) |  | [optional] 

## Methods

### NewCreateContainerGroup

`func NewCreateContainerGroup(name string, container CreateContainer, autostartPolicy bool, restartPolicy ContainerRestartPolicy, replicas int32, ) *CreateContainerGroup`

NewCreateContainerGroup instantiates a new CreateContainerGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerGroupWithDefaults

`func NewCreateContainerGroupWithDefaults() *CreateContainerGroup`

NewCreateContainerGroupWithDefaults instantiates a new CreateContainerGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateContainerGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateContainerGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateContainerGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *CreateContainerGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateContainerGroup) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateContainerGroup) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateContainerGroup) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CreateContainerGroup) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CreateContainerGroup) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetContainer

`func (o *CreateContainerGroup) GetContainer() CreateContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *CreateContainerGroup) GetContainerOk() (*CreateContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *CreateContainerGroup) SetContainer(v CreateContainer)`

SetContainer sets Container field to given value.


### GetAutostartPolicy

`func (o *CreateContainerGroup) GetAutostartPolicy() bool`

GetAutostartPolicy returns the AutostartPolicy field if non-nil, zero value otherwise.

### GetAutostartPolicyOk

`func (o *CreateContainerGroup) GetAutostartPolicyOk() (*bool, bool)`

GetAutostartPolicyOk returns a tuple with the AutostartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostartPolicy

`func (o *CreateContainerGroup) SetAutostartPolicy(v bool)`

SetAutostartPolicy sets AutostartPolicy field to given value.


### GetRestartPolicy

`func (o *CreateContainerGroup) GetRestartPolicy() ContainerRestartPolicy`

GetRestartPolicy returns the RestartPolicy field if non-nil, zero value otherwise.

### GetRestartPolicyOk

`func (o *CreateContainerGroup) GetRestartPolicyOk() (*ContainerRestartPolicy, bool)`

GetRestartPolicyOk returns a tuple with the RestartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartPolicy

`func (o *CreateContainerGroup) SetRestartPolicy(v ContainerRestartPolicy)`

SetRestartPolicy sets RestartPolicy field to given value.


### GetReplicas

`func (o *CreateContainerGroup) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *CreateContainerGroup) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *CreateContainerGroup) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetCountryCodes

`func (o *CreateContainerGroup) GetCountryCodes() []CountryCode`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *CreateContainerGroup) GetCountryCodesOk() (*[]CountryCode, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *CreateContainerGroup) SetCountryCodes(v []CountryCode)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *CreateContainerGroup) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetNetworking

`func (o *CreateContainerGroup) GetNetworking() CreateContainerGroupNetworking`

GetNetworking returns the Networking field if non-nil, zero value otherwise.

### GetNetworkingOk

`func (o *CreateContainerGroup) GetNetworkingOk() (*CreateContainerGroupNetworking, bool)`

GetNetworkingOk returns a tuple with the Networking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworking

`func (o *CreateContainerGroup) SetNetworking(v CreateContainerGroupNetworking)`

SetNetworking sets Networking field to given value.

### HasNetworking

`func (o *CreateContainerGroup) HasNetworking() bool`

HasNetworking returns a boolean if a field has been set.

### GetLivenessProbe

`func (o *CreateContainerGroup) GetLivenessProbe() ContainerGroupLivenessProbe`

GetLivenessProbe returns the LivenessProbe field if non-nil, zero value otherwise.

### GetLivenessProbeOk

`func (o *CreateContainerGroup) GetLivenessProbeOk() (*ContainerGroupLivenessProbe, bool)`

GetLivenessProbeOk returns a tuple with the LivenessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbe

`func (o *CreateContainerGroup) SetLivenessProbe(v ContainerGroupLivenessProbe)`

SetLivenessProbe sets LivenessProbe field to given value.

### HasLivenessProbe

`func (o *CreateContainerGroup) HasLivenessProbe() bool`

HasLivenessProbe returns a boolean if a field has been set.

### GetReadinessProbe

`func (o *CreateContainerGroup) GetReadinessProbe() ContainerGroupReadinessProbe`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *CreateContainerGroup) GetReadinessProbeOk() (*ContainerGroupReadinessProbe, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *CreateContainerGroup) SetReadinessProbe(v ContainerGroupReadinessProbe)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *CreateContainerGroup) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### GetStartupProbe

`func (o *CreateContainerGroup) GetStartupProbe() ContainerGroupStartupProbe`

GetStartupProbe returns the StartupProbe field if non-nil, zero value otherwise.

### GetStartupProbeOk

`func (o *CreateContainerGroup) GetStartupProbeOk() (*ContainerGroupStartupProbe, bool)`

GetStartupProbeOk returns a tuple with the StartupProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupProbe

`func (o *CreateContainerGroup) SetStartupProbe(v ContainerGroupStartupProbe)`

SetStartupProbe sets StartupProbe field to given value.

### HasStartupProbe

`func (o *CreateContainerGroup) HasStartupProbe() bool`

HasStartupProbe returns a boolean if a field has been set.

### GetQueueConnection

`func (o *CreateContainerGroup) GetQueueConnection() ContainerGroupQueueConnection`

GetQueueConnection returns the QueueConnection field if non-nil, zero value otherwise.

### GetQueueConnectionOk

`func (o *CreateContainerGroup) GetQueueConnectionOk() (*ContainerGroupQueueConnection, bool)`

GetQueueConnectionOk returns a tuple with the QueueConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueConnection

`func (o *CreateContainerGroup) SetQueueConnection(v ContainerGroupQueueConnection)`

SetQueueConnection sets QueueConnection field to given value.

### HasQueueConnection

`func (o *CreateContainerGroup) HasQueueConnection() bool`

HasQueueConnection returns a boolean if a field has been set.

### GetQueueAutoscaler

`func (o *CreateContainerGroup) GetQueueAutoscaler() QueueAutoscaler`

GetQueueAutoscaler returns the QueueAutoscaler field if non-nil, zero value otherwise.

### GetQueueAutoscalerOk

`func (o *CreateContainerGroup) GetQueueAutoscalerOk() (*QueueAutoscaler, bool)`

GetQueueAutoscalerOk returns a tuple with the QueueAutoscaler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueAutoscaler

`func (o *CreateContainerGroup) SetQueueAutoscaler(v QueueAutoscaler)`

SetQueueAutoscaler sets QueueAutoscaler field to given value.

### HasQueueAutoscaler

`func (o *CreateContainerGroup) HasQueueAutoscaler() bool`

HasQueueAutoscaler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


