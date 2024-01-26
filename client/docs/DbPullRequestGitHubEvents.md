# DbPullRequestGitHubEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **int32** | Pull request event identifier | 
**PrNumber** | **int32** | Pull request number | 
**PrState** | **string** | Pull request state | 
**PrIsDraft** | **bool** | Pull request is draft | 
**PrIsMerged** | **bool** | Pull request is merged | 
**PrMergeableState** | **string** | Pull request mergeable state | 
**PrIsRebaseable** | **bool** | Pull request is rebaseable | 
**PrTitle** | **string** | Pull request title | 
**PrHeadLabel** | **string** | Pull request source ref | 
**PrBaseLabel** | **string** | Pull request target ref | 
**PrHeadRef** | **string** | Pull request source branch | 
**PrBaseRef** | **string** | Pull request target branch | 
**PrAuthorLogin** | **string** | Pull request author username | 
**PrCreatedAt** | Pointer to **time.Time** | Timestamp representing pr creation date | [optional] 
**PrClosedAt** | **time.Time** | Timestamp representing pr close date | 
**PrMergedAt** | **time.Time** | Timestamp representing pr merge date | 
**PrUpdatedAt** | **time.Time** | Timestamp representing repository last update | 
**PrComments** | **int32** | PR comments | 
**PrAdditions** | **int32** | PR lines added | 
**PrDeletions** | **int32** | PR lines deleted | 
**PrChangedFiles** | **int32** | PR files changed | 
**RepoName** | **string** | Pull request repo full name | 
**PrCommits** | **int32** | Number of commits in the PR | 

## Methods

### NewDbPullRequestGitHubEvents

`func NewDbPullRequestGitHubEvents(eventId int32, prNumber int32, prState string, prIsDraft bool, prIsMerged bool, prMergeableState string, prIsRebaseable bool, prTitle string, prHeadLabel string, prBaseLabel string, prHeadRef string, prBaseRef string, prAuthorLogin string, prClosedAt time.Time, prMergedAt time.Time, prUpdatedAt time.Time, prComments int32, prAdditions int32, prDeletions int32, prChangedFiles int32, repoName string, prCommits int32, ) *DbPullRequestGitHubEvents`

NewDbPullRequestGitHubEvents instantiates a new DbPullRequestGitHubEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbPullRequestGitHubEventsWithDefaults

`func NewDbPullRequestGitHubEventsWithDefaults() *DbPullRequestGitHubEvents`

