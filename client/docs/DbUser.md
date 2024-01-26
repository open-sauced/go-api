# DbUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | User identifier | 
**OpenIssues** | **int32** | Total number of open issues user has across public activity | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing user creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing user last update | [optional] 
**FirstOpenedPrAt** | Pointer to **time.Time** | Timestamp representing user first open PR | [optional] 
**FirstPushedCommitAt** | Pointer to **time.Time** | Timestamp representing user first commit push | [optional] 
**ConnectedAt** | Pointer to **time.Time** | Timestamp representing user logging in to open sauced for the first time | [optional] 
**CampaignStartDate** | Pointer to **time.Time** | Timestamp representing user email campaign start date | [optional] 
**NodeId** | **string** | User GitHub node id | 
**AvatarUrl** | **string** | User GitHub avatar URL | 
**GravatarId** | Pointer to **string** | User GitHub gravatar URL | [optional] 
**Url** | Pointer to **string** | User GitHub profile URL | [optional] 
**Login** | **string** | User unique login name | 
**Email** | **string** | User email address | 
**IsPrivate** | **bool** | Flag indicating whether user opted to have a private profile (beta feature | [default to false]
**IsOpenSaucedMember** | **bool** | Flag indicating app.opensauced user status | [default to false]
**IsOnboarded** | **bool** | Flag indicating user&#39;s onboarding status | [default to false]
**IsWaitlisted** | **bool** | Flag indicating user&#39;s waitlist status | [default to false]
**Role** | **int32** | Insights Role | [default to 10]
**Bio** | Pointer to **string** | User bio information | [optional] 
**Blog** | Pointer to **string** | GitHub blog information | [optional] 
**Name** | Pointer to **string** | User name information | [optional] 
**TwitterUsername** | Pointer to **string** | User Twitter information | [optional] 
**LinkedinUrl** | Pointer to **string** | LinkedIn URL | [optional] 
**GithubSponsorsUrl** | Pointer to **string** | GitHub Sponsors URL | [optional] 
**DiscordUrl** | Pointer to **string** | Discord URL | [optional] 
**Company** | Pointer to **string** | User company information | [optional] 
**Location** | Pointer to **string** | User location information | [optional] 
**DisplayLocalTime** | Pointer to **bool** | User display local time information | [optional] 
**Interests** | Pointer to **string** | User topic interests | [optional] 
**Hireable** | Pointer to **bool** | GitHub user hireable status | [optional] 
**PublicRepos** | **int32** | GitHub user public repos number | 
**PublicGists** | **int32** | GitHub user public gists number | 
**Followers** | **int32** | GitHub user public followers number | 
**Following** | **int32** | GitHub user public following number | 
**Type** | **string** | GitHub user type | [default to "User"]
**DisplayEmail** | Pointer to **bool** | User display public email | [optional] 
**ReceiveCollaboration** | Pointer to **bool** | User receives collaboration requests | [optional] 
**ReceiveProductUpdates** | Pointer to **bool** | User receives product updates through email | [optional] 
**Timezone** | Pointer to **string** | User timezone in UTC | [optional] 
**CouponCode** | Pointer to **string** | Coupon Code | [optional] 
**Languages** | **map[string]interface{}** | GitHub top languages | [default to {}]
**NotificationCount** | **int32** | User notification count | 
**InsightsCount** | **int32** | User insight pages count | 
**HighlightsCount** | **int32** | User highlights count | 
**FollowingCount** | **int32** | User following count | 
**FollowersCount** | **int32** | User followers count | 
**RecentPullRequestsCount** | **int32** | Count of user pull requests within the last 30 days | 
**RecentPullRequestVelocityCount** | **int32** | User average pull request velocity in days over the last 30 days | 
**IsMaintainer** | **bool** | Flag to indicate if user is a maintainer of any repos | 

## Methods

### NewDbUser

`func NewDbUser(id int32, openIssues int32, nodeId string, avatarUrl string, login string, email string, isPrivate bool, isOpenSaucedMember bool, isOnboarded bool, isWaitlisted bool, role int32, publicRepos int32, publicGists int32, followers int32, following int32, type_ string, languages map[string]interface{}, notificationCount int32, insightsCount int32, highlightsCount int32, followingCount int32, followersCount int32, recentPullRequestsCount int32, recentPullRequestVelocityCount int32, isMaintainer bool, ) *DbUser`

NewDbUser instantiates a new DbUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbUserWithDefaults

`func NewDbUserWithDefaults() *DbUser`

NewDbUserWithDefaults instantiates a new DbUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbUser) SetId(v int32)`

SetId sets Id field to given value.


### GetOpenIssues

`func (o *DbUser) GetOpenIssues() int32`

GetOpenIssues returns the OpenIssues field if non-nil, zero value otherwise.

### GetOpenIssuesOk

`func (o *DbUser) GetOpenIssuesOk() (*int32, bool)`

GetOpenIssuesOk returns a tuple with the OpenIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIssues

`func (o *DbUser) SetOpenIssues(v int32)`

SetOpenIssues sets OpenIssues field to given value.


### GetCreatedAt

`func (o *DbUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFirstOpenedPrAt

`func (o *DbUser) GetFirstOpenedPrAt() time.Time`

GetFirstOpenedPrAt returns the FirstOpenedPrAt field if non-nil, zero value otherwise.

### GetFirstOpenedPrAtOk

`func (o *DbUser) GetFirstOpenedPrAtOk() (*time.Time, bool)`

GetFirstOpenedPrAtOk returns a tuple with the FirstOpenedPrAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstOpenedPrAt

`func (o *DbUser) SetFirstOpenedPrAt(v time.Time)`

SetFirstOpenedPrAt sets FirstOpenedPrAt field to given value.

### HasFirstOpenedPrAt

`func (o *DbUser) HasFirstOpenedPrAt() bool`

HasFirstOpenedPrAt returns a boolean if a field has been set.

### GetFirstPushedCommitAt

`func (o *DbUser) GetFirstPushedCommitAt() time.Time`

GetFirstPushedCommitAt returns the FirstPushedCommitAt field if non-nil, zero value otherwise.

### GetFirstPushedCommitAtOk

`func (o *DbUser) GetFirstPushedCommitAtOk() (*time.Time, bool)`

GetFirstPushedCommitAtOk returns a tuple with the FirstPushedCommitAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPushedCommitAt

`func (o *DbUser) SetFirstPushedCommitAt(v time.Time)`

SetFirstPushedCommitAt sets FirstPushedCommitAt field to given value.

### HasFirstPushedCommitAt

`func (o *DbUser) HasFirstPushedCommitAt() bool`

HasFirstPushedCommitAt returns a boolean if a field has been set.

### GetConnectedAt

`func (o *DbUser) GetConnectedAt() time.Time`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *DbUser) GetConnectedAtOk() (*time.Time, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *DbUser) SetConnectedAt(v time.Time)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *DbUser) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetCampaignStartDate

