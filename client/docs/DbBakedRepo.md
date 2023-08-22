# DbBakedRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Baked repository identifier | 
**CloneUrl** | **string** | Repository clone url | 
**RepoId** | **float32** | Repository identifier | 

## Methods

### NewDbBakedRepo

`func NewDbBakedRepo(id float32, cloneUrl string, repoId float32, ) *DbBakedRepo`

NewDbBakedRepo instantiates a new DbBakedRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbBakedRepoWithDefaults

`func NewDbBakedRepoWithDefaults() *DbBakedRepo`

NewDbBakedRepoWithDefaults instantiates a new DbBakedRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbBakedRepo) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbBakedRepo) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbBakedRepo) SetId(v float32)`

SetId sets Id field to given value.


### GetCloneUrl

`func (o *DbBakedRepo) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *DbBakedRepo) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *DbBakedRepo) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.


### GetRepoId

`func (o *DbBakedRepo) GetRepoId() float32`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *DbBakedRepo) GetRepoIdOk() (*float32, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *DbBakedRepo) SetRepoId(v float32)`

SetRepoId sets RepoId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


