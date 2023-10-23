# SupabaseAuthDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Supabase authenticated unique user identifier | 
**UserName** | Pointer to **string** | Supabase authenticated user login name | [optional] 
**Role** | Pointer to **string** | Supabase authenticated user role | [optional] 
**Email** | Pointer to **string** | Supabase authenticated user email | [optional] 
**ConfirmedAt** | Pointer to **string** | Timestamp representing supabase user registration confirmation | [optional] 
**LastSignInAt** | Pointer to **string** | Timestamp representing supabase user last sign in | [optional] 
**CreatedAt** | Pointer to **string** | Timestamp representing supabase user creation | [optional] 
**UpdatedAt** | Pointer to **string** | Timestamp representing supabase user last update | [optional] 
**IsOnboarded** | Pointer to **bool** | Flag indicated whether the user is onboarded | [optional] 
**IsWaitlisted** | Pointer to **bool** | Flag indicated whether the user is waitlisted | [optional] 
**InsightsRole** | Pointer to **int32** | Authenticated User&#39;s Insights Role | [optional] 
**Bio** | Pointer to **string** | User bio information | [optional] 
**Name** | Pointer to **string** | User name information | [optional] 
**Url** | Pointer to **string** | User website | [optional] 
**TwitterUsername** | Pointer to **string** | User Twitter information | [optional] 
**Company** | Pointer to **string** | User company information | [optional] 
**Location** | Pointer to **string** | User location information | [optional] 
**DisplayLocalTime** | Pointer to **bool** | User display local time information | [optional] 
**LinkedinUrl** | Pointer to **string** | LinkedIn URL | [optional] 
**GithubSponsorsUrl** | Pointer to **string** | GitHub Sponsors URL | [optional] 
**DiscordUrl** | Pointer to **string** | Discord URL | [optional] 
**NotificationCount** | Pointer to **int32** | Unread User Notification Count | [optional] 
**InsightsCount** | Pointer to **int32** | Unread Insight Pagees Count | [optional] 
**CouponCode** | Pointer to **string** | Coupon Code | [optional] 

## Methods

### NewSupabaseAuthDto

`func NewSupabaseAuthDto(id string, ) *SupabaseAuthDto`

NewSupabaseAuthDto instantiates a new SupabaseAuthDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupabaseAuthDtoWithDefaults

`func NewSupabaseAuthDtoWithDefaults() *SupabaseAuthDto`

NewSupabaseAuthDtoWithDefaults instantiates a new SupabaseAuthDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupabaseAuthDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupabaseAuthDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupabaseAuthDto) SetId(v string)`

SetId sets Id field to given value.


### GetUserName

`func (o *SupabaseAuthDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SupabaseAuthDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SupabaseAuthDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SupabaseAuthDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetRole

`func (o *SupabaseAuthDto) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SupabaseAuthDto) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SupabaseAuthDto) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SupabaseAuthDto) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetEmail

`func (o *SupabaseAuthDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SupabaseAuthDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SupabaseAuthDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SupabaseAuthDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetConfirmedAt

`func (o *SupabaseAuthDto) GetConfirmedAt() string`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *SupabaseAuthDto) GetConfirmedAtOk() (*string, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *SupabaseAuthDto) SetConfirmedAt(v string)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *SupabaseAuthDto) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### GetLastSignInAt

`func (o *SupabaseAuthDto) GetLastSignInAt() string`

GetLastSignInAt returns the LastSignInAt field if non-nil, zero value otherwise.

### GetLastSignInAtOk

`func (o *SupabaseAuthDto) GetLastSignInAtOk() (*string, bool)`

GetLastSignInAtOk returns a tuple with the LastSignInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSignInAt

`func (o *SupabaseAuthDto) SetLastSignInAt(v string)`

SetLastSignInAt sets LastSignInAt field to given value.

### HasLastSignInAt