`func (o *DbUser) GetCampaignStartDate() time.Time`

GetCampaignStartDate returns the CampaignStartDate field if non-nil, zero value otherwise.

### GetCampaignStartDateOk

`func (o *DbUser) GetCampaignStartDateOk() (*time.Time, bool)`

GetCampaignStartDateOk returns a tuple with the CampaignStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignStartDate

`func (o *DbUser) SetCampaignStartDate(v time.Time)`

SetCampaignStartDate sets CampaignStartDate field to given value.

### HasCampaignStartDate

`func (o *DbUser) HasCampaignStartDate() bool`

HasCampaignStartDate returns a boolean if a field has been set.

### GetNodeId

`func (o *DbUser) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *DbUser) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *DbUser) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetAvatarUrl

`func (o *DbUser) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *DbUser) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *DbUser) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.


### GetGravatarId

`func (o *DbUser) GetGravatarId() string`

GetGravatarId returns the GravatarId field if non-nil, zero value otherwise.

### GetGravatarIdOk

`func (o *DbUser) GetGravatarIdOk() (*string, bool)`

GetGravatarIdOk returns a tuple with the GravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarId

`func (o *DbUser) SetGravatarId(v string)`

SetGravatarId sets GravatarId field to given value.

### HasGravatarId

