# DbPullRequestGitHubEventsHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **time.Time** | Timestamp representing histogram bucket day | 
**Interval** | **float32** | The width in days of the individual time bucket | 
**PrCount** | **int32** | Number of Prs created in day bucket | 
**AcceptedPrs** | **int32** | Number of accepted/merged Prs in bucket | 
**OpenPrs** | **int32** | Number of open, unmerged Prs in bucket | 
**ClosedPrs** | **int32** | Number of closed, unmerged Prs in bucket | 
**DraftPrs** | **int32** | Number of drafted, unmerged Prs in bucket | 
**ActivePrs** | **int32** | Number of active, unmerged Prs in bucket | 
**SpamPrs** | **int32** | Number of Prs marked as spam within bucket | 
**PrVelocity** | **int32** | The average number of days to merge a PR over the time period | 

## Methods

### NewDbPullRequestGitHubEventsHistogram

`func NewDbPullRequestGitHubEventsHistogram(bucket time.Time, interval float32, prCount int32, acceptedPrs int32, openPrs int32, closedPrs int32, draftPrs int32, activePrs int32, spamPrs int32, prVelocity int32, ) *DbPullRequestGitHubEventsHistogram`

NewDbPullRequestGitHubEventsHistogram instantiates a new DbPullRequestGitHubEventsHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbPullRequestGitHubEventsHistogramWithDefaults

`func NewDbPullRequestGitHubEventsHistogramWithDefaults() *DbPullRequestGitHubEventsHistogram`

NewDbPullRequestGitHubEventsHistogramWithDefaults instantiates a new DbPullRequestGitHubEventsHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *DbPullRequestGitHubEventsHistogram) GetBucket() time.Time`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *DbPullRequestGitHubEventsHistogram) GetBucketOk() (*time.Time, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *DbPullRequestGitHubEventsHistogram) SetBucket(v time.Time)`

SetBucket sets Bucket field to given value.


### GetInterval

`func (o *DbPullRequestGitHubEventsHistogram) GetInterval() float32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DbPullRequestGitHubEventsHistogram) GetIntervalOk() (*float32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DbPullRequestGitHubEventsHistogram) SetInterval(v float32)`

SetInterval sets Interval field to given value.


### GetPrCount

`func (o *DbPullRequestGitHubEventsHistogram) GetPrCount() int32`

GetPrCount returns the PrCount field if non-nil, zero value otherwise.

### GetPrCountOk

`func (o *DbPullRequestGitHubEventsHistogram) GetPrCountOk() (*int32, bool)`

GetPrCountOk returns a tuple with the PrCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrCount

`func (o *DbPullRequestGitHubEventsHistogram) SetPrCount(v int32)`

SetPrCount sets PrCount field to given value.


### GetAcceptedPrs

`func (o *DbPullRequestGitHubEventsHistogram) GetAcceptedPrs() int32`

GetAcceptedPrs returns the AcceptedPrs field if non-nil, zero value otherwise.

### GetAcceptedPrsOk

`func (o *DbPullRequestGitHubEventsHistogram) GetAcceptedPrsOk() (*int32, bool)`

GetAcceptedPrsOk returns a tuple with the AcceptedPrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedPrs

`func (o *DbPullRequestGitHubEventsHistogram) SetAcceptedPrs(v int32)`

SetAcceptedPrs sets AcceptedPrs field to given value.


### GetOpenPrs

`func (o *DbPullRequestGitHubEventsHistogram) GetOpenPrs() int32`

GetOpenPrs returns the OpenPrs field if non-nil, zero value otherwise.

### GetOpenPrsOk

`func (o *DbPullRequestGitHubEventsHistogram) GetOpenPrsOk() (*int32, bool)`

GetOpenPrsOk returns a tuple with the OpenPrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrs

`func (o *DbPullRequestGitHubEventsHistogram) SetOpenPrs(v int32)`

SetOpenPrs sets OpenPrs field to given value.


### GetClosedPrs

`func (o *DbPullRequestGitHubEventsHistogram) GetClosedPrs() int32`

GetClosedPrs returns the ClosedPrs field if non-nil, zero value otherwise.

### GetClosedPrsOk

`func (o *DbPullRequestGitHubEventsHistogram) GetClosedPrsOk() (*int32, bool)`

GetClosedPrsOk returns a tuple with the ClosedPrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedPrs

`func (o *DbPullRequestGitHubEventsHistogram) SetClosedPrs(v int32)`

SetClosedPrs sets ClosedPrs field to given value.


### GetDraftPrs

`func (o *DbPullRequestGitHubEventsHistogram) GetDraftPrs() int32`

GetDraftPrs returns the DraftPrs field if non-nil, zero value otherwise.

### GetDraftPrsOk

`func (o *DbPullRequestGitHubEventsHistogram) GetDraftPrsOk() (*int32, bool)`

GetDraftPrsOk returns a tuple with the DraftPrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftPrs

`func (o *DbPullRequestGitHubEventsHistogram) SetDraftPrs(v int32)`

SetDraftPrs sets DraftPrs field to given value.


### GetActivePrs

`func (o *DbPullRequestGitHubEventsHistogram) GetActivePrs() int32`

GetActivePrs returns the ActivePrs field if non-nil, zero value otherwise.

### GetActivePrsOk

`func (o *DbPullRequestGitHubEventsHistogram) GetActivePrsOk() (*int32, bool)`

GetActivePrsOk returns a tuple with the ActivePrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePrs

`func (o *DbPullRequestGitHubEventsHistogram) SetActivePrs(v int32)`

SetActivePrs sets ActivePrs field to given value.


### GetSpamPrs

`func (o *DbPullRequestGitHubEventsHistogram) GetSpamPrs() int32`

GetSpamPrs returns the SpamPrs field if non-nil, zero value otherwise.

### GetSpamPrsOk

`func (o *DbPullRequestGitHubEventsHistogram) GetSpamPrsOk() (*int32, bool)`

GetSpamPrsOk returns a tuple with the SpamPrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamPrs

`func (o *DbPullRequestGitHubEventsHistogram) SetSpamPrs(v int32)`

SetSpamPrs sets SpamPrs field to given value.


### GetPrVelocity

`func (o *DbPullRequestGitHubEventsHistogram) GetPrVelocity() int32`

GetPrVelocity returns the PrVelocity field if non-nil, zero value otherwise.

### GetPrVelocityOk

`func (o *DbPullRequestGitHubEventsHistogram) GetPrVelocityOk() (*int32, bool)`

GetPrVelocityOk returns a tuple with the PrVelocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrVelocity

`func (o *DbPullRequestGitHubEventsHistogram) SetPrVelocity(v int32)`

SetPrVelocity sets PrVelocity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


