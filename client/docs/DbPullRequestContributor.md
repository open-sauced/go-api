# DbPullRequestContributor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorLogin** | **string** | Pull request author username | 
**UserId** | **float32** | Pull request author id | 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing pr last update | [optional] 

## Methods

### NewDbPullRequestContributor

`func NewDbPullRequestContributor(authorLogin string, userId float32, ) *DbPullRequestContributor`

NewDbPullRequestContributor instantiates a new DbPullRequestContributor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbPullRequestContributorWithDefaults

`func NewDbPullRequestContributorWithDefaults() *DbPullRequestContributor`

NewDbPullRequestContributorWithDefaults instantiates a new DbPullRequestContributor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorLogin

`func (o *DbPullRequestContributor) GetAuthorLogin() string`

GetAuthorLogin returns the AuthorLogin field if non-nil, zero value otherwise.

### GetAuthorLoginOk

`func (o *DbPullRequestContributor) GetAuthorLoginOk() (*string, bool)`

GetAuthorLoginOk returns a tuple with the AuthorLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorLogin

`func (o *DbPullRequestContributor) SetAuthorLogin(v string)`

SetAuthorLogin sets AuthorLogin field to given value.


### GetUserId

`func (o *DbPullRequestContributor) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DbPullRequestContributor) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DbPullRequestContributor) SetUserId(v float32)`

SetUserId sets UserId field to given value.


### GetUpdatedAt

`func (o *DbPullRequestContributor) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbPullRequestContributor) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbPullRequestContributor) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbPullRequestContributor) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


