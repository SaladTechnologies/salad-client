# ContainerGroupLivenessProbe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exec** | Pointer to [**ContainerGroupProbeExec**](ContainerGroupProbeExec.md) |  | [optional] 
**FailureThreshold** | **int32** | Number of consecutive failures required to consider the probe as failed | [default to 3]
**Grpc** | Pointer to [**ContainerGroupProbeGrpc**](ContainerGroupProbeGrpc.md) |  | [optional] 
**Http** | Pointer to [**ContainerGroupProbeHttp**](ContainerGroupProbeHttp.md) |  | [optional] 
**InitialDelaySeconds** | **int32** | Number of seconds to wait after container start before initiating liveness probes | [default to 0]
**PeriodSeconds** | **int32** | Frequency in seconds at which the probe should be executed | [default to 10]
**SuccessThreshold** | **int32** | Number of consecutive successes required to consider the probe successful | [default to 1]
**Tcp** | Pointer to [**ContainerGroupProbeTcp**](ContainerGroupProbeTcp.md) |  | [optional] 
**TimeoutSeconds** | **int32** | Number of seconds after which the probe times out if no response is received | [default to 30]

## Methods

### NewContainerGroupLivenessProbe

`func NewContainerGroupLivenessProbe(failureThreshold int32, initialDelaySeconds int32, periodSeconds int32, successThreshold int32, timeoutSeconds int32, ) *ContainerGroupLivenessProbe`

NewContainerGroupLivenessProbe instantiates a new ContainerGroupLivenessProbe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupLivenessProbeWithDefaults

`func NewContainerGroupLivenessProbeWithDefaults() *ContainerGroupLivenessProbe`

NewContainerGroupLivenessProbeWithDefaults instantiates a new ContainerGroupLivenessProbe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExec

`func (o *ContainerGroupLivenessProbe) GetExec() ContainerGroupProbeExec`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *ContainerGroupLivenessProbe) GetExecOk() (*ContainerGroupProbeExec, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *ContainerGroupLivenessProbe) SetExec(v ContainerGroupProbeExec)`

SetExec sets Exec field to given value.

### HasExec

`func (o *ContainerGroupLivenessProbe) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetFailureThreshold

`func (o *ContainerGroupLivenessProbe) GetFailureThreshold() int32`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *ContainerGroupLivenessProbe) GetFailureThresholdOk() (*int32, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *ContainerGroupLivenessProbe) SetFailureThreshold(v int32)`

SetFailureThreshold sets FailureThreshold field to given value.


### GetGrpc

`func (o *ContainerGroupLivenessProbe) GetGrpc() ContainerGroupProbeGrpc`

GetGrpc returns the Grpc field if non-nil, zero value otherwise.

### GetGrpcOk

`func (o *ContainerGroupLivenessProbe) GetGrpcOk() (*ContainerGroupProbeGrpc, bool)`

GetGrpcOk returns a tuple with the Grpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpc

`func (o *ContainerGroupLivenessProbe) SetGrpc(v ContainerGroupProbeGrpc)`

SetGrpc sets Grpc field to given value.

### HasGrpc

`func (o *ContainerGroupLivenessProbe) HasGrpc() bool`

HasGrpc returns a boolean if a field has been set.

### GetHttp

`func (o *ContainerGroupLivenessProbe) GetHttp() ContainerGroupProbeHttp`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *ContainerGroupLivenessProbe) GetHttpOk() (*ContainerGroupProbeHttp, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *ContainerGroupLivenessProbe) SetHttp(v ContainerGroupProbeHttp)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *ContainerGroupLivenessProbe) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetInitialDelaySeconds

`func (o *ContainerGroupLivenessProbe) GetInitialDelaySeconds() int32`

GetInitialDelaySeconds returns the InitialDelaySeconds field if non-nil, zero value otherwise.

### GetInitialDelaySecondsOk

`func (o *ContainerGroupLivenessProbe) GetInitialDelaySecondsOk() (*int32, bool)`

GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDelaySeconds

`func (o *ContainerGroupLivenessProbe) SetInitialDelaySeconds(v int32)`

SetInitialDelaySeconds sets InitialDelaySeconds field to given value.


### GetPeriodSeconds

`func (o *ContainerGroupLivenessProbe) GetPeriodSeconds() int32`

GetPeriodSeconds returns the PeriodSeconds field if non-nil, zero value otherwise.

### GetPeriodSecondsOk

`func (o *ContainerGroupLivenessProbe) GetPeriodSecondsOk() (*int32, bool)`

GetPeriodSecondsOk returns a tuple with the PeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodSeconds

`func (o *ContainerGroupLivenessProbe) SetPeriodSeconds(v int32)`

SetPeriodSeconds sets PeriodSeconds field to given value.


### GetSuccessThreshold

`func (o *ContainerGroupLivenessProbe) GetSuccessThreshold() int32`

GetSuccessThreshold returns the SuccessThreshold field if non-nil, zero value otherwise.

### GetSuccessThresholdOk

`func (o *ContainerGroupLivenessProbe) GetSuccessThresholdOk() (*int32, bool)`

GetSuccessThresholdOk returns a tuple with the SuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessThreshold

`func (o *ContainerGroupLivenessProbe) SetSuccessThreshold(v int32)`

SetSuccessThreshold sets SuccessThreshold field to given value.


### GetTcp

`func (o *ContainerGroupLivenessProbe) GetTcp() ContainerGroupProbeTcp`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *ContainerGroupLivenessProbe) GetTcpOk() (*ContainerGroupProbeTcp, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *ContainerGroupLivenessProbe) SetTcp(v ContainerGroupProbeTcp)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *ContainerGroupLivenessProbe) HasTcp() bool`

HasTcp returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *ContainerGroupLivenessProbe) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *ContainerGroupLivenessProbe) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *ContainerGroupLivenessProbe) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


