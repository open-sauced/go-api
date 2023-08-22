# FindContributorPullRequests200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbPullRequest**](DbPullRequest.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewFindContributorPullRequests200Response

`func NewFindContributorPullRequests200Response(data []DbPullRequest, meta PageMetaDto, ) *FindContributorPullRequests200Response`

NewFindContributorPullRequests200Response instantiates a new FindContributorPullRequests200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindContributorPullRequests200ResponseWithDefaults

`func NewFindContributorPullRequests200ResponseWithDefaults() *FindContributorPullRequests200Response`

NewFindContributorPullRequests200ResponseWithDefaults instantiates a new FindContributorPullRequests200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindContributorPullRequests200Response) GetData() []DbPullRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindContributorPullRequests200Response) GetDataOk() (*[]DbPullRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindContributorPullRequests200Response) SetData(v []DbPullRequest)`

SetData sets Data field to given value.


### GetMeta

`func (o *FindContributorPullRequests200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindContributorPullRequests200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindContributorPullRequests200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


