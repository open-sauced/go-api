# CreateInsightDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Insight Page Name | 
**IsPublic** | **bool** | Insight Page Visibility | 
**Repos** | [**[]RepoInfo**](RepoInfo.md) | An array of repository information objects | 

## Methods

### NewCreateInsightDto

`func NewCreateInsightDto(name string, isPublic bool, repos []RepoInfo, ) *CreateInsightDto`

NewCreateInsightDto instantiates a new CreateInsightDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInsightDtoWithDefaults

`func NewCreateInsightDtoWithDefaults() *CreateInsightDto`

NewCreateInsightDtoWithDefaults instantiates a new CreateInsightDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateInsightDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInsightDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInsightDto) SetName(v string)`

SetName sets Name field to given value.


### GetIsPublic

`func (o *CreateInsightDto) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *CreateInsightDto) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *CreateInsightDto) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetRepos

`func (o *CreateInsightDto) GetRepos() []RepoInfo`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *CreateInsightDto) GetReposOk() (*[]RepoInfo, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *CreateInsightDto) SetRepos(v []RepoInfo)`

SetRepos sets Repos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


