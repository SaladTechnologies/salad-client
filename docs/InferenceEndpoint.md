# InferenceEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The inference endpoint identifier. | 
**Name** | **string** | The inference endpoint name. | 
**OrganizationName** | **string** | The organization name. | 
**DisplayName** | **string** | The display-friendly name of the resource. | 
**Description** | **string** | The detailed description of the resource. | 
**Readme** | **string** | A markdown file containing a detailed description of the inference endpoint | 
**PriceDescription** | **string** | A description of the price | 
**IconUrl** | **string** | The URL of the icon image | 
**InputSchema** | **string** | The input schema | 
**OutputSchema** | **string** | The output schema | 

## Methods

### NewInferenceEndpoint

`func NewInferenceEndpoint(id string, name string, organizationName string, displayName string, description string, readme string, priceDescription string, iconUrl string, inputSchema string, outputSchema string, ) *InferenceEndpoint`

NewInferenceEndpoint instantiates a new InferenceEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInferenceEndpointWithDefaults

`func NewInferenceEndpointWithDefaults() *InferenceEndpoint`

NewInferenceEndpointWithDefaults instantiates a new InferenceEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InferenceEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InferenceEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InferenceEndpoint) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *InferenceEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InferenceEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InferenceEndpoint) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationName

`func (o *InferenceEndpoint) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *InferenceEndpoint) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *InferenceEndpoint) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetDisplayName

`func (o *InferenceEndpoint) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InferenceEndpoint) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InferenceEndpoint) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *InferenceEndpoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InferenceEndpoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InferenceEndpoint) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetReadme

`func (o *InferenceEndpoint) GetReadme() string`

GetReadme returns the Readme field if non-nil, zero value otherwise.

### GetReadmeOk

`func (o *InferenceEndpoint) GetReadmeOk() (*string, bool)`

GetReadmeOk returns a tuple with the Readme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadme

`func (o *InferenceEndpoint) SetReadme(v string)`

SetReadme sets Readme field to given value.


### GetPriceDescription

`func (o *InferenceEndpoint) GetPriceDescription() string`

GetPriceDescription returns the PriceDescription field if non-nil, zero value otherwise.

### GetPriceDescriptionOk

`func (o *InferenceEndpoint) GetPriceDescriptionOk() (*string, bool)`

GetPriceDescriptionOk returns a tuple with the PriceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDescription

`func (o *InferenceEndpoint) SetPriceDescription(v string)`

SetPriceDescription sets PriceDescription field to given value.


### GetIconUrl

`func (o *InferenceEndpoint) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *InferenceEndpoint) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *InferenceEndpoint) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### GetInputSchema

`func (o *InferenceEndpoint) GetInputSchema() string`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *InferenceEndpoint) GetInputSchemaOk() (*string, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *InferenceEndpoint) SetInputSchema(v string)`

SetInputSchema sets InputSchema field to given value.


### GetOutputSchema

`func (o *InferenceEndpoint) GetOutputSchema() string`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *InferenceEndpoint) GetOutputSchemaOk() (*string, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *InferenceEndpoint) SetOutputSchema(v string)`

SetOutputSchema sets OutputSchema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


