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
	"fmt"
)

// PullRequestOrderFieldsEnum the model 'PullRequestOrderFieldsEnum'
type PullRequestOrderFieldsEnum string

// List of PullRequestOrderFieldsEnum
const (
	CREATED_AT PullRequestOrderFieldsEnum = "created_at"
	CLOSED_AT PullRequestOrderFieldsEnum = "closed_at"
	MERGED_AT PullRequestOrderFieldsEnum = "merged_at"
	UPDATED_AT PullRequestOrderFieldsEnum = "updated_at"
)

// All allowed values of PullRequestOrderFieldsEnum enum
var AllowedPullRequestOrderFieldsEnumEnumValues = []PullRequestOrderFieldsEnum{
	"created_at",
	"closed_at",
	"merged_at",
	"updated_at",
}

func (v *PullRequestOrderFieldsEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PullRequestOrderFieldsEnum(value)
	for _, existing := range AllowedPullRequestOrderFieldsEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PullRequestOrderFieldsEnum", value)
}

// NewPullRequestOrderFieldsEnumFromValue returns a pointer to a valid PullRequestOrderFieldsEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPullRequestOrderFieldsEnumFromValue(v string) (*PullRequestOrderFieldsEnum, error) {
	ev := PullRequestOrderFieldsEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PullRequestOrderFieldsEnum: valid values are %v", v, AllowedPullRequestOrderFieldsEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PullRequestOrderFieldsEnum) IsValid() bool {
	for _, existing := range AllowedPullRequestOrderFieldsEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PullRequestOrderFieldsEnum value
func (v PullRequestOrderFieldsEnum) Ptr() *PullRequestOrderFieldsEnum {
	return &v
}

type NullablePullRequestOrderFieldsEnum struct {
	value *PullRequestOrderFieldsEnum
	isSet bool
}

func (v NullablePullRequestOrderFieldsEnum) Get() *PullRequestOrderFieldsEnum {
	return v.value
}

func (v *NullablePullRequestOrderFieldsEnum) Set(val *PullRequestOrderFieldsEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePullRequestOrderFieldsEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePullRequestOrderFieldsEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullRequestOrderFieldsEnum(val *PullRequestOrderFieldsEnum) *NullablePullRequestOrderFieldsEnum {
	return &NullablePullRequestOrderFieldsEnum{value: val, isSet: true}
}

func (v NullablePullRequestOrderFieldsEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullRequestOrderFieldsEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

