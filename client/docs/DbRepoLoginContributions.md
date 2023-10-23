# DbRepoLoginContributions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **string** | User list collaborator&#39;s login | [optional] 
**Commits** | **int32** | Number of commits associated with a user login | 
**PrsCreated** | **int32** | Number of PRs created associated with a user login | 
**PrsReviewed** | **int32** | Number of PRs reviewed by a user login | 
**IssuesCreated** | **int32** | Number of issues created by a user login | 
**Comments** | **int32** | Number of comments associated with a user login | 

## Methods

### NewDbRepoLoginContributions

`func NewDbRepoLoginContributions(commits int32, prsCreated int32, prsReviewed int32, issuesCreated int32, comments int32, ) *DbRepoLoginContributions`

NewDbRepoLoginContributions instantiates a new DbRepoLoginContributions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbRepoLoginContributionsWithDefaults

`func NewDbRepoLoginContributionsWithDefaults() *DbRepoLoginContributions`

NewDbRepoLoginContributionsWithDefaults instantiates a new DbRepoLoginContributions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *DbRepoLoginContributions) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *DbRepoLoginContributions) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *DbRepoLoginContributions) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *DbRepoLoginContributions) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetCommits

`func (o *DbRepoLoginContributions) GetCommits() int32`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *DbRepoLoginContributions) GetCommitsOk() (*int32, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *DbRepoLoginContributions) SetCommits(v int32)`

SetCommits sets Commits field to given value.


### GetPrsCreated

`func (o *DbRepoLoginContributions) GetPrsCreated() int32`

GetPrsCreated returns the PrsCreated field if non-nil, zero value otherwise.

### GetPrsCreatedOk

`func (o *DbRepoLoginContributions) GetPrsCreatedOk() (*int32, bool)`

GetPrsCreatedOk returns a tuple with the PrsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrsCreated

`func (o *DbRepoLoginContributions) SetPrsCreated(v int32)`

SetPrsCreated sets PrsCreated field to given value.


### GetPrsReviewed

`func (o *DbRepoLoginContributions) GetPrsReviewed() int32`

GetPrsReviewed returns the PrsReviewed field if non-nil, zero value otherwise.

### GetPrsReviewedOk

`func (o *DbRepoLoginContributions) GetPrsReviewedOk() (*int32, bool)`

GetPrsReviewedOk returns a tuple with the PrsReviewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrsReviewed

`func (o *DbRepoLoginContributions) SetPrsReviewed(v int32)`

SetPrsReviewed sets PrsReviewed field to given value.


### GetIssuesCreated

`func (o *DbRepoLoginContributions) GetIssuesCreated() int32`

GetIssuesCreated returns the IssuesCreated field if non-nil, zero value otherwise.

### GetIssuesCreatedOk

`func (o *DbRepoLoginContributions) GetIssuesCreatedOk() (*int32, bool)`

GetIssuesCreatedOk returns a tuple with the IssuesCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesCreated

`func (o *DbRepoLoginContributions) SetIssuesCreated(v int32)`

SetIssuesCreated sets IssuesCreated field to given value.


### GetComments

`func (o *DbRepoLoginContributions) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DbRepoLoginContributions) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DbRepoLoginContributions) SetComments(v int32)`

SetComments sets Comments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


