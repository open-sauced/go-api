# DbWorkspaceIssueStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Opened** | **int32** | Number of issues opened in a time range | [default to 0]
**Closed** | **int32** | Number of issues closed in a time range | [default to 0]
**Velocity** | **int32** | Repository average open/close time for issues | [default to 0]

## Methods

### NewDbWorkspaceIssueStats

`func NewDbWorkspaceIssueStats(opened int32, closed int32, velocity int32, ) *DbWorkspaceIssueStats`

NewDbWorkspaceIssueStats instantiates a new DbWorkspaceIssueStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbWorkspaceIssueStatsWithDefaults

`func NewDbWorkspaceIssueStatsWithDefaults() *DbWorkspaceIssueStats`

NewDbWorkspaceIssueStatsWithDefaults instantiates a new DbWorkspaceIssueStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpened

`func (o *DbWorkspaceIssueStats) GetOpened() int32`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *DbWorkspaceIssueStats) GetOpenedOk() (*int32, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *DbWorkspaceIssueStats) SetOpened(v int32)`

SetOpened sets Opened field to given value.


### GetClosed

`func (o *DbWorkspaceIssueStats) GetClosed() int32`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *DbWorkspaceIssueStats) GetClosedOk() (*int32, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *DbWorkspaceIssueStats) SetClosed(v int32)`

SetClosed sets Closed field to given value.


### GetVelocity

`func (o *DbWorkspaceIssueStats) GetVelocity() int32`

GetVelocity returns the Velocity field if non-nil, zero value otherwise.

### GetVelocityOk

`func (o *DbWorkspaceIssueStats) GetVelocityOk() (*int32, bool)`

GetVelocityOk returns a tuple with the Velocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocity

`func (o *DbWorkspaceIssueStats) SetVelocity(v int32)`

SetVelocity sets Velocity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


