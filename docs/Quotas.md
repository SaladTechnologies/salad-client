# Quotas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerGroupsQuotas** | [**ContainerGroupsQuotas**](ContainerGroupsQuotas.md) |  | 
**CreateTime** | Pointer to **time.Time** | The time the resource was created | [optional] 
**UpdateTime** | Pointer to **time.Time** | The time the resource was last updated | [optional] 

## Methods

### NewQuotas

`func NewQuotas(containerGroupsQuotas ContainerGroupsQuotas, ) *Quotas`

NewQuotas instantiates a new Quotas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotasWithDefaults

`func NewQuotasWithDefaults() *Quotas`

NewQuotasWithDefaults instantiates a new Quotas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerGroupsQuotas

`func (o *Quotas) GetContainerGroupsQuotas() ContainerGroupsQuotas`

GetContainerGroupsQuotas returns the ContainerGroupsQuotas field if non-nil, zero value otherwise.

### GetContainerGroupsQuotasOk

`func (o *Quotas) GetContainerGroupsQuotasOk() (*ContainerGroupsQuotas, bool)`

GetContainerGroupsQuotasOk returns a tuple with the ContainerGroupsQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerGroupsQuotas

`func (o *Quotas) SetContainerGroupsQuotas(v ContainerGroupsQuotas)`

SetContainerGroupsQuotas sets ContainerGroupsQuotas field to given value.


### GetCreateTime

`func (o *Quotas) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Quotas) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Quotas) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Quotas) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *Quotas) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Quotas) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Quotas) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *Quotas) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


