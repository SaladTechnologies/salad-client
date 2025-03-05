/*
SaladCloud API

The SaladCloud REST API. Please refer to the [SaladCloud API Documentation](https://docs.salad.com/api-reference) for more details.

API version: 0.9.0-alpha.7
Contact: cloud@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package saladclient

import (
	"encoding/json"
)

// checks if the UpdateContainerGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateContainerGroup{}

// UpdateContainerGroup Represents a request to update a container group
type UpdateContainerGroup struct {
	DisplayName NullableString `json:"display_name,omitempty" validate:"regexp=^[ ,-.0-9A-Za-z]+$"`
	Container *UpdateContainer `json:"container,omitempty"`
	Replicas NullableInt32 `json:"replicas,omitempty"`
	// List of countries nodes must be located in. Remove this field to permit nodes from any country.
	CountryCodes []CountryCode `json:"country_codes,omitempty"`
	Networking *UpdateContainerGroupNetworking `json:"networking,omitempty"`
	LivenessProbe *ContainerGroupLivenessProbe `json:"liveness_probe,omitempty"`
	ReadinessProbe *ContainerGroupReadinessProbe `json:"readiness_probe,omitempty"`
	StartupProbe *ContainerGroupStartupProbe `json:"startup_probe,omitempty"`
	QueueAutoscaler *QueueAutoscaler `json:"queue_autoscaler,omitempty"`
}

// NewUpdateContainerGroup instantiates a new UpdateContainerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateContainerGroup() *UpdateContainerGroup {
	this := UpdateContainerGroup{}
	return &this
}

// NewUpdateContainerGroupWithDefaults instantiates a new UpdateContainerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateContainerGroupWithDefaults() *UpdateContainerGroup {
	this := UpdateContainerGroup{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateContainerGroup) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateContainerGroup) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UpdateContainerGroup) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *UpdateContainerGroup) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *UpdateContainerGroup) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *UpdateContainerGroup) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *UpdateContainerGroup) GetContainer() UpdateContainer {
	if o == nil || IsNil(o.Container) {
		var ret UpdateContainer
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContainerGroup) GetContainerOk() (*UpdateContainer, bool) {
	if o == nil || IsNil(o.Container) {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *UpdateContainerGroup) HasContainer() bool {
	if o != nil && !IsNil(o.Container) {
		return true
	}

	return false
}

// SetContainer gets a reference to the given UpdateContainer and assigns it to the Container field.
func (o *UpdateContainerGroup) SetContainer(v UpdateContainer) {
	o.Container = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateContainerGroup) GetReplicas() int32 {
	if o == nil || IsNil(o.Replicas.Get()) {
		var ret int32
		return ret
	}
	return *o.Replicas.Get()
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateContainerGroup) GetReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Replicas.Get(), o.Replicas.IsSet()
}

// HasReplicas returns a boolean if a field has been set.
func (o *UpdateContainerGroup) HasReplicas() bool {
	if o != nil && o.Replicas.IsSet() {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given NullableInt32 and assigns it to the Replicas field.
func (o *UpdateContainerGroup) SetReplicas(v int32) {
	o.Replicas.Set(&v)
}
// SetReplicasNil sets the value for Replicas to be an explicit nil
func (o *UpdateContainerGroup) SetReplicasNil() {
	o.Replicas.Set(nil)
}

// UnsetReplicas ensures that no value is present for Replicas, not even an explicit nil
func (o *UpdateContainerGroup) UnsetReplicas() {
	o.Replicas.Unset()
}

// GetCountryCodes returns the CountryCodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateContainerGroup) GetCountryCodes() []CountryCode {
	if o == nil {
		var ret []CountryCode
		return ret
	}
	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateContainerGroup) GetCountryCodesOk() ([]CountryCode, bool) {
	if o == nil || IsNil(o.CountryCodes) {
		return nil, false
	}
	return o.CountryCodes, true
}

// HasCountryCodes returns a boolean if a field has been set.
func (o *UpdateContainerGroup) HasCountryCodes() bool {
	if o != nil && !IsNil(o.CountryCodes) {
		return true
	}

	return false
}

// SetCountryCodes gets a reference to the given []CountryCode and assigns it to the CountryCodes field.
func (o *UpdateContainerGroup) SetCountryCodes(v []CountryCode) {
	o.CountryCodes = v
}

// GetNetworking returns the Networking field value if set, zero value otherwise.
func (o *UpdateContainerGroup) GetNetworking() UpdateContainerGroupNetworking {
	if o == nil || IsNil(o.Networking) {
		var ret UpdateContainerGroupNetworking
		return ret
	}
	return *o.Networking
}

// GetNetworkingOk returns a tuple with the Networking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContainerGroup) GetNetworkingOk() (*UpdateContainerGroupNetworking, bool) {
	if o == nil || IsNil(o.Networking) {
		return nil, false
	}
	return o.Networking, true
}

// HasNetworking returns a boolean if a field has been set.
func (o *UpdateContainerGroup) HasNetworking() bool {
	if o != nil && !IsNil(o.Networking) {
		return true
	}

	return false
}

// SetNetworking gets a reference to the given UpdateContainerGroupNetworking and assigns it to the Networking field.
func (o *UpdateContainerGroup) SetNetworking(v UpdateContainerGroupNetworking) {
	o.Networking = &v
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise.
func (o *UpdateContainerGroup) GetLivenessProbe() ContainerGroupLivenessProbe {
	if o == nil || IsNil(o.LivenessProbe) {
		var ret ContainerGroupLivenessProbe
		return ret
	}
	return *o.LivenessProbe
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContainerGroup) GetLivenessProbeOk() (*ContainerGroupLivenessProbe, bool) {
	if o == nil || IsNil(o.LivenessProbe) {
		return nil, false
	}
	return o.LivenessProbe, true
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *UpdateContainerGroup) HasLivenessProbe() bool {
	if o != nil && !IsNil(o.LivenessProbe) {
		return true
	}

	return false
}

// SetLivenessProbe gets a reference to the given ContainerGroupLivenessProbe and assigns it to the LivenessProbe field.
func (o *UpdateContainerGroup) SetLivenessProbe(v ContainerGroupLivenessProbe) {
	o.LivenessProbe = &v
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise.
func (o *UpdateContainerGroup) GetReadinessProbe() ContainerGroupReadinessProbe {
	if o == nil || IsNil(o.ReadinessProbe) {
		var ret ContainerGroupReadinessProbe
		return ret
	}
	return *o.ReadinessProbe
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContainerGroup) GetReadinessProbeOk() (*ContainerGroupReadinessProbe, bool) {
	if o == nil || IsNil(o.ReadinessProbe) {
		return nil, false
	}
	return o.ReadinessProbe, true
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *UpdateContainerGroup) HasReadinessProbe() bool {
	if o != nil && !IsNil(o.ReadinessProbe) {
		return true
	}

	return false
}

// SetReadinessProbe gets a reference to the given ContainerGroupReadinessProbe and assigns it to the ReadinessProbe field.
func (o *UpdateContainerGroup) SetReadinessProbe(v ContainerGroupReadinessProbe) {
	o.ReadinessProbe = &v
}

// GetStartupProbe returns the StartupProbe field value if set, zero value otherwise.
func (o *UpdateContainerGroup) GetStartupProbe() ContainerGroupStartupProbe {
	if o == nil || IsNil(o.StartupProbe) {
		var ret ContainerGroupStartupProbe
		return ret
	}
	return *o.StartupProbe
}

// GetStartupProbeOk returns a tuple with the StartupProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContainerGroup) GetStartupProbeOk() (*ContainerGroupStartupProbe, bool) {
	if o == nil || IsNil(o.StartupProbe) {
		return nil, false
	}
	return o.StartupProbe, true
}

// HasStartupProbe returns a boolean if a field has been set.
func (o *UpdateContainerGroup) HasStartupProbe() bool {
	if o != nil && !IsNil(o.StartupProbe) {
		return true
	}

	return false
}

// SetStartupProbe gets a reference to the given ContainerGroupStartupProbe and assigns it to the StartupProbe field.
func (o *UpdateContainerGroup) SetStartupProbe(v ContainerGroupStartupProbe) {
	o.StartupProbe = &v
}

// GetQueueAutoscaler returns the QueueAutoscaler field value if set, zero value otherwise.
func (o *UpdateContainerGroup) GetQueueAutoscaler() QueueAutoscaler {
	if o == nil || IsNil(o.QueueAutoscaler) {
		var ret QueueAutoscaler
		return ret
	}
	return *o.QueueAutoscaler
}

// GetQueueAutoscalerOk returns a tuple with the QueueAutoscaler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateContainerGroup) GetQueueAutoscalerOk() (*QueueAutoscaler, bool) {
	if o == nil || IsNil(o.QueueAutoscaler) {
		return nil, false
	}
	return o.QueueAutoscaler, true
}

// HasQueueAutoscaler returns a boolean if a field has been set.
func (o *UpdateContainerGroup) HasQueueAutoscaler() bool {
	if o != nil && !IsNil(o.QueueAutoscaler) {
		return true
	}

	return false
}

// SetQueueAutoscaler gets a reference to the given QueueAutoscaler and assigns it to the QueueAutoscaler field.
func (o *UpdateContainerGroup) SetQueueAutoscaler(v QueueAutoscaler) {
	o.QueueAutoscaler = &v
}

func (o UpdateContainerGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateContainerGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	if !IsNil(o.Container) {
		toSerialize["container"] = o.Container
	}
	if o.Replicas.IsSet() {
		toSerialize["replicas"] = o.Replicas.Get()
	}
	if o.CountryCodes != nil {
		toSerialize["country_codes"] = o.CountryCodes
	}
	if !IsNil(o.Networking) {
		toSerialize["networking"] = o.Networking
	}
	if !IsNil(o.LivenessProbe) {
		toSerialize["liveness_probe"] = o.LivenessProbe
	}
	if !IsNil(o.ReadinessProbe) {
		toSerialize["readiness_probe"] = o.ReadinessProbe
	}
	if !IsNil(o.StartupProbe) {
		toSerialize["startup_probe"] = o.StartupProbe
	}
	if !IsNil(o.QueueAutoscaler) {
		toSerialize["queue_autoscaler"] = o.QueueAutoscaler
	}
	return toSerialize, nil
}

type NullableUpdateContainerGroup struct {
	value *UpdateContainerGroup
	isSet bool
}

func (v NullableUpdateContainerGroup) Get() *UpdateContainerGroup {
	return v.value
}

func (v *NullableUpdateContainerGroup) Set(val *UpdateContainerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateContainerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateContainerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateContainerGroup(val *UpdateContainerGroup) *NullableUpdateContainerGroup {
	return &NullableUpdateContainerGroup{value: val, isSet: true}
}

func (v NullableUpdateContainerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateContainerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


