# FindAllTopReposByUsername200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbRepo**](DbRepo.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewFindAllTopReposByUsername200Response

`func NewFindAllTopReposByUsername200Response(data []DbRepo, meta PageMetaDto, ) *FindAllTopReposByUsername200Response`

NewFindAllTopReposByUsername200Response instantiates a new FindAllTopReposByUsername200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindAllTopReposByUsername200ResponseWithDefaults

`func NewFindAllTopReposByUsername200ResponseWithDefaults() *FindAllTopReposByUsername200Response`

NewFindAllTopReposByUsername200ResponseWithDefaults instantiates a new FindAllTopReposByUsername200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindAllTopReposByUsername200Response) GetData() []DbRepo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindAllTopReposByUsername200Response) GetDataOk() (*[]DbRepo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindAllTopReposByUsername200Response) SetData(v []DbRepo)`

SetData sets Data field to given value.


### GetMeta

`func (o *FindAllTopReposByUsername200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindAllTopReposByUsername200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindAllTopReposByUsername200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


