# ListAllBakedRepos200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbBakedRepo**](DbBakedRepo.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewListAllBakedRepos200Response

`func NewListAllBakedRepos200Response(data []DbBakedRepo, meta PageMetaDto, ) *ListAllBakedRepos200Response`

NewListAllBakedRepos200Response instantiates a new ListAllBakedRepos200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllBakedRepos200ResponseWithDefaults

`func NewListAllBakedRepos200ResponseWithDefaults() *ListAllBakedRepos200Response`

NewListAllBakedRepos200ResponseWithDefaults instantiates a new ListAllBakedRepos200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAllBakedRepos200Response) GetData() []DbBakedRepo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAllBakedRepos200Response) GetDataOk() (*[]DbBakedRepo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAllBakedRepos200Response) SetData(v []DbBakedRepo)`

SetData sets Data field to given value.


### GetMeta

`func (o *ListAllBakedRepos200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAllBakedRepos200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAllBakedRepos200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


