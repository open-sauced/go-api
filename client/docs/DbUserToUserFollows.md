# DbUserToUserFollows

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | User Follow identifier | 
**UserId** | **int32** | User identifier | 
**FollowingUserId** | **int32** | User follower identifier | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing user follow creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing user follow last update | [optional] 

## Methods

### NewDbUserToUserFollows

`func NewDbUserToUserFollows(id int32, userId int32, followingUserId int32, ) *DbUserToUserFollows`

NewDbUserToUserFollows instantiates a new DbUserToUserFollows object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbUserToUserFollowsWithDefaults

`func NewDbUserToUserFollowsWithDefaults() *DbUserToUserFollows`

NewDbUserToUserFollowsWithDefaults instantiates a new DbUserToUserFollows object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbUserToUserFollows) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbUserToUserFollows) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbUserToUserFollows) SetId(v int32)`

SetId sets Id field to given value.


### GetUserId

`func (o *DbUserToUserFollows) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DbUserToUserFollows) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DbUserToUserFollows) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetFollowingUserId

`func (o *DbUserToUserFollows) GetFollowingUserId() int32`

GetFollowingUserId returns the FollowingUserId field if non-nil, zero value otherwise.

### GetFollowingUserIdOk

`func (o *DbUserToUserFollows) GetFollowingUserIdOk() (*int32, bool)`

GetFollowingUserIdOk returns a tuple with the FollowingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingUserId

`func (o *DbUserToUserFollows) SetFollowingUserId(v int32)`

SetFollowingUserId sets FollowingUserId field to given value.


### GetCreatedAt

`func (o *DbUserToUserFollows) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbUserToUserFollows) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbUserToUserFollows) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbUserToUserFollows) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbUserToUserFollows) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbUserToUserFollows) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbUserToUserFollows) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbUserToUserFollows) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


