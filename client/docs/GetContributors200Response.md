# GetContributors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbUser**](DbUser.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewGetContributors200Response

`func NewGetContributors200Response(data []DbUser, meta PageMetaDto, ) *GetContributors200Response`

NewGetContributors200Response instantiates a new GetContributors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContributors200ResponseWithDefaults

`func NewGetContributors200ResponseWithDefaults() *GetContributors200Response`

NewGetContributors200ResponseWithDefaults instantiates a new GetContributors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetContributors200Response) GetData() []DbUser`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetContributors200Response) GetDataOk() (*[]DbUser, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetContributors200Response) SetData(v []DbUser)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetContributors200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetContributors200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetContributors200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


