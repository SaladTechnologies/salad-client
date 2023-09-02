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
)

// checks if the RecipeDeploymentList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipeDeploymentList{}

// RecipeDeploymentList Represents a list of recipe deployments
type RecipeDeploymentList struct {
	Items []RecipeDeployment `json:"items"`
}

// NewRecipeDeploymentList instantiates a new RecipeDeploymentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipeDeploymentList(items []RecipeDeployment) *RecipeDeploymentList {
	this := RecipeDeploymentList{}
	this.Items = items
	return &this
}

// NewRecipeDeploymentListWithDefaults instantiates a new RecipeDeploymentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipeDeploymentListWithDefaults() *RecipeDeploymentList {
	this := RecipeDeploymentList{}
	return &this
}

// GetItems returns the Items field value
func (o *RecipeDeploymentList) GetItems() []RecipeDeployment {
	if o == nil {
		var ret []RecipeDeployment
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *RecipeDeploymentList) GetItemsOk() ([]RecipeDeployment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *RecipeDeploymentList) SetItems(v []RecipeDeployment) {
	o.Items = v
}

func (o RecipeDeploymentList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipeDeploymentList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableRecipeDeploymentList struct {
	value *RecipeDeploymentList
	isSet bool
}

func (v NullableRecipeDeploymentList) Get() *RecipeDeploymentList {
	return v.value
}

func (v *NullableRecipeDeploymentList) Set(val *RecipeDeploymentList) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipeDeploymentList) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipeDeploymentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipeDeploymentList(val *RecipeDeploymentList) *NullableRecipeDeploymentList {
	return &NullableRecipeDeploymentList{value: val, isSet: true}
}

func (v NullableRecipeDeploymentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipeDeploymentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