`func (o *DbUser) HasGravatarId() bool`

HasGravatarId returns a boolean if a field has been set.

### GetUrl

`func (o *DbUser) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DbUser) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DbUser) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DbUser) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLogin

`func (o *DbUser) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *DbUser) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *DbUser) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetEmail

`func (o *DbUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DbUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DbUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIsPrivate

`func (o *DbUser) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *DbUser) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *DbUser) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.


### GetIsOpenSaucedMember

`func (o *DbUser) GetIsOpenSaucedMember() bool`

GetIsOpenSaucedMember returns the IsOpenSaucedMember field if non-nil, zero value otherwise.

### GetIsOpenSaucedMemberOk

`func (o *DbUser) GetIsOpenSaucedMemberOk() (*bool, bool)`

GetIsOpenSaucedMemberOk returns a tuple with the IsOpenSaucedMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpenSaucedMember

`func (o *DbUser) SetIsOpenSaucedMember(v bool)`

SetIsOpenSaucedMember sets IsOpenSaucedMember field to given value.


### GetIsOnboarded

`func (o *DbUser) GetIsOnboarded() bool`

GetIsOnboarded returns the IsOnboarded field if non-nil, zero value otherwise.

### GetIsOnboardedOk

`func (o *DbUser) GetIsOnboardedOk() (*bool, bool)`

GetIsOnboardedOk returns a tuple with the IsOnboarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnboarded

`func (o *DbUser) SetIsOnboarded(v bool)`

SetIsOnboarded sets IsOnboarded field to given value.


### GetIsWaitlisted

`func (o *DbUser) GetIsWaitlisted() bool`

GetIsWaitlisted returns the IsWaitlisted field if non-nil, zero value otherwise.

### GetIsWaitlistedOk

`func (o *DbUser) GetIsWaitlistedOk() (*bool, bool)`

GetIsWaitlistedOk returns a tuple with the IsWaitlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWaitlisted

`func (o *DbUser) SetIsWaitlisted(v bool)`

SetIsWaitlisted sets IsWaitlisted field to given value.


### GetRole

`func (o *DbUser) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DbUser) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DbUser) SetRole(v int32)`

SetRole sets Role field to given value.


### GetBio

