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
	"time"
	"bytes"
	"fmt"
)

// checks if the ContainerGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerGroup{}

// ContainerGroup Represents a container group
type ContainerGroup struct {
	Id string `json:"id"`
	Name string `json:"name" validate:"regexp=^[a-z][a-z0-9-]{0,61}[a-z0-9]$"`
	DisplayName string `json:"display_name" validate:"regexp=^[ ,-.0-9A-Za-z]+$"`
	Container Container `json:"container"`
	AutostartPolicy bool `json:"autostart_policy"`
	RestartPolicy ContainerRestartPolicy `json:"restart_policy"`
	Replicas int32 `json:"replicas"`
	CurrentState ContainerGroupState `json:"current_state"`
	// List of countries nodes must be located in. Remove this field to permit nodes from any country.
	CountryCodes []CountryCode `json:"country_codes,omitempty"`
	Networking *ContainerGroupNetworking `json:"networking,omitempty"`
	LivenessProbe *ContainerGroupLivenessProbe `json:"liveness_probe,omitempty"`
	ReadinessProbe *ContainerGroupReadinessProbe `json:"readiness_probe,omitempty"`
	StartupProbe *ContainerGroupStartupProbe `json:"startup_probe,omitempty"`
	QueueConnection *ContainerGroupQueueConnection `json:"queue_connection,omitempty"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	PendingChange bool `json:"pending_change"`
	Version int32 `json:"version"`
	QueueAutoscaler *QueueAutoscaler `json:"queue_autoscaler,omitempty"`
}

type _ContainerGroup ContainerGroup

// NewContainerGroup instantiates a new ContainerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerGroup(id string, name string, displayName string, container Container, autostartPolicy bool, restartPolicy ContainerRestartPolicy, replicas int32, currentState ContainerGroupState, createTime time.Time, updateTime time.Time, pendingChange bool, version int32) *ContainerGroup {
	this := ContainerGroup{}
	this.Id = id
	this.Name = name
	this.DisplayName = displayName
	this.Container = container
	this.AutostartPolicy = autostartPolicy
	this.RestartPolicy = restartPolicy
	this.Replicas = replicas
	this.CurrentState = currentState
	this.CreateTime = createTime
	this.UpdateTime = updateTime
	this.PendingChange = pendingChange
	this.Version = version
	return &this
}

// NewContainerGroupWithDefaults instantiates a new ContainerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerGroupWithDefaults() *ContainerGroup {
	this := ContainerGroup{}
	return &this
}

// GetId returns the Id field value
func (o *ContainerGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ContainerGroup) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ContainerGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerGroup) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *ContainerGroup) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ContainerGroup) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetContainer returns the Container field value
func (o *ContainerGroup) GetContainer() Container {
	if o == nil {
		var ret Container
		return ret
	}

	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetContainerOk() (*Container, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value
func (o *ContainerGroup) SetContainer(v Container) {
	o.Container = v
}

// GetAutostartPolicy returns the AutostartPolicy field value
func (o *ContainerGroup) GetAutostartPolicy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutostartPolicy
}

// GetAutostartPolicyOk returns a tuple with the AutostartPolicy field value
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetAutostartPolicyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutostartPolicy, true
}

// SetAutostartPolicy sets field value
func (o *ContainerGroup) SetAutostartPolicy(v bool) {
	o.AutostartPolicy = v
}

// GetRestartPolicy returns the RestartPolicy field value
func (o *ContainerGroup) GetRestartPolicy() ContainerRestartPolicy {
	if o == nil {
		var ret ContainerRestartPolicy
		return ret
	}

	return o.RestartPolicy
}

// GetRestartPolicyOk returns a tuple with the RestartPolicy field value
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetRestartPolicyOk() (*ContainerRestartPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestartPolicy, true
}

// SetRestartPolicy sets field value
func (o *ContainerGroup) SetRestartPolicy(v ContainerRestartPolicy) {
	o.RestartPolicy = v
}

// GetReplicas returns the Replicas field value
func (o *ContainerGroup) GetReplicas() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replicas, true
}

// SetReplicas sets field value
func (o *ContainerGroup) SetReplicas(v int32) {
	o.Replicas = v
}

// GetCurrentState returns the CurrentState field value
func (o *ContainerGroup) GetCurrentState() ContainerGroupState {
	if o == nil {
		var ret ContainerGroupState
		return ret
	}

	return o.CurrentState
}

// GetCurrentStateOk returns a tuple with the CurrentState field value
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetCurrentStateOk() (*ContainerGroupState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentState, true
}

// SetCurrentState sets field value
func (o *ContainerGroup) SetCurrentState(v ContainerGroupState) {
	o.CurrentState = v
}

// GetCountryCodes returns the CountryCodes field value if set, zero value otherwise.
func (o *ContainerGroup) GetCountryCodes() []CountryCode {
	if o == nil || IsNil(o.CountryCodes) {
		var ret []CountryCode
		return ret
	}
	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetCountryCodesOk() ([]CountryCode, bool) {
	if o == nil || IsNil(o.CountryCodes) {
		return nil, false
	}
	return o.CountryCodes, true
}

// HasCountryCodes returns a boolean if a field has been set.
func (o *ContainerGroup) HasCountryCodes() bool {
	if o != nil && !IsNil(o.CountryCodes) {
		return true
	}

	return false
}

// SetCountryCodes gets a reference to the given []CountryCode and assigns it to the CountryCodes field.
func (o *ContainerGroup) SetCountryCodes(v []CountryCode) {
	o.CountryCodes = v
}

// GetNetworking returns the Networking field value if set, zero value otherwise.
func (o *ContainerGroup) GetNetworking() ContainerGroupNetworking {
	if o == nil || IsNil(o.Networking) {
		var ret ContainerGroupNetworking
		return ret
	}
	return *o.Networking
}

// GetNetworkingOk returns a tuple with the Networking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetNetworkingOk() (*ContainerGroupNetworking, bool) {
	if o == nil || IsNil(o.Networking) {
		return nil, false
	}
	return o.Networking, true
}

// HasNetworking returns a boolean if a field has been set.
func (o *ContainerGroup) HasNetworking() bool {
	if o != nil && !IsNil(o.Networking) {
		return true
	}

	return false
}

// SetNetworking gets a reference to the given ContainerGroupNetworking and assigns it to the Networking field.
func (o *ContainerGroup) SetNetworking(v ContainerGroupNetworking) {
	o.Networking = &v
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise.
func (o *ContainerGroup) GetLivenessProbe() ContainerGroupLivenessProbe {
	if o == nil || IsNil(o.LivenessProbe) {
		var ret ContainerGroupLivenessProbe
		return ret
	}
	return *o.LivenessProbe
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetLivenessProbeOk() (*ContainerGroupLivenessProbe, bool) {
	if o == nil || IsNil(o.LivenessProbe) {
		return nil, false
	}
	return o.LivenessProbe, true
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *ContainerGroup) HasLivenessProbe() bool {
	if o != nil && !IsNil(o.LivenessProbe) {
		return true
	}

	return false
}

// SetLivenessProbe gets a reference to the given ContainerGroupLivenessProbe and assigns it to the LivenessProbe field.
func (o *ContainerGroup) SetLivenessProbe(v ContainerGroupLivenessProbe) {
	o.LivenessProbe = &v
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise.
func (o *ContainerGroup) GetReadinessProbe() ContainerGroupReadinessProbe {
	if o == nil || IsNil(o.ReadinessProbe) {
		var ret ContainerGroupReadinessProbe
		return ret
	}
	return *o.ReadinessProbe
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetReadinessProbeOk() (*ContainerGroupReadinessProbe, bool) {
	if o == nil || IsNil(o.ReadinessProbe) {
		return nil, false
	}
	return o.ReadinessProbe, true
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *ContainerGroup) HasReadinessProbe() bool {
	if o != nil && !IsNil(o.ReadinessProbe) {
		return true
	}

	return false
}

// SetReadinessProbe gets a reference to the given ContainerGroupReadinessProbe and assigns it to the ReadinessProbe field.
func (o *ContainerGroup) SetReadinessProbe(v ContainerGroupReadinessProbe) {
	o.ReadinessProbe = &v
}

// GetStartupProbe returns the StartupProbe field value if set, zero value otherwise.
func (o *ContainerGroup) GetStartupProbe() ContainerGroupStartupProbe {
	if o == nil || IsNil(o.StartupProbe) {
		var ret ContainerGroupStartupProbe
		return ret
	}
	return *o.StartupProbe
}

// GetStartupProbeOk returns a tuple with the StartupProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetStartupProbeOk() (*ContainerGroupStartupProbe, bool) {
	if o == nil || IsNil(o.StartupProbe) {
		return nil, false
	}
	return o.StartupProbe, true
}

// HasStartupProbe returns a boolean if a field has been set.
func (o *ContainerGroup) HasStartupProbe() bool {
	if o != nil && !IsNil(o.StartupProbe) {
		return true
	}

	return false
}

// SetStartupProbe gets a reference to the given ContainerGroupStartupProbe and assigns it to the StartupProbe field.
func (o *ContainerGroup) SetStartupProbe(v ContainerGroupStartupProbe) {
	o.StartupProbe = &v
}

// GetQueueConnection returns the QueueConnection field value if set, zero value otherwise.
func (o *ContainerGroup) GetQueueConnection() ContainerGroupQueueConnection {
	if o == nil || IsNil(o.QueueConnection) {
		var ret ContainerGroupQueueConnection
		return ret
	}
	return *o.QueueConnection
}

// GetQueueConnectionOk returns a tuple with the QueueConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetQueueConnectionOk() (*ContainerGroupQueueConnection, bool) {
	if o == nil || IsNil(o.QueueConnection) {
		return nil, false
	}
	return o.QueueConnection, true
}

// HasQueueConnection returns a boolean if a field has been set.
func (o *ContainerGroup) HasQueueConnection() bool {
	if o != nil && !IsNil(o.QueueConnection) {
		return true
	}

	return false
}

// SetQueueConnection gets a reference to the given ContainerGroupQueueConnection and assigns it to the QueueConnection field.
func (o *ContainerGroup) SetQueueConnection(v ContainerGroupQueueConnection) {
	o.QueueConnection = &v
}

// GetCreateTime returns the CreateTime field value
func (o *ContainerGroup) GetCreateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateTime, true
}

// SetCreateTime sets field value
func (o *ContainerGroup) SetCreateTime(v time.Time) {
	o.CreateTime = v
}

// GetUpdateTime returns the UpdateTime field value
func (o *ContainerGroup) GetUpdateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateTime, true
}

// SetUpdateTime sets field value
func (o *ContainerGroup) SetUpdateTime(v time.Time) {
	o.UpdateTime = v
}

// GetPendingChange returns the PendingChange field value
func (o *ContainerGroup) GetPendingChange() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PendingChange
}

// GetPendingChangeOk returns a tuple with the PendingChange field value
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetPendingChangeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PendingChange, true
}

// SetPendingChange sets field value
func (o *ContainerGroup) SetPendingChange(v bool) {
	o.PendingChange = v
}

// GetVersion returns the Version field value
func (o *ContainerGroup) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ContainerGroup) SetVersion(v int32) {
	o.Version = v
}

// GetQueueAutoscaler returns the QueueAutoscaler field value if set, zero value otherwise.
func (o *ContainerGroup) GetQueueAutoscaler() QueueAutoscaler {
	if o == nil || IsNil(o.QueueAutoscaler) {
		var ret QueueAutoscaler
		return ret
	}
	return *o.QueueAutoscaler
}

// GetQueueAutoscalerOk returns a tuple with the QueueAutoscaler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerGroup) GetQueueAutoscalerOk() (*QueueAutoscaler, bool) {
	if o == nil || IsNil(o.QueueAutoscaler) {
		return nil, false
	}
	return o.QueueAutoscaler, true
}

// HasQueueAutoscaler returns a boolean if a field has been set.
func (o *ContainerGroup) HasQueueAutoscaler() bool {
	if o != nil && !IsNil(o.QueueAutoscaler) {
		return true
	}

	return false
}

// SetQueueAutoscaler gets a reference to the given QueueAutoscaler and assigns it to the QueueAutoscaler field.
func (o *ContainerGroup) SetQueueAutoscaler(v QueueAutoscaler) {
	o.QueueAutoscaler = &v
}

func (o ContainerGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["display_name"] = o.DisplayName
	toSerialize["container"] = o.Container
	toSerialize["autostart_policy"] = o.AutostartPolicy
	toSerialize["restart_policy"] = o.RestartPolicy
	toSerialize["replicas"] = o.Replicas
	toSerialize["current_state"] = o.CurrentState
	if !IsNil(o.CountryCodes) {
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
	if !IsNil(o.QueueConnection) {
		toSerialize["queue_connection"] = o.QueueConnection
	}
	toSerialize["create_time"] = o.CreateTime
	toSerialize["update_time"] = o.UpdateTime
	toSerialize["pending_change"] = o.PendingChange
	toSerialize["version"] = o.Version
	if !IsNil(o.QueueAutoscaler) {
		toSerialize["queue_autoscaler"] = o.QueueAutoscaler
	}
	return toSerialize, nil
}

func (o *ContainerGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"display_name",
		"container",
		"autostart_policy",
		"restart_policy",
		"replicas",
		"current_state",
		"create_time",
		"update_time",
		"pending_change",
		"version",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varContainerGroup := _ContainerGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContainerGroup)

	if err != nil {
		return err
	}

	*o = ContainerGroup(varContainerGroup)

	return err
}

type NullableContainerGroup struct {
	value *ContainerGroup
	isSet bool
}

func (v NullableContainerGroup) Get() *ContainerGroup {
	return v.value
}

func (v *NullableContainerGroup) Set(val *ContainerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerGroup(val *ContainerGroup) *NullableContainerGroup {
	return &NullableContainerGroup{value: val, isSet: true}
}

func (v NullableContainerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


