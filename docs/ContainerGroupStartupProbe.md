# ContainerGroupStartupProbe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tcp** | Pointer to [**ContainerGroupProbeTcp**](ContainerGroupProbeTcp.md) |  | [optional] 
**Http** | Pointer to [**ContainerGroupProbeHttp**](ContainerGroupProbeHttp.md) |  | [optional] 
**Grpc** | Pointer to [**ContainerGroupProbeGrpc**](ContainerGroupProbeGrpc.md) |  | [optional] 
**Exec** | Pointer to [**ContainerGroupProbeExec**](ContainerGroupProbeExec.md) |  | [optional] 
**InitialDelaySeconds** | **int32** |  | [default to 0]
**PeriodSeconds** | **int32** |  | [default to 3]
**TimeoutSeconds** | **int32** |  | [default to 10]
**SuccessThreshold** | **int32** |  | [default to 2]
**FailureThreshold** | **int32** |  | [default to 1200]

## Methods

### NewContainerGroupStartupProbe

`func NewContainerGroupStartupProbe(initialDelaySeconds int32, periodSeconds int32, timeoutSeconds int32, successThreshold int32, failureThreshold int32, ) *ContainerGroupStartupProbe`

NewContainerGroupStartupProbe instantiates a new ContainerGroupStartupProbe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupStartupProbeWithDefaults

`func NewContainerGroupStartupProbeWithDefaults() *ContainerGroupStartupProbe`

NewContainerGroupStartupProbeWithDefaults instantiates a new ContainerGroupStartupProbe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTcp

`func (o *ContainerGroupStartupProbe) GetTcp() ContainerGroupProbeTcp`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *ContainerGroupStartupProbe) GetTcpOk() (*ContainerGroupProbeTcp, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *ContainerGroupStartupProbe) SetTcp(v ContainerGroupProbeTcp)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *ContainerGroupStartupProbe) HasTcp() bool`

HasTcp returns a boolean if a field has been set.

### GetHttp

`func (o *ContainerGroupStartupProbe) GetHttp() ContainerGroupProbeHttp`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *ContainerGroupStartupProbe) GetHttpOk() (*ContainerGroupProbeHttp, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *ContainerGroupStartupProbe) SetHttp(v ContainerGroupProbeHttp)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *ContainerGroupStartupProbe) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetGrpc

`func (o *ContainerGroupStartupProbe) GetGrpc() ContainerGroupProbeGrpc`

GetGrpc returns the Grpc field if non-nil, zero value otherwise.

### GetGrpcOk

`func (o *ContainerGroupStartupProbe) GetGrpcOk() (*ContainerGroupProbeGrpc, bool)`

GetGrpcOk returns a tuple with the Grpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpc

`func (o *ContainerGroupStartupProbe) SetGrpc(v ContainerGroupProbeGrpc)`

SetGrpc sets Grpc field to given value.

### HasGrpc

`func (o *ContainerGroupStartupProbe) HasGrpc() bool`

HasGrpc returns a boolean if a field has been set.

### GetExec

`func (o *ContainerGroupStartupProbe) GetExec() ContainerGroupProbeExec`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *ContainerGroupStartupProbe) GetExecOk() (*ContainerGroupProbeExec, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *ContainerGroupStartupProbe) SetExec(v ContainerGroupProbeExec)`

SetExec sets Exec field to given value.

### HasExec

`func (o *ContainerGroupStartupProbe) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetInitialDelaySeconds

`func (o *ContainerGroupStartupProbe) GetInitialDelaySeconds() int32`

GetInitialDelaySeconds returns the InitialDelaySeconds field if non-nil, zero value otherwise.

### GetInitialDelaySecondsOk

`func (o *ContainerGroupStartupProbe) GetInitialDelaySecondsOk() (*int32, bool)`

GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDelaySeconds

`func (o *ContainerGroupStartupProbe) SetInitialDelaySeconds(v int32)`

SetInitialDelaySeconds sets InitialDelaySeconds field to given value.


### GetPeriodSeconds

`func (o *ContainerGroupStartupProbe) GetPeriodSeconds() int32`

GetPeriodSeconds returns the PeriodSeconds field if non-nil, zero value otherwise.

### GetPeriodSecondsOk

`func (o *ContainerGroupStartupProbe) GetPeriodSecondsOk() (*int32, bool)`

GetPeriodSecondsOk returns a tuple with the PeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodSeconds

`func (o *ContainerGroupStartupProbe) SetPeriodSeconds(v int32)`

SetPeriodSeconds sets PeriodSeconds field to given value.


### GetTimeoutSeconds

`func (o *ContainerGroupStartupProbe) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *ContainerGroupStartupProbe) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *ContainerGroupStartupProbe) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### GetSuccessThreshold

`func (o *ContainerGroupStartupProbe) GetSuccessThreshold() int32`

GetSuccessThreshold returns the SuccessThreshold field if non-nil, zero value otherwise.

### GetSuccessThresholdOk

`func (o *ContainerGroupStartupProbe) GetSuccessThresholdOk() (*int32, bool)`

GetSuccessThresholdOk returns a tuple with the SuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessThreshold

`func (o *ContainerGroupStartupProbe) SetSuccessThreshold(v int32)`

SetSuccessThreshold sets SuccessThreshold field to given value.


### GetFailureThreshold

`func (o *ContainerGroupStartupProbe) GetFailureThreshold() int32`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *ContainerGroupStartupProbe) GetFailureThresholdOk() (*int32, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *ContainerGroupStartupProbe) SetFailureThreshold(v int32)`

SetFailureThreshold sets FailureThreshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


