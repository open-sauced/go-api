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

// checks if the UserOnboardingDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserOnboardingDto{}

// UserOnboardingDto struct for UserOnboardingDto
type UserOnboardingDto struct {
	// An array of interests
	Interests []string `json:"interests"`
	// User timezone in UTC
	Timezone string `json:"timezone"`
}

// NewUserOnboardingDto instantiates a new UserOnboardingDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserOnboardingDto(interests []string, timezone string) *UserOnboardingDto {
	this := UserOnboardingDto{}
	this.Interests = interests
	this.Timezone = timezone
	return &this
}

// NewUserOnboardingDtoWithDefaults instantiates a new UserOnboardingDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserOnboardingDtoWithDefaults() *UserOnboardingDto {
	this := UserOnboardingDto{}
	return &this
}

// GetInterests returns the Interests field value
func (o *UserOnboardingDto) GetInterests() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Interests
}

// GetInterestsOk returns a tuple with the Interests field value
// and a boolean to check if the value has been set.
func (o *UserOnboardingDto) GetInterestsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interests, true
}

// SetInterests sets field value
func (o *UserOnboardingDto) SetInterests(v []string) {
	o.Interests = v
}

// GetTimezone returns the Timezone field value
func (o *UserOnboardingDto) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *UserOnboardingDto) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *UserOnboardingDto) SetTimezone(v string) {
	o.Timezone = v
}

func (o UserOnboardingDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserOnboardingDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["interests"] = o.Interests
	toSerialize["timezone"] = o.Timezone
	return toSerialize, nil
}

type NullableUserOnboardingDto struct {
	value *UserOnboardingDto
	isSet bool
}

func (v NullableUserOnboardingDto) Get() *UserOnboardingDto {
	return v.value
}

func (v *NullableUserOnboardingDto) Set(val *UserOnboardingDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserOnboardingDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserOnboardingDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserOnboardingDto(val *UserOnboardingDto) *NullableUserOnboardingDto {
	return &NullableUserOnboardingDto{value: val, isSet: true}
}

func (v NullableUserOnboardingDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserOnboardingDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

