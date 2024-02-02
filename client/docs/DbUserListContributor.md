# DbUserListContributor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User list contributor identifier | 
**UserId** | **int32** | User identifier | 
**ListId** | **string** | List identifier | 
**Username** | Pointer to **string** | List user source username | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing top repo first index | [optional] 

## Methods

### NewDbUserListContributor

`func NewDbUserListContributor(id string, userId int32, listId string, ) *DbUserListContributor`

NewDbUserListContributor instantiates a new DbUserListContributor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbUserListContributorWithDefaults

`func NewDbUserListContributorWithDefaults() *DbUserListContributor`

NewDbUserListContributorWithDefaults instantiates a new DbUserListContributor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbUserListContributor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbUserListContributor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbUserListContributor) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *DbUserListContributor) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DbUserListContributor) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DbUserListContributor) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetListId

`func (o *DbUserListContributor) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *DbUserListContributor) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *DbUserListContributor) SetListId(v string)`

SetListId sets ListId field to given value.


### GetUsername

`func (o *DbUserListContributor) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DbUserListContributor) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DbUserListContributor) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DbUserListContributor) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DbUserListContributor) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbUserListContributor) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbUserListContributor) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbUserListContributor) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