`func (o *DbUser) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *DbUser) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *DbUser) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *DbUser) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetBlog

`func (o *DbUser) GetBlog() string`

GetBlog returns the Blog field if non-nil, zero value otherwise.

### GetBlogOk

`func (o *DbUser) GetBlogOk() (*string, bool)`

GetBlogOk returns a tuple with the Blog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlog

`func (o *DbUser) SetBlog(v string)`

SetBlog sets Blog field to given value.

### HasBlog

`func (o *DbUser) HasBlog() bool`

HasBlog returns a boolean if a field has been set.

### GetName

`func (o *DbUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DbUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTwitterUsername

`func (o *DbUser) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *DbUser) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *DbUser) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *DbUser) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### GetLinkedinUrl

`func (o *DbUser) GetLinkedinUrl() string`

GetLinkedinUrl returns the LinkedinUrl field if non-nil, zero value otherwise.

### GetLinkedinUrlOk

`func (o *DbUser) GetLinkedinUrlOk() (*string, bool)`

GetLinkedinUrlOk returns a tuple with the LinkedinUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedinUrl

`func (o *DbUser) SetLinkedinUrl(v string)`

SetLinkedinUrl sets LinkedinUrl field to given value.

### HasLinkedinUrl

`func (o *DbUser) HasLinkedinUrl() bool`

HasLinkedinUrl returns a boolean if a field has been set.

### GetGithubSponsorsUrl

`func (o *DbUser) GetGithubSponsorsUrl() string`

GetGithubSponsorsUrl returns the GithubSponsorsUrl field if non-nil, zero value otherwise.

### GetGithubSponsorsUrlOk

`func (o *DbUser) GetGithubSponsorsUrlOk() (*string, bool)`

GetGithubSponsorsUrlOk returns a tuple with the GithubSponsorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubSponsorsUrl

`func (o *DbUser) SetGithubSponsorsUrl(v string)`

SetGithubSponsorsUrl sets GithubSponsorsUrl field to given value.

### HasGithubSponsorsUrl

`func (o *DbUser) HasGithubSponsorsUrl() bool`

HasGithubSponsorsUrl returns a boolean if a field has been set.

### GetDiscordUrl

`func (o *DbUser) GetDiscordUrl() string`

GetDiscordUrl returns the DiscordUrl field if non-nil, zero value otherwise.

### GetDiscordUrlOk

`func (o *DbUser) GetDiscordUrlOk() (*string, bool)`

GetDiscordUrlOk returns a tuple with the DiscordUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordUrl

`func (o *DbUser) SetDiscordUrl(v string)`

SetDiscordUrl sets DiscordUrl field to given value.

### HasDiscordUrl

`func (o *DbUser) HasDiscordUrl() bool`

HasDiscordUrl returns a boolean if a field has been set.

### GetCompany

`func (o *DbUser) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *DbUser) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *DbUser) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *DbUser) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetLocation

`func (o *DbUser) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DbUser) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DbUser) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DbUser) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDisplayLocalTime

`func (o *DbUser) GetDisplayLocalTime() bool`

GetDisplayLocalTime returns the DisplayLocalTime field if non-nil, zero value otherwise.

### GetDisplayLocalTimeOk

`func (o *DbUser) GetDisplayLocalTimeOk() (*bool, bool)`

GetDisplayLocalTimeOk returns a tuple with the DisplayLocalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLocalTime

`func (o *DbUser) SetDisplayLocalTime(v bool)`

SetDisplayLocalTime sets DisplayLocalTime field to given value.

### HasDisplayLocalTime

`func (o *DbUser) HasDisplayLocalTime() bool`

HasDisplayLocalTime returns a boolean if a field has been set.

### GetInterests

`func (o *DbUser) GetInterests() string`

GetInterests returns the Interests field if non-nil, zero value otherwise.

### GetInterestsOk

`func (o *DbUser) GetInterestsOk() (*string, bool)`

GetInterestsOk returns a tuple with the Interests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterests

`func (o *DbUser) SetInterests(v string)`

SetInterests sets Interests field to given value.

### HasInterests

`func (o *DbUser) HasInterests() bool`

HasInterests returns a boolean if a field has been set.

### GetHireable

`func (o *DbUser) GetHireable() bool`

GetHireable returns the Hireable field if non-nil, zero value otherwise.

### GetHireableOk

`func (o *DbUser) GetHireableOk() (*bool, bool)`

GetHireableOk returns a tuple with the Hireable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHireable

`func (o *DbUser) SetHireable(v bool)`

SetHireable sets Hireable field to given value.

### HasHireable

`func (o *DbUser) HasHireable() bool`

HasHireable returns a boolean if a field has been set.

### GetPublicRepos

`func (o *DbUser) GetPublicRepos() int32`

GetPublicRepos returns the PublicRepos field if non-nil, zero value otherwise.

### GetPublicReposOk

`func (o *DbUser) GetPublicReposOk() (*int32, bool)`

GetPublicReposOk returns a tuple with the PublicRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicRepos

`func (o *DbUser) SetPublicRepos(v int32)`

SetPublicRepos sets PublicRepos field to given value.


### GetPublicGists

`func (o *DbUser) GetPublicGists() int32`

GetPublicGists returns the PublicGists field if non-nil, zero value otherwise.

### GetPublicGistsOk

`func (o *DbUser) GetPublicGistsOk() (*int32, bool)`

GetPublicGistsOk returns a tuple with the PublicGists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicGists

`func (o *DbUser) SetPublicGists(v int32)`

SetPublicGists sets PublicGists field to given value.


### GetFollowers

`func (o *DbUser) GetFollowers() int32`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *DbUser) GetFollowersOk() (*int32, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *DbUser) SetFollowers(v int32)`

