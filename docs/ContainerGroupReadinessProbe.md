# ContainerGroupReadinessProbe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exec** | Pointer to [**ContainerGroupProbeExec**](ContainerGroupProbeExec.md) |  | [optional] 
**FailureThreshold** | **int32** | The number of consecutive failures required to consider the probe failed. After this many consecutive failures, the container is marked as not ready. | [default to 3]
**Grpc** | Pointer to [**ContainerGroupProbeGrpc**](ContainerGroupProbeGrpc.md) |  | [optional] 
**Http** | Pointer to [**ContainerGroupProbeHttp**](ContainerGroupProbeHttp.md) |  | [optional] 
**InitialDelaySeconds** | **int32** | The time in seconds to wait after the container starts before initiating the first probe. This allows time for the application to initialize before being tested. | [default to 0]
**PeriodSeconds** | **int32** | How frequently (in seconds) the probe should be executed during the container&#39;s lifetime. Specifies the interval between consecutive probe executions. | [default to 1]
**SuccessThreshold** | **int32** | The minimum consecutive successes required to consider the probe successful after it has failed. Defines how many successful probe results are needed to transition from failure to success. | [default to 1]
**Tcp** | Pointer to [**ContainerGroupProbeTcp**](ContainerGroupProbeTcp.md) |  | [optional] 
**TimeoutSeconds** | **int32** | The maximum time in seconds that the probe has to complete. If the probe doesn&#39;t return a result before the timeout, it&#39;s considered failed. | [default to 1]

## Methods

### NewContainerGroupReadinessProbe

`func NewContainerGroupReadinessProbe(failureThreshold int32, initialDelaySeconds int32, periodSeconds int32, successThreshold int32, timeoutSeconds int32, ) *ContainerGroupReadinessProbe`

NewContainerGroupReadinessProbe instantiates a new ContainerGroupReadinessProbe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupReadinessProbeWithDefaults

`func NewContainerGroupReadinessProbeWithDefaults() *ContainerGroupReadinessProbe`

NewContainerGroupReadinessProbeWithDefaults instantiates a new ContainerGroupReadinessProbe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExec

`func (o *ContainerGroupReadinessProbe) GetExec() ContainerGroupProbeExec`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *ContainerGroupReadinessProbe) GetExecOk() (*ContainerGroupProbeExec, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *ContainerGroupReadinessProbe) SetExec(v ContainerGroupProbeExec)`

SetExec sets Exec field to given value.

### HasExec

`func (o *ContainerGroupReadinessProbe) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetFailureThreshold

`func (o *ContainerGroupReadinessProbe) GetFailureThreshold() int32`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *ContainerGroupReadinessProbe) GetFailureThresholdOk() (*int32, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *ContainerGroupReadinessProbe) SetFailureThreshold(v int32)`

SetFailureThreshold sets FailureThreshold field to given value.


### GetGrpc

`func (o *ContainerGroupReadinessProbe) GetGrpc() ContainerGroupProbeGrpc`

GetGrpc returns the Grpc field if non-nil, zero value otherwise.

### GetGrpcOk

`func (o *ContainerGroupReadinessProbe) GetGrpcOk() (*ContainerGroupProbeGrpc, bool)`

GetGrpcOk returns a tuple with the Grpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrpc

`func (o *ContainerGroupReadinessProbe) SetGrpc(v ContainerGroupProbeGrpc)`

SetGrpc sets Grpc field to given value.

### HasGrpc

`func (o *ContainerGroupReadinessProbe) HasGrpc() bool`

HasGrpc returns a boolean if a field has been set.

### GetHttp

`func (o *ContainerGroupReadinessProbe) GetHttp() ContainerGroupProbeHttp`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *ContainerGroupReadinessProbe) GetHttpOk() (*ContainerGroupProbeHttp, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *ContainerGroupReadinessProbe) SetHttp(v ContainerGroupProbeHttp)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *ContainerGroupReadinessProbe) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetInitialDelaySeconds

`func (o *ContainerGroupReadinessProbe) GetInitialDelaySeconds() int32`

GetInitialDelaySeconds returns the InitialDelaySeconds field if non-nil, zero value otherwise.

### GetInitialDelaySecondsOk

`func (o *ContainerGroupReadinessProbe) GetInitialDelaySecondsOk() (*int32, bool)`

GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDelaySeconds

`func (o *ContainerGroupReadinessProbe) SetInitialDelaySeconds(v int32)`

SetInitialDelaySeconds sets InitialDelaySeconds field to given value.


### GetPeriodSeconds

`func (o *ContainerGroupReadinessProbe) GetPeriodSeconds() int32`

GetPeriodSeconds returns the PeriodSeconds field if non-nil, zero value otherwise.

### GetPeriodSecondsOk

`func (o *ContainerGroupReadinessProbe) GetPeriodSecondsOk() (*int32, bool)`

GetPeriodSecondsOk returns a tuple with the PeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodSeconds

`func (o *ContainerGroupReadinessProbe) SetPeriodSeconds(v int32)`

SetPeriodSeconds sets PeriodSeconds field to given value.


### GetSuccessThreshold

`func (o *ContainerGroupReadinessProbe) GetSuccessThreshold() int32`

GetSuccessThreshold returns the SuccessThreshold field if non-nil, zero value otherwise.

### GetSuccessThresholdOk

`func (o *ContainerGroupReadinessProbe) GetSuccessThresholdOk() (*int32, bool)`

GetSuccessThresholdOk returns a tuple with the SuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessThreshold

`func (o *ContainerGroupReadinessProbe) SetSuccessThreshold(v int32)`

SetSuccessThreshold sets SuccessThreshold field to given value.


### GetTcp

`func (o *ContainerGroupReadinessProbe) GetTcp() ContainerGroupProbeTcp`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *ContainerGroupReadinessProbe) GetTcpOk() (*ContainerGroupProbeTcp, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *ContainerGroupReadinessProbe) SetTcp(v ContainerGroupProbeTcp)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *ContainerGroupReadinessProbe) HasTcp() bool`

HasTcp returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *ContainerGroupReadinessProbe) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *ContainerGroupReadinessProbe) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *ContainerGroupReadinessProbe) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


