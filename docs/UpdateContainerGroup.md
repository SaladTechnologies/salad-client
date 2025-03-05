# UpdateContainerGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Container** | Pointer to [**UpdateContainer**](UpdateContainer.md) |  | [optional] 
**Replicas** | Pointer to **NullableInt32** |  | [optional] 
**CountryCodes** | Pointer to [**[]CountryCode**](CountryCode.md) | List of countries nodes must be located in. Remove this field to permit nodes from any country. | [optional] 
**Networking** | Pointer to [**UpdateContainerGroupNetworking**](UpdateContainerGroupNetworking.md) |  | [optional] 
**LivenessProbe** | Pointer to [**ContainerGroupLivenessProbe**](ContainerGroupLivenessProbe.md) |  | [optional] 
**ReadinessProbe** | Pointer to [**ContainerGroupReadinessProbe**](ContainerGroupReadinessProbe.md) |  | [optional] 
**StartupProbe** | Pointer to [**ContainerGroupStartupProbe**](ContainerGroupStartupProbe.md) |  | [optional] 
**QueueAutoscaler** | Pointer to [**QueueAutoscaler**](QueueAutoscaler.md) |  | [optional] 

## Methods

### NewUpdateContainerGroup

`func NewUpdateContainerGroup() *UpdateContainerGroup`

NewUpdateContainerGroup instantiates a new UpdateContainerGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContainerGroupWithDefaults

`func NewUpdateContainerGroupWithDefaults() *UpdateContainerGroup`

NewUpdateContainerGroupWithDefaults instantiates a new UpdateContainerGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UpdateContainerGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateContainerGroup) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateContainerGroup) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateContainerGroup) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UpdateContainerGroup) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UpdateContainerGroup) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetContainer

`func (o *UpdateContainerGroup) GetContainer() UpdateContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *UpdateContainerGroup) GetContainerOk() (*UpdateContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *UpdateContainerGroup) SetContainer(v UpdateContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *UpdateContainerGroup) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetReplicas

`func (o *UpdateContainerGroup) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *UpdateContainerGroup) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *UpdateContainerGroup) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *UpdateContainerGroup) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### SetReplicasNil

`func (o *UpdateContainerGroup) SetReplicasNil(b bool)`

 SetReplicasNil sets the value for Replicas to be an explicit nil

### UnsetReplicas
`func (o *UpdateContainerGroup) UnsetReplicas()`

UnsetReplicas ensures that no value is present for Replicas, not even an explicit nil
### GetCountryCodes

`func (o *UpdateContainerGroup) GetCountryCodes() []CountryCode`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *UpdateContainerGroup) GetCountryCodesOk() (*[]CountryCode, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *UpdateContainerGroup) SetCountryCodes(v []CountryCode)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *UpdateContainerGroup) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### SetCountryCodesNil

`func (o *UpdateContainerGroup) SetCountryCodesNil(b bool)`

 SetCountryCodesNil sets the value for CountryCodes to be an explicit nil

### UnsetCountryCodes
`func (o *UpdateContainerGroup) UnsetCountryCodes()`

UnsetCountryCodes ensures that no value is present for CountryCodes, not even an explicit nil
### GetNetworking

`func (o *UpdateContainerGroup) GetNetworking() UpdateContainerGroupNetworking`

GetNetworking returns the Networking field if non-nil, zero value otherwise.

### GetNetworkingOk

`func (o *UpdateContainerGroup) GetNetworkingOk() (*UpdateContainerGroupNetworking, bool)`

GetNetworkingOk returns a tuple with the Networking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworking

`func (o *UpdateContainerGroup) SetNetworking(v UpdateContainerGroupNetworking)`

SetNetworking sets Networking field to given value.

### HasNetworking

`func (o *UpdateContainerGroup) HasNetworking() bool`

HasNetworking returns a boolean if a field has been set.

### GetLivenessProbe

`func (o *UpdateContainerGroup) GetLivenessProbe() ContainerGroupLivenessProbe`

GetLivenessProbe returns the LivenessProbe field if non-nil, zero value otherwise.

### GetLivenessProbeOk

`func (o *UpdateContainerGroup) GetLivenessProbeOk() (*ContainerGroupLivenessProbe, bool)`

GetLivenessProbeOk returns a tuple with the LivenessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbe

`func (o *UpdateContainerGroup) SetLivenessProbe(v ContainerGroupLivenessProbe)`

SetLivenessProbe sets LivenessProbe field to given value.

### HasLivenessProbe

`func (o *UpdateContainerGroup) HasLivenessProbe() bool`

HasLivenessProbe returns a boolean if a field has been set.

### GetReadinessProbe

`func (o *UpdateContainerGroup) GetReadinessProbe() ContainerGroupReadinessProbe`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *UpdateContainerGroup) GetReadinessProbeOk() (*ContainerGroupReadinessProbe, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *UpdateContainerGroup) SetReadinessProbe(v ContainerGroupReadinessProbe)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *UpdateContainerGroup) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### GetStartupProbe

`func (o *UpdateContainerGroup) GetStartupProbe() ContainerGroupStartupProbe`

GetStartupProbe returns the StartupProbe field if non-nil, zero value otherwise.

### GetStartupProbeOk

`func (o *UpdateContainerGroup) GetStartupProbeOk() (*ContainerGroupStartupProbe, bool)`

GetStartupProbeOk returns a tuple with the StartupProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupProbe

`func (o *UpdateContainerGroup) SetStartupProbe(v ContainerGroupStartupProbe)`

SetStartupProbe sets StartupProbe field to given value.

### HasStartupProbe

`func (o *UpdateContainerGroup) HasStartupProbe() bool`

HasStartupProbe returns a boolean if a field has been set.

### GetQueueAutoscaler

`func (o *UpdateContainerGroup) GetQueueAutoscaler() QueueAutoscaler`

GetQueueAutoscaler returns the QueueAutoscaler field if non-nil, zero value otherwise.

### GetQueueAutoscalerOk

`func (o *UpdateContainerGroup) GetQueueAutoscalerOk() (*QueueAutoscaler, bool)`

GetQueueAutoscalerOk returns a tuple with the QueueAutoscaler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueAutoscaler

`func (o *UpdateContainerGroup) SetQueueAutoscaler(v QueueAutoscaler)`

SetQueueAutoscaler sets QueueAutoscaler field to given value.

### HasQueueAutoscaler

`func (o *UpdateContainerGroup) HasQueueAutoscaler() bool`

HasQueueAutoscaler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


