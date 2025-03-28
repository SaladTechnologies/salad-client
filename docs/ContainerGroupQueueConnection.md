# ContainerGroupQueueConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The endpoint path for accessing the queue service, relative to the base URL of the queue server. | 
**Port** | **int32** | The network port number used to connect to the queue service. Must be a valid TCP/IP port between 1 and 65535. | 
**QueueName** | **string** | Unique identifier for the queue. Must start with a lowercase letter, can contain lowercase letters, numbers, and hyphens, and must end with a letter or number. | 

## Methods

### NewContainerGroupQueueConnection

`func NewContainerGroupQueueConnection(path string, port int32, queueName string, ) *ContainerGroupQueueConnection`

NewContainerGroupQueueConnection instantiates a new ContainerGroupQueueConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGroupQueueConnectionWithDefaults

`func NewContainerGroupQueueConnectionWithDefaults() *ContainerGroupQueueConnection`

NewContainerGroupQueueConnectionWithDefaults instantiates a new ContainerGroupQueueConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ContainerGroupQueueConnection) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ContainerGroupQueueConnection) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ContainerGroupQueueConnection) SetPath(v string)`

SetPath sets Path field to given value.


### GetPort

`func (o *ContainerGroupQueueConnection) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ContainerGroupQueueConnection) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ContainerGroupQueueConnection) SetPort(v int32)`

SetPort sets Port field to given value.


### GetQueueName

`func (o *ContainerGroupQueueConnection) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *ContainerGroupQueueConnection) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *ContainerGroupQueueConnection) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


