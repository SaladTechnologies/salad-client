# InferenceEndpointCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InferenceEndpoint**](InferenceEndpoint.md) | The list of inference endpoints. | 
**Page** | **int32** | The page number. | 
**PageSize** | **int32** | The maximum number of items per page. | 
**TotalSize** | **int32** | The total number of items in the collection. | 

## Methods

### NewInferenceEndpointCollection

`func NewInferenceEndpointCollection(items []InferenceEndpoint, page int32, pageSize int32, totalSize int32, ) *InferenceEndpointCollection`

NewInferenceEndpointCollection instantiates a new InferenceEndpointCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInferenceEndpointCollectionWithDefaults

`func NewInferenceEndpointCollectionWithDefaults() *InferenceEndpointCollection`

NewInferenceEndpointCollectionWithDefaults instantiates a new InferenceEndpointCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InferenceEndpointCollection) GetItems() []InferenceEndpoint`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InferenceEndpointCollection) GetItemsOk() (*[]InferenceEndpoint, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InferenceEndpointCollection) SetItems(v []InferenceEndpoint)`

SetItems sets Items field to given value.


### GetPage

`func (o *InferenceEndpointCollection) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *InferenceEndpointCollection) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *InferenceEndpointCollection) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *InferenceEndpointCollection) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *InferenceEndpointCollection) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *InferenceEndpointCollection) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotalSize

`func (o *InferenceEndpointCollection) GetTotalSize() int32`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *InferenceEndpointCollection) GetTotalSizeOk() (*int32, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *InferenceEndpointCollection) SetTotalSize(v int32)`

SetTotalSize sets TotalSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


