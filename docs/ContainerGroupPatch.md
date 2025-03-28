# ContainerGroupPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The display name for the container group. If null is provided, the display name will be set to the container group name. | [optional] 
**Container** | Pointer to [**UpdateContainer**](UpdateContainer.md) |  | [optional] 
**Replicas** | Pointer to **NullableInt32** | The desired number of instances for your container group deployment. | [optional] 
**CountryCodes** | Pointer to [**[]CountryCode**](CountryCode.md) | List of countries nodes must be located in. Remove this field to permit nodes from any country. | [optional] 
**Networking** | Pointer to [**UpdateContainerGroupNetworking**](UpdateContainerGroupNetworking.md) |  | [optional] 
**LivenessProbe** | Pointer to [**ContainerGroupLivenessProbe**](ContainerGroupLivenessProbe.md) |  | [optional] 
**ReadinessProbe** | Pointer to [**ContainerGroupReadinessProbe**](ContainerGroupReadinessProbe.md) |  | [optional] 
**StartupProbe** | Pointer to [**ContainerGroupStartupProbe**](ContainerGroupStartupProbe.md) |  | [optional] 
**QueueAutoscaler** | Pointer to [**ContainerGroupQueueAutoscaler**](ContainerGroupQueueAutoscaler.md) |  | [optional] 

## Methods

### NewContainerGroupPatch

`func NewContainerGroupPatch() *ContainerGroupPatch`

NewContainerGroupPatch instantiates a new ContainerGroupPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupPatchWithDefaults

`func NewContainerGroupPatchWithDefaults() *ContainerGroupPatch`

NewContainerGroupPatchWithDefaults instantiates a new ContainerGroupPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ContainerGroupPatch) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ContainerGroupPatch) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ContainerGroupPatch) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ContainerGroupPatch) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ContainerGroupPatch) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ContainerGroupPatch) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetContainer

`func (o *ContainerGroupPatch) GetContainer() UpdateContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ContainerGroupPatch) GetContainerOk() (*UpdateContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ContainerGroupPatch) SetContainer(v UpdateContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ContainerGroupPatch) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetReplicas

`func (o *ContainerGroupPatch) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *ContainerGroupPatch) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *ContainerGroupPatch) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *ContainerGroupPatch) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### SetReplicasNil

`func (o *ContainerGroupPatch) SetReplicasNil(b bool)`

 SetReplicasNil sets the value for Replicas to be an explicit nil

### UnsetReplicas
`func (o *ContainerGroupPatch) UnsetReplicas()`

UnsetReplicas ensures that no value is present for Replicas, not even an explicit nil
### GetCountryCodes

`func (o *ContainerGroupPatch) GetCountryCodes() []CountryCode`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *ContainerGroupPatch) GetCountryCodesOk() (*[]CountryCode, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *ContainerGroupPatch) SetCountryCodes(v []CountryCode)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *ContainerGroupPatch) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### SetCountryCodesNil

`func (o *ContainerGroupPatch) SetCountryCodesNil(b bool)`

 SetCountryCodesNil sets the value for CountryCodes to be an explicit nil

### UnsetCountryCodes
`func (o *ContainerGroupPatch) UnsetCountryCodes()`

UnsetCountryCodes ensures that no value is present for CountryCodes, not even an explicit nil
### GetNetworking

`func (o *ContainerGroupPatch) GetNetworking() UpdateContainerGroupNetworking`

GetNetworking returns the Networking field if non-nil, zero value otherwise.

### GetNetworkingOk

`func (o *ContainerGroupPatch) GetNetworkingOk() (*UpdateContainerGroupNetworking, bool)`

GetNetworkingOk returns a tuple with the Networking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworking

`func (o *ContainerGroupPatch) SetNetworking(v UpdateContainerGroupNetworking)`

SetNetworking sets Networking field to given value.

### HasNetworking

`func (o *ContainerGroupPatch) HasNetworking() bool`

HasNetworking returns a boolean if a field has been set.

### GetLivenessProbe

`func (o *ContainerGroupPatch) GetLivenessProbe() ContainerGroupLivenessProbe`

GetLivenessProbe returns the LivenessProbe field if non-nil, zero value otherwise.

### GetLivenessProbeOk

`func (o *ContainerGroupPatch) GetLivenessProbeOk() (*ContainerGroupLivenessProbe, bool)`

GetLivenessProbeOk returns a tuple with the LivenessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbe

`func (o *ContainerGroupPatch) SetLivenessProbe(v ContainerGroupLivenessProbe)`

SetLivenessProbe sets LivenessProbe field to given value.

### HasLivenessProbe

`func (o *ContainerGroupPatch) HasLivenessProbe() bool`

HasLivenessProbe returns a boolean if a field has been set.

### GetReadinessProbe

`func (o *ContainerGroupPatch) GetReadinessProbe() ContainerGroupReadinessProbe`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *ContainerGroupPatch) GetReadinessProbeOk() (*ContainerGroupReadinessProbe, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *ContainerGroupPatch) SetReadinessProbe(v ContainerGroupReadinessProbe)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *ContainerGroupPatch) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### GetStartupProbe

`func (o *ContainerGroupPatch) GetStartupProbe() ContainerGroupStartupProbe`

GetStartupProbe returns the StartupProbe field if non-nil, zero value otherwise.

### GetStartupProbeOk

`func (o *ContainerGroupPatch) GetStartupProbeOk() (*ContainerGroupStartupProbe, bool)`

GetStartupProbeOk returns a tuple with the StartupProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupProbe

`func (o *ContainerGroupPatch) SetStartupProbe(v ContainerGroupStartupProbe)`

SetStartupProbe sets StartupProbe field to given value.

### HasStartupProbe

`func (o *ContainerGroupPatch) HasStartupProbe() bool`

HasStartupProbe returns a boolean if a field has been set.

### GetQueueAutoscaler

`func (o *ContainerGroupPatch) GetQueueAutoscaler() ContainerGroupQueueAutoscaler`

GetQueueAutoscaler returns the QueueAutoscaler field if non-nil, zero value otherwise.

### GetQueueAutoscalerOk

`func (o *ContainerGroupPatch) GetQueueAutoscalerOk() (*ContainerGroupQueueAutoscaler, bool)`

GetQueueAutoscalerOk returns a tuple with the QueueAutoscaler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueAutoscaler

`func (o *ContainerGroupPatch) SetQueueAutoscaler(v ContainerGroupQueueAutoscaler)`

SetQueueAutoscaler sets QueueAutoscaler field to given value.

### HasQueueAutoscaler

`func (o *ContainerGroupPatch) HasQueueAutoscaler() bool`

HasQueueAutoscaler returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