NewDbPullRequestGitHubEventsWithDefaults instantiates a new DbPullRequestGitHubEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *DbPullRequestGitHubEvents) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *DbPullRequestGitHubEvents) GetEventIdOk() (*int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *DbPullRequestGitHubEvents) SetEventId(v int32)`

SetEventId sets EventId field to given value.


### GetPrNumber

`func (o *DbPullRequestGitHubEvents) GetPrNumber() int32`

GetPrNumber returns the PrNumber field if non-nil, zero value otherwise.

### GetPrNumberOk

`func (o *DbPullRequestGitHubEvents) GetPrNumberOk() (*int32, bool)`

GetPrNumberOk returns a tuple with the PrNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrNumber

`func (o *DbPullRequestGitHubEvents) SetPrNumber(v int32)`

SetPrNumber sets PrNumber field to given value.


### GetPrState

`func (o *DbPullRequestGitHubEvents) GetPrState() string`

GetPrState returns the PrState field if non-nil, zero value otherwise.

### GetPrStateOk

`func (o *DbPullRequestGitHubEvents) GetPrStateOk() (*string, bool)`

GetPrStateOk returns a tuple with the PrState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrState

`func (o *DbPullRequestGitHubEvents) SetPrState(v string)`

SetPrState sets PrState field to given value.


### GetPrIsDraft

`func (o *DbPullRequestGitHubEvents) GetPrIsDraft() bool`

GetPrIsDraft returns the PrIsDraft field if non-nil, zero value otherwise.

### GetPrIsDraftOk

`func (o *DbPullRequestGitHubEvents) GetPrIsDraftOk() (*bool, bool)`

GetPrIsDraftOk returns a tuple with the PrIsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrIsDraft

`func (o *DbPullRequestGitHubEvents) SetPrIsDraft(v bool)`

SetPrIsDraft sets PrIsDraft field to given value.


### GetPrIsMerged

`func (o *DbPullRequestGitHubEvents) GetPrIsMerged() bool`

GetPrIsMerged returns the PrIsMerged field if non-nil, zero value otherwise.

### GetPrIsMergedOk

`func (o *DbPullRequestGitHubEvents) GetPrIsMergedOk() (*bool, bool)`

GetPrIsMergedOk returns a tuple with the PrIsMerged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrIsMerged

`func (o *DbPullRequestGitHubEvents) SetPrIsMerged(v bool)`

SetPrIsMerged sets PrIsMerged field to given value.


### GetPrMergeableState

`func (o *DbPullRequestGitHubEvents) GetPrMergeableState() string`

GetPrMergeableState returns the PrMergeableState field if non-nil, zero value otherwise.

### GetPrMergeableStateOk

`func (o *DbPullRequestGitHubEvents) GetPrMergeableStateOk() (*string, bool)`

GetPrMergeableStateOk returns a tuple with the PrMergeableState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrMergeableState

`func (o *DbPullRequestGitHubEvents) SetPrMergeableState(v string)`

SetPrMergeableState sets PrMergeableState field to given value.


### GetPrIsRebaseable

`func (o *DbPullRequestGitHubEvents) GetPrIsRebaseable() bool`

GetPrIsRebaseable returns the PrIsRebaseable field if non-nil, zero value otherwise.

### GetPrIsRebaseableOk

`func (o *DbPullRequestGitHubEvents) GetPrIsRebaseableOk() (*bool, bool)`

GetPrIsRebaseableOk returns a tuple with the PrIsRebaseable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrIsRebaseable

`func (o *DbPullRequestGitHubEvents) SetPrIsRebaseable(v bool)`

SetPrIsRebaseable sets PrIsRebaseable field to given value.


### GetPrTitle

`func (o *DbPullRequestGitHubEvents) GetPrTitle() string`

GetPrTitle returns the PrTitle field if non-nil, zero value otherwise.

### GetPrTitleOk

`func (o *DbPullRequestGitHubEvents) GetPrTitleOk() (*string, bool)`

GetPrTitleOk returns a tuple with the PrTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrTitle

`func (o *DbPullRequestGitHubEvents) SetPrTitle(v string)`

SetPrTitle sets PrTitle field to given value.


### GetPrHeadLabel

`func (o *DbPullRequestGitHubEvents) GetPrHeadLabel() string`

GetPrHeadLabel returns the PrHeadLabel field if non-nil, zero value otherwise.

### GetPrHeadLabelOk

`func (o *DbPullRequestGitHubEvents) GetPrHeadLabelOk() (*string, bool)`

GetPrHeadLabelOk returns a tuple with the PrHeadLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrHeadLabel

`func (o *DbPullRequestGitHubEvents) SetPrHeadLabel(v string)`

SetPrHeadLabel sets PrHeadLabel field to given value.


### GetPrBaseLabel

`func (o *DbPullRequestGitHubEvents) GetPrBaseLabel() string`

GetPrBaseLabel returns the PrBaseLabel field if non-nil, zero value otherwise.

### GetPrBaseLabelOk

`func (o *DbPullRequestGitHubEvents) GetPrBaseLabelOk() (*string, bool)`

GetPrBaseLabelOk returns a tuple with the PrBaseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrBaseLabel

`func (o *DbPullRequestGitHubEvents) SetPrBaseLabel(v string)`

SetPrBaseLabel sets PrBaseLabel field to given value.


### GetPrHeadRef

`func (o *DbPullRequestGitHubEvents) GetPrHeadRef() string`

GetPrHeadRef returns the PrHeadRef field if non-nil, zero value otherwise.

### GetPrHeadRefOk

`func (o *DbPullRequestGitHubEvents) GetPrHeadRefOk() (*string, bool)`

GetPrHeadRefOk returns a tuple with the PrHeadRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrHeadRef

`func (o *DbPullRequestGitHubEvents) SetPrHeadRef(v string)`

SetPrHeadRef sets PrHeadRef field to given value.


### GetPrBaseRef

`func (o *DbPullRequestGitHubEvents) GetPrBaseRef() string`

GetPrBaseRef returns the PrBaseRef field if non-nil, zero value otherwise.

### GetPrBaseRefOk

`func (o *DbPullRequestGitHubEvents) GetPrBaseRefOk() (*string, bool)`

GetPrBaseRefOk returns a tuple with the PrBaseRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrBaseRef

`func (o *DbPullRequestGitHubEvents) SetPrBaseRef(v string)`

SetPrBaseRef sets PrBaseRef field to given value.


### GetPrAuthorLogin

`func (o *DbPullRequestGitHubEvents) GetPrAuthorLogin() string`

GetPrAuthorLogin returns the PrAuthorLogin field if non-nil, zero value otherwise.

### GetPrAuthorLoginOk

`func (o *DbPullRequestGitHubEvents) GetPrAuthorLoginOk() (*string, bool)`

GetPrAuthorLoginOk returns a tuple with the PrAuthorLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrAuthorLogin

`func (o *DbPullRequestGitHubEvents) SetPrAuthorLogin(v string)`

SetPrAuthorLogin sets PrAuthorLogin field to given value.


### GetPrCreatedAt

`func (o *DbPullRequestGitHubEvents) GetPrCreatedAt() time.Time`

GetPrCreatedAt returns the PrCreatedAt field if non-nil, zero value otherwise.

### GetPrCreatedAtOk

`func (o *DbPullRequestGitHubEvents) GetPrCreatedAtOk() (*time.Time, bool)`

GetPrCreatedAtOk returns a tuple with the PrCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrCreatedAt

`func (o *DbPullRequestGitHubEvents) SetPrCreatedAt(v time.Time)`

SetPrCreatedAt sets PrCreatedAt field to given value.

### HasPrCreatedAt

`func (o *DbPullRequestGitHubEvents) HasPrCreatedAt() bool`

HasPrCreatedAt returns a boolean if a field has been set.

### GetPrClosedAt

`func (o *DbPullRequestGitHubEvents) GetPrClosedAt() time.Time`

GetPrClosedAt returns the PrClosedAt field if non-nil, zero value otherwise.

### GetPrClosedAtOk

`func (o *DbPullRequestGitHubEvents) GetPrClosedAtOk() (*time.Time, bool)`

GetPrClosedAtOk returns a tuple with the PrClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrClosedAt

`func (o *DbPullRequestGitHubEvents) SetPrClosedAt(v time.Time)`

SetPrClosedAt sets PrClosedAt field to given value.


### GetPrMergedAt

`func (o *DbPullRequestGitHubEvents) GetPrMergedAt() time.Time`

GetPrMergedAt returns the PrMergedAt field if non-nil, zero value otherwise.

### GetPrMergedAtOk

`func (o *DbPullRequestGitHubEvents) GetPrMergedAtOk() (*time.Time, bool)`

GetPrMergedAtOk returns a tuple with the PrMergedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrMergedAt

`func (o *DbPullRequestGitHubEvents) SetPrMergedAt(v time.Time)`

SetPrMergedAt sets PrMergedAt field to given value.


### GetPrUpdatedAt

`func (o *DbPullRequestGitHubEvents) GetPrUpdatedAt() time.Time`

GetPrUpdatedAt returns the PrUpdatedAt field if non-nil, zero value otherwise.

### GetPrUpdatedAtOk

`func (o *DbPullRequestGitHubEvents) GetPrUpdatedAtOk() (*time.Time, bool)`

GetPrUpdatedAtOk returns a tuple with the PrUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrUpdatedAt

`func (o *DbPullRequestGitHubEvents) SetPrUpdatedAt(v time.Time)`

SetPrUpdatedAt sets PrUpdatedAt field to given value.


### GetPrComments

`func (o *DbPullRequestGitHubEvents) GetPrComments() int32`

GetPrComments returns the PrComments field if non-nil, zero value otherwise.

### GetPrCommentsOk

`func (o *DbPullRequestGitHubEvents) GetPrCommentsOk() (*int32, bool)`

GetPrCommentsOk returns a tuple with the PrComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrComments

`func (o *DbPullRequestGitHubEvents) SetPrComments(v int32)`

SetPrComments sets PrComments field to given value.


### GetPrAdditions

`func (o *DbPullRequestGitHubEvents) GetPrAdditions() int32`

GetPrAdditions returns the PrAdditions field if non-nil, zero value otherwise.

### GetPrAdditionsOk

`func (o *DbPullRequestGitHubEvents) GetPrAdditionsOk() (*int32, bool)`

GetPrAdditionsOk returns a tuple with the PrAdditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrAdditions

`func (o *DbPullRequestGitHubEvents) SetPrAdditions(v int32)`

SetPrAdditions sets PrAdditions field to given value.


### GetPrDeletions

`func (o *DbPullRequestGitHubEvents) GetPrDeletions() int32`

GetPrDeletions returns the PrDeletions field if non-nil, zero value otherwise.

### GetPrDeletionsOk

`func (o *DbPullRequestGitHubEvents) GetPrDeletionsOk() (*int32, bool)`

GetPrDeletionsOk returns a tuple with the PrDeletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrDeletions

`func (o *DbPullRequestGitHubEvents) SetPrDeletions(v int32)`

SetPrDeletions sets PrDeletions field to given value.


### GetPrChangedFiles

`func (o *DbPullRequestGitHubEvents) GetPrChangedFiles() int32`

GetPrChangedFiles returns the PrChangedFiles field if non-nil, zero value otherwise.

### GetPrChangedFilesOk

`func (o *DbPullRequestGitHubEvents) GetPrChangedFilesOk() (*int32, bool)`

GetPrChangedFilesOk returns a tuple with the PrChangedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrChangedFiles

`func (o *DbPullRequestGitHubEvents) SetPrChangedFiles(v int32)`

SetPrChangedFiles sets PrChangedFiles field to given value.


### GetRepoName

`func (o *DbPullRequestGitHubEvents) GetRepoName() string`

GetRepoName returns the RepoName field if non-nil, zero value otherwise.

### GetRepoNameOk

`func (o *DbPullRequestGitHubEvents) GetRepoNameOk() (*string, bool)`

GetRepoNameOk returns a tuple with the RepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoName

`func (o *DbPullRequestGitHubEvents) SetRepoName(v string)`

SetRepoName sets RepoName field to given value.


### GetPrCommits

`func (o *DbPullRequestGitHubEvents) GetPrCommits() int32`

GetPrCommits returns the PrCommits field if non-nil, zero value otherwise.

### GetPrCommitsOk

`func (o *DbPullRequestGitHubEvents) GetPrCommitsOk() (*int32, bool)`

GetPrCommitsOk returns a tuple with the PrCommits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrCommits

`func (o *DbPullRequestGitHubEvents) SetPrCommits(v int32)`

SetPrCommits sets PrCommits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


