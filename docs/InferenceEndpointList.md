# InferenceEndpointList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InferenceEndpoint**](InferenceEndpoint.md) | The list of inference endpoints. | 
**Page** | **int32** | The page number. | 
**PageSize** | **int32** | The maximum number of items per page. | 
**TotalSize** | **int32** | The total number of items in the collection. | 

## Methods

### NewInferenceEndpointList

`func NewInferenceEndpointList(items []InferenceEndpoint, page int32, pageSize int32, totalSize int32, ) *InferenceEndpointList`

NewInferenceEndpointList instantiates a new InferenceEndpointList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInferenceEndpointListWithDefaults

`func NewInferenceEndpointListWithDefaults() *InferenceEndpointList`

NewInferenceEndpointListWithDefaults instantiates a new InferenceEndpointList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InferenceEndpointList) GetItems() []InferenceEndpoint`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InferenceEndpointList) GetItemsOk() (*[]InferenceEndpoint, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InferenceEndpointList) SetItems(v []InferenceEndpoint)`

SetItems sets Items field to given value.


### GetPage

`func (o *InferenceEndpointList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *InferenceEndpointList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *InferenceEndpointList) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *InferenceEndpointList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *InferenceEndpointList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *InferenceEndpointList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotalSize

`func (o *InferenceEndpointList) GetTotalSize() int32`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *InferenceEndpointList) GetTotalSizeOk() (*int32, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *InferenceEndpointList) SetTotalSize(v int32)`

SetTotalSize sets TotalSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


