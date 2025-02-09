# ContainerLogging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Axiom** | Pointer to [**NullableContainerLoggingAxiom**](ContainerLoggingAxiom.md) |  | [optional] 
**Datadog** | Pointer to [**NullableContainerLoggingDatadog**](ContainerLoggingDatadog.md) |  | [optional] 
**NewRelic** | Pointer to [**NullableContainerLoggingNewRelic**](ContainerLoggingNewRelic.md) |  | [optional] 
**Splunk** | Pointer to [**NullableContainerLoggingSplunk**](ContainerLoggingSplunk.md) |  | [optional] 
**Tcp** | Pointer to [**NullableContainerLoggingTcp**](ContainerLoggingTcp.md) |  | [optional] 
**Http** | Pointer to [**NullableContainerLoggingHttp**](ContainerLoggingHttp.md) |  | [optional] 

## Methods

### NewContainerLogging

`func NewContainerLogging() *ContainerLogging`

NewContainerLogging instantiates a new ContainerLogging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerLoggingWithDefaults

`func NewContainerLoggingWithDefaults() *ContainerLogging`

NewContainerLoggingWithDefaults instantiates a new ContainerLogging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxiom

`func (o *ContainerLogging) GetAxiom() ContainerLoggingAxiom`

GetAxiom returns the Axiom field if non-nil, zero value otherwise.

### GetAxiomOk

`func (o *ContainerLogging) GetAxiomOk() (*ContainerLoggingAxiom, bool)`

GetAxiomOk returns a tuple with the Axiom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxiom

`func (o *ContainerLogging) SetAxiom(v ContainerLoggingAxiom)`

SetAxiom sets Axiom field to given value.

### HasAxiom

`func (o *ContainerLogging) HasAxiom() bool`

HasAxiom returns a boolean if a field has been set.

### SetAxiomNil

`func (o *ContainerLogging) SetAxiomNil(b bool)`

 SetAxiomNil sets the value for Axiom to be an explicit nil

### UnsetAxiom
`func (o *ContainerLogging) UnsetAxiom()`

UnsetAxiom ensures that no value is present for Axiom, not even an explicit nil
### GetDatadog

`func (o *ContainerLogging) GetDatadog() ContainerLoggingDatadog`

GetDatadog returns the Datadog field if non-nil, zero value otherwise.

### GetDatadogOk

`func (o *ContainerLogging) GetDatadogOk() (*ContainerLoggingDatadog, bool)`

GetDatadogOk returns a tuple with the Datadog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatadog

`func (o *ContainerLogging) SetDatadog(v ContainerLoggingDatadog)`

SetDatadog sets Datadog field to given value.

### HasDatadog

`func (o *ContainerLogging) HasDatadog() bool`

HasDatadog returns a boolean if a field has been set.

### SetDatadogNil

`func (o *ContainerLogging) SetDatadogNil(b bool)`

 SetDatadogNil sets the value for Datadog to be an explicit nil

### UnsetDatadog
`func (o *ContainerLogging) UnsetDatadog()`

UnsetDatadog ensures that no value is present for Datadog, not even an explicit nil
### GetNewRelic

`func (o *ContainerLogging) GetNewRelic() ContainerLoggingNewRelic`

GetNewRelic returns the NewRelic field if non-nil, zero value otherwise.

### GetNewRelicOk

`func (o *ContainerLogging) GetNewRelicOk() (*ContainerLoggingNewRelic, bool)`

GetNewRelicOk returns a tuple with the NewRelic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRelic

`func (o *ContainerLogging) SetNewRelic(v ContainerLoggingNewRelic)`

SetNewRelic sets NewRelic field to given value.

### HasNewRelic

`func (o *ContainerLogging) HasNewRelic() bool`

HasNewRelic returns a boolean if a field has been set.

### SetNewRelicNil

`func (o *ContainerLogging) SetNewRelicNil(b bool)`

 SetNewRelicNil sets the value for NewRelic to be an explicit nil

### UnsetNewRelic
`func (o *ContainerLogging) UnsetNewRelic()`

UnsetNewRelic ensures that no value is present for NewRelic, not even an explicit nil
### GetSplunk

`func (o *ContainerLogging) GetSplunk() ContainerLoggingSplunk`

GetSplunk returns the Splunk field if non-nil, zero value otherwise.

### GetSplunkOk

`func (o *ContainerLogging) GetSplunkOk() (*ContainerLoggingSplunk, bool)`

GetSplunkOk returns a tuple with the Splunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplunk

`func (o *ContainerLogging) SetSplunk(v ContainerLoggingSplunk)`

SetSplunk sets Splunk field to given value.

### HasSplunk

`func (o *ContainerLogging) HasSplunk() bool`

HasSplunk returns a boolean if a field has been set.

### SetSplunkNil

`func (o *ContainerLogging) SetSplunkNil(b bool)`

 SetSplunkNil sets the value for Splunk to be an explicit nil

### UnsetSplunk
`func (o *ContainerLogging) UnsetSplunk()`

UnsetSplunk ensures that no value is present for Splunk, not even an explicit nil
### GetTcp

`func (o *ContainerLogging) GetTcp() ContainerLoggingTcp`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *ContainerLogging) GetTcpOk() (*ContainerLoggingTcp, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *ContainerLogging) SetTcp(v ContainerLoggingTcp)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *ContainerLogging) HasTcp() bool`

HasTcp returns a boolean if a field has been set.

### SetTcpNil

`func (o *ContainerLogging) SetTcpNil(b bool)`

 SetTcpNil sets the value for Tcp to be an explicit nil

### UnsetTcp
`func (o *ContainerLogging) UnsetTcp()`

UnsetTcp ensures that no value is present for Tcp, not even an explicit nil
### GetHttp

`func (o *ContainerLogging) GetHttp() ContainerLoggingHttp`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *ContainerLogging) GetHttpOk() (*ContainerLoggingHttp, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *ContainerLogging) SetHttp(v ContainerLoggingHttp)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *ContainerLogging) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### SetHttpNil

`func (o *ContainerLogging) SetHttpNil(b bool)`

 SetHttpNil sets the value for Http to be an explicit nil

### UnsetHttp
`func (o *ContainerLogging) UnsetHttp()`

UnsetHttp ensures that no value is present for Http, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


