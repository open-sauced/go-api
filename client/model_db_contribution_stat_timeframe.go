/*
@open-sauced/api.opensauced.pizza

 ## Swagger-UI API Documentation  This REST API can be used to create, read, update or delete data from the Open Sauced community platform. The Swagger-UI provides useful information to get started and an overview of all available resources. Each API route is clickable and has their own detailed description on how to use it. The base URL for the API is [api.opensauced.pizza](https://api.opensauced.pizza).  [comment]: # (TODO: add bearer auth information)  ## Rate limiting  Every IP address is allowed to perform 5000 requests per hour. This is measured by saving the date of the initial request and counting all requests in the next hour. When an IP address goes over the limit, HTTP status code 429 is returned. The returned HTTP headers of any API request show the current rate limit status:  header | description --- | --- `X-RateLimit-Limit` | The maximum number of requests allowed per hour `X-RateLimit-Remaining` | The number of requests remaining in the current rate limit window `X-RateLimit-Reset` | The date and time at which the current rate limit window resets in [UTC epoch seconds](https://en.wikipedia.org/wiki/Unix_time)  [comment]: # (TODO: add pagination information)  ## Common response codes  Each route shows for each method which data they expect and which they will respond when the call succeeds. The table below shows most common response codes you can receive from our endpoints.  code | condition --- | --- [`200`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200) | The [`GET`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET) request was handled successfully. The response provides the requested data. [`201`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/201) | The [`POST`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST) request was handled successfully. The response provides the created data. [`204`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/204) | The [`PATCH`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PATCH) or [`DELETE`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/DELETE) request was handled successfully. The response provides no data, generally. [`400`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400) | The server will not process the request due to something that is perceived to be a client error. Check the provided error for mote information. [`401`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401) | The request requires user authentication. Check the provided error for more information. [`403`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403) | The request was valid, but the server is refusing user access. Check the provided error for more information. [`404`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404) | The requested resource could not be found. Check the provided error for more information. [`429`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) | The current API Key made too many requests in the last hour. Check [Rate limiting](#ratelimiting) for more information.  ## Additional links

API version: 1
Contact: hello@opensauced.pizza
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the DbContributionStatTimeframe type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbContributionStatTimeframe{}

// DbContributionStatTimeframe struct for DbContributionStatTimeframe
type DbContributionStatTimeframe struct {
	// The ISO timestamp for the start of the time frame
	TimeStart string `json:"time_start"`
	// The ISO timestamp for the end of the time frame
	TimeEnd string `json:"time_end"`
	// Number of commits associated with a user login
	Commits int32 `json:"commits"`
	// Number of PRs created associated with a user login
	PrsCreated int32 `json:"prs_created"`
	// Number of PRs reviewed by a user login
	PrsReviewed int32 `json:"prs_reviewed"`
	// Number of issues created by a user login
	IssuesCreated int32 `json:"issues_created"`
	// Number of comments associated with a user login
	Comments int32 `json:"comments"`
}

// NewDbContributionStatTimeframe instantiates a new DbContributionStatTimeframe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbContributionStatTimeframe(timeStart string, timeEnd string, commits int32, prsCreated int32, prsReviewed int32, issuesCreated int32, comments int32) *DbContributionStatTimeframe {
	this := DbContributionStatTimeframe{}
	this.TimeStart = timeStart
	this.TimeEnd = timeEnd
	this.Commits = commits
	this.PrsCreated = prsCreated
	this.PrsReviewed = prsReviewed
	this.IssuesCreated = issuesCreated
	this.Comments = comments
	return &this
}

// NewDbContributionStatTimeframeWithDefaults instantiates a new DbContributionStatTimeframe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbContributionStatTimeframeWithDefaults() *DbContributionStatTimeframe {
	this := DbContributionStatTimeframe{}
	return &this
}

// GetTimeStart returns the TimeStart field value
func (o *DbContributionStatTimeframe) GetTimeStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeStart
}

// GetTimeStartOk returns a tuple with the TimeStart field value
// and a boolean to check if the value has been set.
func (o *DbContributionStatTimeframe) GetTimeStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStart, true
}

// SetTimeStart sets field value
func (o *DbContributionStatTimeframe) SetTimeStart(v string) {
	o.TimeStart = v
}

// GetTimeEnd returns the TimeEnd field value
func (o *DbContributionStatTimeframe) GetTimeEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeEnd
}

// GetTimeEndOk returns a tuple with the TimeEnd field value
// and a boolean to check if the value has been set.
func (o *DbContributionStatTimeframe) GetTimeEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeEnd, true
}

// SetTimeEnd sets field value
func (o *DbContributionStatTimeframe) SetTimeEnd(v string) {
	o.TimeEnd = v
}

// GetCommits returns the Commits field value
func (o *DbContributionStatTimeframe) GetCommits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Commits
}

// GetCommitsOk returns a tuple with the Commits field value
// and a boolean to check if the value has been set.
func (o *DbContributionStatTimeframe) GetCommitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commits, true
}

// SetCommits sets field value
func (o *DbContributionStatTimeframe) SetCommits(v int32) {
	o.Commits = v
}

// GetPrsCreated returns the PrsCreated field value
func (o *DbContributionStatTimeframe) GetPrsCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrsCreated
}

// GetPrsCreatedOk returns a tuple with the PrsCreated field value
// and a boolean to check if the value has been set.
func (o *DbContributionStatTimeframe) GetPrsCreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrsCreated, true
}

// SetPrsCreated sets field value
func (o *DbContributionStatTimeframe) SetPrsCreated(v int32) {
	o.PrsCreated = v
}

// GetPrsReviewed returns the PrsReviewed field value
func (o *DbContributionStatTimeframe) GetPrsReviewed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrsReviewed
}

// GetPrsReviewedOk returns a tuple with the PrsReviewed field value
// and a boolean to check if the value has been set.
func (o *DbContributionStatTimeframe) GetPrsReviewedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrsReviewed, true
}

// SetPrsReviewed sets field value
func (o *DbContributionStatTimeframe) SetPrsReviewed(v int32) {
	o.PrsReviewed = v
}

// GetIssuesCreated returns the IssuesCreated field value
func (o *DbContributionStatTimeframe) GetIssuesCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IssuesCreated
}

// GetIssuesCreatedOk returns a tuple with the IssuesCreated field value
// and a boolean to check if the value has been set.
func (o *DbContributionStatTimeframe) GetIssuesCreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuesCreated, true
}

// SetIssuesCreated sets field value
func (o *DbContributionStatTimeframe) SetIssuesCreated(v int32) {
	o.IssuesCreated = v
}

// GetComments returns the Comments field value
func (o *DbContributionStatTimeframe) GetComments() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *DbContributionStatTimeframe) GetCommentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comments, true
}

// SetComments sets field value
func (o *DbContributionStatTimeframe) SetComments(v int32) {
	o.Comments = v
}

func (o DbContributionStatTimeframe) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbContributionStatTimeframe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["time_start"] = o.TimeStart
	toSerialize["time_end"] = o.TimeEnd
	toSerialize["commits"] = o.Commits
	toSerialize["prs_created"] = o.PrsCreated
	toSerialize["prs_reviewed"] = o.PrsReviewed
	toSerialize["issues_created"] = o.IssuesCreated
	toSerialize["comments"] = o.Comments
	return toSerialize, nil
}

type NullableDbContributionStatTimeframe struct {
	value *DbContributionStatTimeframe
	isSet bool
}

func (v NullableDbContributionStatTimeframe) Get() *DbContributionStatTimeframe {
	return v.value
}

func (v *NullableDbContributionStatTimeframe) Set(val *DbContributionStatTimeframe) {
	v.value = val
	v.isSet = true
}

func (v NullableDbContributionStatTimeframe) IsSet() bool {
	return v.isSet
}

func (v *NullableDbContributionStatTimeframe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbContributionStatTimeframe(val *DbContributionStatTimeframe) *NullableDbContributionStatTimeframe {
	return &NullableDbContributionStatTimeframe{value: val, isSet: true}
}

func (v NullableDbContributionStatTimeframe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbContributionStatTimeframe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
