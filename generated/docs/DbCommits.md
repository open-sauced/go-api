# DbCommits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Commit identifier | 
**CommitHash** | **string** | Hash for the commit | 
**CommitDate** | **string** | Date for the commit | 
**BakedRepoId** | **float32** | Baked repo identifier | 
**CommitAuthorId** | **float32** | Commit author identifier | 

## Methods

### NewDbCommits

`func NewDbCommits(id float32, commitHash string, commitDate string, bakedRepoId float32, commitAuthorId float32, ) *DbCommits`

NewDbCommits instantiates a new DbCommits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbCommitsWithDefaults

`func NewDbCommitsWithDefaults() *DbCommits`

NewDbCommitsWithDefaults instantiates a new DbCommits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbCommits) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbCommits) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbCommits) SetId(v float32)`

SetId sets Id field to given value.


### GetCommitHash

`func (o *DbCommits) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *DbCommits) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *DbCommits) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.


### GetCommitDate

`func (o *DbCommits) GetCommitDate() string`

GetCommitDate returns the CommitDate field if non-nil, zero value otherwise.

### GetCommitDateOk

`func (o *DbCommits) GetCommitDateOk() (*string, bool)`

GetCommitDateOk returns a tuple with the CommitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitDate

`func (o *DbCommits) SetCommitDate(v string)`

SetCommitDate sets CommitDate field to given value.


### GetBakedRepoId

`func (o *DbCommits) GetBakedRepoId() float32`

GetBakedRepoId returns the BakedRepoId field if non-nil, zero value otherwise.

### GetBakedRepoIdOk

`func (o *DbCommits) GetBakedRepoIdOk() (*float32, bool)`

GetBakedRepoIdOk returns a tuple with the BakedRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBakedRepoId

`func (o *DbCommits) SetBakedRepoId(v float32)`

SetBakedRepoId sets BakedRepoId field to given value.


### GetCommitAuthorId

`func (o *DbCommits) GetCommitAuthorId() float32`

GetCommitAuthorId returns the CommitAuthorId field if non-nil, zero value otherwise.

### GetCommitAuthorIdOk

`func (o *DbCommits) GetCommitAuthorIdOk() (*float32, bool)`

GetCommitAuthorIdOk returns a tuple with the CommitAuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitAuthorId

`func (o *DbCommits) SetCommitAuthorId(v float32)`

SetCommitAuthorId sets CommitAuthorId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


