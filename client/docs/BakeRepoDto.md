# BakeRepoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Repo clone URL | 
**Wait** | **bool** | Option to wait for Pizza service to finish baking repo | 

## Methods

### NewBakeRepoDto

`func NewBakeRepoDto(url string, wait bool, ) *BakeRepoDto`

NewBakeRepoDto instantiates a new BakeRepoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBakeRepoDtoWithDefaults

`func NewBakeRepoDtoWithDefaults() *BakeRepoDto`

NewBakeRepoDtoWithDefaults instantiates a new BakeRepoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *BakeRepoDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BakeRepoDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BakeRepoDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetWait

`func (o *BakeRepoDto) GetWait() bool`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *BakeRepoDto) GetWaitOk() (*bool, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *BakeRepoDto) SetWait(v bool)`

SetWait sets Wait field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


