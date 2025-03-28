# ContainerGroupInstanceCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | [**[]ContainerGroupInstance**](ContainerGroupInstance.md) | An array of container group instances, each representing a deployed container group with its current state and configuration information. | 

## Methods

### NewContainerGroupInstanceCollection

`func NewContainerGroupInstanceCollection(instances []ContainerGroupInstance, ) *ContainerGroupInstanceCollection`

NewContainerGroupInstanceCollection instantiates a new ContainerGroupInstanceCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupInstanceCollectionWithDefaults

`func NewContainerGroupInstanceCollectionWithDefaults() *ContainerGroupInstanceCollection`

NewContainerGroupInstanceCollectionWithDefaults instantiates a new ContainerGroupInstanceCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *ContainerGroupInstanceCollection) GetInstances() []ContainerGroupInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ContainerGroupInstanceCollection) GetInstancesOk() (*[]ContainerGroupInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ContainerGroupInstanceCollection) SetInstances(v []ContainerGroupInstance)`

SetInstances sets Instances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