SetFollowers sets Followers field to given value.


### GetFollowing

`func (o *DbUser) GetFollowing() int32`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *DbUser) GetFollowingOk() (*int32, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *DbUser) SetFollowing(v int32)`

SetFollowing sets Following field to given value.


### GetType

`func (o *DbUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DbUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DbUser) SetType(v string)`

SetType sets Type field to given value.


### GetDisplayEmail

`func (o *DbUser) GetDisplayEmail() bool`

GetDisplayEmail returns the DisplayEmail field if non-nil, zero value otherwise.

### GetDisplayEmailOk

`func (o *DbUser) GetDisplayEmailOk() (*bool, bool)`

GetDisplayEmailOk returns a tuple with the DisplayEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayEmail

`func (o *DbUser) SetDisplayEmail(v bool)`

SetDisplayEmail sets DisplayEmail field to given value.

### HasDisplayEmail

`func (o *DbUser) HasDisplayEmail() bool`

HasDisplayEmail returns a boolean if a field has been set.

### GetReceiveCollaboration

`func (o *DbUser) GetReceiveCollaboration() bool`

GetReceiveCollaboration returns the ReceiveCollaboration field if non-nil, zero value otherwise.

### GetReceiveCollaborationOk

`func (o *DbUser) GetReceiveCollaborationOk() (*bool, bool)`

GetReceiveCollaborationOk returns a tuple with the ReceiveCollaboration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveCollaboration

`func (o *DbUser) SetReceiveCollaboration(v bool)`

SetReceiveCollaboration sets ReceiveCollaboration field to given value.

### HasReceiveCollaboration

`func (o *DbUser) HasReceiveCollaboration() bool`

HasReceiveCollaboration returns a boolean if a field has been set.

### GetReceiveProductUpdates

`func (o *DbUser) GetReceiveProductUpdates() bool`

GetReceiveProductUpdates returns the ReceiveProductUpdates field if non-nil, zero value otherwise.

### GetReceiveProductUpdatesOk

`func (o *DbUser) GetReceiveProductUpdatesOk() (*bool, bool)`

GetReceiveProductUpdatesOk returns a tuple with the ReceiveProductUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveProductUpdates

`func (o *DbUser) SetReceiveProductUpdates(v bool)`

SetReceiveProductUpdates sets ReceiveProductUpdates field to given value.

### HasReceiveProductUpdates

`func (o *DbUser) HasReceiveProductUpdates() bool`

HasReceiveProductUpdates returns a boolean if a field has been set.

### GetTimezone

`func (o *DbUser) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *DbUser) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *DbUser) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *DbUser) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetCouponCode

`func (o *DbUser) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *DbUser) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *DbUser) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *DbUser) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### GetLanguages

`func (o *DbUser) GetLanguages() map[string]interface{}`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *DbUser) GetLanguagesOk() (*map[string]interface{}, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *DbUser) SetLanguages(v map[string]interface{})`

SetLanguages sets Languages field to given value.


### GetNotificationCount

`func (o *DbUser) GetNotificationCount() int32`

GetNotificationCount returns the NotificationCount field if non-nil, zero value otherwise.

### GetNotificationCountOk

`func (o *DbUser) GetNotificationCountOk() (*int32, bool)`

