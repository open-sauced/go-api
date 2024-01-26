# GetUserListContributorHighlightedRepos200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbUserHighlightRepo**](DbUserHighlightRepo.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewGetUserListContributorHighlightedRepos200Response

`func NewGetUserListContributorHighlightedRepos200Response(data []DbUserHighlightRepo, meta PageMetaDto, ) *GetUserListContributorHighlightedRepos200Response`

NewGetUserListContributorHighlightedRepos200Response instantiates a new GetUserListContributorHighlightedRepos200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserListContributorHighlightedRepos200ResponseWithDefaults

`func NewGetUserListContributorHighlightedRepos200ResponseWithDefaults() *GetUserListContributorHighlightedRepos200Response`

NewGetUserListContributorHighlightedRepos200ResponseWithDefaults instantiates a new GetUserListContributorHighlightedRepos200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetUserListContributorHighlightedRepos200Response) GetData() []DbUserHighlightRepo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserListContributorHighlightedRepos200Response) GetDataOk() (*[]DbUserHighlightRepo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserListContributorHighlightedRepos200Response) SetData(v []DbUserHighlightRepo)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetUserListContributorHighlightedRepos200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetUserListContributorHighlightedRepos200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetUserListContributorHighlightedRepos200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


