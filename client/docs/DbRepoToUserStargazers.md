# DbRepoToUserStargazers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Stargaze identifier | 
**UserId** | **int32** | User identifier | 
**RepoId** | **int32** | Repository identifier | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing stargaze creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing stargaze last update | [optional] 

## Methods

### NewDbRepoToUserStargazers

`func NewDbRepoToUserStargazers(id int32, userId int32, repoId int32, ) *DbRepoToUserStargazers`

NewDbRepoToUserStargazers instantiates a new DbRepoToUserStargazers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbRepoToUserStargazersWithDefaults

`func NewDbRepoToUserStargazersWithDefaults() *DbRepoToUserStargazers`

NewDbRepoToUserStargazersWithDefaults instantiates a new DbRepoToUserStargazers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbRepoToUserStargazers) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbRepoToUserStargazers) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbRepoToUserStargazers) SetId(v int32)`

SetId sets Id field to given value.


### GetUserId

`func (o *DbRepoToUserStargazers) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DbRepoToUserStargazers) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DbRepoToUserStargazers) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetRepoId

`func (o *DbRepoToUserStargazers) GetRepoId() int32`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *DbRepoToUserStargazers) GetRepoIdOk() (*int32, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *DbRepoToUserStargazers) SetRepoId(v int32)`

SetRepoId sets RepoId field to given value.


### GetCreatedAt

`func (o *DbRepoToUserStargazers) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbRepoToUserStargazers) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbRepoToUserStargazers) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbRepoToUserStargazers) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbRepoToUserStargazers) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbRepoToUserStargazers) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbRepoToUserStargazers) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbRepoToUserStargazers) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


