# FindAllHighlightRepos200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbUserHighlightRepo**](DbUserHighlightRepo.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewFindAllHighlightRepos200Response

`func NewFindAllHighlightRepos200Response(data []DbUserHighlightRepo, meta PageMetaDto, ) *FindAllHighlightRepos200Response`

NewFindAllHighlightRepos200Response instantiates a new FindAllHighlightRepos200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindAllHighlightRepos200ResponseWithDefaults

`func NewFindAllHighlightRepos200ResponseWithDefaults() *FindAllHighlightRepos200Response`

NewFindAllHighlightRepos200ResponseWithDefaults instantiates a new FindAllHighlightRepos200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindAllHighlightRepos200Response) GetData() []DbUserHighlightRepo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindAllHighlightRepos200Response) GetDataOk() (*[]DbUserHighlightRepo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindAllHighlightRepos200Response) SetData(v []DbUserHighlightRepo)`

SetData sets Data field to given value.


### GetMeta

`func (o *FindAllHighlightRepos200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindAllHighlightRepos200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindAllHighlightRepos200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


