# GpuClassPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | [**NullableContainerGroupPriority**](ContainerGroupPriority.md) |  | 
**Price** | **string** | The price | 

## Methods

### NewGpuClassPrice

`func NewGpuClassPrice(priority NullableContainerGroupPriority, price string, ) *GpuClassPrice`

NewGpuClassPrice instantiates a new GpuClassPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpuClassPriceWithDefaults

`func NewGpuClassPriceWithDefaults() *GpuClassPrice`

NewGpuClassPriceWithDefaults instantiates a new GpuClassPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *GpuClassPrice) GetPriority() ContainerGroupPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GpuClassPrice) GetPriorityOk() (*ContainerGroupPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GpuClassPrice) SetPriority(v ContainerGroupPriority)`

SetPriority sets Priority field to given value.


### SetPriorityNil

`func (o *GpuClassPrice) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *GpuClassPrice) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetPrice

`func (o *GpuClassPrice) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GpuClassPrice) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GpuClassPrice) SetPrice(v string)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


