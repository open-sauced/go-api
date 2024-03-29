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

// checks if the OmitTypeClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OmitTypeClass{}

// OmitTypeClass struct for OmitTypeClass
type OmitTypeClass struct {
	// Notification identifier
	Id int32 `json:"id"`
	// User ID
	UserId int32 `json:"user_id"`
	// From User ID
	FromUserId *int32 `json:"from_user_id,omitempty"`
	// User notification type
	Type string `json:"type"`
	// User notification message
	Message *string `json:"message,omitempty"`
	// Timestamp representing db-user-notification creation
	NotifiedAt *time.Time `json:"notified_at,omitempty"`
	// Notification Source ID (highlight, user, invite)
	MetaId *string `json:"meta_id,omitempty"`
}

// NewOmitTypeClass instantiates a new OmitTypeClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOmitTypeClass(id int32, userId int32, type_ string) *OmitTypeClass {
	this := OmitTypeClass{}
	this.Id = id
	this.UserId = userId
	this.Type = type_
	return &this
}

// NewOmitTypeClassWithDefaults instantiates a new OmitTypeClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOmitTypeClassWithDefaults() *OmitTypeClass {
	this := OmitTypeClass{}
	return &this
}

// GetId returns the Id field value
func (o *OmitTypeClass) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OmitTypeClass) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OmitTypeClass) SetId(v int32) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *OmitTypeClass) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *OmitTypeClass) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *OmitTypeClass) SetUserId(v int32) {
	o.UserId = v
}

// GetFromUserId returns the FromUserId field value if set, zero value otherwise.
func (o *OmitTypeClass) GetFromUserId() int32 {
	if o == nil || IsNil(o.FromUserId) {
		var ret int32
		return ret
	}
	return *o.FromUserId
}

// GetFromUserIdOk returns a tuple with the FromUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmitTypeClass) GetFromUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FromUserId) {
		return nil, false
	}
	return o.FromUserId, true
}

// HasFromUserId returns a boolean if a field has been set.
func (o *OmitTypeClass) HasFromUserId() bool {
	if o != nil && !IsNil(o.FromUserId) {
		return true
	}

	return false
}

// SetFromUserId gets a reference to the given int32 and assigns it to the FromUserId field.
func (o *OmitTypeClass) SetFromUserId(v int32) {
	o.FromUserId = &v
}

// GetType returns the Type field value
func (o *OmitTypeClass) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OmitTypeClass) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OmitTypeClass) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OmitTypeClass) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmitTypeClass) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OmitTypeClass) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OmitTypeClass) SetMessage(v string) {
	o.Message = &v
}

// GetNotifiedAt returns the NotifiedAt field value if set, zero value otherwise.
func (o *OmitTypeClass) GetNotifiedAt() time.Time {
	if o == nil || IsNil(o.NotifiedAt) {
		var ret time.Time
		return ret
	}
	return *o.NotifiedAt
}

// GetNotifiedAtOk returns a tuple with the NotifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmitTypeClass) GetNotifiedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NotifiedAt) {
		return nil, false
	}
	return o.NotifiedAt, true
}

// HasNotifiedAt returns a boolean if a field has been set.
func (o *OmitTypeClass) HasNotifiedAt() bool {
	if o != nil && !IsNil(o.NotifiedAt) {
		return true
	}

	return false
}

// SetNotifiedAt gets a reference to the given time.Time and assigns it to the NotifiedAt field.
func (o *OmitTypeClass) SetNotifiedAt(v time.Time) {
	o.NotifiedAt = &v
}

// GetMetaId returns the MetaId field value if set, zero value otherwise.
func (o *OmitTypeClass) GetMetaId() string {
	if o == nil || IsNil(o.MetaId) {
		var ret string
		return ret
	}
	return *o.MetaId
}

// GetMetaIdOk returns a tuple with the MetaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OmitTypeClass) GetMetaIdOk() (*string, bool) {
	if o == nil || IsNil(o.MetaId) {
		return nil, false
	}
	return o.MetaId, true
}

// HasMetaId returns a boolean if a field has been set.
func (o *OmitTypeClass) HasMetaId() bool {
	if o != nil && !IsNil(o.MetaId) {
		return true
	}

	return false
}

// SetMetaId gets a reference to the given string and assigns it to the MetaId field.
func (o *OmitTypeClass) SetMetaId(v string) {
	o.MetaId = &v
}

func (o OmitTypeClass) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OmitTypeClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	if !IsNil(o.FromUserId) {
		toSerialize["from_user_id"] = o.FromUserId
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.NotifiedAt) {
		toSerialize["notified_at"] = o.NotifiedAt
	}
	if !IsNil(o.MetaId) {
		toSerialize["meta_id"] = o.MetaId
	}
	return toSerialize, nil
}

type NullableOmitTypeClass struct {
	value *OmitTypeClass
	isSet bool
}

func (v NullableOmitTypeClass) Get() *OmitTypeClass {
	return v.value
}

func (v *NullableOmitTypeClass) Set(val *OmitTypeClass) {
	v.value = val
	v.isSet = true
}

func (v NullableOmitTypeClass) IsSet() bool {
	return v.isSet
}

func (v *NullableOmitTypeClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOmitTypeClass(val *OmitTypeClass) *NullableOmitTypeClass {
	return &NullableOmitTypeClass{value: val, isSet: true}
}

func (v NullableOmitTypeClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOmitTypeClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
