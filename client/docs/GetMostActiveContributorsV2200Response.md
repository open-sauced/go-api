# GetMostActiveContributorsV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbUserListContributorStat**](DbUserListContributorStat.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewGetMostActiveContributorsV2200Response

`func NewGetMostActiveContributorsV2200Response(data []DbUserListContributorStat, meta PageMetaDto, ) *GetMostActiveContributorsV2200Response`

NewGetMostActiveContributorsV2200Response instantiates a new GetMostActiveContributorsV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMostActiveContributorsV2200ResponseWithDefaults

`func NewGetMostActiveContributorsV2200ResponseWithDefaults() *GetMostActiveContributorsV2200Response`

NewGetMostActiveContributorsV2200ResponseWithDefaults instantiates a new GetMostActiveContributorsV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetMostActiveContributorsV2200Response) GetData() []DbUserListContributorStat`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMostActiveContributorsV2200Response) GetDataOk() (*[]DbUserListContributorStat, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMostActiveContributorsV2200Response) SetData(v []DbUserListContributorStat)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetMostActiveContributorsV2200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetMostActiveContributorsV2200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetMostActiveContributorsV2200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