`func (o *SupabaseAuthDto) HasLastSignInAt() bool`

HasLastSignInAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SupabaseAuthDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SupabaseAuthDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SupabaseAuthDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SupabaseAuthDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SupabaseAuthDto) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SupabaseAuthDto) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SupabaseAuthDto) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SupabaseAuthDto) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetIsOnboarded

`func (o *SupabaseAuthDto) GetIsOnboarded() bool`

GetIsOnboarded returns the IsOnboarded field if non-nil, zero value otherwise.

### GetIsOnboardedOk

`func (o *SupabaseAuthDto) GetIsOnboardedOk() (*bool, bool)`

GetIsOnboardedOk returns a tuple with the IsOnboarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnboarded

`func (o *SupabaseAuthDto) SetIsOnboarded(v bool)`

SetIsOnboarded sets IsOnboarded field to given value.

### HasIsOnboarded

`func (o *SupabaseAuthDto) HasIsOnboarded() bool`

HasIsOnboarded returns a boolean if a field has been set.

### GetIsWaitlisted

`func (o *SupabaseAuthDto) GetIsWaitlisted() bool`

GetIsWaitlisted returns the IsWaitlisted field if non-nil, zero value otherwise.

### GetIsWaitlistedOk

`func (o *SupabaseAuthDto) GetIsWaitlistedOk() (*bool, bool)`

GetIsWaitlistedOk returns a tuple with the IsWaitlisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWaitlisted

`func (o *SupabaseAuthDto) SetIsWaitlisted(v bool)`

SetIsWaitlisted sets IsWaitlisted field to given value.

### HasIsWaitlisted

`func (o *SupabaseAuthDto) HasIsWaitlisted() bool`

HasIsWaitlisted returns a boolean if a field has been set.

### GetInsightsRole

`func (o *SupabaseAuthDto) GetInsightsRole() int32`

GetInsightsRole returns the InsightsRole field if non-nil, zero value otherwise.

### GetInsightsRoleOk

`func (o *SupabaseAuthDto) GetInsightsRoleOk() (*int32, bool)`

GetInsightsRoleOk returns a tuple with the InsightsRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsRole

`func (o *SupabaseAuthDto) SetInsightsRole(v int32)`

SetInsightsRole sets InsightsRole field to given value.

### HasInsightsRole

`func (o *SupabaseAuthDto) HasInsightsRole() bool`

HasInsightsRole returns a boolean if a field has been set.

### GetBio

`func (o *SupabaseAuthDto) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *SupabaseAuthDto) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *SupabaseAuthDto) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *SupabaseAuthDto) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetName

`func (o *SupabaseAuthDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupabaseAuthDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupabaseAuthDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SupabaseAuthDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *SupabaseAuthDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SupabaseAuthDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SupabaseAuthDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SupabaseAuthDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTwitterUsername

`func (o *SupabaseAuthDto) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *SupabaseAuthDto) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *SupabaseAuthDto) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *SupabaseAuthDto) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### GetCompany

`func (o *SupabaseAuthDto) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *SupabaseAuthDto) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *SupabaseAuthDto) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *SupabaseAuthDto) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetLocation

