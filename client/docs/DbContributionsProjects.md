# DbContributionsProjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoName** | **string** | The org/repo name | 
**Commits** | **int32** | Number of commits within the time range | 
**PrsCreated** | **int32** | Number of PRs created for the project within the time range | 
**PrsReviewed** | **int32** | Number of PRs reviewed for the project within the time range | 
**IssuesCreated** | **int32** | Number of issues for the project within the time range | 
**CommitComments** | **int32** | Number of commit comments for the project within the time range | 
**IssueComments** | **int32** | Number of issue comments for the project within the time range | 
**PrReviewComments** | **int32** | Number of pr review comments for the project within the time range | 
**Comments** | **int32** | Number of total comments for the project within the time range | 
**TotalContributions** | **int32** | Number of total contributions for the project within the time range | 

## Methods

### NewDbContributionsProjects

`func NewDbContributionsProjects(repoName string, commits int32, prsCreated int32, prsReviewed int32, issuesCreated int32, commitComments int32, issueComments int32, prReviewComments int32, comments int32, totalContributions int32, ) *DbContributionsProjects`

NewDbContributionsProjects instantiates a new DbContributionsProjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbContributionsProjectsWithDefaults

`func NewDbContributionsProjectsWithDefaults() *DbContributionsProjects`

NewDbContributionsProjectsWithDefaults instantiates a new DbContributionsProjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepoName

`func (o *DbContributionsProjects) GetRepoName() string`

GetRepoName returns the RepoName field if non-nil, zero value otherwise.

### GetRepoNameOk

`func (o *DbContributionsProjects) GetRepoNameOk() (*string, bool)`

GetRepoNameOk returns a tuple with the RepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoName

`func (o *DbContributionsProjects) SetRepoName(v string)`

SetRepoName sets RepoName field to given value.


### GetCommits

`func (o *DbContributionsProjects) GetCommits() int32`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *DbContributionsProjects) GetCommitsOk() (*int32, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *DbContributionsProjects) SetCommits(v int32)`

SetCommits sets Commits field to given value.


### GetPrsCreated

`func (o *DbContributionsProjects) GetPrsCreated() int32`

GetPrsCreated returns the PrsCreated field if non-nil, zero value otherwise.

### GetPrsCreatedOk

`func (o *DbContributionsProjects) GetPrsCreatedOk() (*int32, bool)`

GetPrsCreatedOk returns a tuple with the PrsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrsCreated

`func (o *DbContributionsProjects) SetPrsCreated(v int32)`

SetPrsCreated sets PrsCreated field to given value.


### GetPrsReviewed

`func (o *DbContributionsProjects) GetPrsReviewed() int32`

GetPrsReviewed returns the PrsReviewed field if non-nil, zero value otherwise.

### GetPrsReviewedOk

`func (o *DbContributionsProjects) GetPrsReviewedOk() (*int32, bool)`

GetPrsReviewedOk returns a tuple with the PrsReviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrsReviewed

`func (o *DbContributionsProjects) SetPrsReviewed(v int32)`

SetPrsReviewed sets PrsReviewed field to given value.


### GetIssuesCreated

`func (o *DbContributionsProjects) GetIssuesCreated() int32`

GetIssuesCreated returns the IssuesCreated field if non-nil, zero value otherwise.

### GetIssuesCreatedOk

`func (o *DbContributionsProjects) GetIssuesCreatedOk() (*int32, bool)`

GetIssuesCreatedOk returns a tuple with the IssuesCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesCreated

`func (o *DbContributionsProjects) SetIssuesCreated(v int32)`

SetIssuesCreated sets IssuesCreated field to given value.


### GetCommitComments

`func (o *DbContributionsProjects) GetCommitComments() int32`

GetCommitComments returns the CommitComments field if non-nil, zero value otherwise.

### GetCommitCommentsOk

`func (o *DbContributionsProjects) GetCommitCommentsOk() (*int32, bool)`

GetCommitCommentsOk returns a tuple with the CommitComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitComments

`func (o *DbContributionsProjects) SetCommitComments(v int32)`

SetCommitComments sets CommitComments field to given value.


### GetIssueComments

`func (o *DbContributionsProjects) GetIssueComments() int32`

GetIssueComments returns the IssueComments field if non-nil, zero value otherwise.

### GetIssueCommentsOk

`func (o *DbContributionsProjects) GetIssueCommentsOk() (*int32, bool)`

GetIssueCommentsOk returns a tuple with the IssueComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueComments

`func (o *DbContributionsProjects) SetIssueComments(v int32)`

SetIssueComments sets IssueComments field to given value.


### GetPrReviewComments

`func (o *DbContributionsProjects) GetPrReviewComments() int32`

GetPrReviewComments returns the PrReviewComments field if non-nil, zero value otherwise.

### GetPrReviewCommentsOk

`func (o *DbContributionsProjects) GetPrReviewCommentsOk() (*int32, bool)`

GetPrReviewCommentsOk returns a tuple with the PrReviewComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrReviewComments

`func (o *DbContributionsProjects) SetPrReviewComments(v int32)`

SetPrReviewComments sets PrReviewComments field to given value.


### GetComments

`func (o *DbContributionsProjects) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DbContributionsProjects) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DbContributionsProjects) SetComments(v int32)`

SetComments sets Comments field to given value.


### GetTotalContributions

`func (o *DbContributionsProjects) GetTotalContributions() int32`

GetTotalContributions returns the TotalContributions field if non-nil, zero value otherwise.

### GetTotalContributionsOk

`func (o *DbContributionsProjects) GetTotalContributionsOk() (*int32, bool)`

GetTotalContributionsOk returns a tuple with the TotalContributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalContributions

`func (o *DbContributionsProjects) SetTotalContributions(v int32)`

SetTotalContributions sets TotalContributions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


