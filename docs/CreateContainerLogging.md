# CreateContainerLogging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Axiom** | Pointer to [**ContainerLoggingAxiom**](ContainerLoggingAxiom.md) |  | [optional] 
**Datadog** | Pointer to [**ContainerLoggingDatadog**](ContainerLoggingDatadog.md) |  | [optional] 
**Http** | Pointer to [**CreateContainerLoggingHttp**](CreateContainerLoggingHttp.md) |  | [optional] 
**NewRelic** | Pointer to [**ContainerLoggingNewRelic**](ContainerLoggingNewRelic.md) |  | [optional] 
**Splunk** | Pointer to [**ContainerLoggingSplunk**](ContainerLoggingSplunk.md) |  | [optional] 
**Tcp** | Pointer to [**ContainerLoggingTcp**](ContainerLoggingTcp.md) |  | [optional] 

## Methods

### NewCreateContainerLogging

`func NewCreateContainerLogging() *CreateContainerLogging`

NewCreateContainerLogging instantiates a new CreateContainerLogging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerLoggingWithDefaults

`func NewCreateContainerLoggingWithDefaults() *CreateContainerLogging`

NewCreateContainerLoggingWithDefaults instantiates a new CreateContainerLogging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxiom

`func (o *CreateContainerLogging) GetAxiom() ContainerLoggingAxiom`

GetAxiom returns the Axiom field if non-nil, zero value otherwise.

### GetAxiomOk

`func (o *CreateContainerLogging) GetAxiomOk() (*ContainerLoggingAxiom, bool)`

GetAxiomOk returns a tuple with the Axiom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxiom

`func (o *CreateContainerLogging) SetAxiom(v ContainerLoggingAxiom)`

SetAxiom sets Axiom field to given value.

### HasAxiom

`func (o *CreateContainerLogging) HasAxiom() bool`

HasAxiom returns a boolean if a field has been set.

### GetDatadog

`func (o *CreateContainerLogging) GetDatadog() ContainerLoggingDatadog`

GetDatadog returns the Datadog field if non-nil, zero value otherwise.

### GetDatadogOk

`func (o *CreateContainerLogging) GetDatadogOk() (*ContainerLoggingDatadog, bool)`

GetDatadogOk returns a tuple with the Datadog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatadog

`func (o *CreateContainerLogging) SetDatadog(v ContainerLoggingDatadog)`

SetDatadog sets Datadog field to given value.

### HasDatadog

`func (o *CreateContainerLogging) HasDatadog() bool`

HasDatadog returns a boolean if a field has been set.

### GetHttp

`func (o *CreateContainerLogging) GetHttp() CreateContainerLoggingHttp`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *CreateContainerLogging) GetHttpOk() (*CreateContainerLoggingHttp, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *CreateContainerLogging) SetHttp(v CreateContainerLoggingHttp)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *CreateContainerLogging) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetNewRelic

`func (o *CreateContainerLogging) GetNewRelic() ContainerLoggingNewRelic`

GetNewRelic returns the NewRelic field if non-nil, zero value otherwise.

### GetNewRelicOk

`func (o *CreateContainerLogging) GetNewRelicOk() (*ContainerLoggingNewRelic, bool)`

GetNewRelicOk returns a tuple with the NewRelic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRelic

`func (o *CreateContainerLogging) SetNewRelic(v ContainerLoggingNewRelic)`

SetNewRelic sets NewRelic field to given value.

### HasNewRelic

`func (o *CreateContainerLogging) HasNewRelic() bool`

HasNewRelic returns a boolean if a field has been set.

### GetSplunk

`func (o *CreateContainerLogging) GetSplunk() ContainerLoggingSplunk`

GetSplunk returns the Splunk field if non-nil, zero value otherwise.

### GetSplunkOk

`func (o *CreateContainerLogging) GetSplunkOk() (*ContainerLoggingSplunk, bool)`

GetSplunkOk returns a tuple with the Splunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplunk

`func (o *CreateContainerLogging) SetSplunk(v ContainerLoggingSplunk)`

SetSplunk sets Splunk field to given value.

### HasSplunk

`func (o *CreateContainerLogging) HasSplunk() bool`

HasSplunk returns a boolean if a field has been set.

### GetTcp

`func (o *CreateContainerLogging) GetTcp() ContainerLoggingTcp`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *CreateContainerLogging) GetTcpOk() (*ContainerLoggingTcp, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *CreateContainerLogging) SetTcp(v ContainerLoggingTcp)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *CreateContainerLogging) HasTcp() bool`

HasTcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


