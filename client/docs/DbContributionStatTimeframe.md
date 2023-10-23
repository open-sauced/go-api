# DbContributionStatTimeframe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeStart** | **string** | The ISO timestamp for the start of the time frame | 
**TimeEnd** | **string** | The ISO timestamp for the end of the time frame | 
**Commits** | **int32** | Number of commits associated with a user login | 
**PrsCreated** | **int32** | Number of PRs created associated with a user login | 
**PrsReviewed** | **int32** | Number of PRs reviewed by a user login | 
**IssuesCreated** | **int32** | Number of issues created by a user login | 
**Comments** | **int32** | Number of comments associated with a user login | 

## Methods

### NewDbContributionStatTimeframe

`func NewDbContributionStatTimeframe(timeStart string, timeEnd string, commits int32, prsCreated int32, prsReviewed int32, issuesCreated int32, comments int32, ) *DbContributionStatTimeframe`

NewDbContributionStatTimeframe instantiates a new DbContributionStatTimeframe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbContributionStatTimeframeWithDefaults

`func NewDbContributionStatTimeframeWithDefaults() *DbContributionStatTimeframe`

NewDbContributionStatTimeframeWithDefaults instantiates a new DbContributionStatTimeframe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeStart

`func (o *DbContributionStatTimeframe) GetTimeStart() string`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *DbContributionStatTimeframe) GetTimeStartOk() (*string, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *DbContributionStatTimeframe) SetTimeStart(v string)`

SetTimeStart sets TimeStart field to given value.


### GetTimeEnd

`func (o *DbContributionStatTimeframe) GetTimeEnd() string`

GetTimeEnd returns the TimeEnd field if non-nil, zero value otherwise.

### GetTimeEndOk

`func (o *DbContributionStatTimeframe) GetTimeEndOk() (*string, bool)`

GetTimeEndOk returns a tuple with the TimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEnd

`func (o *DbContributionStatTimeframe) SetTimeEnd(v string)`

SetTimeEnd sets TimeEnd field to given value.


### GetCommits

`func (o *DbContributionStatTimeframe) GetCommits() int32`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *DbContributionStatTimeframe) GetCommitsOk() (*int32, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *DbContributionStatTimeframe) SetCommits(v int32)`

SetCommits sets Commits field to given value.


### GetPrsCreated

`func (o *DbContributionStatTimeframe) GetPrsCreated() int32`

GetPrsCreated returns the PrsCreated field if non-nil, zero value otherwise.

### GetPrsCreatedOk

`func (o *DbContributionStatTimeframe) GetPrsCreatedOk() (*int32, bool)`

GetPrsCreatedOk returns a tuple with the PrsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrsCreated

`func (o *DbContributionStatTimeframe) SetPrsCreated(v int32)`

SetPrsCreated sets PrsCreated field to given value.


### GetPrsReviewed

`func (o *DbContributionStatTimeframe) GetPrsReviewed() int32`

GetPrsReviewed returns the PrsReviewed field if non-nil, zero value otherwise.

### GetPrsReviewedOk

`func (o *DbContributionStatTimeframe) GetPrsReviewedOk() (*int32, bool)`

GetPrsReviewedOk returns a tuple with the PrsReviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrsReviewed

`func (o *DbContributionStatTimeframe) SetPrsReviewed(v int32)`

SetPrsReviewed sets PrsReviewed field to given value.


### GetIssuesCreated

`func (o *DbContributionStatTimeframe) GetIssuesCreated() int32`

GetIssuesCreated returns the IssuesCreated field if non-nil, zero value otherwise.

### GetIssuesCreatedOk

`func (o *DbContributionStatTimeframe) GetIssuesCreatedOk() (*int32, bool)`

GetIssuesCreatedOk returns a tuple with the IssuesCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesCreated

`func (o *DbContributionStatTimeframe) SetIssuesCreated(v int32)`

SetIssuesCreated sets IssuesCreated field to given value.


### GetComments

`func (o *DbContributionStatTimeframe) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DbContributionStatTimeframe) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DbContributionStatTimeframe) SetComments(v int32)`

SetComments sets Comments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


