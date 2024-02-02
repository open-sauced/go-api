# DbWorkspacePrStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Opened** | **int32** | Number of PRs opened in a time range | [default to 0]
**Merged** | **int32** | Number of PRs closed in a time range | [default to 0]
**Velocity** | **int32** | Repository average open/close time for PRs | [default to 0]

## Methods

### NewDbWorkspacePrStats

`func NewDbWorkspacePrStats(opened int32, merged int32, velocity int32, ) *DbWorkspacePrStats`

NewDbWorkspacePrStats instantiates a new DbWorkspacePrStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbWorkspacePrStatsWithDefaults

`func NewDbWorkspacePrStatsWithDefaults() *DbWorkspacePrStats`

NewDbWorkspacePrStatsWithDefaults instantiates a new DbWorkspacePrStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpened

`func (o *DbWorkspacePrStats) GetOpened() int32`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *DbWorkspacePrStats) GetOpenedOk() (*int32, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *DbWorkspacePrStats) SetOpened(v int32)`

SetOpened sets Opened field to given value.


### GetMerged

`func (o *DbWorkspacePrStats) GetMerged() int32`

GetMerged returns the Merged field if non-nil, zero value otherwise.

### GetMergedOk

`func (o *DbWorkspacePrStats) GetMergedOk() (*int32, bool)`

GetMergedOk returns a tuple with the Merged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerged

`func (o *DbWorkspacePrStats) SetMerged(v int32)`

SetMerged sets Merged field to given value.


### GetVelocity

`func (o *DbWorkspacePrStats) GetVelocity() int32`

GetVelocity returns the Velocity field if non-nil, zero value otherwise.

### GetVelocityOk

`func (o *DbWorkspacePrStats) GetVelocityOk() (*int32, bool)`

GetVelocityOk returns a tuple with the Velocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocity

`func (o *DbWorkspacePrStats) SetVelocity(v int32)`

SetVelocity sets Velocity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


