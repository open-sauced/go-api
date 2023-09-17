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
	"time"
)

// checks if the DbUserListContributor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbUserListContributor{}

// DbUserListContributor struct for DbUserListContributor
type DbUserListContributor struct {
	// User list contributor identifier
	Id string `json:"id"`
	// User identifier
	UserId int32 `json:"user_id"`
	// List identifier
	ListId string `json:"list_id"`
	// Timestamp representing top repo first index
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// User list collaborator's login
	Login *string `json:"login,omitempty"`
}

// NewDbUserListContributor instantiates a new DbUserListContributor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbUserListContributor(id string, userId int32, listId string) *DbUserListContributor {
	this := DbUserListContributor{}
	this.Id = id
	this.UserId = userId
	this.ListId = listId
	return &this
}

// NewDbUserListContributorWithDefaults instantiates a new DbUserListContributor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbUserListContributorWithDefaults() *DbUserListContributor {
	this := DbUserListContributor{}
	return &this
}

// GetId returns the Id field value
func (o *DbUserListContributor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DbUserListContributor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DbUserListContributor) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *DbUserListContributor) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *DbUserListContributor) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *DbUserListContributor) SetUserId(v int32) {
	o.UserId = v
}

// GetListId returns the ListId field value
func (o *DbUserListContributor) GetListId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListId
}

// GetListIdOk returns a tuple with the ListId field value
// and a boolean to check if the value has been set.
func (o *DbUserListContributor) GetListIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListId, true
}

// SetListId sets field value
func (o *DbUserListContributor) SetListId(v string) {
	o.ListId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DbUserListContributor) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbUserListContributor) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DbUserListContributor) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DbUserListContributor) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *DbUserListContributor) GetLogin() string {
	if o == nil || IsNil(o.Login) {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbUserListContributor) GetLoginOk() (*string, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *DbUserListContributor) HasLogin() bool {
	if o != nil && !IsNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *DbUserListContributor) SetLogin(v string) {
	o.Login = &v
}

func (o DbUserListContributor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbUserListContributor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	toSerialize["list_id"] = o.ListId
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Login) {
		toSerialize["login"] = o.Login
	}
	return toSerialize, nil
}

type NullableDbUserListContributor struct {
	value *DbUserListContributor
	isSet bool
}

func (v NullableDbUserListContributor) Get() *DbUserListContributor {
	return v.value
}

func (v *NullableDbUserListContributor) Set(val *DbUserListContributor) {
	v.value = val
	v.isSet = true
}

func (v NullableDbUserListContributor) IsSet() bool {
	return v.isSet
}

func (v *NullableDbUserListContributor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbUserListContributor(val *DbUserListContributor) *NullableDbUserListContributor {
	return &NullableDbUserListContributor{value: val, isSet: true}
}

func (v NullableDbUserListContributor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbUserListContributor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