GetNotificationCountOk returns a tuple with the NotificationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCount

`func (o *DbUser) SetNotificationCount(v int32)`

SetNotificationCount sets NotificationCount field to given value.


### GetInsightsCount

`func (o *DbUser) GetInsightsCount() int32`

GetInsightsCount returns the InsightsCount field if non-nil, zero value otherwise.

### GetInsightsCountOk

`func (o *DbUser) GetInsightsCountOk() (*int32, bool)`

GetInsightsCountOk returns a tuple with the InsightsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsCount

`func (o *DbUser) SetInsightsCount(v int32)`

SetInsightsCount sets InsightsCount field to given value.


### GetHighlightsCount

`func (o *DbUser) GetHighlightsCount() int32`

GetHighlightsCount returns the HighlightsCount field if non-nil, zero value otherwise.

### GetHighlightsCountOk

`func (o *DbUser) GetHighlightsCountOk() (*int32, bool)`

GetHighlightsCountOk returns a tuple with the HighlightsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightsCount

`func (o *DbUser) SetHighlightsCount(v int32)`

SetHighlightsCount sets HighlightsCount field to given value.


### GetFollowingCount

`func (o *DbUser) GetFollowingCount() int32`

GetFollowingCount returns the FollowingCount field if non-nil, zero value otherwise.

### GetFollowingCountOk

`func (o *DbUser) GetFollowingCountOk() (*int32, bool)`

GetFollowingCountOk returns a tuple with the FollowingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingCount

`func (o *DbUser) SetFollowingCount(v int32)`

SetFollowingCount sets FollowingCount field to given value.


### GetFollowersCount

`func (o *DbUser) GetFollowersCount() int32`

GetFollowersCount returns the FollowersCount field if non-nil, zero value otherwise.

### GetFollowersCountOk

`func (o *DbUser) GetFollowersCountOk() (*int32, bool)`

GetFollowersCountOk returns a tuple with the FollowersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersCount

`func (o *DbUser) SetFollowersCount(v int32)`

SetFollowersCount sets FollowersCount field to given value.


### GetRecentPullRequestsCount

`func (o *DbUser) GetRecentPullRequestsCount() int32`

GetRecentPullRequestsCount returns the RecentPullRequestsCount field if non-nil, zero value otherwise.

### GetRecentPullRequestsCountOk

`func (o *DbUser) GetRecentPullRequestsCountOk() (*int32, bool)`

GetRecentPullRequestsCountOk returns a tuple with the RecentPullRequestsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentPullRequestsCount

`func (o *DbUser) SetRecentPullRequestsCount(v int32)`

SetRecentPullRequestsCount sets RecentPullRequestsCount field to given value.


### GetRecentPullRequestVelocityCount

`func (o *DbUser) GetRecentPullRequestVelocityCount() int32`

GetRecentPullRequestVelocityCount returns the RecentPullRequestVelocityCount field if non-nil, zero value otherwise.

### GetRecentPullRequestVelocityCountOk

`func (o *DbUser) GetRecentPullRequestVelocityCountOk() (*int32, bool)`

GetRecentPullRequestVelocityCountOk returns a tuple with the RecentPullRequestVelocityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentPullRequestVelocityCount

`func (o *DbUser) SetRecentPullRequestVelocityCount(v int32)`

SetRecentPullRequestVelocityCount sets RecentPullRequestVelocityCount field to given value.


### GetIsMaintainer

`func (o *DbUser) GetIsMaintainer() bool`

GetIsMaintainer returns the IsMaintainer field if non-nil, zero value otherwise.

### GetIsMaintainerOk

`func (o *DbUser) GetIsMaintainerOk() (*bool, bool)`

GetIsMaintainerOk returns a tuple with the IsMaintainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaintainer

`func (o *DbUser) SetIsMaintainer(v bool)`

SetIsMaintainer sets IsMaintainer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


