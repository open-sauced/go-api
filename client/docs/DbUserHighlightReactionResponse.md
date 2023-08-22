# DbUserHighlightReactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmojiId** | **string** | Emoji identifier | 
**ReactionCount** | Pointer to **float32** |  | [optional] 

## Methods

### NewDbUserHighlightReactionResponse

`func NewDbUserHighlightReactionResponse(emojiId string, ) *DbUserHighlightReactionResponse`

NewDbUserHighlightReactionResponse instantiates a new DbUserHighlightReactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbUserHighlightReactionResponseWithDefaults

`func NewDbUserHighlightReactionResponseWithDefaults() *DbUserHighlightReactionResponse`

NewDbUserHighlightReactionResponseWithDefaults instantiates a new DbUserHighlightReactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmojiId

`func (o *DbUserHighlightReactionResponse) GetEmojiId() string`

GetEmojiId returns the EmojiId field if non-nil, zero value otherwise.

### GetEmojiIdOk

`func (o *DbUserHighlightReactionResponse) GetEmojiIdOk() (*string, bool)`

GetEmojiIdOk returns a tuple with the EmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiId

`func (o *DbUserHighlightReactionResponse) SetEmojiId(v string)`

SetEmojiId sets EmojiId field to given value.


### GetReactionCount

`func (o *DbUserHighlightReactionResponse) GetReactionCount() float32`

GetReactionCount returns the ReactionCount field if non-nil, zero value otherwise.

### GetReactionCountOk

`func (o *DbUserHighlightReactionResponse) GetReactionCountOk() (*float32, bool)`

GetReactionCountOk returns a tuple with the ReactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionCount

`func (o *DbUserHighlightReactionResponse) SetReactionCount(v float32)`

SetReactionCount sets ReactionCount field to given value.

### HasReactionCount

`func (o *DbUserHighlightReactionResponse) HasReactionCount() bool`

HasReactionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


