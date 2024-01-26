# DbPushGitHubEventsHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **time.Time** | Timestamp representing histogram bucket day | 
**PushesCount** | **int32** | Number of pushes in day bucket | 

## Methods

### NewDbPushGitHubEventsHistogram

`func NewDbPushGitHubEventsHistogram(bucket time.Time, pushesCount int32, ) *DbPushGitHubEventsHistogram`

NewDbPushGitHubEventsHistogram instantiates a new DbPushGitHubEventsHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbPushGitHubEventsHistogramWithDefaults

`func NewDbPushGitHubEventsHistogramWithDefaults() *DbPushGitHubEventsHistogram`

NewDbPushGitHubEventsHistogramWithDefaults instantiates a new DbPushGitHubEventsHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *DbPushGitHubEventsHistogram) GetBucket() time.Time`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *DbPushGitHubEventsHistogram) GetBucketOk() (*time.Time, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *DbPushGitHubEventsHistogram) SetBucket(v time.Time)`

SetBucket sets Bucket field to given value.


### GetPushesCount

`func (o *DbPushGitHubEventsHistogram) GetPushesCount() int32`

GetPushesCount returns the PushesCount field if non-nil, zero value otherwise.

### GetPushesCountOk

`func (o *DbPushGitHubEventsHistogram) GetPushesCountOk() (*int32, bool)`

GetPushesCountOk returns a tuple with the PushesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushesCount

`func (o *DbPushGitHubEventsHistogram) SetPushesCount(v int32)`

SetPushesCount sets PushesCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


