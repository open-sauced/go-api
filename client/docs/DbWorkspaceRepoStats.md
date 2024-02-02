# DbWorkspaceRepoStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stars** | **int32** | Number of stars for workspace repos | [default to 0]
**Forks** | **int32** | Number of forks for repos | [default to 0]
**Health** | **int32** | Repository average health | [default to 0]

## Methods

### NewDbWorkspaceRepoStats

`func NewDbWorkspaceRepoStats(stars int32, forks int32, health int32, ) *DbWorkspaceRepoStats`

NewDbWorkspaceRepoStats instantiates a new DbWorkspaceRepoStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbWorkspaceRepoStatsWithDefaults

`func NewDbWorkspaceRepoStatsWithDefaults() *DbWorkspaceRepoStats`

NewDbWorkspaceRepoStatsWithDefaults instantiates a new DbWorkspaceRepoStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStars

`func (o *DbWorkspaceRepoStats) GetStars() int32`

GetStars returns the Stars field if non-nil, zero value otherwise.

### GetStarsOk

`func (o *DbWorkspaceRepoStats) GetStarsOk() (*int32, bool)`

GetStarsOk returns a tuple with the Stars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStars

`func (o *DbWorkspaceRepoStats) SetStars(v int32)`

SetStars sets Stars field to given value.


### GetForks

`func (o *DbWorkspaceRepoStats) GetForks() int32`

GetForks returns the Forks field if non-nil, zero value otherwise.

### GetForksOk

`func (o *DbWorkspaceRepoStats) GetForksOk() (*int32, bool)`

GetForksOk returns a tuple with the Forks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForks

`func (o *DbWorkspaceRepoStats) SetForks(v int32)`

SetForks sets Forks field to given value.


### GetHealth

`func (o *DbWorkspaceRepoStats) GetHealth() int32`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *DbWorkspaceRepoStats) GetHealthOk() (*int32, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *DbWorkspaceRepoStats) SetHealth(v int32)`

SetHealth sets Health field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


