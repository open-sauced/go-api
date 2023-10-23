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

// checks if the DbContributionsProjects type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbContributionsProjects{}

// DbContributionsProjects struct for DbContributionsProjects
type DbContributionsProjects struct {
	// The org name of the repo
	OrgId string `json:"org_id"`
	// The project name of the repo
	ProjectId string `json:"project_id"`
	// The ID of the associated repo in the OpenSauced API context
	RepoId int32 `json:"repo_id"`
	// The number of contributions associated with a org/repo
	Contributions int32 `json:"contributions"`
}

// NewDbContributionsProjects instantiates a new DbContributionsProjects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbContributionsProjects(orgId string, projectId string, repoId int32, contributions int32) *DbContributionsProjects {
	this := DbContributionsProjects{}
	this.OrgId = orgId
	this.ProjectId = projectId
	this.RepoId = repoId
	this.Contributions = contributions
	return &this
}

// NewDbContributionsProjectsWithDefaults instantiates a new DbContributionsProjects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbContributionsProjectsWithDefaults() *DbContributionsProjects {
	this := DbContributionsProjects{}
	return &this
}

// GetOrgId returns the OrgId field value
func (o *DbContributionsProjects) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *DbContributionsProjects) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *DbContributionsProjects) SetOrgId(v string) {
	o.OrgId = v
}

// GetProjectId returns the ProjectId field value
func (o *DbContributionsProjects) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DbContributionsProjects) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DbContributionsProjects) SetProjectId(v string) {
	o.ProjectId = v
}

// GetRepoId returns the RepoId field value
func (o *DbContributionsProjects) GetRepoId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RepoId
}

// GetRepoIdOk returns a tuple with the RepoId field value
// and a boolean to check if the value has been set.
func (o *DbContributionsProjects) GetRepoIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepoId, true
}

// SetRepoId sets field value
func (o *DbContributionsProjects) SetRepoId(v int32) {
	o.RepoId = v
}

// GetContributions returns the Contributions field value
func (o *DbContributionsProjects) GetContributions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Contributions
}

// GetContributionsOk returns a tuple with the Contributions field value
// and a boolean to check if the value has been set.
func (o *DbContributionsProjects) GetContributionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contributions, true
}

// SetContributions sets field value
func (o *DbContributionsProjects) SetContributions(v int32) {
	o.Contributions = v
}

func (o DbContributionsProjects) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbContributionsProjects) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["org_id"] = o.OrgId
	toSerialize["project_id"] = o.ProjectId
	toSerialize["repo_id"] = o.RepoId
	toSerialize["contributions"] = o.Contributions
	return toSerialize, nil
}

type NullableDbContributionsProjects struct {
	value *DbContributionsProjects
	isSet bool
}

func (v NullableDbContributionsProjects) Get() *DbContributionsProjects {
	return v.value
}

func (v *NullableDbContributionsProjects) Set(val *DbContributionsProjects) {
	v.value = val
	v.isSet = true
}

func (v NullableDbContributionsProjects) IsSet() bool {
	return v.isSet
}

func (v *NullableDbContributionsProjects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbContributionsProjects(val *DbContributionsProjects) *NullableDbContributionsProjects {
	return &NullableDbContributionsProjects{value: val, isSet: true}
}

func (v NullableDbContributionsProjects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbContributionsProjects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}