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

// CreateRecipeDeploymentProblemType the model 'CreateRecipeDeploymentProblemType'
type CreateRecipeDeploymentProblemType string

// List of CreateRecipeDeploymentProblemType
const (
	BILLING_INFORMATION_REQUIRED CreateRecipeDeploymentProblemType = "billing_information_required"
	CREATED_QUOTA_EXCEEDED CreateRecipeDeploymentProblemType = "created_quota_exceeded"
	NAME_CONFLICT CreateRecipeDeploymentProblemType = "name_conflict"
	PROJECT_NOT_FOUND CreateRecipeDeploymentProblemType = "project_not_found"
	RECIPE_NOT_FOUND CreateRecipeDeploymentProblemType = "recipe_not_found"
	CREATED_INSTANCE_QUOTA_EXCEEDED CreateRecipeDeploymentProblemType = "created_instance_quota_exceeded"
)

// All allowed values of CreateRecipeDeploymentProblemType enum
var AllowedCreateRecipeDeploymentProblemTypeEnumValues = []CreateRecipeDeploymentProblemType{
	"billing_information_required",
	"created_quota_exceeded",
	"name_conflict",
	"project_not_found",
	"unexpected_error",
	"organization_not_found",
	"recipe_not_found",
	"created_instance_quota_exceeded",
	"null",
}

func (v *CreateRecipeDeploymentProblemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreateRecipeDeploymentProblemType(value)
	for _, existing := range AllowedCreateRecipeDeploymentProblemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CreateRecipeDeploymentProblemType", value)
}

// NewCreateRecipeDeploymentProblemTypeFromValue returns a pointer to a valid CreateRecipeDeploymentProblemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreateRecipeDeploymentProblemTypeFromValue(v string) (*CreateRecipeDeploymentProblemType, error) {
	ev := CreateRecipeDeploymentProblemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreateRecipeDeploymentProblemType: valid values are %v", v, AllowedCreateRecipeDeploymentProblemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreateRecipeDeploymentProblemType) IsValid() bool {
	for _, existing := range AllowedCreateRecipeDeploymentProblemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreateRecipeDeploymentProblemType value
func (v CreateRecipeDeploymentProblemType) Ptr() *CreateRecipeDeploymentProblemType {
	return &v
}

type NullableCreateRecipeDeploymentProblemType struct {
	value *CreateRecipeDeploymentProblemType
	isSet bool
}

func (v NullableCreateRecipeDeploymentProblemType) Get() *CreateRecipeDeploymentProblemType {
	return v.value
}

func (v *NullableCreateRecipeDeploymentProblemType) Set(val *CreateRecipeDeploymentProblemType) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecipeDeploymentProblemType) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecipeDeploymentProblemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecipeDeploymentProblemType(val *CreateRecipeDeploymentProblemType) *NullableCreateRecipeDeploymentProblemType {
	return &NullableCreateRecipeDeploymentProblemType{value: val, isSet: true}
}

func (v NullableCreateRecipeDeploymentProblemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecipeDeploymentProblemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

