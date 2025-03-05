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
	"bytes"
	"fmt"
)

// checks if the QueueJobList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueueJobList{}

// QueueJobList Represents a list of queue jobs
type QueueJobList struct {
	Items []QueueJob `json:"items"`
}

type _QueueJobList QueueJobList

// NewQueueJobList instantiates a new QueueJobList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueJobList(items []QueueJob) *QueueJobList {
	this := QueueJobList{}
	this.Items = items
	return &this
}

// NewQueueJobListWithDefaults instantiates a new QueueJobList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueJobListWithDefaults() *QueueJobList {
	this := QueueJobList{}
	return &this
}

// GetItems returns the Items field value
func (o *QueueJobList) GetItems() []QueueJob {
	if o == nil {
		var ret []QueueJob
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *QueueJobList) GetItemsOk() ([]QueueJob, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *QueueJobList) SetItems(v []QueueJob) {
	o.Items = v
}

func (o QueueJobList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueueJobList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *QueueJobList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
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

	varQueueJobList := _QueueJobList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueueJobList)

	if err != nil {
		return err
	}

	*o = QueueJobList(varQueueJobList)

	return err
}

type NullableQueueJobList struct {
	value *QueueJobList
	isSet bool
}

func (v NullableQueueJobList) Get() *QueueJobList {
	return v.value
}

func (v *NullableQueueJobList) Set(val *QueueJobList) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueJobList) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueJobList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueJobList(val *QueueJobList) *NullableQueueJobList {
	return &NullableQueueJobList{value: val, isSet: true}
}

func (v NullableQueueJobList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueJobList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


