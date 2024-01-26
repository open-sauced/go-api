# DbContributionStatTimeframe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** | The ISO timestamp for the given time bucket | 
**Commits** | **int32** | Number of commits within the time range | 
**PrsCreated** | **int32** | Number of PRs created within the time range | 
**PrsReviewed** | **int32** | Number of PRs reviewed within the time range | 
**IssuesCreated** | **int32** | Number of issues within the time range | 
**CommitComments** | **int32** | Number of commit comments within the time range | 
**IssueComments** | **int32** | Number of issue comments within the time range | 
**PrReviewComments** | **int32** | Number of pr review comments within the time range | 
**Comments** | **int32** | Number of total comments within the time range | 
**TotalContributions** | **int32** | Number of total contributions for a user login | 

## Methods

### NewDbContributionStatTimeframe

`func NewDbContributionStatTimeframe(bucket string, commits int32, prsCreated int32, prsReviewed int32, issuesCreated int32, commitComments int32, issueComments int32, prReviewComments int32, comments int32, totalContributions int32, ) *DbContributionStatTimeframe`

NewDbContributionStatTimeframe instantiates a new DbContributionStatTimeframe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbContributionStatTimeframeWithDefaults

`func NewDbContributionStatTimeframeWithDefaults() *DbContributionStatTimeframe`

NewDbContributionStatTimeframeWithDefaults instantiates a new DbContributionStatTimeframe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *DbContributionStatTimeframe) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *DbContributionStatTimeframe) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *DbContributionStatTimeframe) SetBucket(v string)`

SetBucket sets Bucket field to given value.


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


### GetCommitComments

`func (o *DbContributionStatTimeframe) GetCommitComments() int32`

GetCommitComments returns the CommitComments field if non-nil, zero value otherwise.

### GetCommitCommentsOk

`func (o *DbContributionStatTimeframe) GetCommitCommentsOk() (*int32, bool)`

GetCommitCommentsOk returns a tuple with the CommitComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitComments

`func (o *DbContributionStatTimeframe) SetCommitComments(v int32)`

SetCommitComments sets CommitComments field to given value.


### GetIssueComments

`func (o *DbContributionStatTimeframe) GetIssueComments() int32`

GetIssueComments returns the IssueComments field if non-nil, zero value otherwise.

### GetIssueCommentsOk

`func (o *DbContributionStatTimeframe) GetIssueCommentsOk() (*int32, bool)`

GetIssueCommentsOk returns a tuple with the IssueComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueComments

`func (o *DbContributionStatTimeframe) SetIssueComments(v int32)`

SetIssueComments sets IssueComments field to given value.


### GetPrReviewComments

`func (o *DbContributionStatTimeframe) GetPrReviewComments() int32`

GetPrReviewComments returns the PrReviewComments field if non-nil, zero value otherwise.

### GetPrReviewCommentsOk

`func (o *DbContributionStatTimeframe) GetPrReviewCommentsOk() (*int32, bool)`

GetPrReviewCommentsOk returns a tuple with the PrReviewComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrReviewComments

`func (o *DbContributionStatTimeframe) SetPrReviewComments(v int32)`

SetPrReviewComments sets PrReviewComments field to given value.


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


### GetTotalContributions

`func (o *DbContributionStatTimeframe) GetTotalContributions() int32`

GetTotalContributions returns the TotalContributions field if non-nil, zero value otherwise.

### GetTotalContributionsOk

`func (o *DbContributionStatTimeframe) GetTotalContributionsOk() (*int32, bool)`

GetTotalContributionsOk returns a tuple with the TotalContributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalContributions

`func (o *DbContributionStatTimeframe) SetTotalContributions(v int32)`

SetTotalContributions sets TotalContributions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


