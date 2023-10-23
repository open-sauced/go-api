# DbContributorCategoryTimeframe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeStart** | **string** | The ISO timestamp for the start of the time frame | 
**TimeEnd** | **string** | The ISO timestamp for the end of the time frame | 
**All** | **int32** | Number of all contributors (active, new, and alumni) | 
**Active** | **int32** | Number of contributors who contributed in current time frame and previous time frame | 
**New** | **int32** | Number of contributors who are new to contributing (contributed in current time frame but not the previous time frame) | 
**Alumni** | **int32** | Number of contributors who did not contribute in current time frame but did in the previous time frame | 

## Methods

### NewDbContributorCategoryTimeframe

`func NewDbContributorCategoryTimeframe(timeStart string, timeEnd string, all int32, active int32, new int32, alumni int32, ) *DbContributorCategoryTimeframe`

NewDbContributorCategoryTimeframe instantiates a new DbContributorCategoryTimeframe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbContributorCategoryTimeframeWithDefaults

`func NewDbContributorCategoryTimeframeWithDefaults() *DbContributorCategoryTimeframe`

NewDbContributorCategoryTimeframeWithDefaults instantiates a new DbContributorCategoryTimeframe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeStart

`func (o *DbContributorCategoryTimeframe) GetTimeStart() string`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *DbContributorCategoryTimeframe) GetTimeStartOk() (*string, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *DbContributorCategoryTimeframe) SetTimeStart(v string)`

SetTimeStart sets TimeStart field to given value.


### GetTimeEnd

`func (o *DbContributorCategoryTimeframe) GetTimeEnd() string`

GetTimeEnd returns the TimeEnd field if non-nil, zero value otherwise.

### GetTimeEndOk

`func (o *DbContributorCategoryTimeframe) GetTimeEndOk() (*string, bool)`

GetTimeEndOk returns a tuple with the TimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEnd

`func (o *DbContributorCategoryTimeframe) SetTimeEnd(v string)`

SetTimeEnd sets TimeEnd field to given value.


### GetAll

`func (o *DbContributorCategoryTimeframe) GetAll() int32`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *DbContributorCategoryTimeframe) GetAllOk() (*int32, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *DbContributorCategoryTimeframe) SetAll(v int32)`

SetAll sets All field to given value.


### GetActive

`func (o *DbContributorCategoryTimeframe) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DbContributorCategoryTimeframe) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DbContributorCategoryTimeframe) SetActive(v int32)`

SetActive sets Active field to given value.


### GetNew

`func (o *DbContributorCategoryTimeframe) GetNew() int32`

GetNew returns the New field if non-nil, zero value otherwise.

### GetNewOk

`func (o *DbContributorCategoryTimeframe) GetNewOk() (*int32, bool)`

GetNewOk returns a tuple with the New field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNew

`func (o *DbContributorCategoryTimeframe) SetNew(v int32)`

SetNew sets New field to given value.


### GetAlumni

`func (o *DbContributorCategoryTimeframe) GetAlumni() int32`

GetAlumni returns the Alumni field if non-nil, zero value otherwise.

### GetAlumniOk

`func (o *DbContributorCategoryTimeframe) GetAlumniOk() (*int32, bool)`

GetAlumniOk returns a tuple with the Alumni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlumni

`func (o *DbContributorCategoryTimeframe) SetAlumni(v int32)`

SetAlumni sets Alumni field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


