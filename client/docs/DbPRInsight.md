# DbPRInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Repository identifier | 
**Interval** | **float32** | Selected interval in numerical days, goes back with number, 0 means today | [default to 0]
**Day** | **float32** | Selected interval computed date in human readable format | 
**AllPrs** | **float32** | PR Type: all requests count | 
**AcceptedPrs** | **float32** | PR Type: accepted requests count | 
**SpamPrs** | **float32** | PR Type: spam requests count | 

## Methods

### NewDbPRInsight

`func NewDbPRInsight(id float32, interval float32, day float32, allPrs float32, acceptedPrs float32, spamPrs float32, ) *DbPRInsight`

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

`func (o *DbPRInsight) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbPRInsight) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbPRInsight) SetId(v float32)`

SetId sets Id field to given value.


### GetInterval

`func (o *DbPRInsight) GetInterval() float32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DbPRInsight) GetIntervalOk() (*float32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DbPRInsight) SetInterval(v float32)`

SetInterval sets Interval field to given value.


### GetDay

`func (o *DbPRInsight) GetDay() float32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *DbPRInsight) GetDayOk() (*float32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *DbPRInsight) SetDay(v float32)`

SetDay sets Day field to given value.


### GetAllPrs

`func (o *DbPRInsight) GetAllPrs() float32`

GetAllPrs returns the AllPrs field if non-nil, zero value otherwise.

### GetAllPrsOk

`func (o *DbPRInsight) GetAllPrsOk() (*float32, bool)`

GetAllPrsOk returns a tuple with the AllPrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPrs

`func (o *DbPRInsight) SetAllPrs(v float32)`

SetAllPrs sets AllPrs field to given value.


### GetAcceptedPrs

`func (o *DbPRInsight) GetAcceptedPrs() float32`

GetAcceptedPrs returns the AcceptedPrs field if non-nil, zero value otherwise.

### GetAcceptedPrsOk

`func (o *DbPRInsight) GetAcceptedPrsOk() (*float32, bool)`

GetAcceptedPrsOk returns a tuple with the AcceptedPrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedPrs

`func (o *DbPRInsight) SetAcceptedPrs(v float32)`

SetAcceptedPrs sets AcceptedPrs field to given value.


### GetSpamPrs

`func (o *DbPRInsight) GetSpamPrs() float32`

GetSpamPrs returns the SpamPrs field if non-nil, zero value otherwise.

### GetSpamPrsOk

`func (o *DbPRInsight) GetSpamPrsOk() (*float32, bool)`

GetSpamPrsOk returns a tuple with the SpamPrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamPrs

`func (o *DbPRInsight) SetSpamPrs(v float32)`

SetSpamPrs sets SpamPrs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


