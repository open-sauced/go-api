# ListAllCommitsByBakedRepoId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbCommits**](DbCommits.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewListAllCommitsByBakedRepoId200Response

`func NewListAllCommitsByBakedRepoId200Response(data []DbCommits, meta PageMetaDto, ) *ListAllCommitsByBakedRepoId200Response`

NewListAllCommitsByBakedRepoId200Response instantiates a new ListAllCommitsByBakedRepoId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllCommitsByBakedRepoId200ResponseWithDefaults

`func NewListAllCommitsByBakedRepoId200ResponseWithDefaults() *ListAllCommitsByBakedRepoId200Response`

NewListAllCommitsByBakedRepoId200ResponseWithDefaults instantiates a new ListAllCommitsByBakedRepoId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAllCommitsByBakedRepoId200Response) GetData() []DbCommits`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAllCommitsByBakedRepoId200Response) GetDataOk() (*[]DbCommits, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAllCommitsByBakedRepoId200Response) SetData(v []DbCommits)`

SetData sets Data field to given value.


### GetMeta

`func (o *ListAllCommitsByBakedRepoId200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAllCommitsByBakedRepoId200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAllCommitsByBakedRepoId200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


