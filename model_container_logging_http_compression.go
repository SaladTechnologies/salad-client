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
	"fmt"
)

// ContainerLoggingHttpCompression The compression algorithm to apply to logs before transmission
type ContainerLoggingHttpCompression string

// List of ContainerLoggingHttpCompression
const (
	CONTAINERLOGGINGHTTPCOMPRESSION_NONE ContainerLoggingHttpCompression = "none"
	CONTAINERLOGGINGHTTPCOMPRESSION_GZIP ContainerLoggingHttpCompression = "gzip"
)

// All allowed values of ContainerLoggingHttpCompression enum
var AllowedContainerLoggingHttpCompressionEnumValues = []ContainerLoggingHttpCompression{
	"none",
	"gzip",
}

func (v *ContainerLoggingHttpCompression) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContainerLoggingHttpCompression(value)
	for _, existing := range AllowedContainerLoggingHttpCompressionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContainerLoggingHttpCompression", value)
}

// NewContainerLoggingHttpCompressionFromValue returns a pointer to a valid ContainerLoggingHttpCompression
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContainerLoggingHttpCompressionFromValue(v string) (*ContainerLoggingHttpCompression, error) {
	ev := ContainerLoggingHttpCompression(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContainerLoggingHttpCompression: valid values are %v", v, AllowedContainerLoggingHttpCompressionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContainerLoggingHttpCompression) IsValid() bool {
	for _, existing := range AllowedContainerLoggingHttpCompressionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContainerLoggingHttpCompression value
func (v ContainerLoggingHttpCompression) Ptr() *ContainerLoggingHttpCompression {
	return &v
}

type NullableContainerLoggingHttpCompression struct {
	value *ContainerLoggingHttpCompression
	isSet bool
}

func (v NullableContainerLoggingHttpCompression) Get() *ContainerLoggingHttpCompression {
	return v.value
}

func (v *NullableContainerLoggingHttpCompression) Set(val *ContainerLoggingHttpCompression) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerLoggingHttpCompression) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerLoggingHttpCompression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerLoggingHttpCompression(val *ContainerLoggingHttpCompression) *NullableContainerLoggingHttpCompression {
	return &NullableContainerLoggingHttpCompression{value: val, isSet: true}
}

func (v NullableContainerLoggingHttpCompression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerLoggingHttpCompression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

