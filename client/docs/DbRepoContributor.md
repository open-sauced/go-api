# DbRepoContributor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | **string** | Repo contributor login | 
**Commits** | **int32** | Number of commits for login in repo within the time range | 
**PrsCreated** | **int32** | Number of PRs created for login in repo within the time range | 
**PrsReviewed** | **int32** | Number of PRs reviewed for login in repo within the time range | 
**IssuesCreated** | **int32** | Number of issues created for login in repo within the time range | 
**CommitComments** | **int32** | Number of commit comments for login in repo within the time range | 
**IssueComments** | **int32** | Number of issue comments for login in repo within the time range | 
**PrReviewComments** | **int32** | Number of pr review comments for login in repo within the time range | 
**Comments** | **int32** | Number of total comments for login in repo within the time range | 
**TotalContributions** | **int32** | Number of total contributions for a login in repo within the time range | 

## Methods

### NewDbRepoContributor

`func NewDbRepoContributor(login string, commits int32, prsCreated int32, prsReviewed int32, issuesCreated int32, commitComments int32, issueComments int32, prReviewComments int32, comments int32, totalContributions int32, ) *DbRepoContributor`

NewDbRepoContributor instantiates a new DbRepoContributor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbRepoContributorWithDefaults

`func NewDbRepoContributorWithDefaults() *DbRepoContributor`

NewDbRepoContributorWithDefaults instantiates a new DbRepoContributor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *DbRepoContributor) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *DbRepoContributor) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *DbRepoContributor) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetCommits

`func (o *DbRepoContributor) GetCommits() int32`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *DbRepoContributor) GetCommitsOk() (*int32, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *DbRepoContributor) SetCommits(v int32)`

SetCommits sets Commits field to given value.


### GetPrsCreated

`func (o *DbRepoContributor) GetPrsCreated() int32`

GetPrsCreated returns the PrsCreated field if non-nil, zero value otherwise.

### GetPrsCreatedOk

`func (o *DbRepoContributor) GetPrsCreatedOk() (*int32, bool)`

GetPrsCreatedOk returns a tuple with the PrsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrsCreated

`func (o *DbRepoContributor) SetPrsCreated(v int32)`

SetPrsCreated sets PrsCreated field to given value.


### GetPrsReviewed

`func (o *DbRepoContributor) GetPrsReviewed() int32`

GetPrsReviewed returns the PrsReviewed field if non-nil, zero value otherwise.

### GetPrsReviewedOk

`func (o *DbRepoContributor) GetPrsReviewedOk() (*int32, bool)`

GetPrsReviewedOk returns a tuple with the PrsReviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrsReviewed

`func (o *DbRepoContributor) SetPrsReviewed(v int32)`

SetPrsReviewed sets PrsReviewed field to given value.


### GetIssuesCreated

`func (o *DbRepoContributor) GetIssuesCreated() int32`

GetIssuesCreated returns the IssuesCreated field if non-nil, zero value otherwise.

### GetIssuesCreatedOk

`func (o *DbRepoContributor) GetIssuesCreatedOk() (*int32, bool)`

GetIssuesCreatedOk returns a tuple with the IssuesCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesCreated

`func (o *DbRepoContributor) SetIssuesCreated(v int32)`

SetIssuesCreated sets IssuesCreated field to given value.


### GetCommitComments

`func (o *DbRepoContributor) GetCommitComments() int32`

GetCommitComments returns the CommitComments field if non-nil, zero value otherwise.

### GetCommitCommentsOk

`func (o *DbRepoContributor) GetCommitCommentsOk() (*int32, bool)`

GetCommitCommentsOk returns a tuple with the CommitComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitComments

`func (o *DbRepoContributor) SetCommitComments(v int32)`

SetCommitComments sets CommitComments field to given value.


### GetIssueComments

`func (o *DbRepoContributor) GetIssueComments() int32`

GetIssueComments returns the IssueComments field if non-nil, zero value otherwise.

### GetIssueCommentsOk

`func (o *DbRepoContributor) GetIssueCommentsOk() (*int32, bool)`

GetIssueCommentsOk returns a tuple with the IssueComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueComments

`func (o *DbRepoContributor) SetIssueComments(v int32)`

SetIssueComments sets IssueComments field to given value.


### GetPrReviewComments

`func (o *DbRepoContributor) GetPrReviewComments() int32`

GetPrReviewComments returns the PrReviewComments field if non-nil, zero value otherwise.

### GetPrReviewCommentsOk

`func (o *DbRepoContributor) GetPrReviewCommentsOk() (*int32, bool)`

GetPrReviewCommentsOk returns a tuple with the PrReviewComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrReviewComments

`func (o *DbRepoContributor) SetPrReviewComments(v int32)`

SetPrReviewComments sets PrReviewComments field to given value.


### GetComments

`func (o *DbRepoContributor) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DbRepoContributor) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DbRepoContributor) SetComments(v int32)`

SetComments sets Comments field to given value.


### GetTotalContributions

`func (o *DbRepoContributor) GetTotalContributions() int32`

GetTotalContributions returns the TotalContributions field if non-nil, zero value otherwise.

### GetTotalContributionsOk

`func (o *DbRepoContributor) GetTotalContributionsOk() (*int32, bool)`

GetTotalContributionsOk returns a tuple with the TotalContributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalContributions

`func (o *DbRepoContributor) SetTotalContributions(v int32)`

SetTotalContributions sets TotalContributions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


