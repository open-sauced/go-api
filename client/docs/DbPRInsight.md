# DbPRInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Repository identifier | 
**Interval** | **int32** | Selected interval in numerical days, goes back with number, 0 means today | [default to 0]
**Day** | **string** | Selected interval computed date in human readable format | 
**AllPrs** | **int32** | PR Type: all requests count | 
**AcceptedPrs** | **int32** | PR Type: accepted requests count | 
**SpamPrs** | **int32** | PR Type: spam requests count | 

## Methods

### NewDbPRInsight

`func NewDbPRInsight(id int32, interval int32, day string, allPrs int32, acceptedPrs int32, spamPrs int32, ) *DbPRInsight`

NewDbPRInsight instantiates a new DbPRInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbPRInsightWithDefaults

`func NewDbPRInsightWithDefaults() *DbPRInsight`

NewDbPRInsightWithDefaults instantiates a new DbPRInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbPRInsight) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbPRInsight) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbPRInsight) SetId(v int32)`

SetId sets Id field to given value.


### GetInterval

`func (o *DbPRInsight) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DbPRInsight) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DbPRInsight) SetInterval(v int32)`

SetInterval sets Interval field to given value.


### GetDay

`func (o *DbPRInsight) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *DbPRInsight) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *DbPRInsight) SetDay(v string)`

SetDay sets Day field to given value.


### GetAllPrs

`func (o *DbPRInsight) GetAllPrs() int32`

GetAllPrs returns the AllPrs field if non-nil, zero value otherwise.

### GetAllPrsOk

`func (o *DbPRInsight) GetAllPrsOk() (*int32, bool)`

GetAllPrsOk returns a tuple with the AllPrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPrs

`func (o *DbPRInsight) SetAllPrs(v int32)`

SetAllPrs sets AllPrs field to given value.


### GetAcceptedPrs

`func (o *DbPRInsight) GetAcceptedPrs() int32`

GetAcceptedPrs returns the AcceptedPrs field if non-nil, zero value otherwise.

### GetAcceptedPrsOk

`func (o *DbPRInsight) GetAcceptedPrsOk() (*int32, bool)`

GetAcceptedPrsOk returns a tuple with the AcceptedPrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedPrs

`func (o *DbPRInsight) SetAcceptedPrs(v int32)`

SetAcceptedPrs sets AcceptedPrs field to given value.


### GetSpamPrs

`func (o *DbPRInsight) GetSpamPrs() int32`

GetSpamPrs returns the SpamPrs field if non-nil, zero value otherwise.

### GetSpamPrsOk

`func (o *DbPRInsight) GetSpamPrsOk() (*int32, bool)`

GetSpamPrsOk returns a tuple with the SpamPrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamPrs

`func (o *DbPRInsight) SetSpamPrs(v int32)`

SetSpamPrs sets SpamPrs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


