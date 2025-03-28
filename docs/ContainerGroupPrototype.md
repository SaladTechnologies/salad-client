# ContainerGroupPrototype

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutostartPolicy** | **bool** | Determines whether the container group should start automatically when created (true) or remain stopped until manually started (false) | 
**Container** | [**CreateContainer**](CreateContainer.md) |  | 
**CountryCodes** | Pointer to [**[]CountryCode**](CountryCode.md) | List of countries nodes must be located in. Remove this field to permit nodes from any country. | [optional] 
**DisplayName** | Pointer to **string** | Human-readable name for the container group that can include spaces and special characters, used for display purposes | [optional] 
**LivenessProbe** | Pointer to [**ContainerGroupLivenessProbe**](ContainerGroupLivenessProbe.md) |  | [optional] 
**Name** | **string** | Unique identifier for the container group that must follow DNS naming conventions (lowercase alphanumeric with hyphens) | 
**Networking** | Pointer to [**CreateContainerGroupNetworking**](CreateContainerGroupNetworking.md) |  | [optional] 
**QueueAutoscaler** | Pointer to [**ContainerGroupQueueAutoscaler**](ContainerGroupQueueAutoscaler.md) |  | [optional] 
**QueueConnection** | Pointer to [**ContainerGroupQueueConnection**](ContainerGroupQueueConnection.md) |  | [optional] 
**ReadinessProbe** | Pointer to [**ContainerGroupReadinessProbe**](ContainerGroupReadinessProbe.md) |  | [optional] 
**Replicas** | **int32** | Number of container instances to deploy and maintain for this container group | 
**RestartPolicy** | [**ContainerRestartPolicy**](ContainerRestartPolicy.md) |  | 
**StartupProbe** | Pointer to [**ContainerGroupStartupProbe**](ContainerGroupStartupProbe.md) |  | [optional] 

## Methods

### NewContainerGroupPrototype

`func NewContainerGroupPrototype(autostartPolicy bool, container CreateContainer, name string, replicas int32, restartPolicy ContainerRestartPolicy, ) *ContainerGroupPrototype`

NewContainerGroupPrototype instantiates a new ContainerGroupPrototype object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupPrototypeWithDefaults

`func NewContainerGroupPrototypeWithDefaults() *ContainerGroupPrototype`

NewContainerGroupPrototypeWithDefaults instantiates a new ContainerGroupPrototype object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutostartPolicy

`func (o *ContainerGroupPrototype) GetAutostartPolicy() bool`

GetAutostartPolicy returns the AutostartPolicy field if non-nil, zero value otherwise.

### GetAutostartPolicyOk

`func (o *ContainerGroupPrototype) GetAutostartPolicyOk() (*bool, bool)`

GetAutostartPolicyOk returns a tuple with the AutostartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostartPolicy

`func (o *ContainerGroupPrototype) SetAutostartPolicy(v bool)`

SetAutostartPolicy sets AutostartPolicy field to given value.


### GetContainer

`func (o *ContainerGroupPrototype) GetContainer() CreateContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ContainerGroupPrototype) GetContainerOk() (*CreateContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ContainerGroupPrototype) SetContainer(v CreateContainer)`

SetContainer sets Container field to given value.


### GetCountryCodes

`func (o *ContainerGroupPrototype) GetCountryCodes() []CountryCode`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *ContainerGroupPrototype) GetCountryCodesOk() (*[]CountryCode, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *ContainerGroupPrototype) SetCountryCodes(v []CountryCode)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *ContainerGroupPrototype) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetDisplayName

