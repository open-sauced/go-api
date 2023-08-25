# DbUserHighlightReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Reaction identifier | 
**HighlightId** | **int32** | Highlight identifier | 
**UserId** | **int32** | User identifier | 
**EmojiId** | **string** | Emoji identifier | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing highlight reaction creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing highlight reaction last update | [optional] 
**ReactionCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewDbUserHighlightReaction

`func NewDbUserHighlightReaction(id string, highlightId int32, userId int32, emojiId string, ) *DbUserHighlightReaction`

NewDbUserHighlightReaction instantiates a new DbUserHighlightReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbUserHighlightReactionWithDefaults

`func NewDbUserHighlightReactionWithDefaults() *DbUserHighlightReaction`

NewDbUserHighlightReactionWithDefaults instantiates a new DbUserHighlightReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbUserHighlightReaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbUserHighlightReaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbUserHighlightReaction) SetId(v string)`

SetId sets Id field to given value.


### GetHighlightId

`func (o *DbUserHighlightReaction) GetHighlightId() int32`

GetHighlightId returns the HighlightId field if non-nil, zero value otherwise.

### GetHighlightIdOk

`func (o *DbUserHighlightReaction) GetHighlightIdOk() (*int32, bool)`

GetHighlightIdOk returns a tuple with the HighlightId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightId

`func (o *DbUserHighlightReaction) SetHighlightId(v int32)`

SetHighlightId sets HighlightId field to given value.


### GetUserId

`func (o *DbUserHighlightReaction) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DbUserHighlightReaction) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DbUserHighlightReaction) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetEmojiId

`func (o *DbUserHighlightReaction) GetEmojiId() string`

GetEmojiId returns the EmojiId field if non-nil, zero value otherwise.

### GetEmojiIdOk

`func (o *DbUserHighlightReaction) GetEmojiIdOk() (*string, bool)`

GetEmojiIdOk returns a tuple with the EmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiId

`func (o *DbUserHighlightReaction) SetEmojiId(v string)`

SetEmojiId sets EmojiId field to given value.


### GetCreatedAt

`func (o *DbUserHighlightReaction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbUserHighlightReaction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbUserHighlightReaction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbUserHighlightReaction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbUserHighlightReaction) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbUserHighlightReaction) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbUserHighlightReaction) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbUserHighlightReaction) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReactionCount

`func (o *DbUserHighlightReaction) GetReactionCount() int32`

GetReactionCount returns the ReactionCount field if non-nil, zero value otherwise.

### GetReactionCountOk

`func (o *DbUserHighlightReaction) GetReactionCountOk() (*int32, bool)`

GetReactionCountOk returns a tuple with the ReactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionCount

`func (o *DbUserHighlightReaction) SetReactionCount(v int32)`

SetReactionCount sets ReactionCount field to given value.

### HasReactionCount

`func (o *DbUserHighlightReaction) HasReactionCount() bool`

HasReactionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


