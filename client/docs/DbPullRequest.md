# DbPullRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Pull request identifier | 
**Number** | **int32** | Pull request number | 
**State** | **string** | Pull request state | 
**Draft** | **bool** | Pull request is draft | 
**Merged** | **bool** | Pull request is merged | 
**Mergeable** | **bool** | Pull request is mergeable | 
**Rebaseable** | **bool** | Pull request is rebaseable | 
**Title** | **string** | Pull request title | 
**SourceLabel** | Pointer to **string** | Pull request source ref | [optional] 
**TargetLabel** | Pointer to **string** | Pull request target ref | [optional] 
**SourceBranch** | Pointer to **string** | Pull request source branch | [optional] 
**TargetBranch** | Pointer to **string** | Pull request target branch | [optional] 
**Labels** | Pointer to **string** | Pull request labels | [optional] 
**LabelNames** | Pointer to **[]string** | Pull request label names | [optional] 
**AuthorLogin** | **string** | Pull request author username | 
**AuthorAvatar** | **string** | Pull request author avatar | 
**AssigneeLogin** | Pointer to **string** | Pull request assignee username | [optional] 
**AssigneeAvatar** | Pointer to **string** | Pull request assignee avatar | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing pr creation date | [optional] 
**ClosedAt** | Pointer to **time.Time** | Timestamp representing pr close date | [optional] 
**MergedAt** | Pointer to **time.Time** | Timestamp representing pr merge date | [optional] 
**MergedByLogin** | Pointer to **string** | Pull request merged by username | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing repository last update | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | Timestamp representing internal last update | [optional] 
**Comments** | Pointer to **int32** | PR comments | [optional] 
**Additions** | Pointer to **int32** | PR lines added | [optional] 
**Deletions** | Pointer to **int32** | PR lines deleted | [optional] 
**ChangedFiles** | Pointer to **int32** | PR files changed | [optional] 
**FullName** | Pointer to **string** | Pull request repo full name | [optional] 
**Commits** | Pointer to **int32** | Number of commits in the PR | [optional] 

## Methods

### NewDbPullRequest

`func NewDbPullRequest(id int32, number int32, state string, draft bool, merged bool, mergeable bool, rebaseable bool, title string, authorLogin string, authorAvatar string, ) *DbPullRequest`

NewDbPullRequest instantiates a new DbPullRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbPullRequestWithDefaults

`func NewDbPullRequestWithDefaults() *DbPullRequest`

NewDbPullRequestWithDefaults instantiates a new DbPullRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbPullRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbPullRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbPullRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetNumber

`func (o *DbPullRequest) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *DbPullRequest) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *DbPullRequest) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetState

`func (o *DbPullRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DbPullRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DbPullRequest) SetState(v string)`

SetState sets State field to given value.


### GetDraft

`func (o *DbPullRequest) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *DbPullRequest) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *DbPullRequest) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetMerged

`func (o *DbPullRequest) GetMerged() bool`

GetMerged returns the Merged field if non-nil, zero value otherwise.

### GetMergedOk

`func (o *DbPullRequest) GetMergedOk() (*bool, bool)`

GetMergedOk returns a tuple with the Merged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerged

`func (o *DbPullRequest) SetMerged(v bool)`

SetMerged sets Merged field to given value.


### GetMergeable

`func (o *DbPullRequest) GetMergeable() bool`

GetMergeable returns the Mergeable field if non-nil, zero value otherwise.

### GetMergeableOk

`func (o *DbPullRequest) GetMergeableOk() (*bool, bool)`

GetMergeableOk returns a tuple with the Mergeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeable

`func (o *DbPullRequest) SetMergeable(v bool)`

SetMergeable sets Mergeable field to given value.


### GetRebaseable

`func (o *DbPullRequest) GetRebaseable() bool`

GetRebaseable returns the Rebaseable field if non-nil, zero value otherwise.

### GetRebaseableOk

`func (o *DbPullRequest) GetRebaseableOk() (*bool, bool)`

GetRebaseableOk returns a tuple with the Rebaseable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebaseable

`func (o *DbPullRequest) SetRebaseable(v bool)`

SetRebaseable sets Rebaseable field to given value.


### GetTitle