`func (o *ContainerGroupPrototype) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ContainerGroupPrototype) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ContainerGroupPrototype) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ContainerGroupPrototype) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetLivenessProbe

`func (o *ContainerGroupPrototype) GetLivenessProbe() ContainerGroupLivenessProbe`

GetLivenessProbe returns the LivenessProbe field if non-nil, zero value otherwise.

### GetLivenessProbeOk

`func (o *ContainerGroupPrototype) GetLivenessProbeOk() (*ContainerGroupLivenessProbe, bool)`

GetLivenessProbeOk returns a tuple with the LivenessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbe

`func (o *ContainerGroupPrototype) SetLivenessProbe(v ContainerGroupLivenessProbe)`

SetLivenessProbe sets LivenessProbe field to given value.

### HasLivenessProbe

`func (o *ContainerGroupPrototype) HasLivenessProbe() bool`

HasLivenessProbe returns a boolean if a field has been set.

### GetName

`func (o *ContainerGroupPrototype) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerGroupPrototype) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerGroupPrototype) SetName(v string)`

SetName sets Name field to given value.


### GetNetworking

`func (o *ContainerGroupPrototype) GetNetworking() CreateContainerGroupNetworking`

GetNetworking returns the Networking field if non-nil, zero value otherwise.

### GetNetworkingOk

`func (o *ContainerGroupPrototype) GetNetworkingOk() (*CreateContainerGroupNetworking, bool)`

GetNetworkingOk returns a tuple with the Networking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworking

`func (o *ContainerGroupPrototype) SetNetworking(v CreateContainerGroupNetworking)`

SetNetworking sets Networking field to given value.

### HasNetworking

`func (o *ContainerGroupPrototype) HasNetworking() bool`

HasNetworking returns a boolean if a field has been set.

### GetQueueAutoscaler

`func (o *ContainerGroupPrototype) GetQueueAutoscaler() ContainerGroupQueueAutoscaler`

GetQueueAutoscaler returns the QueueAutoscaler field if non-nil, zero value otherwise.

### GetQueueAutoscalerOk

`func (o *ContainerGroupPrototype) GetQueueAutoscalerOk() (*ContainerGroupQueueAutoscaler, bool)`

GetQueueAutoscalerOk returns a tuple with the QueueAutoscaler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueAutoscaler

`func (o *ContainerGroupPrototype) SetQueueAutoscaler(v ContainerGroupQueueAutoscaler)`

SetQueueAutoscaler sets QueueAutoscaler field to given value.

### HasQueueAutoscaler

`func (o *ContainerGroupPrototype) HasQueueAutoscaler() bool`

HasQueueAutoscaler returns a boolean if a field has been set.

### GetQueueConnection

`func (o *ContainerGroupPrototype) GetQueueConnection() ContainerGroupQueueConnection`

GetQueueConnection returns the QueueConnection field if non-nil, zero value otherwise.

### GetQueueConnectionOk

`func (o *ContainerGroupPrototype) GetQueueConnectionOk() (*ContainerGroupQueueConnection, bool)`

GetQueueConnectionOk returns a tuple with the QueueConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueConnection

`func (o *ContainerGroupPrototype) SetQueueConnection(v ContainerGroupQueueConnection)`

SetQueueConnection sets QueueConnection field to given value.

### HasQueueConnection

`func (o *ContainerGroupPrototype) HasQueueConnection() bool`

HasQueueConnection returns a boolean if a field has been set.

### GetReadinessProbe

`func (o *ContainerGroupPrototype) GetReadinessProbe() ContainerGroupReadinessProbe`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *ContainerGroupPrototype) GetReadinessProbeOk() (*ContainerGroupReadinessProbe, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *ContainerGroupPrototype) SetReadinessProbe(v ContainerGroupReadinessProbe)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *ContainerGroupPrototype) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### GetReplicas

`func (o *ContainerGroupPrototype) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *ContainerGroupPrototype) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *ContainerGroupPrototype) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetRestartPolicy

`func (o *ContainerGroupPrototype) GetRestartPolicy() ContainerRestartPolicy`

GetRestartPolicy returns the RestartPolicy field if non-nil, zero value otherwise.

### GetRestartPolicyOk

`func (o *ContainerGroupPrototype) GetRestartPolicyOk() (*ContainerRestartPolicy, bool)`

GetRestartPolicyOk returns a tuple with the RestartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartPolicy

`func (o *ContainerGroupPrototype) SetRestartPolicy(v ContainerRestartPolicy)`

SetRestartPolicy sets RestartPolicy field to given value.


### GetStartupProbe

`func (o *ContainerGroupPrototype) GetStartupProbe() ContainerGroupStartupProbe`

GetStartupProbe returns the StartupProbe field if non-nil, zero value otherwise.

### GetStartupProbeOk

`func (o *ContainerGroupPrototype) GetStartupProbeOk() (*ContainerGroupStartupProbe, bool)`

GetStartupProbeOk returns a tuple with the StartupProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupProbe

`func (o *ContainerGroupPrototype) SetStartupProbe(v ContainerGroupStartupProbe)`

SetStartupProbe sets StartupProbe field to given value.

### HasStartupProbe

`func (o *ContainerGroupPrototype) HasStartupProbe() bool`

HasStartupProbe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


