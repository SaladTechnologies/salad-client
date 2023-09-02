/*
SaladCloud Public API

The SaladCloud Public API.

API version: 1.0.0-alpha.56
Contact: support@salad.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// StopContainerGroupProblemType the model 'StopContainerGroupProblemType'
type StopContainerGroupProblemType string

// List of StopContainerGroupProblemType
const (
	NULL StopContainerGroupProblemType = "null"
)

// All allowed values of StopContainerGroupProblemType enum
var AllowedStopContainerGroupProblemTypeEnumValues = []StopContainerGroupProblemType{
	"null",
}

func (v *StopContainerGroupProblemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StopContainerGroupProblemType(value)
	for _, existing := range AllowedStopContainerGroupProblemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StopContainerGroupProblemType", value)
}

// NewStopContainerGroupProblemTypeFromValue returns a pointer to a valid StopContainerGroupProblemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStopContainerGroupProblemTypeFromValue(v string) (*StopContainerGroupProblemType, error) {
	ev := StopContainerGroupProblemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StopContainerGroupProblemType: valid values are %v", v, AllowedStopContainerGroupProblemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StopContainerGroupProblemType) IsValid() bool {
	for _, existing := range AllowedStopContainerGroupProblemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StopContainerGroupProblemType value
func (v StopContainerGroupProblemType) Ptr() *StopContainerGroupProblemType {
	return &v
}

type NullableStopContainerGroupProblemType struct {
	value *StopContainerGroupProblemType
	isSet bool
}

func (v NullableStopContainerGroupProblemType) Get() *StopContainerGroupProblemType {
	return v.value
}

func (v *NullableStopContainerGroupProblemType) Set(val *StopContainerGroupProblemType) {
	v.value = val
	v.isSet = true
}

func (v NullableStopContainerGroupProblemType) IsSet() bool {
	return v.isSet
}

func (v *NullableStopContainerGroupProblemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopContainerGroupProblemType(val *StopContainerGroupProblemType) *NullableStopContainerGroupProblemType {
	return &NullableStopContainerGroupProblemType{value: val, isSet: true}
}

func (v NullableStopContainerGroupProblemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopContainerGroupProblemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

