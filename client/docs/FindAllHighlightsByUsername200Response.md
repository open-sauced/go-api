# FindAllHighlightsByUsername200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbUserHighlight**](DbUserHighlight.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewFindAllHighlightsByUsername200Response

`func NewFindAllHighlightsByUsername200Response(data []DbUserHighlight, meta PageMetaDto, ) *FindAllHighlightsByUsername200Response`

NewFindAllHighlightsByUsername200Response instantiates a new FindAllHighlightsByUsername200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindAllHighlightsByUsername200ResponseWithDefaults

`func NewFindAllHighlightsByUsername200ResponseWithDefaults() *FindAllHighlightsByUsername200Response`

NewFindAllHighlightsByUsername200ResponseWithDefaults instantiates a new FindAllHighlightsByUsername200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindAllHighlightsByUsername200Response) GetData() []DbUserHighlight`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindAllHighlightsByUsername200Response) GetDataOk() (*[]DbUserHighlight, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindAllHighlightsByUsername200Response) SetData(v []DbUserHighlight)`

SetData sets Data field to given value.


### GetMeta

`func (o *FindAllHighlightsByUsername200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindAllHighlightsByUsername200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindAllHighlightsByUsername200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


