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

// DeleteRecipeDeploymentProblemType the model 'DeleteRecipeDeploymentProblemType'
type DeleteRecipeDeploymentProblemType string


// All allowed values of DeleteRecipeDeploymentProblemType enum
var AllowedDeleteRecipeDeploymentProblemTypeEnumValues = []DeleteRecipeDeploymentProblemType{
	"null",
}

func (v *DeleteRecipeDeploymentProblemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeleteRecipeDeploymentProblemType(value)
	for _, existing := range AllowedDeleteRecipeDeploymentProblemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeleteRecipeDeploymentProblemType", value)
}

// NewDeleteRecipeDeploymentProblemTypeFromValue returns a pointer to a valid DeleteRecipeDeploymentProblemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeleteRecipeDeploymentProblemTypeFromValue(v string) (*DeleteRecipeDeploymentProblemType, error) {
	ev := DeleteRecipeDeploymentProblemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeleteRecipeDeploymentProblemType: valid values are %v", v, AllowedDeleteRecipeDeploymentProblemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeleteRecipeDeploymentProblemType) IsValid() bool {
	for _, existing := range AllowedDeleteRecipeDeploymentProblemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeleteRecipeDeploymentProblemType value
func (v DeleteRecipeDeploymentProblemType) Ptr() *DeleteRecipeDeploymentProblemType {
	return &v
}

type NullableDeleteRecipeDeploymentProblemType struct {
	value *DeleteRecipeDeploymentProblemType
	isSet bool
}

func (v NullableDeleteRecipeDeploymentProblemType) Get() *DeleteRecipeDeploymentProblemType {
	return v.value
}

func (v *NullableDeleteRecipeDeploymentProblemType) Set(val *DeleteRecipeDeploymentProblemType) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRecipeDeploymentProblemType) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRecipeDeploymentProblemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRecipeDeploymentProblemType(val *DeleteRecipeDeploymentProblemType) *NullableDeleteRecipeDeploymentProblemType {
	return &NullableDeleteRecipeDeploymentProblemType{value: val, isSet: true}
}

func (v NullableDeleteRecipeDeploymentProblemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRecipeDeploymentProblemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