`func (o *SupabaseAuthDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SupabaseAuthDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SupabaseAuthDto) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SupabaseAuthDto) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDisplayLocalTime

`func (o *SupabaseAuthDto) GetDisplayLocalTime() bool`

GetDisplayLocalTime returns the DisplayLocalTime field if non-nil, zero value otherwise.

### GetDisplayLocalTimeOk

`func (o *SupabaseAuthDto) GetDisplayLocalTimeOk() (*bool, bool)`

GetDisplayLocalTimeOk returns a tuple with the DisplayLocalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLocalTime

`func (o *SupabaseAuthDto) SetDisplayLocalTime(v bool)`

SetDisplayLocalTime sets DisplayLocalTime field to given value.

### HasDisplayLocalTime

`func (o *SupabaseAuthDto) HasDisplayLocalTime() bool`

HasDisplayLocalTime returns a boolean if a field has been set.

### GetLinkedinUrl

`func (o *SupabaseAuthDto) GetLinkedinUrl() string`

GetLinkedinUrl returns the LinkedinUrl field if non-nil, zero value otherwise.

### GetLinkedinUrlOk

`func (o *SupabaseAuthDto) GetLinkedinUrlOk() (*string, bool)`

GetLinkedinUrlOk returns a tuple with the LinkedinUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedinUrl

`func (o *SupabaseAuthDto) SetLinkedinUrl(v string)`

SetLinkedinUrl sets LinkedinUrl field to given value.

### HasLinkedinUrl

`func (o *SupabaseAuthDto) HasLinkedinUrl() bool`

HasLinkedinUrl returns a boolean if a field has been set.

### GetGithubSponsorsUrl

`func (o *SupabaseAuthDto) GetGithubSponsorsUrl() string`

GetGithubSponsorsUrl returns the GithubSponsorsUrl field if non-nil, zero value otherwise.

### GetGithubSponsorsUrlOk

`func (o *SupabaseAuthDto) GetGithubSponsorsUrlOk() (*string, bool)`

GetGithubSponsorsUrlOk returns a tuple with the GithubSponsorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubSponsorsUrl

`func (o *SupabaseAuthDto) SetGithubSponsorsUrl(v string)`

SetGithubSponsorsUrl sets GithubSponsorsUrl field to given value.

### HasGithubSponsorsUrl

`func (o *SupabaseAuthDto) HasGithubSponsorsUrl() bool`

HasGithubSponsorsUrl returns a boolean if a field has been set.

### GetDiscordUrl

`func (o *SupabaseAuthDto) GetDiscordUrl() string`

GetDiscordUrl returns the DiscordUrl field if non-nil, zero value otherwise.

### GetDiscordUrlOk

`func (o *SupabaseAuthDto) GetDiscordUrlOk() (*string, bool)`

GetDiscordUrlOk returns a tuple with the DiscordUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordUrl

`func (o *SupabaseAuthDto) SetDiscordUrl(v string)`

SetDiscordUrl sets DiscordUrl field to given value.

### HasDiscordUrl

`func (o *SupabaseAuthDto) HasDiscordUrl() bool`

HasDiscordUrl returns a boolean if a field has been set.

### GetNotificationCount

`func (o *SupabaseAuthDto) GetNotificationCount() int32`

GetNotificationCount returns the NotificationCount field if non-nil, zero value otherwise.

### GetNotificationCountOk

`func (o *SupabaseAuthDto) GetNotificationCountOk() (*int32, bool)`

GetNotificationCountOk returns a tuple with the NotificationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationCount

`func (o *SupabaseAuthDto) SetNotificationCount(v int32)`

SetNotificationCount sets NotificationCount field to given value.

### HasNotificationCount

`func (o *SupabaseAuthDto) HasNotificationCount() bool`

HasNotificationCount returns a boolean if a field has been set.

### GetInsightsCount

`func (o *SupabaseAuthDto) GetInsightsCount() int32`

GetInsightsCount returns the InsightsCount field if non-nil, zero value otherwise.

### GetInsightsCountOk

`func (o *SupabaseAuthDto) GetInsightsCountOk() (*int32, bool)`

GetInsightsCountOk returns a tuple with the InsightsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsCount

`func (o *SupabaseAuthDto) SetInsightsCount(v int32)`

SetInsightsCount sets InsightsCount field to given value.

### HasInsightsCount

`func (o *SupabaseAuthDto) HasInsightsCount() bool`

HasInsightsCount returns a boolean if a field has been set.

### GetCouponCode

`func (o *SupabaseAuthDto) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *SupabaseAuthDto) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *SupabaseAuthDto) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *SupabaseAuthDto) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


