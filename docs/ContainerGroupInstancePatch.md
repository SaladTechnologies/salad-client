# ContainerGroupInstancePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletionCost** | Pointer to **NullableInt32** | The cost of deleting the container group instance | [optional] 

## Methods

### NewContainerGroupInstancePatch

`func NewContainerGroupInstancePatch() *ContainerGroupInstancePatch`

NewContainerGroupInstancePatch instantiates a new ContainerGroupInstancePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupInstancePatchWithDefaults

`func NewContainerGroupInstancePatchWithDefaults() *ContainerGroupInstancePatch`

NewContainerGroupInstancePatchWithDefaults instantiates a new ContainerGroupInstancePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletionCost

`func (o *ContainerGroupInstancePatch) GetDeletionCost() int32`

GetDeletionCost returns the DeletionCost field if non-nil, zero value otherwise.

### GetDeletionCostOk

`func (o *ContainerGroupInstancePatch) GetDeletionCostOk() (*int32, bool)`

GetDeletionCostOk returns a tuple with the DeletionCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionCost

`func (o *ContainerGroupInstancePatch) SetDeletionCost(v int32)`

SetDeletionCost sets DeletionCost field to given value.

### HasDeletionCost

`func (o *ContainerGroupInstancePatch) HasDeletionCost() bool`

HasDeletionCost returns a boolean if a field has been set.

### SetDeletionCostNil

`func (o *ContainerGroupInstancePatch) SetDeletionCostNil(b bool)`

 SetDeletionCostNil sets the value for DeletionCost to be an explicit nil

### UnsetDeletionCost
`func (o *ContainerGroupInstancePatch) UnsetDeletionCost()`

UnsetDeletionCost ensures that no value is present for DeletionCost, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


