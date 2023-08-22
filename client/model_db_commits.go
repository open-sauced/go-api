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

// checks if the DbCommits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbCommits{}

// DbCommits struct for DbCommits
type DbCommits struct {
	// Commit identifier
	Id float32 `json:"id"`
	// Hash for the commit
	CommitHash string `json:"commit_hash"`
	// Date for the commit
	CommitDate string `json:"commit_date"`
	// Baked repo identifier
	BakedRepoId float32 `json:"baked_repo_id"`
	// Commit author identifier
	CommitAuthorId float32 `json:"commit_author_id"`
}

// NewDbCommits instantiates a new DbCommits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbCommits(id float32, commitHash string, commitDate string, bakedRepoId float32, commitAuthorId float32) *DbCommits {
	this := DbCommits{}
	this.Id = id
	this.CommitHash = commitHash
	this.CommitDate = commitDate
	this.BakedRepoId = bakedRepoId
	this.CommitAuthorId = commitAuthorId
	return &this
}

// NewDbCommitsWithDefaults instantiates a new DbCommits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbCommitsWithDefaults() *DbCommits {
	this := DbCommits{}
	return &this
}

// GetId returns the Id field value
func (o *DbCommits) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DbCommits) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DbCommits) SetId(v float32) {
	o.Id = v
}

// GetCommitHash returns the CommitHash field value
func (o *DbCommits) GetCommitHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitHash
}

// GetCommitHashOk returns a tuple with the CommitHash field value
// and a boolean to check if the value has been set.
func (o *DbCommits) GetCommitHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitHash, true
}

// SetCommitHash sets field value
func (o *DbCommits) SetCommitHash(v string) {
	o.CommitHash = v
}

// GetCommitDate returns the CommitDate field value
func (o *DbCommits) GetCommitDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitDate
}

// GetCommitDateOk returns a tuple with the CommitDate field value
// and a boolean to check if the value has been set.
func (o *DbCommits) GetCommitDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitDate, true
}

// SetCommitDate sets field value
func (o *DbCommits) SetCommitDate(v string) {
	o.CommitDate = v
}

// GetBakedRepoId returns the BakedRepoId field value
func (o *DbCommits) GetBakedRepoId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BakedRepoId
}

// GetBakedRepoIdOk returns a tuple with the BakedRepoId field value
// and a boolean to check if the value has been set.
func (o *DbCommits) GetBakedRepoIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BakedRepoId, true
}

// SetBakedRepoId sets field value
func (o *DbCommits) SetBakedRepoId(v float32) {
	o.BakedRepoId = v
}

// GetCommitAuthorId returns the CommitAuthorId field value
func (o *DbCommits) GetCommitAuthorId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CommitAuthorId
}

// GetCommitAuthorIdOk returns a tuple with the CommitAuthorId field value
// and a boolean to check if the value has been set.
func (o *DbCommits) GetCommitAuthorIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitAuthorId, true
}

// SetCommitAuthorId sets field value
func (o *DbCommits) SetCommitAuthorId(v float32) {
	o.CommitAuthorId = v
}

func (o DbCommits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbCommits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["commit_hash"] = o.CommitHash
	toSerialize["commit_date"] = o.CommitDate
	toSerialize["baked_repo_id"] = o.BakedRepoId
	toSerialize["commit_author_id"] = o.CommitAuthorId
	return toSerialize, nil
}

type NullableDbCommits struct {
	value *DbCommits
	isSet bool
}

func (v NullableDbCommits) Get() *DbCommits {
	return v.value
}

func (v *NullableDbCommits) Set(val *DbCommits) {
	v.value = val
	v.isSet = true
}

func (v NullableDbCommits) IsSet() bool {
	return v.isSet
}

func (v *NullableDbCommits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbCommits(val *DbCommits) *NullableDbCommits {
	return &NullableDbCommits{value: val, isSet: true}
}

func (v NullableDbCommits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbCommits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

