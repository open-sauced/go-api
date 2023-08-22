# FindAllByRepoId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbContribution**](DbContribution.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewFindAllByRepoId200Response

`func NewFindAllByRepoId200Response(data []DbContribution, meta PageMetaDto, ) *FindAllByRepoId200Response`

NewFindAllByRepoId200Response instantiates a new FindAllByRepoId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindAllByRepoId200ResponseWithDefaults

`func NewFindAllByRepoId200ResponseWithDefaults() *FindAllByRepoId200Response`

NewFindAllByRepoId200ResponseWithDefaults instantiates a new FindAllByRepoId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindAllByRepoId200Response) GetData() []DbContribution`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindAllByRepoId200Response) GetDataOk() (*[]DbContribution, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindAllByRepoId200Response) SetData(v []DbContribution)`

SetData sets Data field to given value.


### GetMeta

`func (o *FindAllByRepoId200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindAllByRepoId200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindAllByRepoId200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


