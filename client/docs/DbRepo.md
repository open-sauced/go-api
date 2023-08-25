# DbRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Repository identifier | 
**UserId** | **int32** | Owner user identifier | 
**Size** | **int32** | Total size in bytes | 
**Issues** | **int32** | Total number of issues | 
**Stars** | **int32** | Total number of stars | 
**Forks** | **int32** | Total number of forks | 
**Watchers** | **int32** | Total number of watchers | 
**Subscribers** | **int32** | Total number of subscribers | 
**Network** | **int32** | Total number of network usages | 
**IsFork** | **bool** | Flag indicating repo is a fork | 
**IsPrivate** | **bool** | Flag indicating repo is private | 
**IsTemplate** | **bool** | Flag indicating repo is a template | 
**IsArchived** | **bool** | Flag indicating repo is archived | 
**IsDisabled** | **bool** | Flag indicating repo is disabled | 
**HasIssues** | **bool** | Flag indicating repo has issues enabled | 
**HasProjects** | **bool** | Flag indicating repo has projects enabled | 
**HasDownloads** | **bool** | Flag indicating repo has downloads enabled | 
**HasWiki** | **bool** | Flag indicating repo has wiki enabled | 
**HasPages** | **bool** | Flag indicating repo has pages enabled | 
**HasDiscussions** | **bool** | Flag indicating repo has discussions enabled | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing repository creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing repository last update | [optional] 
**PushedAt** | Pointer to **time.Time** | Timestamp representing repository last push | [optional] 
**DefaultBranch** | **string** | Repository default branch | 
**NodeId** | **string** | Repository GQL node id | 
**GitUrl** | **string** | Repository git url | 
**SshUrl** | **string** | Repository ssh url | 
**CloneUrl** | **string** | Repository clone url | 
**SvnUrl** | **string** | Repository svn url | 
**MirrorUrl** | Pointer to **string** | Repository mirror url | [optional] 
**Name** | **string** | Repository unique name | 
**FullName** | **string** | Repository full name | 
**Description** | **string** | Repository short description | 
**Language** | **string** | Repository programming language | 
**License** | **string** | Repository SPDX license | 
**Url** | **string** | Repository GitHub linked URL | 
**Homepage** | **string** | Repository GitHub homepage | 
**Topics** | **[]string** | Repository GitHub topics | 
**OpenPrsCount** | Pointer to **int32** | Repository number of open PRs | [optional] 
**ClosedPrsCount** | Pointer to **int32** | Repository number of closed PRs | [optional] 
**MergedPrsCount** | Pointer to **int32** | Repository number of merged PRs | [optional] 
**DraftPrsCount** | Pointer to **int32** | Repository number of draft PRs | [optional] 
**SpamPrsCount** | Pointer to **int32** | Repository number of spam PRs | [optional] 
**PrVelocityCount** | Pointer to **int32** | Repository average open/close time for PRs | [optional] 
**PrActiveCount** | Pointer to **int32** | Number of non-closed PRs updated within the day range | [optional] 

## Methods

### NewDbRepo

`func NewDbRepo(id int32, userId int32, size int32, issues int32, stars int32, forks int32, watchers int32, subscribers int32, network int32, isFork bool, isPrivate bool, isTemplate bool, isArchived bool, isDisabled bool, hasIssues bool, hasProjects bool, hasDownloads bool, hasWiki bool, hasPages bool, hasDiscussions bool, defaultBranch string, nodeId string, gitUrl string, sshUrl string, cloneUrl string, svnUrl string, name string, fullName string, description string, language string, license string, url string, homepage string, topics []string, ) *DbRepo`

NewDbRepo instantiates a new DbRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbRepoWithDefaults

`func NewDbRepoWithDefaults() *DbRepo`

NewDbRepoWithDefaults instantiates a new DbRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbRepo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbRepo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbRepo) SetId(v int32)`

SetId sets Id field to given value.


### GetUserId

`func (o *DbRepo) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DbRepo) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DbRepo) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetSize

`func (o *DbRepo) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DbRepo) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DbRepo) SetSize(v int32)`

SetSize sets Size field to given value.


### GetIssues

