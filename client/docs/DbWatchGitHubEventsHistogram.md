# DbWatchGitHubEventsHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **time.Time** | Timestamp representing histogram bucket day | 
**StarCount** | **int32** | Number of stars in day bucket | 

## Methods

### NewDbWatchGitHubEventsHistogram

`func NewDbWatchGitHubEventsHistogram(bucket time.Time, starCount int32, ) *DbWatchGitHubEventsHistogram`

NewDbWatchGitHubEventsHistogram instantiates a new DbWatchGitHubEventsHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbWatchGitHubEventsHistogramWithDefaults

`func NewDbWatchGitHubEventsHistogramWithDefaults() *DbWatchGitHubEventsHistogram`

NewDbWatchGitHubEventsHistogramWithDefaults instantiates a new DbWatchGitHubEventsHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *DbWatchGitHubEventsHistogram) GetBucket() time.Time`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *DbWatchGitHubEventsHistogram) GetBucketOk() (*time.Time, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *DbWatchGitHubEventsHistogram) SetBucket(v time.Time)`

SetBucket sets Bucket field to given value.


### GetStarCount

`func (o *DbWatchGitHubEventsHistogram) GetStarCount() int32`

GetStarCount returns the StarCount field if non-nil, zero value otherwise.

### GetStarCountOk

`func (o *DbWatchGitHubEventsHistogram) GetStarCountOk() (*int32, bool)`

GetStarCountOk returns a tuple with the StarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarCount

`func (o *DbWatchGitHubEventsHistogram) SetStarCount(v int32)`

SetStarCount sets StarCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


