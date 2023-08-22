# FindAllEmojis200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbEmoji**](DbEmoji.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewFindAllEmojis200Response

`func NewFindAllEmojis200Response(data []DbEmoji, meta PageMetaDto, ) *FindAllEmojis200Response`

NewFindAllEmojis200Response instantiates a new FindAllEmojis200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindAllEmojis200ResponseWithDefaults

`func NewFindAllEmojis200ResponseWithDefaults() *FindAllEmojis200Response`

NewFindAllEmojis200ResponseWithDefaults instantiates a new FindAllEmojis200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindAllEmojis200Response) GetData() []DbEmoji`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindAllEmojis200Response) GetDataOk() (*[]DbEmoji, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindAllEmojis200Response) SetData(v []DbEmoji)`

SetData sets Data field to given value.


### GetMeta

`func (o *FindAllEmojis200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindAllEmojis200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindAllEmojis200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


