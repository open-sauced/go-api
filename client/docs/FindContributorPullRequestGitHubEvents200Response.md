# FindContributorPullRequestGitHubEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbPullRequestGitHubEvents**](DbPullRequestGitHubEvents.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewFindContributorPullRequestGitHubEvents200Response

`func NewFindContributorPullRequestGitHubEvents200Response(data []DbPullRequestGitHubEvents, meta PageMetaDto, ) *FindContributorPullRequestGitHubEvents200Response`

NewFindContributorPullRequestGitHubEvents200Response instantiates a new FindContributorPullRequestGitHubEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindContributorPullRequestGitHubEvents200ResponseWithDefaults

`func NewFindContributorPullRequestGitHubEvents200ResponseWithDefaults() *FindContributorPullRequestGitHubEvents200Response`

NewFindContributorPullRequestGitHubEvents200ResponseWithDefaults instantiates a new FindContributorPullRequestGitHubEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindContributorPullRequestGitHubEvents200Response) GetData() []DbPullRequestGitHubEvents`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindContributorPullRequestGitHubEvents200Response) GetDataOk() (*[]DbPullRequestGitHubEvents, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindContributorPullRequestGitHubEvents200Response) SetData(v []DbPullRequestGitHubEvents)`

SetData sets Data field to given value.


### GetMeta

`func (o *FindContributorPullRequestGitHubEvents200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindContributorPullRequestGitHubEvents200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindContributorPullRequestGitHubEvents200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


