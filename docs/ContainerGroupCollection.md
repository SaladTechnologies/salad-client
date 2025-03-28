# ContainerGroupCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ContainerGroup**](ContainerGroup.md) | An array containing container group objects. Each object represents a discrete container group with its own properties, configuration, and status. | 

## Methods

### NewContainerGroupCollection

`func NewContainerGroupCollection(items []ContainerGroup, ) *ContainerGroupCollection`

NewContainerGroupCollection instantiates a new ContainerGroupCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupCollectionWithDefaults

`func NewContainerGroupCollectionWithDefaults() *ContainerGroupCollection`

NewContainerGroupCollectionWithDefaults instantiates a new ContainerGroupCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ContainerGroupCollection) GetItems() []ContainerGroup`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContainerGroupCollection) GetItemsOk() (*[]ContainerGroup, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContainerGroupCollection) SetItems(v []ContainerGroup)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


