# GetUserListContributors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbUserListContributor**](DbUserListContributor.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewGetUserListContributors200Response

`func NewGetUserListContributors200Response(data []DbUserListContributor, meta PageMetaDto, ) *GetUserListContributors200Response`

NewGetUserListContributors200Response instantiates a new GetUserListContributors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserListContributors200ResponseWithDefaults

`func NewGetUserListContributors200ResponseWithDefaults() *GetUserListContributors200Response`

NewGetUserListContributors200ResponseWithDefaults instantiates a new GetUserListContributors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetUserListContributors200Response) GetData() []DbUserListContributor`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserListContributors200Response) GetDataOk() (*[]DbUserListContributor, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserListContributors200Response) SetData(v []DbUserListContributor)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetUserListContributors200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetUserListContributors200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetUserListContributors200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


