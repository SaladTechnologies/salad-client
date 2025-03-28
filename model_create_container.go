/*
SaladCloud API

The SaladCloud REST API. Please refer to the [SaladCloud API Documentation](https://docs.salad.com/api-reference) for more details.

API version: 0.9.0-alpha.11
Contact: cloud@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package saladclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContainer{}

// CreateContainer Configuration for creating a container within a container group. Defines the container image, resource requirements, environment variables, and other settings needed to deploy and run the container.
type CreateContainer struct {
	// Pass a command (and optional arguments) to override the ENTRYPOINT and CMD of a container image. Each element in the array represents a command segment or argument.
	Command []string `json:"command,omitempty"`
	// Key-value pairs of environment variables to set within the container. These variables will be available to processes running inside the container.
	EnvironmentVariables map[string]string `json:"environment_variables,omitempty"`
	// The container image.
	Image string `json:"image" validate:"regexp=^.*$"`
	// The container image caching.
	ImageCaching *bool `json:"image_caching,omitempty"`
	Logging *CreateContainerLogging `json:"logging,omitempty"`
	Priority NullableContainerGroupPriority `json:"priority,omitempty"`
	RegistryAuthentication *ContainerRegistryAuthentication `json:"registry_authentication,omitempty"`
	Resources ContainerResourceRequirements `json:"resources"`
}

type _CreateContainer CreateContainer

// NewCreateContainer instantiates a new CreateContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContainer(image string, resources ContainerResourceRequirements) *CreateContainer {
	this := CreateContainer{}
	this.Image = image
	this.Resources = resources
	return &this
}

// NewCreateContainerWithDefaults instantiates a new CreateContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContainerWithDefaults() *CreateContainer {
	this := CreateContainer{}
	return &this
}

// GetCommand returns the Command field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContainer) GetCommand() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContainer) GetCommandOk() ([]string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *CreateContainer) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given []string and assigns it to the Command field.
func (o *CreateContainer) SetCommand(v []string) {
	o.Command = v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise.
func (o *CreateContainer) GetEnvironmentVariables() map[string]string {
	if o == nil || IsNil(o.EnvironmentVariables) {
		var ret map[string]string
		return ret
	}
	return o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainer) GetEnvironmentVariablesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.EnvironmentVariables) {
		return map[string]string{}, false
	}
	return o.EnvironmentVariables, true
}

// HasEnvironmentVariables returns a boolean if a field has been set.
func (o *CreateContainer) HasEnvironmentVariables() bool {
	if o != nil && !IsNil(o.EnvironmentVariables) {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given map[string]string and assigns it to the EnvironmentVariables field.
func (o *CreateContainer) SetEnvironmentVariables(v map[string]string) {
	o.EnvironmentVariables = v
}

// GetImage returns the Image field value
func (o *CreateContainer) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *CreateContainer) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *CreateContainer) SetImage(v string) {
	o.Image = v
}

// GetImageCaching returns the ImageCaching field value if set, zero value otherwise.
func (o *CreateContainer) GetImageCaching() bool {
	if o == nil || IsNil(o.ImageCaching) {
		var ret bool
		return ret
	}
	return *o.ImageCaching
}

// GetImageCachingOk returns a tuple with the ImageCaching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainer) GetImageCachingOk() (*bool, bool) {
	if o == nil || IsNil(o.ImageCaching) {
		return nil, false
	}
	return o.ImageCaching, true
}

// HasImageCaching returns a boolean if a field has been set.
func (o *CreateContainer) HasImageCaching() bool {
	if o != nil && !IsNil(o.ImageCaching) {
		return true
	}

	return false
}

// SetImageCaching gets a reference to the given bool and assigns it to the ImageCaching field.
func (o *CreateContainer) SetImageCaching(v bool) {
	o.ImageCaching = &v
}

// GetLogging returns the Logging field value if set, zero value otherwise.
func (o *CreateContainer) GetLogging() CreateContainerLogging {
	if o == nil || IsNil(o.Logging) {
		var ret CreateContainerLogging
		return ret
	}
	return *o.Logging
}

// GetLoggingOk returns a tuple with the Logging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainer) GetLoggingOk() (*CreateContainerLogging, bool) {
	if o == nil || IsNil(o.Logging) {
		return nil, false
	}
	return o.Logging, true
}

// HasLogging returns a boolean if a field has been set.
func (o *CreateContainer) HasLogging() bool {
	if o != nil && !IsNil(o.Logging) {
		return true
	}

	return false
}

// SetLogging gets a reference to the given CreateContainerLogging and assigns it to the Logging field.
func (o *CreateContainer) SetLogging(v CreateContainerLogging) {
	o.Logging = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContainer) GetPriority() ContainerGroupPriority {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret ContainerGroupPriority
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContainer) GetPriorityOk() (*ContainerGroupPriority, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *CreateContainer) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableContainerGroupPriority and assigns it to the Priority field.
func (o *CreateContainer) SetPriority(v ContainerGroupPriority) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *CreateContainer) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *CreateContainer) UnsetPriority() {
	o.Priority.Unset()
}

// GetRegistryAuthentication returns the RegistryAuthentication field value if set, zero value otherwise.
func (o *CreateContainer) GetRegistryAuthentication() ContainerRegistryAuthentication {
	if o == nil || IsNil(o.RegistryAuthentication) {
		var ret ContainerRegistryAuthentication
		return ret
	}
	return *o.RegistryAuthentication
}

// GetRegistryAuthenticationOk returns a tuple with the RegistryAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContainer) GetRegistryAuthenticationOk() (*ContainerRegistryAuthentication, bool) {
	if o == nil || IsNil(o.RegistryAuthentication) {
		return nil, false
	}
	return o.RegistryAuthentication, true
}

// HasRegistryAuthentication returns a boolean if a field has been set.
func (o *CreateContainer) HasRegistryAuthentication() bool {
	if o != nil && !IsNil(o.RegistryAuthentication) {
		return true
	}

	return false
}

// SetRegistryAuthentication gets a reference to the given ContainerRegistryAuthentication and assigns it to the RegistryAuthentication field.
func (o *CreateContainer) SetRegistryAuthentication(v ContainerRegistryAuthentication) {
	o.RegistryAuthentication = &v
}

// GetResources returns the Resources field value
func (o *CreateContainer) GetResources() ContainerResourceRequirements {
	if o == nil {
		var ret ContainerResourceRequirements
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *CreateContainer) GetResourcesOk() (*ContainerResourceRequirements, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *CreateContainer) SetResources(v ContainerResourceRequirements) {
	o.Resources = v
}

func (o CreateContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if !IsNil(o.EnvironmentVariables) {
		toSerialize["environment_variables"] = o.EnvironmentVariables
	}
	toSerialize["image"] = o.Image
	if !IsNil(o.ImageCaching) {
		toSerialize["image_caching"] = o.ImageCaching
	}
	if !IsNil(o.Logging) {
		toSerialize["logging"] = o.Logging
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	if !IsNil(o.RegistryAuthentication) {
		toSerialize["registry_authentication"] = o.RegistryAuthentication
	}
	toSerialize["resources"] = o.Resources
	return toSerialize, nil
}

func (o *CreateContainer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"image",
		"resources",
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

	varCreateContainer := _CreateContainer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateContainer)

	if err != nil {
		return err
	}

	*o = CreateContainer(varCreateContainer)

	return err
}

type NullableCreateContainer struct {
	value *CreateContainer
	isSet bool
}

func (v NullableCreateContainer) Get() *CreateContainer {
	return v.value
}

func (v *NullableCreateContainer) Set(val *CreateContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContainer(val *CreateContainer) *NullableCreateContainer {
	return &NullableCreateContainer{value: val, isSet: true}
}

func (v NullableCreateContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


