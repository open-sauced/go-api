# DbEndorsement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Endorsement identifier | 
**CreatorUserId** | Pointer to **int32** | Endorsement Creator User ID | [optional] 
**RecipientUserId** | Pointer to **int32** | Endorsement Recipient User ID | [optional] 
**RepoId** | **int32** | Repository ID | 
**SourceCommentUrl** | Pointer to **string** | Endorsement Source Comment URL | [optional] 
**SourceContextUrl** | **string** | Endorsement Source Context URL | 
**Type** | **string** | Endorsement Type | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing endorsement creation | [optional] 

## Methods

### NewDbEndorsement

`func NewDbEndorsement(id string, repoId int32, sourceContextUrl string, type_ string, ) *DbEndorsement`

NewDbEndorsement instantiates a new DbEndorsement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbEndorsementWithDefaults

`func NewDbEndorsementWithDefaults() *DbEndorsement`

NewDbEndorsementWithDefaults instantiates a new DbEndorsement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbEndorsement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbEndorsement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbEndorsement) SetId(v string)`

SetId sets Id field to given value.


### GetCreatorUserId

`func (o *DbEndorsement) GetCreatorUserId() int32`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *DbEndorsement) GetCreatorUserIdOk() (*int32, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *DbEndorsement) SetCreatorUserId(v int32)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *DbEndorsement) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### GetRecipientUserId

`func (o *DbEndorsement) GetRecipientUserId() int32`

GetRecipientUserId returns the RecipientUserId field if non-nil, zero value otherwise.

### GetRecipientUserIdOk

`func (o *DbEndorsement) GetRecipientUserIdOk() (*int32, bool)`

GetRecipientUserIdOk returns a tuple with the RecipientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientUserId

`func (o *DbEndorsement) SetRecipientUserId(v int32)`

SetRecipientUserId sets RecipientUserId field to given value.

### HasRecipientUserId

`func (o *DbEndorsement) HasRecipientUserId() bool`

HasRecipientUserId returns a boolean if a field has been set.

### GetRepoId

`func (o *DbEndorsement) GetRepoId() int32`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *DbEndorsement) GetRepoIdOk() (*int32, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *DbEndorsement) SetRepoId(v int32)`

SetRepoId sets RepoId field to given value.


### GetSourceCommentUrl

`func (o *DbEndorsement) GetSourceCommentUrl() string`

GetSourceCommentUrl returns the SourceCommentUrl field if non-nil, zero value otherwise.

### GetSourceCommentUrlOk

`func (o *DbEndorsement) GetSourceCommentUrlOk() (*string, bool)`

GetSourceCommentUrlOk returns a tuple with the SourceCommentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCommentUrl

`func (o *DbEndorsement) SetSourceCommentUrl(v string)`

SetSourceCommentUrl sets SourceCommentUrl field to given value.

### HasSourceCommentUrl

`func (o *DbEndorsement) HasSourceCommentUrl() bool`

HasSourceCommentUrl returns a boolean if a field has been set.

### GetSourceContextUrl

`func (o *DbEndorsement) GetSourceContextUrl() string`

GetSourceContextUrl returns the SourceContextUrl field if non-nil, zero value otherwise.

### GetSourceContextUrlOk

`func (o *DbEndorsement) GetSourceContextUrlOk() (*string, bool)`

GetSourceContextUrlOk returns a tuple with the SourceContextUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceContextUrl

`func (o *DbEndorsement) SetSourceContextUrl(v string)`

SetSourceContextUrl sets SourceContextUrl field to given value.


### GetType

`func (o *DbEndorsement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DbEndorsement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DbEndorsement) SetType(v string)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *DbEndorsement) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbEndorsement) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbEndorsement) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbEndorsement) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


