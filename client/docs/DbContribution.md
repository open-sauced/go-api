# DbContribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Contribution identifier | 
**RepoId** | **int32** | Repository identifier | 
**Count** | **int32** | Total number of contributed pull requests | 
**LastMergedAt** | **time.Time** | Timestamp representing last contribution merge to the default branch | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing contribution creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing contribution last update | [optional] 
**Contributor** | **string** | Contributor unique user name | 
**Url** | **string** | Contribution GitHub origin URL | 

## Methods

### NewDbContribution

`func NewDbContribution(id int32, repoId int32, count int32, lastMergedAt time.Time, contributor string, url string, ) *DbContribution`

NewDbContribution instantiates a new DbContribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbContributionWithDefaults

`func NewDbContributionWithDefaults() *DbContribution`

NewDbContributionWithDefaults instantiates a new DbContribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbContribution) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbContribution) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbContribution) SetId(v int32)`

SetId sets Id field to given value.


### GetRepoId

`func (o *DbContribution) GetRepoId() int32`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *DbContribution) GetRepoIdOk() (*int32, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *DbContribution) SetRepoId(v int32)`

SetRepoId sets RepoId field to given value.


### GetCount

`func (o *DbContribution) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DbContribution) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DbContribution) SetCount(v int32)`

SetCount sets Count field to given value.


### GetLastMergedAt

`func (o *DbContribution) GetLastMergedAt() time.Time`

GetLastMergedAt returns the LastMergedAt field if non-nil, zero value otherwise.

### GetLastMergedAtOk

`func (o *DbContribution) GetLastMergedAtOk() (*time.Time, bool)`

GetLastMergedAtOk returns a tuple with the LastMergedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMergedAt

`func (o *DbContribution) SetLastMergedAt(v time.Time)`

SetLastMergedAt sets LastMergedAt field to given value.


### GetCreatedAt

`func (o *DbContribution) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbContribution) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbContribution) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbContribution) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbContribution) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbContribution) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbContribution) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbContribution) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetContributor

`func (o *DbContribution) GetContributor() string`

GetContributor returns the Contributor field if non-nil, zero value otherwise.

### GetContributorOk

`func (o *DbContribution) GetContributorOk() (*string, bool)`

GetContributorOk returns a tuple with the Contributor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributor

`func (o *DbContribution) SetContributor(v string)`

SetContributor sets Contributor field to given value.


### GetUrl

`func (o *DbContribution) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DbContribution) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DbContribution) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


