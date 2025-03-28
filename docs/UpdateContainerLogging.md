# UpdateContainerLogging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Axiom** | Pointer to [**ContainerLoggingAxiom**](ContainerLoggingAxiom.md) |  | [optional] 
**Datadog** | Pointer to [**ContainerLoggingDatadog**](ContainerLoggingDatadog.md) |  | [optional] 
**Http** | Pointer to [**ContainerLoggingHttp**](ContainerLoggingHttp.md) |  | [optional] 
**NewRelic** | Pointer to [**ContainerLoggingNewRelic**](ContainerLoggingNewRelic.md) |  | [optional] 
**Splunk** | Pointer to [**ContainerLoggingSplunk**](ContainerLoggingSplunk.md) |  | [optional] 
**Tcp** | Pointer to [**ContainerLoggingTcp**](ContainerLoggingTcp.md) |  | [optional] 

## Methods

### NewUpdateContainerLogging

`func NewUpdateContainerLogging() *UpdateContainerLogging`

NewUpdateContainerLogging instantiates a new UpdateContainerLogging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContainerLoggingWithDefaults

`func NewUpdateContainerLoggingWithDefaults() *UpdateContainerLogging`

NewUpdateContainerLoggingWithDefaults instantiates a new UpdateContainerLogging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxiom

`func (o *UpdateContainerLogging) GetAxiom() ContainerLoggingAxiom`

GetAxiom returns the Axiom field if non-nil, zero value otherwise.

### GetAxiomOk

`func (o *UpdateContainerLogging) GetAxiomOk() (*ContainerLoggingAxiom, bool)`

GetAxiomOk returns a tuple with the Axiom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxiom

`func (o *UpdateContainerLogging) SetAxiom(v ContainerLoggingAxiom)`

SetAxiom sets Axiom field to given value.

### HasAxiom

`func (o *UpdateContainerLogging) HasAxiom() bool`

HasAxiom returns a boolean if a field has been set.

### GetDatadog

`func (o *UpdateContainerLogging) GetDatadog() ContainerLoggingDatadog`

GetDatadog returns the Datadog field if non-nil, zero value otherwise.

### GetDatadogOk

`func (o *UpdateContainerLogging) GetDatadogOk() (*ContainerLoggingDatadog, bool)`

GetDatadogOk returns a tuple with the Datadog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatadog

`func (o *UpdateContainerLogging) SetDatadog(v ContainerLoggingDatadog)`

SetDatadog sets Datadog field to given value.

### HasDatadog

`func (o *UpdateContainerLogging) HasDatadog() bool`

HasDatadog returns a boolean if a field has been set.

### GetHttp

`func (o *UpdateContainerLogging) GetHttp() ContainerLoggingHttp`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *UpdateContainerLogging) GetHttpOk() (*ContainerLoggingHttp, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *UpdateContainerLogging) SetHttp(v ContainerLoggingHttp)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *UpdateContainerLogging) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetNewRelic

`func (o *UpdateContainerLogging) GetNewRelic() ContainerLoggingNewRelic`

GetNewRelic returns the NewRelic field if non-nil, zero value otherwise.

### GetNewRelicOk

`func (o *UpdateContainerLogging) GetNewRelicOk() (*ContainerLoggingNewRelic, bool)`

GetNewRelicOk returns a tuple with the NewRelic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRelic

`func (o *UpdateContainerLogging) SetNewRelic(v ContainerLoggingNewRelic)`

SetNewRelic sets NewRelic field to given value.

### HasNewRelic

`func (o *UpdateContainerLogging) HasNewRelic() bool`

HasNewRelic returns a boolean if a field has been set.

### GetSplunk

`func (o *UpdateContainerLogging) GetSplunk() ContainerLoggingSplunk`

GetSplunk returns the Splunk field if non-nil, zero value otherwise.

### GetSplunkOk

`func (o *UpdateContainerLogging) GetSplunkOk() (*ContainerLoggingSplunk, bool)`

GetSplunkOk returns a tuple with the Splunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplunk

`func (o *UpdateContainerLogging) SetSplunk(v ContainerLoggingSplunk)`

SetSplunk sets Splunk field to given value.

### HasSplunk

`func (o *UpdateContainerLogging) HasSplunk() bool`

HasSplunk returns a boolean if a field has been set.

### GetTcp

`func (o *UpdateContainerLogging) GetTcp() ContainerLoggingTcp`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *UpdateContainerLogging) GetTcpOk() (*ContainerLoggingTcp, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *UpdateContainerLogging) SetTcp(v ContainerLoggingTcp)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *UpdateContainerLogging) HasTcp() bool`

HasTcp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


