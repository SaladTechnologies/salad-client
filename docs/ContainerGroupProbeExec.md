# ContainerGroupProbeExec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **[]string** | The command to execute inside the container. Exit status of 0 is considered successful, any other exit status is considered failure. | 

## Methods

### NewContainerGroupProbeExec

`func NewContainerGroupProbeExec(command []string, ) *ContainerGroupProbeExec`

NewContainerGroupProbeExec instantiates a new ContainerGroupProbeExec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupProbeExecWithDefaults

`func NewContainerGroupProbeExecWithDefaults() *ContainerGroupProbeExec`

NewContainerGroupProbeExecWithDefaults instantiates a new ContainerGroupProbeExec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *ContainerGroupProbeExec) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ContainerGroupProbeExec) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ContainerGroupProbeExec) SetCommand(v []string)`

SetCommand sets Command field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


