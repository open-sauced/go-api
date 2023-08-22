# SearchAllPullRequestContributors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbPullRequestContributor**](DbPullRequestContributor.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewSearchAllPullRequestContributors200Response

`func NewSearchAllPullRequestContributors200Response(data []DbPullRequestContributor, meta PageMetaDto, ) *SearchAllPullRequestContributors200Response`

NewSearchAllPullRequestContributors200Response instantiates a new SearchAllPullRequestContributors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAllPullRequestContributors200ResponseWithDefaults

`func NewSearchAllPullRequestContributors200ResponseWithDefaults() *SearchAllPullRequestContributors200Response`

NewSearchAllPullRequestContributors200ResponseWithDefaults instantiates a new SearchAllPullRequestContributors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SearchAllPullRequestContributors200Response) GetData() []DbPullRequestContributor`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SearchAllPullRequestContributors200Response) GetDataOk() (*[]DbPullRequestContributor, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SearchAllPullRequestContributors200Response) SetData(v []DbPullRequestContributor)`

SetData sets Data field to given value.


### GetMeta

`func (o *SearchAllPullRequestContributors200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SearchAllPullRequestContributors200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SearchAllPullRequestContributors200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


