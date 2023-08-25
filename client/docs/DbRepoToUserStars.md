# DbRepoToUserStars

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Star identifier | 
**UserId** | **int32** | User identifier | 
**RepoId** | **int32** | Repository identifier | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing star creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing star last update | [optional] 

## Methods

### NewDbRepoToUserStars

`func NewDbRepoToUserStars(id int32, userId int32, repoId int32, ) *DbRepoToUserStars`

NewDbRepoToUserStars instantiates a new DbRepoToUserStars object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbRepoToUserStarsWithDefaults

`func NewDbRepoToUserStarsWithDefaults() *DbRepoToUserStars`

NewDbRepoToUserStarsWithDefaults instantiates a new DbRepoToUserStars object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbRepoToUserStars) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbRepoToUserStars) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbRepoToUserStars) SetId(v int32)`

SetId sets Id field to given value.


### GetUserId

`func (o *DbRepoToUserStars) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DbRepoToUserStars) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DbRepoToUserStars) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetRepoId

`func (o *DbRepoToUserStars) GetRepoId() int32`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *DbRepoToUserStars) GetRepoIdOk() (*int32, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *DbRepoToUserStars) SetRepoId(v int32)`

SetRepoId sets RepoId field to given value.


### GetCreatedAt

`func (o *DbRepoToUserStars) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbRepoToUserStars) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbRepoToUserStars) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbRepoToUserStars) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbRepoToUserStars) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbRepoToUserStars) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbRepoToUserStars) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbRepoToUserStars) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