`func (o *DbRepo) GetIssues() int32`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *DbRepo) GetIssuesOk() (*int32, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *DbRepo) SetIssues(v int32)`

SetIssues sets Issues field to given value.


### GetStars

`func (o *DbRepo) GetStars() int32`

GetStars returns the Stars field if non-nil, zero value otherwise.

### GetStarsOk

`func (o *DbRepo) GetStarsOk() (*int32, bool)`

GetStarsOk returns a tuple with the Stars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStars

`func (o *DbRepo) SetStars(v int32)`

SetStars sets Stars field to given value.


### GetForks

`func (o *DbRepo) GetForks() int32`

GetForks returns the Forks field if non-nil, zero value otherwise.

### GetForksOk

`func (o *DbRepo) GetForksOk() (*int32, bool)`

GetForksOk returns a tuple with the Forks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForks

`func (o *DbRepo) SetForks(v int32)`

SetForks sets Forks field to given value.


### GetWatchers

`func (o *DbRepo) GetWatchers() int32`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *DbRepo) GetWatchersOk() (*int32, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *DbRepo) SetWatchers(v int32)`

SetWatchers sets Watchers field to given value.


### GetSubscribers

`func (o *DbRepo) GetSubscribers() int32`

GetSubscribers returns the Subscribers field if non-nil, zero value otherwise.

### GetSubscribersOk

`func (o *DbRepo) GetSubscribersOk() (*int32, bool)`

GetSubscribersOk returns a tuple with the Subscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribers

`func (o *DbRepo) SetSubscribers(v int32)`

SetSubscribers sets Subscribers field to given value.


### GetNetwork

`func (o *DbRepo) GetNetwork() int32`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DbRepo) GetNetworkOk() (*int32, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DbRepo) SetNetwork(v int32)`

SetNetwork sets Network field to given value.


### GetIsFork

`func (o *DbRepo) GetIsFork() bool`

GetIsFork returns the IsFork field if non-nil, zero value otherwise.

### GetIsForkOk

`func (o *DbRepo) GetIsForkOk() (*bool, bool)`

GetIsForkOk returns a tuple with the IsFork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFork

`func (o *DbRepo) SetIsFork(v bool)`

SetIsFork sets IsFork field to given value.


### GetIsPrivate

`func (o *DbRepo) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *DbRepo) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *DbRepo) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.


### GetIsTemplate

`func (o *DbRepo) GetIsTemplate() bool`

GetIsTemplate returns the IsTemplate field if non-nil, zero value otherwise.

### GetIsTemplateOk

`func (o *DbRepo) GetIsTemplateOk() (*bool, bool)`

GetIsTemplateOk returns a tuple with the IsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTemplate

`func (o *DbRepo) SetIsTemplate(v bool)`

SetIsTemplate sets IsTemplate field to given value.


### GetIsArchived

`func (o *DbRepo) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *DbRepo) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *DbRepo) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.


### GetIsDisabled

`func (o *DbRepo) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *DbRepo) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *DbRepo) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.


### GetHasIssues

`func (o *DbRepo) GetHasIssues() bool`

GetHasIssues returns the HasIssues field if non-nil, zero value otherwise.

### GetHasIssuesOk

`func (o *DbRepo) GetHasIssuesOk() (*bool, bool)`

GetHasIssuesOk returns a tuple with the HasIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIssues

`func (o *DbRepo) SetHasIssues(v bool)`

SetHasIssues sets HasIssues field to given value.


### GetHasProjects

`func (o *DbRepo) GetHasProjects() bool`

GetHasProjects returns the HasProjects field if non-nil, zero value otherwise.

### GetHasProjectsOk

`func (o *DbRepo) GetHasProjectsOk() (*bool, bool)`

GetHasProjectsOk returns a tuple with the HasProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProjects

`func (o *DbRepo) SetHasProjects(v bool)`

SetHasProjects sets HasProjects field to given value.


### GetHasDownloads

`func (o *DbRepo) GetHasDownloads() bool`

GetHasDownloads returns the HasDownloads field if non-nil, zero value otherwise.

### GetHasDownloadsOk

`func (o *DbRepo) GetHasDownloadsOk() (*bool, bool)`

GetHasDownloadsOk returns a tuple with the HasDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDownloads

`func (o *DbRepo) SetHasDownloads(v bool)`

SetHasDownloads sets HasDownloads field to given value.


### GetHasWiki

`func (o *DbRepo) GetHasWiki() bool`

GetHasWiki returns the HasWiki field if non-nil, zero value otherwise.

### GetHasWikiOk

`func (o *DbRepo) GetHasWikiOk() (*bool, bool)`

GetHasWikiOk returns a tuple with the HasWiki field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWiki

`func (o *DbRepo) SetHasWiki(v bool)`

SetHasWiki sets HasWiki field to given value.


### GetHasPages

