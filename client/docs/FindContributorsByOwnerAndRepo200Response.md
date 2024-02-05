# FindContributorsByOwnerAndRepo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbRepoContributor**](DbRepoContributor.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewFindContributorsByOwnerAndRepo200Response

`func NewFindContributorsByOwnerAndRepo200Response(data []DbRepoContributor, meta PageMetaDto, ) *FindContributorsByOwnerAndRepo200Response`

NewFindContributorsByOwnerAndRepo200Response instantiates a new FindContributorsByOwnerAndRepo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindContributorsByOwnerAndRepo200ResponseWithDefaults

`func NewFindContributorsByOwnerAndRepo200ResponseWithDefaults() *FindContributorsByOwnerAndRepo200Response`

NewFindContributorsByOwnerAndRepo200ResponseWithDefaults instantiates a new FindContributorsByOwnerAndRepo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindContributorsByOwnerAndRepo200Response) GetData() []DbRepoContributor`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindContributorsByOwnerAndRepo200Response) GetDataOk() (*[]DbRepoContributor, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindContributorsByOwnerAndRepo200Response) SetData(v []DbRepoContributor)`

SetData sets Data field to given value.


### GetMeta

`func (o *FindContributorsByOwnerAndRepo200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindContributorsByOwnerAndRepo200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindContributorsByOwnerAndRepo200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


