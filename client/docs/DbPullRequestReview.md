# DbPullRequestReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Pull request review identifier | 
**ReviewerLogin** | **string** | Pull request reviewer username | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing pr review creation date | [optional] 
**PublishedAt** | Pointer to **time.Time** | Timestamp representing pr review published date | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing pr review update date | [optional] 
**State** | Pointer to **string** | Pull request review state | [optional] 
**PullRequest** | [**DbPullRequest**](DbPullRequest.md) |  | 

## Methods

### NewDbPullRequestReview

`func NewDbPullRequestReview(id int32, reviewerLogin string, pullRequest DbPullRequest, ) *DbPullRequestReview`

NewDbPullRequestReview instantiates a new DbPullRequestReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbPullRequestReviewWithDefaults

`func NewDbPullRequestReviewWithDefaults() *DbPullRequestReview`

NewDbPullRequestReviewWithDefaults instantiates a new DbPullRequestReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbPullRequestReview) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbPullRequestReview) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbPullRequestReview) SetId(v int32)`

SetId sets Id field to given value.


### GetReviewerLogin

`func (o *DbPullRequestReview) GetReviewerLogin() string`

GetReviewerLogin returns the ReviewerLogin field if non-nil, zero value otherwise.

### GetReviewerLoginOk

`func (o *DbPullRequestReview) GetReviewerLoginOk() (*string, bool)`

GetReviewerLoginOk returns a tuple with the ReviewerLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerLogin

`func (o *DbPullRequestReview) SetReviewerLogin(v string)`

SetReviewerLogin sets ReviewerLogin field to given value.


### GetCreatedAt

`func (o *DbPullRequestReview) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbPullRequestReview) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbPullRequestReview) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbPullRequestReview) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPublishedAt

`func (o *DbPullRequestReview) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *DbPullRequestReview) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *DbPullRequestReview) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *DbPullRequestReview) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbPullRequestReview) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbPullRequestReview) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbPullRequestReview) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbPullRequestReview) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetState

`func (o *DbPullRequestReview) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DbPullRequestReview) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DbPullRequestReview) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DbPullRequestReview) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPullRequest

`func (o *DbPullRequestReview) GetPullRequest() DbPullRequest`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *DbPullRequestReview) GetPullRequestOk() (*DbPullRequest, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *DbPullRequestReview) SetPullRequest(v DbPullRequest)`

SetPullRequest sets PullRequest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