`func (o *DbRepo) GetHasPages() bool`

GetHasPages returns the HasPages field if non-nil, zero value otherwise.

### GetHasPagesOk

`func (o *DbRepo) GetHasPagesOk() (*bool, bool)`

GetHasPagesOk returns a tuple with the HasPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPages

`func (o *DbRepo) SetHasPages(v bool)`

SetHasPages sets HasPages field to given value.


### GetHasDiscussions

`func (o *DbRepo) GetHasDiscussions() bool`

GetHasDiscussions returns the HasDiscussions field if non-nil, zero value otherwise.

### GetHasDiscussionsOk

`func (o *DbRepo) GetHasDiscussionsOk() (*bool, bool)`

GetHasDiscussionsOk returns a tuple with the HasDiscussions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDiscussions

`func (o *DbRepo) SetHasDiscussions(v bool)`

SetHasDiscussions sets HasDiscussions field to given value.


### GetCreatedAt

`func (o *DbRepo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbRepo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbRepo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbRepo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbRepo) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbRepo) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbRepo) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbRepo) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPushedAt

`func (o *DbRepo) GetPushedAt() time.Time`

GetPushedAt returns the PushedAt field if non-nil, zero value otherwise.

### GetPushedAtOk

`func (o *DbRepo) GetPushedAtOk() (*time.Time, bool)`

GetPushedAtOk returns a tuple with the PushedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAt

`func (o *DbRepo) SetPushedAt(v time.Time)`

SetPushedAt sets PushedAt field to given value.

### HasPushedAt

`func (o *DbRepo) HasPushedAt() bool`

HasPushedAt returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *DbRepo) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *DbRepo) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *DbRepo) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.


### GetNodeId

`func (o *DbRepo) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *DbRepo) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *DbRepo) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetGitUrl

`func (o *DbRepo) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *DbRepo) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *DbRepo) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.


### GetSshUrl

`func (o *DbRepo) GetSshUrl() string`

GetSshUrl returns the SshUrl field if non-nil, zero value otherwise.

### GetSshUrlOk

`func (o *DbRepo) GetSshUrlOk() (*string, bool)`

GetSshUrlOk returns a tuple with the SshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUrl

`func (o *DbRepo) SetSshUrl(v string)`

SetSshUrl sets SshUrl field to given value.


### GetCloneUrl

`func (o *DbRepo) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *DbRepo) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *DbRepo) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.


### GetSvnUrl

`func (o *DbRepo) GetSvnUrl() string`

GetSvnUrl returns the SvnUrl field if non-nil, zero value otherwise.

### GetSvnUrlOk

`func (o *DbRepo) GetSvnUrlOk() (*string, bool)`

GetSvnUrlOk returns a tuple with the SvnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvnUrl

`func (o *DbRepo) SetSvnUrl(v string)`

SetSvnUrl sets SvnUrl field to given value.


### GetMirrorUrl

`func (o *DbRepo) GetMirrorUrl() string`

GetMirrorUrl returns the MirrorUrl field if non-nil, zero value otherwise.

### GetMirrorUrlOk

`func (o *DbRepo) GetMirrorUrlOk() (*string, bool)`

GetMirrorUrlOk returns a tuple with the MirrorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorUrl

`func (o *DbRepo) SetMirrorUrl(v string)`

SetMirrorUrl sets MirrorUrl field to given value.

### HasMirrorUrl

`func (o *DbRepo) HasMirrorUrl() bool`

HasMirrorUrl returns a boolean if a field has been set.

### GetName

`func (o *DbRepo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbRepo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbRepo) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *DbRepo) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *DbRepo) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *DbRepo) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetDescription

`func (o *DbRepo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DbRepo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DbRepo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLanguage

`func (o *DbRepo) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *DbRepo) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *DbRepo) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLicense

`func (o *DbRepo) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *DbRepo) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *DbRepo) SetLicense(v string)`

SetLicense sets License field to given value.


### GetUrl

`func (o *DbRepo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DbRepo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DbRepo) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHomepage

