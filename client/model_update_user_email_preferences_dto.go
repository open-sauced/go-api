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

// checks if the UpdateUserEmailPreferencesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserEmailPreferencesDto{}

// UpdateUserEmailPreferencesDto struct for UpdateUserEmailPreferencesDto
type UpdateUserEmailPreferencesDto struct {
	// User Display Public Email
	DisplayEmail bool `json:"display_email"`
	// User Recieve Collaboration Requests
	ReceiveCollaboration bool `json:"receive_collaboration"`
}

// NewUpdateUserEmailPreferencesDto instantiates a new UpdateUserEmailPreferencesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserEmailPreferencesDto(displayEmail bool, receiveCollaboration bool) *UpdateUserEmailPreferencesDto {
	this := UpdateUserEmailPreferencesDto{}
	this.DisplayEmail = displayEmail
	this.ReceiveCollaboration = receiveCollaboration
	return &this
}

// NewUpdateUserEmailPreferencesDtoWithDefaults instantiates a new UpdateUserEmailPreferencesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserEmailPreferencesDtoWithDefaults() *UpdateUserEmailPreferencesDto {
	this := UpdateUserEmailPreferencesDto{}
	return &this
}

// GetDisplayEmail returns the DisplayEmail field value
func (o *UpdateUserEmailPreferencesDto) GetDisplayEmail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisplayEmail
}

// GetDisplayEmailOk returns a tuple with the DisplayEmail field value
// and a boolean to check if the value has been set.
func (o *UpdateUserEmailPreferencesDto) GetDisplayEmailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayEmail, true
}

// SetDisplayEmail sets field value
func (o *UpdateUserEmailPreferencesDto) SetDisplayEmail(v bool) {
	o.DisplayEmail = v
}

// GetReceiveCollaboration returns the ReceiveCollaboration field value
func (o *UpdateUserEmailPreferencesDto) GetReceiveCollaboration() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ReceiveCollaboration
}

// GetReceiveCollaborationOk returns a tuple with the ReceiveCollaboration field value
// and a boolean to check if the value has been set.
func (o *UpdateUserEmailPreferencesDto) GetReceiveCollaborationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiveCollaboration, true
}

// SetReceiveCollaboration sets field value
func (o *UpdateUserEmailPreferencesDto) SetReceiveCollaboration(v bool) {
	o.ReceiveCollaboration = v
}

func (o UpdateUserEmailPreferencesDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserEmailPreferencesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["display_email"] = o.DisplayEmail
	toSerialize["receive_collaboration"] = o.ReceiveCollaboration
	return toSerialize, nil
}

type NullableUpdateUserEmailPreferencesDto struct {
	value *UpdateUserEmailPreferencesDto
	isSet bool
}

func (v NullableUpdateUserEmailPreferencesDto) Get() *UpdateUserEmailPreferencesDto {
	return v.value
}

func (v *NullableUpdateUserEmailPreferencesDto) Set(val *UpdateUserEmailPreferencesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserEmailPreferencesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserEmailPreferencesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserEmailPreferencesDto(val *UpdateUserEmailPreferencesDto) *NullableUpdateUserEmailPreferencesDto {
	return &NullableUpdateUserEmailPreferencesDto{value: val, isSet: true}
}

func (v NullableUpdateUserEmailPreferencesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserEmailPreferencesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
