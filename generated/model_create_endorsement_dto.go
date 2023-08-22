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

// checks if the CreateEndorsementDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEndorsementDto{}

// CreateEndorsementDto struct for CreateEndorsementDto
type CreateEndorsementDto struct {
	// Endorsement Creator User ID
	CreatorUserId float32 `json:"creator_user_id"`
	// Endorsement Recipient User ID
	RecipientUserId float32 `json:"recipient_user_id"`
	// Repository ID
	RepoId float32 `json:"repo_id"`
	// Endorsement Source Comment URL
	SourceCommentUrl *string `json:"source_comment_url,omitempty"`
	// Endorsement Source Context URL
	SourceContextUrl string `json:"source_context_url"`
	// Endorsement Type
	Type string `json:"type"`
}

// NewCreateEndorsementDto instantiates a new CreateEndorsementDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEndorsementDto(creatorUserId float32, recipientUserId float32, repoId float32, sourceContextUrl string, type_ string) *CreateEndorsementDto {
	this := CreateEndorsementDto{}
	this.CreatorUserId = creatorUserId
	this.RecipientUserId = recipientUserId
	this.RepoId = repoId
	var sourceCommentUrl string = ""
	this.SourceCommentUrl = &sourceCommentUrl
	this.SourceContextUrl = sourceContextUrl
	this.Type = type_
	return &this
}

// NewCreateEndorsementDtoWithDefaults instantiates a new CreateEndorsementDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEndorsementDtoWithDefaults() *CreateEndorsementDto {
	this := CreateEndorsementDto{}
	var sourceCommentUrl string = ""
	this.SourceCommentUrl = &sourceCommentUrl
	return &this
}

// GetCreatorUserId returns the CreatorUserId field value
func (o *CreateEndorsementDto) GetCreatorUserId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatorUserId
}

// GetCreatorUserIdOk returns a tuple with the CreatorUserId field value
// and a boolean to check if the value has been set.
func (o *CreateEndorsementDto) GetCreatorUserIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorUserId, true
}

// SetCreatorUserId sets field value
func (o *CreateEndorsementDto) SetCreatorUserId(v float32) {
	o.CreatorUserId = v
}

// GetRecipientUserId returns the RecipientUserId field value
func (o *CreateEndorsementDto) GetRecipientUserId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RecipientUserId
}

// GetRecipientUserIdOk returns a tuple with the RecipientUserId field value
// and a boolean to check if the value has been set.
func (o *CreateEndorsementDto) GetRecipientUserIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientUserId, true
}

// SetRecipientUserId sets field value
func (o *CreateEndorsementDto) SetRecipientUserId(v float32) {
	o.RecipientUserId = v
}

// GetRepoId returns the RepoId field value
func (o *CreateEndorsementDto) GetRepoId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RepoId
}

// GetRepoIdOk returns a tuple with the RepoId field value
// and a boolean to check if the value has been set.
func (o *CreateEndorsementDto) GetRepoIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepoId, true
}

// SetRepoId sets field value
func (o *CreateEndorsementDto) SetRepoId(v float32) {
	o.RepoId = v
}

// GetSourceCommentUrl returns the SourceCommentUrl field value if set, zero value otherwise.
func (o *CreateEndorsementDto) GetSourceCommentUrl() string {
	if o == nil || IsNil(o.SourceCommentUrl) {
		var ret string
		return ret
	}
	return *o.SourceCommentUrl
}

// GetSourceCommentUrlOk returns a tuple with the SourceCommentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEndorsementDto) GetSourceCommentUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SourceCommentUrl) {
		return nil, false
	}
	return o.SourceCommentUrl, true
}

// HasSourceCommentUrl returns a boolean if a field has been set.
func (o *CreateEndorsementDto) HasSourceCommentUrl() bool {
	if o != nil && !IsNil(o.SourceCommentUrl) {
		return true
	}

	return false
}

// SetSourceCommentUrl gets a reference to the given string and assigns it to the SourceCommentUrl field.
func (o *CreateEndorsementDto) SetSourceCommentUrl(v string) {
	o.SourceCommentUrl = &v
}

// GetSourceContextUrl returns the SourceContextUrl field value
func (o *CreateEndorsementDto) GetSourceContextUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceContextUrl
}

// GetSourceContextUrlOk returns a tuple with the SourceContextUrl field value
// and a boolean to check if the value has been set.
func (o *CreateEndorsementDto) GetSourceContextUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceContextUrl, true
}

// SetSourceContextUrl sets field value
func (o *CreateEndorsementDto) SetSourceContextUrl(v string) {
	o.SourceContextUrl = v
}

// GetType returns the Type field value
func (o *CreateEndorsementDto) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateEndorsementDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateEndorsementDto) SetType(v string) {
	o.Type = v
}

func (o CreateEndorsementDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEndorsementDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["creator_user_id"] = o.CreatorUserId
	toSerialize["recipient_user_id"] = o.RecipientUserId
	toSerialize["repo_id"] = o.RepoId
	if !IsNil(o.SourceCommentUrl) {
		toSerialize["source_comment_url"] = o.SourceCommentUrl
	}
	toSerialize["source_context_url"] = o.SourceContextUrl
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCreateEndorsementDto struct {
	value *CreateEndorsementDto
	isSet bool
}

func (v NullableCreateEndorsementDto) Get() *CreateEndorsementDto {
	return v.value
}

func (v *NullableCreateEndorsementDto) Set(val *CreateEndorsementDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEndorsementDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEndorsementDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEndorsementDto(val *CreateEndorsementDto) *NullableCreateEndorsementDto {
	return &NullableCreateEndorsementDto{value: val, isSet: true}
}

func (v NullableCreateEndorsementDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEndorsementDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


