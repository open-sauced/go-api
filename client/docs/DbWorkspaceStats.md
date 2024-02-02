# DbWorkspaceStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PullRequests** | [**DbWorkspacePrStats**](DbWorkspacePrStats.md) |  | 
**Issues** | [**DbWorkspaceIssueStats**](DbWorkspaceIssueStats.md) |  | 
**Repos** | [**DbWorkspaceRepoStats**](DbWorkspaceRepoStats.md) |  | 

## Methods

### NewDbWorkspaceStats

`func NewDbWorkspaceStats(pullRequests DbWorkspacePrStats, issues DbWorkspaceIssueStats, repos DbWorkspaceRepoStats, ) *DbWorkspaceStats`

NewDbWorkspaceStats instantiates a new DbWorkspaceStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbWorkspaceStatsWithDefaults

`func NewDbWorkspaceStatsWithDefaults() *DbWorkspaceStats`

NewDbWorkspaceStatsWithDefaults instantiates a new DbWorkspaceStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPullRequests

`func (o *DbWorkspaceStats) GetPullRequests() DbWorkspacePrStats`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *DbWorkspaceStats) GetPullRequestsOk() (*DbWorkspacePrStats, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *DbWorkspaceStats) SetPullRequests(v DbWorkspacePrStats)`

SetPullRequests sets PullRequests field to given value.


### GetIssues

`func (o *DbWorkspaceStats) GetIssues() DbWorkspaceIssueStats`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *DbWorkspaceStats) GetIssuesOk() (*DbWorkspaceIssueStats, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *DbWorkspaceStats) SetIssues(v DbWorkspaceIssueStats)`

SetIssues sets Issues field to given value.


### GetRepos

`func (o *DbWorkspaceStats) GetRepos() DbWorkspaceRepoStats`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *DbWorkspaceStats) GetReposOk() (*DbWorkspaceRepoStats, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *DbWorkspaceStats) SetRepos(v DbWorkspaceRepoStats)`

SetRepos sets Repos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


