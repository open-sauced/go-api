/*
@open-sauced/api.opensauced.pizza

 ## Swagger-UI API Documentation  This REST API can be used to create, read, update or delete data from the Open Sauced community platform. The Swagger-UI provides useful information to get started and an overview of all available resources. Each API route is clickable and has their own detailed description on how to use it. The base URL for the API is [api.opensauced.pizza](https://api.opensauced.pizza).  [comment]: # (TODO: add bearer auth information)  ## Rate limiting  Every IP address is allowed to perform 5000 requests per hour. This is measured by saving the date of the initial request and counting all requests in the next hour. When an IP address goes over the limit, HTTP status code 429 is returned. The returned HTTP headers of any API request show the current rate limit status:  header | description --- | --- `X-RateLimit-Limit` | The maximum number of requests allowed per hour `X-RateLimit-Remaining` | The number of requests remaining in the current rate limit window `X-RateLimit-Reset` | The date and time at which the current rate limit window resets in [UTC epoch seconds](https://en.wikipedia.org/wiki/Unix_time)  [comment]: # (TODO: add pagination information)  ## Common response codes  Each route shows for each method which data they expect and which they will respond when the call succeeds. The table below shows most common response codes you can receive from our endpoints.  code | condition --- | --- [`200`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200) | The [`GET`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET) request was handled successfully. The response provides the requested data. [`201`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/201) | The [`POST`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST) request was handled successfully. The response provides the created data. [`204`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/204) | The [`PATCH`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PATCH) or [`DELETE`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/DELETE) request was handled successfully. The response provides no data, generally. [`400`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400) | The server will not process the request due to something that is perceived to be a client error. Check the provided error for mote information. [`401`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401) | The request requires user authentication. Check the provided error for more information. [`403`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403) | The request was valid, but the server is refusing user access. Check the provided error for more information. [`404`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404) | The requested resource could not be found. Check the provided error for more information. [`429`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) | The current API Key made too many requests in the last hour. Check [Rate limiting](#ratelimiting) for more information.  ## Additional links

API version: 2
Contact: hello@opensauced.pizza
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the DbPullRequestReviewGitHubEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbPullRequestReviewGitHubEvents{}

// DbPullRequestReviewGitHubEvents struct for DbPullRequestReviewGitHubEvents
type DbPullRequestReviewGitHubEvents struct {
	// Pull request event identifier
	EventId int32 `json:"event_id"`
	// Pull request number
	PrNumber int32 `json:"pr_number"`
	// Pull request state
	PrState string `json:"pr_state"`
	// Pull request is draft
	PrIsDraft bool `json:"pr_is_draft"`
	// Pull request is merged
	PrIsMerged bool `json:"pr_is_merged"`
	// Pull request mergeable state
	PrMergeableState string `json:"pr_mergeable_state"`
	// Pull request is rebaseable
	PrIsRebaseable bool `json:"pr_is_rebaseable"`
	// Pull request title
	PrTitle string `json:"pr_title"`
	// Pull request source ref
	PrHeadLabel string `json:"pr_head_label"`
	// Pull request target ref
	PrBaseLabel string `json:"pr_base_label"`
	// Pull request source branch
	PrHeadRef string `json:"pr_head_ref"`
	// Pull request target branch
	PrBaseRef string `json:"pr_base_ref"`
	// Pull request author username
	PrAuthorLogin string `json:"pr_author_login"`
	// Timestamp representing pr creation date
	PrCreatedAt *time.Time `json:"pr_created_at,omitempty"`
	// Timestamp representing pr close date
	PrClosedAt time.Time `json:"pr_closed_at"`
	// Timestamp representing pr merge date
	PrMergedAt time.Time `json:"pr_merged_at"`
	// Timestamp representing repository last update
	PrUpdatedAt time.Time `json:"pr_updated_at"`
	// PR comments
	PrComments int32 `json:"pr_comments"`
	// PR lines added
	PrAdditions int32 `json:"pr_additions"`
	// PR lines deleted
	PrDeletions int32 `json:"pr_deletions"`
	// PR files changed
	PrChangedFiles int32 `json:"pr_changed_files"`
	// Pull request repo full name
	RepoName string `json:"repo_name"`
	// Number of commits in the PR
	PrCommits int32 `json:"pr_commits"`
	// Pull request review body
	PrReviewBody string `json:"pr_review_body"`
}

// NewDbPullRequestReviewGitHubEvents instantiates a new DbPullRequestReviewGitHubEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbPullRequestReviewGitHubEvents(eventId int32, prNumber int32, prState string, prIsDraft bool, prIsMerged bool, prMergeableState string, prIsRebaseable bool, prTitle string, prHeadLabel string, prBaseLabel string, prHeadRef string, prBaseRef string, prAuthorLogin string, prClosedAt time.Time, prMergedAt time.Time, prUpdatedAt time.Time, prComments int32, prAdditions int32, prDeletions int32, prChangedFiles int32, repoName string, prCommits int32, prReviewBody string) *DbPullRequestReviewGitHubEvents {
	this := DbPullRequestReviewGitHubEvents{}
	this.EventId = eventId
	this.PrNumber = prNumber
	this.PrState = prState
	this.PrIsDraft = prIsDraft
	this.PrIsMerged = prIsMerged
	this.PrMergeableState = prMergeableState
	this.PrIsRebaseable = prIsRebaseable
	this.PrTitle = prTitle
	this.PrHeadLabel = prHeadLabel
	this.PrBaseLabel = prBaseLabel
	this.PrHeadRef = prHeadRef
	this.PrBaseRef = prBaseRef
	this.PrAuthorLogin = prAuthorLogin
	this.PrClosedAt = prClosedAt
	this.PrMergedAt = prMergedAt
	this.PrUpdatedAt = prUpdatedAt
	this.PrComments = prComments
	this.PrAdditions = prAdditions
	this.PrDeletions = prDeletions
	this.PrChangedFiles = prChangedFiles
	this.RepoName = repoName
	this.PrCommits = prCommits
	this.PrReviewBody = prReviewBody
	return &this
}

// NewDbPullRequestReviewGitHubEventsWithDefaults instantiates a new DbPullRequestReviewGitHubEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbPullRequestReviewGitHubEventsWithDefaults() *DbPullRequestReviewGitHubEvents {
	this := DbPullRequestReviewGitHubEvents{}
	return &this
}

// GetEventId returns the EventId field value
func (o *DbPullRequestReviewGitHubEvents) GetEventId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetEventIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *DbPullRequestReviewGitHubEvents) SetEventId(v int32) {
	o.EventId = v
}

// GetPrNumber returns the PrNumber field value
func (o *DbPullRequestReviewGitHubEvents) GetPrNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrNumber
}

// GetPrNumberOk returns a tuple with the PrNumber field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrNumber, true
}

// SetPrNumber sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrNumber(v int32) {
	o.PrNumber = v
}

// GetPrState returns the PrState field value
func (o *DbPullRequestReviewGitHubEvents) GetPrState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrState
}

// GetPrStateOk returns a tuple with the PrState field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrState, true
}

// SetPrState sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrState(v string) {
	o.PrState = v
}

// GetPrIsDraft returns the PrIsDraft field value
func (o *DbPullRequestReviewGitHubEvents) GetPrIsDraft() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PrIsDraft
}

// GetPrIsDraftOk returns a tuple with the PrIsDraft field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrIsDraftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrIsDraft, true
}

// SetPrIsDraft sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrIsDraft(v bool) {
	o.PrIsDraft = v
}

// GetPrIsMerged returns the PrIsMerged field value
func (o *DbPullRequestReviewGitHubEvents) GetPrIsMerged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PrIsMerged
}

// GetPrIsMergedOk returns a tuple with the PrIsMerged field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrIsMergedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrIsMerged, true
}

// SetPrIsMerged sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrIsMerged(v bool) {
	o.PrIsMerged = v
}

// GetPrMergeableState returns the PrMergeableState field value
func (o *DbPullRequestReviewGitHubEvents) GetPrMergeableState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrMergeableState
}

// GetPrMergeableStateOk returns a tuple with the PrMergeableState field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrMergeableStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrMergeableState, true
}

// SetPrMergeableState sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrMergeableState(v string) {
	o.PrMergeableState = v
}

// GetPrIsRebaseable returns the PrIsRebaseable field value
func (o *DbPullRequestReviewGitHubEvents) GetPrIsRebaseable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PrIsRebaseable
}

// GetPrIsRebaseableOk returns a tuple with the PrIsRebaseable field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrIsRebaseableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrIsRebaseable, true
}

// SetPrIsRebaseable sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrIsRebaseable(v bool) {
	o.PrIsRebaseable = v
}

// GetPrTitle returns the PrTitle field value
func (o *DbPullRequestReviewGitHubEvents) GetPrTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrTitle
}

// GetPrTitleOk returns a tuple with the PrTitle field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrTitle, true
}

// SetPrTitle sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrTitle(v string) {
	o.PrTitle = v
}

// GetPrHeadLabel returns the PrHeadLabel field value
func (o *DbPullRequestReviewGitHubEvents) GetPrHeadLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrHeadLabel
}

// GetPrHeadLabelOk returns a tuple with the PrHeadLabel field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrHeadLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrHeadLabel, true
}

// SetPrHeadLabel sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrHeadLabel(v string) {
	o.PrHeadLabel = v
}

// GetPrBaseLabel returns the PrBaseLabel field value
func (o *DbPullRequestReviewGitHubEvents) GetPrBaseLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrBaseLabel
}

// GetPrBaseLabelOk returns a tuple with the PrBaseLabel field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrBaseLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrBaseLabel, true
}

// SetPrBaseLabel sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrBaseLabel(v string) {
	o.PrBaseLabel = v
}

// GetPrHeadRef returns the PrHeadRef field value
func (o *DbPullRequestReviewGitHubEvents) GetPrHeadRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrHeadRef
}

// GetPrHeadRefOk returns a tuple with the PrHeadRef field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrHeadRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrHeadRef, true
}

// SetPrHeadRef sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrHeadRef(v string) {
	o.PrHeadRef = v
}

// GetPrBaseRef returns the PrBaseRef field value
func (o *DbPullRequestReviewGitHubEvents) GetPrBaseRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrBaseRef
}

// GetPrBaseRefOk returns a tuple with the PrBaseRef field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrBaseRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrBaseRef, true
}

// SetPrBaseRef sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrBaseRef(v string) {
	o.PrBaseRef = v
}

// GetPrAuthorLogin returns the PrAuthorLogin field value
func (o *DbPullRequestReviewGitHubEvents) GetPrAuthorLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrAuthorLogin
}

// GetPrAuthorLoginOk returns a tuple with the PrAuthorLogin field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrAuthorLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrAuthorLogin, true
}

// SetPrAuthorLogin sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrAuthorLogin(v string) {
	o.PrAuthorLogin = v
}

// GetPrCreatedAt returns the PrCreatedAt field value if set, zero value otherwise.
func (o *DbPullRequestReviewGitHubEvents) GetPrCreatedAt() time.Time {
	if o == nil || IsNil(o.PrCreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.PrCreatedAt
}

// GetPrCreatedAtOk returns a tuple with the PrCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PrCreatedAt) {
		return nil, false
	}
	return o.PrCreatedAt, true
}

// HasPrCreatedAt returns a boolean if a field has been set.
func (o *DbPullRequestReviewGitHubEvents) HasPrCreatedAt() bool {
	if o != nil && !IsNil(o.PrCreatedAt) {
		return true
	}

	return false
}

// SetPrCreatedAt gets a reference to the given time.Time and assigns it to the PrCreatedAt field.
func (o *DbPullRequestReviewGitHubEvents) SetPrCreatedAt(v time.Time) {
	o.PrCreatedAt = &v
}

// GetPrClosedAt returns the PrClosedAt field value
func (o *DbPullRequestReviewGitHubEvents) GetPrClosedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PrClosedAt
}

// GetPrClosedAtOk returns a tuple with the PrClosedAt field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrClosedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrClosedAt, true
}

// SetPrClosedAt sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrClosedAt(v time.Time) {
	o.PrClosedAt = v
}

// GetPrMergedAt returns the PrMergedAt field value
func (o *DbPullRequestReviewGitHubEvents) GetPrMergedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PrMergedAt
}

// GetPrMergedAtOk returns a tuple with the PrMergedAt field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrMergedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrMergedAt, true
}

// SetPrMergedAt sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrMergedAt(v time.Time) {
	o.PrMergedAt = v
}

// GetPrUpdatedAt returns the PrUpdatedAt field value
func (o *DbPullRequestReviewGitHubEvents) GetPrUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PrUpdatedAt
}

// GetPrUpdatedAtOk returns a tuple with the PrUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrUpdatedAt, true
}

// SetPrUpdatedAt sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrUpdatedAt(v time.Time) {
	o.PrUpdatedAt = v
}

// GetPrComments returns the PrComments field value
func (o *DbPullRequestReviewGitHubEvents) GetPrComments() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrComments
}

// GetPrCommentsOk returns a tuple with the PrComments field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrCommentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrComments, true
}

// SetPrComments sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrComments(v int32) {
	o.PrComments = v
}

// GetPrAdditions returns the PrAdditions field value
func (o *DbPullRequestReviewGitHubEvents) GetPrAdditions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrAdditions
}

// GetPrAdditionsOk returns a tuple with the PrAdditions field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrAdditionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrAdditions, true
}

// SetPrAdditions sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrAdditions(v int32) {
	o.PrAdditions = v
}

// GetPrDeletions returns the PrDeletions field value
func (o *DbPullRequestReviewGitHubEvents) GetPrDeletions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrDeletions
}

// GetPrDeletionsOk returns a tuple with the PrDeletions field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrDeletionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrDeletions, true
}

// SetPrDeletions sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrDeletions(v int32) {
	o.PrDeletions = v
}

// GetPrChangedFiles returns the PrChangedFiles field value
func (o *DbPullRequestReviewGitHubEvents) GetPrChangedFiles() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrChangedFiles
}

// GetPrChangedFilesOk returns a tuple with the PrChangedFiles field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrChangedFilesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrChangedFiles, true
}

// SetPrChangedFiles sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrChangedFiles(v int32) {
	o.PrChangedFiles = v
}

// GetRepoName returns the RepoName field value
func (o *DbPullRequestReviewGitHubEvents) GetRepoName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepoName
}

// GetRepoNameOk returns a tuple with the RepoName field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetRepoNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepoName, true
}

// SetRepoName sets field value
func (o *DbPullRequestReviewGitHubEvents) SetRepoName(v string) {
	o.RepoName = v
}

// GetPrCommits returns the PrCommits field value
func (o *DbPullRequestReviewGitHubEvents) GetPrCommits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrCommits
}

// GetPrCommitsOk returns a tuple with the PrCommits field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrCommitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrCommits, true
}

// SetPrCommits sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrCommits(v int32) {
	o.PrCommits = v
}

// GetPrReviewBody returns the PrReviewBody field value
func (o *DbPullRequestReviewGitHubEvents) GetPrReviewBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrReviewBody
}

// GetPrReviewBodyOk returns a tuple with the PrReviewBody field value
// and a boolean to check if the value has been set.
func (o *DbPullRequestReviewGitHubEvents) GetPrReviewBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrReviewBody, true
}

// SetPrReviewBody sets field value
func (o *DbPullRequestReviewGitHubEvents) SetPrReviewBody(v string) {
	o.PrReviewBody = v
}

func (o DbPullRequestReviewGitHubEvents) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbPullRequestReviewGitHubEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event_id"] = o.EventId
	toSerialize["pr_number"] = o.PrNumber
	toSerialize["pr_state"] = o.PrState
	toSerialize["pr_is_draft"] = o.PrIsDraft
	toSerialize["pr_is_merged"] = o.PrIsMerged
	toSerialize["pr_mergeable_state"] = o.PrMergeableState
	toSerialize["pr_is_rebaseable"] = o.PrIsRebaseable
	toSerialize["pr_title"] = o.PrTitle
	toSerialize["pr_head_label"] = o.PrHeadLabel
	toSerialize["pr_base_label"] = o.PrBaseLabel
	toSerialize["pr_head_ref"] = o.PrHeadRef
	toSerialize["pr_base_ref"] = o.PrBaseRef
	toSerialize["pr_author_login"] = o.PrAuthorLogin
	if !IsNil(o.PrCreatedAt) {
		toSerialize["pr_created_at"] = o.PrCreatedAt
	}
	toSerialize["pr_closed_at"] = o.PrClosedAt
	toSerialize["pr_merged_at"] = o.PrMergedAt
	toSerialize["pr_updated_at"] = o.PrUpdatedAt
	toSerialize["pr_comments"] = o.PrComments
	toSerialize["pr_additions"] = o.PrAdditions
	toSerialize["pr_deletions"] = o.PrDeletions
	toSerialize["pr_changed_files"] = o.PrChangedFiles
	toSerialize["repo_name"] = o.RepoName
	toSerialize["pr_commits"] = o.PrCommits
	toSerialize["pr_review_body"] = o.PrReviewBody
	return toSerialize, nil
}

type NullableDbPullRequestReviewGitHubEvents struct {
	value *DbPullRequestReviewGitHubEvents
	isSet bool
}

func (v NullableDbPullRequestReviewGitHubEvents) Get() *DbPullRequestReviewGitHubEvents {
	return v.value
}

func (v *NullableDbPullRequestReviewGitHubEvents) Set(val *DbPullRequestReviewGitHubEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableDbPullRequestReviewGitHubEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableDbPullRequestReviewGitHubEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbPullRequestReviewGitHubEvents(val *DbPullRequestReviewGitHubEvents) *NullableDbPullRequestReviewGitHubEvents {
	return &NullableDbPullRequestReviewGitHubEvents{value: val, isSet: true}
}

func (v NullableDbPullRequestReviewGitHubEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbPullRequestReviewGitHubEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
