# VotedRepoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Voted** | **bool** | If the user has voted for the repo | 
**Data** | [**NullableDbRepoToUserVotes**](DbRepoToUserVotes.md) |  | 

## Methods

### NewVotedRepoDto

`func NewVotedRepoDto(voted bool, data NullableDbRepoToUserVotes, ) *VotedRepoDto`

NewVotedRepoDto instantiates a new VotedRepoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVotedRepoDtoWithDefaults

`func NewVotedRepoDtoWithDefaults() *VotedRepoDto`

NewVotedRepoDtoWithDefaults instantiates a new VotedRepoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoted

`func (o *VotedRepoDto) GetVoted() bool`

GetVoted returns the Voted field if non-nil, zero value otherwise.

### GetVotedOk

`func (o *VotedRepoDto) GetVotedOk() (*bool, bool)`

GetVotedOk returns a tuple with the Voted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoted

`func (o *VotedRepoDto) SetVoted(v bool)`

SetVoted sets Voted field to given value.


### GetData

`func (o *VotedRepoDto) GetData() DbRepoToUserVotes`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VotedRepoDto) GetDataOk() (*DbRepoToUserVotes, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VotedRepoDto) SetData(v DbRepoToUserVotes)`

SetData sets Data field to given value.


### SetDataNil

`func (o *VotedRepoDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *VotedRepoDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


