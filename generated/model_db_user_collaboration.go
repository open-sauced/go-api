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
	"time"
)

// checks if the DbUserCollaboration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbUserCollaboration{}

// DbUserCollaboration struct for DbUserCollaboration
type DbUserCollaboration struct {
	// User Collaboration identifier
	Id string `json:"id"`
	// Collaboration Receipient User ID
	UserId *float32 `json:"user_id,omitempty"`
	// Collaboration Request User ID
	RequestUserId *float32 `json:"request_user_id,omitempty"`
	// Collaboration Request Message
	Message string `json:"message"`
	// Collaboration Status
	Status string `json:"status"`
	// Timestamp representing user collaboration creation
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp representing user collaboration last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Timestamp representing user collaboration deletion
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Timestamp representing collaboration request email sent date
	RequestEmailedAt *time.Time `json:"request_emailed_at,omitempty"`
	// Timestamp representing collaboration acceptance email sent date
	CollaborationEmailedAt *time.Time `json:"collaboration_emailed_at,omitempty"`
}

// NewDbUserCollaboration instantiates a new DbUserCollaboration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbUserCollaboration(id string, message string, status string) *DbUserCollaboration {
	this := DbUserCollaboration{}
	this.Id = id
	this.Message = message
	this.Status = status
	return &this
}

// NewDbUserCollaborationWithDefaults instantiates a new DbUserCollaboration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbUserCollaborationWithDefaults() *DbUserCollaboration {
	this := DbUserCollaboration{}
	return &this
}

// GetId returns the Id field value
func (o *DbUserCollaboration) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DbUserCollaboration) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DbUserCollaboration) SetId(v string) {
	o.Id = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *DbUserCollaboration) GetUserId() float32 {
	if o == nil || IsNil(o.UserId) {
		var ret float32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbUserCollaboration) GetUserIdOk() (*float32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *DbUserCollaboration) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given float32 and assigns it to the UserId field.
func (o *DbUserCollaboration) SetUserId(v float32) {
	o.UserId = &v
}

// GetRequestUserId returns the RequestUserId field value if set, zero value otherwise.
func (o *DbUserCollaboration) GetRequestUserId() float32 {
	if o == nil || IsNil(o.RequestUserId) {
		var ret float32
		return ret
	}
	return *o.RequestUserId
}

// GetRequestUserIdOk returns a tuple with the RequestUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbUserCollaboration) GetRequestUserIdOk() (*float32, bool) {
	if o == nil || IsNil(o.RequestUserId) {
		return nil, false
	}
	return o.RequestUserId, true
}

// HasRequestUserId returns a boolean if a field has been set.
func (o *DbUserCollaboration) HasRequestUserId() bool {
	if o != nil && !IsNil(o.RequestUserId) {
		return true
	}

	return false
}

// SetRequestUserId gets a reference to the given float32 and assigns it to the RequestUserId field.
func (o *DbUserCollaboration) SetRequestUserId(v float32) {
	o.RequestUserId = &v
}

// GetMessage returns the Message field value
func (o *DbUserCollaboration) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *DbUserCollaboration) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *DbUserCollaboration) SetMessage(v string) {
	o.Message = v
}

// GetStatus returns the Status field value
func (o *DbUserCollaboration) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DbUserCollaboration) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DbUserCollaboration) SetStatus(v string) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DbUserCollaboration) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbUserCollaboration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DbUserCollaboration) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DbUserCollaboration) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DbUserCollaboration) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbUserCollaboration) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DbUserCollaboration) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DbUserCollaboration) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *DbUserCollaboration) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbUserCollaboration) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *DbUserCollaboration) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *DbUserCollaboration) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetRequestEmailedAt returns the RequestEmailedAt field value if set, zero value otherwise.
func (o *DbUserCollaboration) GetRequestEmailedAt() time.Time {
	if o == nil || IsNil(o.RequestEmailedAt) {
		var ret time.Time
		return ret
	}
	return *o.RequestEmailedAt
}

// GetRequestEmailedAtOk returns a tuple with the RequestEmailedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbUserCollaboration) GetRequestEmailedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RequestEmailedAt) {
		return nil, false
	}
	return o.RequestEmailedAt, true
}

// HasRequestEmailedAt returns a boolean if a field has been set.
func (o *DbUserCollaboration) HasRequestEmailedAt() bool {
	if o != nil && !IsNil(o.RequestEmailedAt) {
		return true
	}

	return false
}

// SetRequestEmailedAt gets a reference to the given time.Time and assigns it to the RequestEmailedAt field.
func (o *DbUserCollaboration) SetRequestEmailedAt(v time.Time) {
	o.RequestEmailedAt = &v
}

// GetCollaborationEmailedAt returns the CollaborationEmailedAt field value if set, zero value otherwise.
func (o *DbUserCollaboration) GetCollaborationEmailedAt() time.Time {
	if o == nil || IsNil(o.CollaborationEmailedAt) {
		var ret time.Time
		return ret
	}
	return *o.CollaborationEmailedAt
}

// GetCollaborationEmailedAtOk returns a tuple with the CollaborationEmailedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbUserCollaboration) GetCollaborationEmailedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CollaborationEmailedAt) {
		return nil, false
	}
	return o.CollaborationEmailedAt, true
}

// HasCollaborationEmailedAt returns a boolean if a field has been set.
func (o *DbUserCollaboration) HasCollaborationEmailedAt() bool {
	if o != nil && !IsNil(o.CollaborationEmailedAt) {
		return true
	}

	return false
}

// SetCollaborationEmailedAt gets a reference to the given time.Time and assigns it to the CollaborationEmailedAt field.
func (o *DbUserCollaboration) SetCollaborationEmailedAt(v time.Time) {
	o.CollaborationEmailedAt = &v
}

func (o DbUserCollaboration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbUserCollaboration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.RequestUserId) {
		toSerialize["request_user_id"] = o.RequestUserId
	}
	toSerialize["message"] = o.Message
	toSerialize["status"] = o.Status
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.RequestEmailedAt) {
		toSerialize["request_emailed_at"] = o.RequestEmailedAt
	}
	if !IsNil(o.CollaborationEmailedAt) {
		toSerialize["collaboration_emailed_at"] = o.CollaborationEmailedAt
	}
	return toSerialize, nil
}

type NullableDbUserCollaboration struct {
	value *DbUserCollaboration
	isSet bool
}

func (v NullableDbUserCollaboration) Get() *DbUserCollaboration {
	return v.value
}

func (v *NullableDbUserCollaboration) Set(val *DbUserCollaboration) {
	v.value = val
	v.isSet = true
}

func (v NullableDbUserCollaboration) IsSet() bool {
	return v.isSet
}

func (v *NullableDbUserCollaboration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbUserCollaboration(val *DbUserCollaboration) *NullableDbUserCollaboration {
	return &NullableDbUserCollaboration{value: val, isSet: true}
}

func (v NullableDbUserCollaboration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbUserCollaboration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


