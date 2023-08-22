# UpdateUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | User Profile Name | 
**Email** | **string** | User Profile Email | 
**Bio** | Pointer to **string** | User Profile Bio | [optional] 
**Url** | Pointer to **string** | User Profile URL | [optional] 
**TwitterUsername** | Pointer to **string** | User Profile Twitter Username | [optional] 
**Company** | Pointer to **string** | User Profile Company | [optional] 
**Location** | Pointer to **string** | User Profile Location | [optional] 
**DisplayLocalTime** | Pointer to **bool** | Display user local time in profile | [optional] 
**Timezone** | Pointer to **string** | User timezone in UTC | [optional] 
**LinkedinUrl** | Pointer to **string** | LinkedIn URL | [optional] 
**GithubSponsorsUrl** | Pointer to **string** | GitHub Sponsors URL | [optional] 
**DiscordUrl** | Pointer to **string** | Discord URL | [optional] 

## Methods

### NewUpdateUserDto

`func NewUpdateUserDto(name string, email string, ) *UpdateUserDto`

NewUpdateUserDto instantiates a new UpdateUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserDtoWithDefaults

`func NewUpdateUserDtoWithDefaults() *UpdateUserDto`

NewUpdateUserDtoWithDefaults instantiates a new UpdateUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateUserDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUserDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUserDto) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *UpdateUserDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUserDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUserDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetBio

`func (o *UpdateUserDto) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *UpdateUserDto) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *UpdateUserDto) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *UpdateUserDto) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateUserDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateUserDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateUserDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateUserDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTwitterUsername

`func (o *UpdateUserDto) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *UpdateUserDto) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *UpdateUserDto) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *UpdateUserDto) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### GetCompany

`func (o *UpdateUserDto) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UpdateUserDto) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UpdateUserDto) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UpdateUserDto) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetLocation

`func (o *UpdateUserDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UpdateUserDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UpdateUserDto) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UpdateUserDto) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDisplayLocalTime

`func (o *UpdateUserDto) GetDisplayLocalTime() bool`

GetDisplayLocalTime returns the DisplayLocalTime field if non-nil, zero value otherwise.

### GetDisplayLocalTimeOk

`func (o *UpdateUserDto) GetDisplayLocalTimeOk() (*bool, bool)`

GetDisplayLocalTimeOk returns a tuple with the DisplayLocalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLocalTime

`func (o *UpdateUserDto) SetDisplayLocalTime(v bool)`

SetDisplayLocalTime sets DisplayLocalTime field to given value.

### HasDisplayLocalTime

`func (o *UpdateUserDto) HasDisplayLocalTime() bool`

HasDisplayLocalTime returns a boolean if a field has been set.

### GetTimezone

`func (o *UpdateUserDto) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdateUserDto) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdateUserDto) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UpdateUserDto) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetLinkedinUrl

`func (o *UpdateUserDto) GetLinkedinUrl() string`

GetLinkedinUrl returns the LinkedinUrl field if non-nil, zero value otherwise.

### GetLinkedinUrlOk

`func (o *UpdateUserDto) GetLinkedinUrlOk() (*string, bool)`

GetLinkedinUrlOk returns a tuple with the LinkedinUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedinUrl

`func (o *UpdateUserDto) SetLinkedinUrl(v string)`

SetLinkedinUrl sets LinkedinUrl field to given value.

### HasLinkedinUrl

`func (o *UpdateUserDto) HasLinkedinUrl() bool`

HasLinkedinUrl returns a boolean if a field has been set.

### GetGithubSponsorsUrl

`func (o *UpdateUserDto) GetGithubSponsorsUrl() string`

GetGithubSponsorsUrl returns the GithubSponsorsUrl field if non-nil, zero value otherwise.

### GetGithubSponsorsUrlOk

`func (o *UpdateUserDto) GetGithubSponsorsUrlOk() (*string, bool)`

GetGithubSponsorsUrlOk returns a tuple with the GithubSponsorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubSponsorsUrl

`func (o *UpdateUserDto) SetGithubSponsorsUrl(v string)`

SetGithubSponsorsUrl sets GithubSponsorsUrl field to given value.

### HasGithubSponsorsUrl

`func (o *UpdateUserDto) HasGithubSponsorsUrl() bool`

HasGithubSponsorsUrl returns a boolean if a field has been set.

### GetDiscordUrl

`func (o *UpdateUserDto) GetDiscordUrl() string`

GetDiscordUrl returns the DiscordUrl field if non-nil, zero value otherwise.

### GetDiscordUrlOk

`func (o *UpdateUserDto) GetDiscordUrlOk() (*string, bool)`

GetDiscordUrlOk returns a tuple with the DiscordUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordUrl

`func (o *UpdateUserDto) SetDiscordUrl(v string)`

SetDiscordUrl sets DiscordUrl field to given value.

### HasDiscordUrl

`func (o *UpdateUserDto) HasDiscordUrl() bool`

HasDiscordUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