`func (o *DbPullRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DbPullRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DbPullRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSourceLabel

`func (o *DbPullRequest) GetSourceLabel() string`

GetSourceLabel returns the SourceLabel field if non-nil, zero value otherwise.

### GetSourceLabelOk

`func (o *DbPullRequest) GetSourceLabelOk() (*string, bool)`

GetSourceLabelOk returns a tuple with the SourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLabel

`func (o *DbPullRequest) SetSourceLabel(v string)`

SetSourceLabel sets SourceLabel field to given value.

### HasSourceLabel

`func (o *DbPullRequest) HasSourceLabel() bool`

HasSourceLabel returns a boolean if a field has been set.

### GetTargetLabel

`func (o *DbPullRequest) GetTargetLabel() string`

GetTargetLabel returns the TargetLabel field if non-nil, zero value otherwise.

### GetTargetLabelOk

`func (o *DbPullRequest) GetTargetLabelOk() (*string, bool)`

GetTargetLabelOk returns a tuple with the TargetLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLabel

`func (o *DbPullRequest) SetTargetLabel(v string)`

SetTargetLabel sets TargetLabel field to given value.

### HasTargetLabel

`func (o *DbPullRequest) HasTargetLabel() bool`

HasTargetLabel returns a boolean if a field has been set.

### GetSourceBranch

`func (o *DbPullRequest) GetSourceBranch() string`

GetSourceBranch returns the SourceBranch field if non-nil, zero value otherwise.

### GetSourceBranchOk

`func (o *DbPullRequest) GetSourceBranchOk() (*string, bool)`

GetSourceBranchOk returns a tuple with the SourceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBranch

`func (o *DbPullRequest) SetSourceBranch(v string)`

SetSourceBranch sets SourceBranch field to given value.

### HasSourceBranch

`func (o *DbPullRequest) HasSourceBranch() bool`

HasSourceBranch returns a boolean if a field has been set.

### GetTargetBranch

`func (o *DbPullRequest) GetTargetBranch() string`

GetTargetBranch returns the TargetBranch field if non-nil, zero value otherwise.

### GetTargetBranchOk

`func (o *DbPullRequest) GetTargetBranchOk() (*string, bool)`

GetTargetBranchOk returns a tuple with the TargetBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBranch

`func (o *DbPullRequest) SetTargetBranch(v string)`

SetTargetBranch sets TargetBranch field to given value.

### HasTargetBranch

`func (o *DbPullRequest) HasTargetBranch() bool`

HasTargetBranch returns a boolean if a field has been set.

### GetLabels

`func (o *DbPullRequest) GetLabels() string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *DbPullRequest) GetLabelsOk() (*string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *DbPullRequest) SetLabels(v string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *DbPullRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLabelNames

`func (o *DbPullRequest) GetLabelNames() []string`

GetLabelNames returns the LabelNames field if non-nil, zero value otherwise.

### GetLabelNamesOk

`func (o *DbPullRequest) GetLabelNamesOk() (*[]string, bool)`

GetLabelNamesOk returns a tuple with the LabelNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelNames

`func (o *DbPullRequest) SetLabelNames(v []string)`

SetLabelNames sets LabelNames field to given value.

### HasLabelNames

`func (o *DbPullRequest) HasLabelNames() bool`

HasLabelNames returns a boolean if a field has been set.

### GetAuthorLogin

`func (o *DbPullRequest) GetAuthorLogin() string`

GetAuthorLogin returns the AuthorLogin field if non-nil, zero value otherwise.

### GetAuthorLoginOk

`func (o *DbPullRequest) GetAuthorLoginOk() (*string, bool)`

GetAuthorLoginOk returns a tuple with the AuthorLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorLogin

`func (o *DbPullRequest) SetAuthorLogin(v string)`

SetAuthorLogin sets AuthorLogin field to given value.


### GetAuthorAvatar

`func (o *DbPullRequest) GetAuthorAvatar() string`

GetAuthorAvatar returns the AuthorAvatar field if non-nil, zero value otherwise.

### GetAuthorAvatarOk

`func (o *DbPullRequest) GetAuthorAvatarOk() (*string, bool)`

GetAuthorAvatarOk returns a tuple with the AuthorAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorAvatar

`func (o *DbPullRequest) SetAuthorAvatar(v string)`

SetAuthorAvatar sets AuthorAvatar field to given value.


### GetAssigneeLogin

`func (o *DbPullRequest) GetAssigneeLogin() string`

GetAssigneeLogin returns the AssigneeLogin field if non-nil, zero value otherwise.

### GetAssigneeLoginOk

`func (o *DbPullRequest) GetAssigneeLoginOk() (*string, bool)`

GetAssigneeLoginOk returns a tuple with the AssigneeLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeLogin

`func (o *DbPullRequest) SetAssigneeLogin(v string)`

SetAssigneeLogin sets AssigneeLogin field to given value.

### HasAssigneeLogin

`func (o *DbPullRequest) HasAssigneeLogin() bool`

HasAssigneeLogin returns a boolean if a field has been set.

### GetAssigneeAvatar

`func (o *DbPullRequest) GetAssigneeAvatar() string`

GetAssigneeAvatar returns the AssigneeAvatar field if non-nil, zero value otherwise.

### GetAssigneeAvatarOk

`func (o *DbPullRequest) GetAssigneeAvatarOk() (*string, bool)`

GetAssigneeAvatarOk returns a tuple with the AssigneeAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeAvatar

`func (o *DbPullRequest) SetAssigneeAvatar(v string)`

SetAssigneeAvatar sets AssigneeAvatar field to given value.

### HasAssigneeAvatar

`func (o *DbPullRequest) HasAssigneeAvatar() bool`

HasAssigneeAvatar returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DbPullRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbPullRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbPullRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbPullRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetClosedAt

`func (o *DbPullRequest) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *DbPullRequest) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *DbPullRequest) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *DbPullRequest) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.

### GetMergedAt

`func (o *DbPullRequest) GetMergedAt() time.Time`

GetMergedAt returns the MergedAt field if non-nil, zero value otherwise.

### GetMergedAtOk

`func (o *DbPullRequest) GetMergedAtOk() (*time.Time, bool)`

GetMergedAtOk returns a tuple with the MergedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedAt

`func (o *DbPullRequest) SetMergedAt(v time.Time)`

SetMergedAt sets MergedAt field to given value.

### HasMergedAt

`func (o *DbPullRequest) HasMergedAt() bool`

HasMergedAt returns a boolean if a field has been set.

### GetMergedByLogin

`func (o *DbPullRequest) GetMergedByLogin() string`

GetMergedByLogin returns the MergedByLogin field if non-nil, zero value otherwise.

### GetMergedByLoginOk

`func (o *DbPullRequest) GetMergedByLoginOk() (*string, bool)`

GetMergedByLoginOk returns a tuple with the MergedByLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedByLogin

`func (o *DbPullRequest) SetMergedByLogin(v string)`

SetMergedByLogin sets MergedByLogin field to given value.

### HasMergedByLogin

`func (o *DbPullRequest) HasMergedByLogin() bool`

HasMergedByLogin returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbPullRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbPullRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbPullRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbPullRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *DbPullRequest) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *DbPullRequest) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *DbPullRequest) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *DbPullRequest) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetComments

`func (o *DbPullRequest) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DbPullRequest) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DbPullRequest) SetComments(v int32)`

SetComments sets Comments field to given value.

### HasComments

`func (o *DbPullRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetAdditions

`func (o *DbPullRequest) GetAdditions() int32`

GetAdditions returns the Additions field if non-nil, zero value otherwise.

### GetAdditionsOk

`func (o *DbPullRequest) GetAdditionsOk() (*int32, bool)`

GetAdditionsOk returns a tuple with the Additions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditions

`func (o *DbPullRequest) SetAdditions(v int32)`

SetAdditions sets Additions field to given value.

### HasAdditions

`func (o *DbPullRequest) HasAdditions() bool`

HasAdditions returns a boolean if a field has been set.

### GetDeletions

`func (o *DbPullRequest) GetDeletions() int32`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *DbPullRequest) GetDeletionsOk() (*int32, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *DbPullRequest) SetDeletions(v int32)`

SetDeletions sets Deletions field to given value.

### HasDeletions

`func (o *DbPullRequest) HasDeletions() bool`

HasDeletions returns a boolean if a field has been set.

### GetChangedFiles

`func (o *DbPullRequest) GetChangedFiles() int32`

GetChangedFiles returns the ChangedFiles field if non-nil, zero value otherwise.

### GetChangedFilesOk

`func (o *DbPullRequest) GetChangedFilesOk() (*int32, bool)`

GetChangedFilesOk returns a tuple with the ChangedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedFiles

`func (o *DbPullRequest) SetChangedFiles(v int32)`

SetChangedFiles sets ChangedFiles field to given value.

### HasChangedFiles

`func (o *DbPullRequest) HasChangedFiles() bool`

HasChangedFiles returns a boolean if a field has been set.

### GetFullName

`func (o *DbPullRequest) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *DbPullRequest) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *DbPullRequest) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *DbPullRequest) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetCommits

`func (o *DbPullRequest) GetCommits() int32`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *DbPullRequest) GetCommitsOk() (*int32, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *DbPullRequest) SetCommits(v int32)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *DbPullRequest) HasCommits() bool`

HasCommits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


