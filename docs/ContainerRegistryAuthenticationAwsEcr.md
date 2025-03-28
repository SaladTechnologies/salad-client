# ContainerRegistryAuthenticationAwsEcr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | **string** | AWS access key ID used for ECR authentication | 
**SecretAccessKey** | **string** | AWS secret access key used for ECR authentication | 

## Methods

### NewContainerRegistryAuthenticationAwsEcr

`func NewContainerRegistryAuthenticationAwsEcr(accessKeyId string, secretAccessKey string, ) *ContainerRegistryAuthenticationAwsEcr`

NewContainerRegistryAuthenticationAwsEcr instantiates a new ContainerRegistryAuthenticationAwsEcr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRegistryAuthenticationAwsEcrWithDefaults

`func NewContainerRegistryAuthenticationAwsEcrWithDefaults() *ContainerRegistryAuthenticationAwsEcr`

NewContainerRegistryAuthenticationAwsEcrWithDefaults instantiates a new ContainerRegistryAuthenticationAwsEcr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *ContainerRegistryAuthenticationAwsEcr) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *ContainerRegistryAuthenticationAwsEcr) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *ContainerRegistryAuthenticationAwsEcr) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.


### GetSecretAccessKey

`func (o *ContainerRegistryAuthenticationAwsEcr) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *ContainerRegistryAuthenticationAwsEcr) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *ContainerRegistryAuthenticationAwsEcr) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


