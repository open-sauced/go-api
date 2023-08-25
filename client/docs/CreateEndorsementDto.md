# CreateEndorsementDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatorUserId** | **int32** | Endorsement Creator User ID | 
**RecipientUserId** | **int32** | Endorsement Recipient User ID | 
**RepoId** | **int32** | Repository ID | 
**SourceCommentUrl** | Pointer to **string** | Endorsement Source Comment URL | [optional] [default to ""]
**SourceContextUrl** | **string** | Endorsement Source Context URL | 
**Type** | **string** | Endorsement Type | 

## Methods

### NewCreateEndorsementDto

`func NewCreateEndorsementDto(creatorUserId int32, recipientUserId int32, repoId int32, sourceContextUrl string, type_ string, ) *CreateEndorsementDto`

NewCreateEndorsementDto instantiates a new CreateEndorsementDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEndorsementDtoWithDefaults

`func NewCreateEndorsementDtoWithDefaults() *CreateEndorsementDto`

NewCreateEndorsementDtoWithDefaults instantiates a new CreateEndorsementDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatorUserId

`func (o *CreateEndorsementDto) GetCreatorUserId() int32`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *CreateEndorsementDto) GetCreatorUserIdOk() (*int32, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *CreateEndorsementDto) SetCreatorUserId(v int32)`

SetCreatorUserId sets CreatorUserId field to given value.


### GetRecipientUserId

`func (o *CreateEndorsementDto) GetRecipientUserId() int32`

GetRecipientUserId returns the RecipientUserId field if non-nil, zero value otherwise.

### GetRecipientUserIdOk

`func (o *CreateEndorsementDto) GetRecipientUserIdOk() (*int32, bool)`

GetRecipientUserIdOk returns a tuple with the RecipientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientUserId

`func (o *CreateEndorsementDto) SetRecipientUserId(v int32)`

SetRecipientUserId sets RecipientUserId field to given value.


### GetRepoId

`func (o *CreateEndorsementDto) GetRepoId() int32`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *CreateEndorsementDto) GetRepoIdOk() (*int32, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *CreateEndorsementDto) SetRepoId(v int32)`

SetRepoId sets RepoId field to given value.


### GetSourceCommentUrl

`func (o *CreateEndorsementDto) GetSourceCommentUrl() string`

GetSourceCommentUrl returns the SourceCommentUrl field if non-nil, zero value otherwise.

### GetSourceCommentUrlOk

`func (o *CreateEndorsementDto) GetSourceCommentUrlOk() (*string, bool)`

GetSourceCommentUrlOk returns a tuple with the SourceCommentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCommentUrl

`func (o *CreateEndorsementDto) SetSourceCommentUrl(v string)`

SetSourceCommentUrl sets SourceCommentUrl field to given value.

### HasSourceCommentUrl

`func (o *CreateEndorsementDto) HasSourceCommentUrl() bool`

HasSourceCommentUrl returns a boolean if a field has been set.

### GetSourceContextUrl

`func (o *CreateEndorsementDto) GetSourceContextUrl() string`

GetSourceContextUrl returns the SourceContextUrl field if non-nil, zero value otherwise.

### GetSourceContextUrlOk

`func (o *CreateEndorsementDto) GetSourceContextUrlOk() (*string, bool)`

GetSourceContextUrlOk returns a tuple with the SourceContextUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceContextUrl

`func (o *CreateEndorsementDto) SetSourceContextUrl(v string)`

SetSourceContextUrl sets SourceContextUrl field to given value.


### GetType

`func (o *CreateEndorsementDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateEndorsementDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateEndorsementDto) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


