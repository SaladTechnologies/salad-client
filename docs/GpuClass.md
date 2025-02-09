# GpuClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier | 
**Name** | **string** | The GPU class name | 
**Prices** | [**[]GpuClassPrice**](GpuClassPrice.md) | The list of prices for each container group priority | 
**IsHighDemand** | Pointer to **bool** | Whether the GPU class is in high demand | [optional] 

## Methods

### NewGpuClass

`func NewGpuClass(id string, name string, prices []GpuClassPrice, ) *GpuClass`

NewGpuClass instantiates a new GpuClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGpuClassWithDefaults

`func NewGpuClassWithDefaults() *GpuClass`

NewGpuClassWithDefaults instantiates a new GpuClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GpuClass) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GpuClass) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GpuClass) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GpuClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GpuClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GpuClass) SetName(v string)`

SetName sets Name field to given value.


### GetPrices

`func (o *GpuClass) GetPrices() []GpuClassPrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *GpuClass) GetPricesOk() (*[]GpuClassPrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *GpuClass) SetPrices(v []GpuClassPrice)`

SetPrices sets Prices field to given value.


### GetIsHighDemand

`func (o *GpuClass) GetIsHighDemand() bool`

GetIsHighDemand returns the IsHighDemand field if non-nil, zero value otherwise.

### GetIsHighDemandOk

`func (o *GpuClass) GetIsHighDemandOk() (*bool, bool)`

GetIsHighDemandOk returns a tuple with the IsHighDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHighDemand

`func (o *GpuClass) SetIsHighDemand(v bool)`

SetIsHighDemand sets IsHighDemand field to given value.

### HasIsHighDemand

`func (o *GpuClass) HasIsHighDemand() bool`

HasIsHighDemand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


