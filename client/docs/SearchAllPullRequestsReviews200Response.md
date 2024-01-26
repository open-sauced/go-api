# SearchAllPullRequestsReviews200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbPullRequestReviewGitHubEvents**](DbPullRequestReviewGitHubEvents.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewSearchAllPullRequestsReviews200Response

`func NewSearchAllPullRequestsReviews200Response(data []DbPullRequestReviewGitHubEvents, meta PageMetaDto, ) *SearchAllPullRequestsReviews200Response`

NewSearchAllPullRequestsReviews200Response instantiates a new SearchAllPullRequestsReviews200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAllPullRequestsReviews200ResponseWithDefaults

`func NewSearchAllPullRequestsReviews200ResponseWithDefaults() *SearchAllPullRequestsReviews200Response`

NewSearchAllPullRequestsReviews200ResponseWithDefaults instantiates a new SearchAllPullRequestsReviews200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SearchAllPullRequestsReviews200Response) GetData() []DbPullRequestReviewGitHubEvents`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SearchAllPullRequestsReviews200Response) GetDataOk() (*[]DbPullRequestReviewGitHubEvents, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SearchAllPullRequestsReviews200Response) SetData(v []DbPullRequestReviewGitHubEvents)`

SetData sets Data field to given value.


### GetMeta

`func (o *SearchAllPullRequestsReviews200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SearchAllPullRequestsReviews200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SearchAllPullRequestsReviews200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


