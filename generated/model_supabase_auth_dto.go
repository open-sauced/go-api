/*
@open-sauced/api.opensauced.pizza

 ## Swagger-UI API Documentation  This REST API can be used to create, read, update or delete data from the Open Sauced community platform. The Swagger-UI provides useful information to get started and an overview of all available resources. Each API route is clickable and has their own detailed description on how to use it. The base URL for the API is [https://api.opensauced.pizza](https://api.opensauced.pizza).  [comment]: # (TODO: add bearer auth information)  ## Rate limiting  Every IP address is allowed to perform 5000 requests per hour. This is measured by saving the date of the initial request and counting all requests in the next hour. When an IP address goes over the limit, HTTP status code 429 is returned. The returned HTTP headers of any API request show the current rate limit status:  header | description --- | --- `X-RateLimit-Limit` | The maximum number of requests allowed per hour `X-RateLimit-Remaining` | The number of requests remaining in the current rate limit window `X-RateLimit-Reset` | The date and time at which the current rate limit window resets in [UTC epoch seconds](https://en.wikipedia.org/wiki/Unix_time)  [comment]: # (TODO: add pagination information)  ## Common response codes  Each route shows for each method which data they expect and which they will respond when the call succeeds. The table below shows most common response codes you can receive from our endpoints.  code | condition --- | --- [`200`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200) | The [`GET`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET) request was handled successfully. The response provides the requested data. [`201`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/201) | The [`POST`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST) request was handled successfully. The response provides the created data. [`204`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/204) | The [`PATCH`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PATCH) or [`DELETE`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/DELETE) request was handled successfully. The response provides no data, generally. [`400`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400) | The server will not process the request due to something that is perceived to be a client error. Check the provided error for mote information. [`401`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401) | The request requires user authentication. Check the provided error for more information. [`403`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403) | The request was valid, but the server is refusing user access. Check the provided error for more information. [`404`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404) | The requested resource could not be found. Check the provided error for more information. [`429`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) | The current API Key made too many requests in the last hour. Check [Rate limiting](#ratelimiting) for more information.  ## Additional links

API version: 1
Contact: hello@opensauced.pizza
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SupabaseAuthDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupabaseAuthDto{}

// SupabaseAuthDto struct for SupabaseAuthDto
type SupabaseAuthDto struct {
	// Supabase authenticated unique user identifier
	Id string `json:"id"`
	// Supabase authenticated user login name
	UserName *string `json:"user_name,omitempty"`
	// Supabase authenticated user role
	Role *string `json:"role,omitempty"`
	// Supabase authenticated user email
	Email *string `json:"email,omitempty"`
	// Timestamp representing supabase user registration confirmation
	ConfirmedAt *string `json:"confirmed_at,omitempty"`
	// Timestamp representing supabase user last sign in
	LastSignInAt *string `json:"last_sign_in_at,omitempty"`
	// Timestamp representing supabase user creation
	CreatedAt *string `json:"created_at,omitempty"`
	// Timestamp representing supabase user last update
	UpdatedAt *string `json:"updated_at,omitempty"`
	// Flag indicated whether the user is onboarded
	IsOnboarded *bool `json:"is_onboarded,omitempty"`
	// Flag indicated whether the user is waitlisted
	IsWaitlisted *bool `json:"is_waitlisted,omitempty"`
	// Authenticated User's Insights Role
	InsightsRole *float32 `json:"insights_role,omitempty"`
	// User bio information
	Bio *string `json:"bio,omitempty"`
	// User name information
	Name *string `json:"name,omitempty"`
	// User website
	Url *string `json:"url,omitempty"`
	// User Twitter information
	TwitterUsername *string `json:"twitter_username,omitempty"`
	// User company information
	Company *string `json:"company,omitempty"`
	// User location information
	Location *string `json:"location,omitempty"`
	// User display local time information
	DisplayLocalTime *bool `json:"display_local_time,omitempty"`
	// LinkedIn URL
	LinkedinUrl *string `json:"linkedin_url,omitempty"`
	// GitHub Sponsors URL
	GithubSponsorsUrl *string `json:"github_sponsors_url,omitempty"`
	// Discord URL
	DiscordUrl *string `json:"discord_url,omitempty"`
	// Unread User Notification Count
	NotificationCount *float32 `json:"notification_count,omitempty"`
}

// NewSupabaseAuthDto instantiates a new SupabaseAuthDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupabaseAuthDto(id string) *SupabaseAuthDto {
	this := SupabaseAuthDto{}
	this.Id = id
	return &this
}

// NewSupabaseAuthDtoWithDefaults instantiates a new SupabaseAuthDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupabaseAuthDtoWithDefaults() *SupabaseAuthDto {
	this := SupabaseAuthDto{}
	return &this
}

// GetId returns the Id field value
func (o *SupabaseAuthDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SupabaseAuthDto) SetId(v string) {
	o.Id = v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *SupabaseAuthDto) SetUserName(v string) {
	o.UserName = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *SupabaseAuthDto) SetRole(v string) {
	o.Role = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SupabaseAuthDto) SetEmail(v string) {
	o.Email = &v
}

// GetConfirmedAt returns the ConfirmedAt field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetConfirmedAt() string {
	if o == nil || IsNil(o.ConfirmedAt) {
		var ret string
		return ret
	}
	return *o.ConfirmedAt
}

// GetConfirmedAtOk returns a tuple with the ConfirmedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetConfirmedAtOk() (*string, bool) {
	if o == nil || IsNil(o.ConfirmedAt) {
		return nil, false
	}
	return o.ConfirmedAt, true
}

// HasConfirmedAt returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasConfirmedAt() bool {
	if o != nil && !IsNil(o.ConfirmedAt) {
		return true
	}

	return false
}

// SetConfirmedAt gets a reference to the given string and assigns it to the ConfirmedAt field.
func (o *SupabaseAuthDto) SetConfirmedAt(v string) {
	o.ConfirmedAt = &v
}

// GetLastSignInAt returns the LastSignInAt field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetLastSignInAt() string {
	if o == nil || IsNil(o.LastSignInAt) {
		var ret string
		return ret
	}
	return *o.LastSignInAt
}

// GetLastSignInAtOk returns a tuple with the LastSignInAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetLastSignInAtOk() (*string, bool) {
	if o == nil || IsNil(o.LastSignInAt) {
		return nil, false
	}
	return o.LastSignInAt, true
}

// HasLastSignInAt returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasLastSignInAt() bool {
	if o != nil && !IsNil(o.LastSignInAt) {
		return true
	}

	return false
}

// SetLastSignInAt gets a reference to the given string and assigns it to the LastSignInAt field.
func (o *SupabaseAuthDto) SetLastSignInAt(v string) {
	o.LastSignInAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SupabaseAuthDto) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *SupabaseAuthDto) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetIsOnboarded returns the IsOnboarded field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetIsOnboarded() bool {
	if o == nil || IsNil(o.IsOnboarded) {
		var ret bool
		return ret
	}
	return *o.IsOnboarded
}

// GetIsOnboardedOk returns a tuple with the IsOnboarded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetIsOnboardedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOnboarded) {
		return nil, false
	}
	return o.IsOnboarded, true
}

// HasIsOnboarded returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasIsOnboarded() bool {
	if o != nil && !IsNil(o.IsOnboarded) {
		return true
	}

	return false
}

// SetIsOnboarded gets a reference to the given bool and assigns it to the IsOnboarded field.
func (o *SupabaseAuthDto) SetIsOnboarded(v bool) {
	o.IsOnboarded = &v
}

// GetIsWaitlisted returns the IsWaitlisted field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetIsWaitlisted() bool {
	if o == nil || IsNil(o.IsWaitlisted) {
		var ret bool
		return ret
	}
	return *o.IsWaitlisted
}

// GetIsWaitlistedOk returns a tuple with the IsWaitlisted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetIsWaitlistedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWaitlisted) {
		return nil, false
	}
	return o.IsWaitlisted, true
}

// HasIsWaitlisted returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasIsWaitlisted() bool {
	if o != nil && !IsNil(o.IsWaitlisted) {
		return true
	}

	return false
}

// SetIsWaitlisted gets a reference to the given bool and assigns it to the IsWaitlisted field.
func (o *SupabaseAuthDto) SetIsWaitlisted(v bool) {
	o.IsWaitlisted = &v
}

// GetInsightsRole returns the InsightsRole field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetInsightsRole() float32 {
	if o == nil || IsNil(o.InsightsRole) {
		var ret float32
		return ret
	}
	return *o.InsightsRole
}

// GetInsightsRoleOk returns a tuple with the InsightsRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetInsightsRoleOk() (*float32, bool) {
	if o == nil || IsNil(o.InsightsRole) {
		return nil, false
	}
	return o.InsightsRole, true
}

// HasInsightsRole returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasInsightsRole() bool {
	if o != nil && !IsNil(o.InsightsRole) {
		return true
	}

	return false
}

// SetInsightsRole gets a reference to the given float32 and assigns it to the InsightsRole field.
func (o *SupabaseAuthDto) SetInsightsRole(v float32) {
	o.InsightsRole = &v
}

// GetBio returns the Bio field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *SupabaseAuthDto) SetBio(v string) {
	o.Bio = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SupabaseAuthDto) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SupabaseAuthDto) SetUrl(v string) {
	o.Url = &v
}

// GetTwitterUsername returns the TwitterUsername field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetTwitterUsername() string {
	if o == nil || IsNil(o.TwitterUsername) {
		var ret string
		return ret
	}
	return *o.TwitterUsername
}

// GetTwitterUsernameOk returns a tuple with the TwitterUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetTwitterUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.TwitterUsername) {
		return nil, false
	}
	return o.TwitterUsername, true
}

// HasTwitterUsername returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasTwitterUsername() bool {
	if o != nil && !IsNil(o.TwitterUsername) {
		return true
	}

	return false
}

// SetTwitterUsername gets a reference to the given string and assigns it to the TwitterUsername field.
func (o *SupabaseAuthDto) SetTwitterUsername(v string) {
	o.TwitterUsername = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *SupabaseAuthDto) SetCompany(v string) {
	o.Company = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *SupabaseAuthDto) SetLocation(v string) {
	o.Location = &v
}

// GetDisplayLocalTime returns the DisplayLocalTime field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetDisplayLocalTime() bool {
	if o == nil || IsNil(o.DisplayLocalTime) {
		var ret bool
		return ret
	}
	return *o.DisplayLocalTime
}

// GetDisplayLocalTimeOk returns a tuple with the DisplayLocalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetDisplayLocalTimeOk() (*bool, bool) {
	if o == nil || IsNil(o.DisplayLocalTime) {
		return nil, false
	}
	return o.DisplayLocalTime, true
}

// HasDisplayLocalTime returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasDisplayLocalTime() bool {
	if o != nil && !IsNil(o.DisplayLocalTime) {
		return true
	}

	return false
}

// SetDisplayLocalTime gets a reference to the given bool and assigns it to the DisplayLocalTime field.
func (o *SupabaseAuthDto) SetDisplayLocalTime(v bool) {
	o.DisplayLocalTime = &v
}

// GetLinkedinUrl returns the LinkedinUrl field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetLinkedinUrl() string {
	if o == nil || IsNil(o.LinkedinUrl) {
		var ret string
		return ret
	}
	return *o.LinkedinUrl
}

// GetLinkedinUrlOk returns a tuple with the LinkedinUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetLinkedinUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LinkedinUrl) {
		return nil, false
	}
	return o.LinkedinUrl, true
}

// HasLinkedinUrl returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasLinkedinUrl() bool {
	if o != nil && !IsNil(o.LinkedinUrl) {
		return true
	}

	return false
}

// SetLinkedinUrl gets a reference to the given string and assigns it to the LinkedinUrl field.
func (o *SupabaseAuthDto) SetLinkedinUrl(v string) {
	o.LinkedinUrl = &v
}

// GetGithubSponsorsUrl returns the GithubSponsorsUrl field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetGithubSponsorsUrl() string {
	if o == nil || IsNil(o.GithubSponsorsUrl) {
		var ret string
		return ret
	}
	return *o.GithubSponsorsUrl
}

// GetGithubSponsorsUrlOk returns a tuple with the GithubSponsorsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetGithubSponsorsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.GithubSponsorsUrl) {
		return nil, false
	}
	return o.GithubSponsorsUrl, true
}

// HasGithubSponsorsUrl returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasGithubSponsorsUrl() bool {
	if o != nil && !IsNil(o.GithubSponsorsUrl) {
		return true
	}

	return false
}

// SetGithubSponsorsUrl gets a reference to the given string and assigns it to the GithubSponsorsUrl field.
func (o *SupabaseAuthDto) SetGithubSponsorsUrl(v string) {
	o.GithubSponsorsUrl = &v
}

// GetDiscordUrl returns the DiscordUrl field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetDiscordUrl() string {
	if o == nil || IsNil(o.DiscordUrl) {
		var ret string
		return ret
	}
	return *o.DiscordUrl
}

// GetDiscordUrlOk returns a tuple with the DiscordUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetDiscordUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DiscordUrl) {
		return nil, false
	}
	return o.DiscordUrl, true
}

// HasDiscordUrl returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasDiscordUrl() bool {
	if o != nil && !IsNil(o.DiscordUrl) {
		return true
	}

	return false
}

// SetDiscordUrl gets a reference to the given string and assigns it to the DiscordUrl field.
func (o *SupabaseAuthDto) SetDiscordUrl(v string) {
	o.DiscordUrl = &v
}

// GetNotificationCount returns the NotificationCount field value if set, zero value otherwise.
func (o *SupabaseAuthDto) GetNotificationCount() float32 {
	if o == nil || IsNil(o.NotificationCount) {
		var ret float32
		return ret
	}
	return *o.NotificationCount
}

// GetNotificationCountOk returns a tuple with the NotificationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupabaseAuthDto) GetNotificationCountOk() (*float32, bool) {
	if o == nil || IsNil(o.NotificationCount) {
		return nil, false
	}
	return o.NotificationCount, true
}

// HasNotificationCount returns a boolean if a field has been set.
func (o *SupabaseAuthDto) HasNotificationCount() bool {
	if o != nil && !IsNil(o.NotificationCount) {
		return true
	}

	return false
}

// SetNotificationCount gets a reference to the given float32 and assigns it to the NotificationCount field.
func (o *SupabaseAuthDto) SetNotificationCount(v float32) {
	o.NotificationCount = &v
}

func (o SupabaseAuthDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupabaseAuthDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.UserName) {
		toSerialize["user_name"] = o.UserName
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.ConfirmedAt) {
		toSerialize["confirmed_at"] = o.ConfirmedAt
	}
	if !IsNil(o.LastSignInAt) {
		toSerialize["last_sign_in_at"] = o.LastSignInAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.IsOnboarded) {
		toSerialize["is_onboarded"] = o.IsOnboarded
	}
	if !IsNil(o.IsWaitlisted) {
		toSerialize["is_waitlisted"] = o.IsWaitlisted
	}
	if !IsNil(o.InsightsRole) {
		toSerialize["insights_role"] = o.InsightsRole
	}
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.TwitterUsername) {
		toSerialize["twitter_username"] = o.TwitterUsername
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.DisplayLocalTime) {
		toSerialize["display_local_time"] = o.DisplayLocalTime
	}
	if !IsNil(o.LinkedinUrl) {
		toSerialize["linkedin_url"] = o.LinkedinUrl
	}
	if !IsNil(o.GithubSponsorsUrl) {
		toSerialize["github_sponsors_url"] = o.GithubSponsorsUrl
	}
	if !IsNil(o.DiscordUrl) {
		toSerialize["discord_url"] = o.DiscordUrl
	}
	if !IsNil(o.NotificationCount) {
		toSerialize["notification_count"] = o.NotificationCount
	}
	return toSerialize, nil
}

type NullableSupabaseAuthDto struct {
	value *SupabaseAuthDto
	isSet bool
}

func (v NullableSupabaseAuthDto) Get() *SupabaseAuthDto {
	return v.value
}

func (v *NullableSupabaseAuthDto) Set(val *SupabaseAuthDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSupabaseAuthDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSupabaseAuthDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupabaseAuthDto(val *SupabaseAuthDto) *NullableSupabaseAuthDto {
	return &NullableSupabaseAuthDto{value: val, isSet: true}
}

func (v NullableSupabaseAuthDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupabaseAuthDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