`func (o *DbRepo) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *DbRepo) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *DbRepo) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.


### GetTopics

`func (o *DbRepo) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *DbRepo) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *DbRepo) SetTopics(v []string)`

SetTopics sets Topics field to given value.


### GetOpenPrsCount

`func (o *DbRepo) GetOpenPrsCount() int32`

GetOpenPrsCount returns the OpenPrsCount field if non-nil, zero value otherwise.

### GetOpenPrsCountOk

`func (o *DbRepo) GetOpenPrsCountOk() (*int32, bool)`

GetOpenPrsCountOk returns a tuple with the OpenPrsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrsCount

`func (o *DbRepo) SetOpenPrsCount(v int32)`

SetOpenPrsCount sets OpenPrsCount field to given value.

### HasOpenPrsCount

`func (o *DbRepo) HasOpenPrsCount() bool`

HasOpenPrsCount returns a boolean if a field has been set.

### GetClosedPrsCount

`func (o *DbRepo) GetClosedPrsCount() int32`

GetClosedPrsCount returns the ClosedPrsCount field if non-nil, zero value otherwise.

### GetClosedPrsCountOk

`func (o *DbRepo) GetClosedPrsCountOk() (*int32, bool)`

GetClosedPrsCountOk returns a tuple with the ClosedPrsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedPrsCount

`func (o *DbRepo) SetClosedPrsCount(v int32)`

SetClosedPrsCount sets ClosedPrsCount field to given value.

### HasClosedPrsCount

`func (o *DbRepo) HasClosedPrsCount() bool`

HasClosedPrsCount returns a boolean if a field has been set.

### GetMergedPrsCount

`func (o *DbRepo) GetMergedPrsCount() int32`

GetMergedPrsCount returns the MergedPrsCount field if non-nil, zero value otherwise.

### GetMergedPrsCountOk

`func (o *DbRepo) GetMergedPrsCountOk() (*int32, bool)`

GetMergedPrsCountOk returns a tuple with the MergedPrsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedPrsCount

`func (o *DbRepo) SetMergedPrsCount(v int32)`

SetMergedPrsCount sets MergedPrsCount field to given value.

### HasMergedPrsCount

`func (o *DbRepo) HasMergedPrsCount() bool`

HasMergedPrsCount returns a boolean if a field has been set.

### GetDraftPrsCount

`func (o *DbRepo) GetDraftPrsCount() int32`

GetDraftPrsCount returns the DraftPrsCount field if non-nil, zero value otherwise.

### GetDraftPrsCountOk

`func (o *DbRepo) GetDraftPrsCountOk() (*int32, bool)`

GetDraftPrsCountOk returns a tuple with the DraftPrsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftPrsCount

`func (o *DbRepo) SetDraftPrsCount(v int32)`

SetDraftPrsCount sets DraftPrsCount field to given value.

### HasDraftPrsCount

`func (o *DbRepo) HasDraftPrsCount() bool`

HasDraftPrsCount returns a boolean if a field has been set.

### GetSpamPrsCount

`func (o *DbRepo) GetSpamPrsCount() int32`

GetSpamPrsCount returns the SpamPrsCount field if non-nil, zero value otherwise.

### GetSpamPrsCountOk

`func (o *DbRepo) GetSpamPrsCountOk() (*int32, bool)`

GetSpamPrsCountOk returns a tuple with the SpamPrsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamPrsCount

`func (o *DbRepo) SetSpamPrsCount(v int32)`

SetSpamPrsCount sets SpamPrsCount field to given value.

### HasSpamPrsCount

`func (o *DbRepo) HasSpamPrsCount() bool`

HasSpamPrsCount returns a boolean if a field has been set.

### GetPrVelocityCount

`func (o *DbRepo) GetPrVelocityCount() int32`

GetPrVelocityCount returns the PrVelocityCount field if non-nil, zero value otherwise.

### GetPrVelocityCountOk

`func (o *DbRepo) GetPrVelocityCountOk() (*int32, bool)`

GetPrVelocityCountOk returns a tuple with the PrVelocityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrVelocityCount

`func (o *DbRepo) SetPrVelocityCount(v int32)`

SetPrVelocityCount sets PrVelocityCount field to given value.

### HasPrVelocityCount

`func (o *DbRepo) HasPrVelocityCount() bool`

HasPrVelocityCount returns a boolean if a field has been set.

### GetPrActiveCount

`func (o *DbRepo) GetPrActiveCount() int32`

GetPrActiveCount returns the PrActiveCount field if non-nil, zero value otherwise.

### GetPrActiveCountOk

`func (o *DbRepo) GetPrActiveCountOk() (*int32, bool)`

GetPrActiveCountOk returns a tuple with the PrActiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrActiveCount

`func (o *DbRepo) SetPrActiveCount(v int32)`

SetPrActiveCount sets PrActiveCount field to given value.

### HasPrActiveCount

`func (o *DbRepo) HasPrActiveCount() bool`

HasPrActiveCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


