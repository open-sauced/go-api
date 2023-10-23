# GetMostActiveContributors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbUserListContributorStat**](DbUserListContributorStat.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewGetMostActiveContributors200Response

`func NewGetMostActiveContributors200Response(data []DbUserListContributorStat, meta PageMetaDto, ) *GetMostActiveContributors200Response`

NewGetMostActiveContributors200Response instantiates a new GetMostActiveContributors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMostActiveContributors200ResponseWithDefaults

`func NewGetMostActiveContributors200ResponseWithDefaults() *GetMostActiveContributors200Response`

NewGetMostActiveContributors200ResponseWithDefaults instantiates a new GetMostActiveContributors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetMostActiveContributors200Response) GetData() []DbUserListContributorStat`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMostActiveContributors200Response) GetDataOk() (*[]DbUserListContributorStat, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMostActiveContributors200Response) SetData(v []DbUserListContributorStat)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetMostActiveContributors200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetMostActiveContributors200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetMostActiveContributors200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


