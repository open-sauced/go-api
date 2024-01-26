# DbInsightRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Insight identifier | 
**InsightId** | **int32** | Insight ID | 
**RepoId** | **int32** | Repo ID | 
**FullName** | **string** | Repo Full Name | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing insight repo creation | [optional] 
**DeletedAt** | Pointer to **time.Time** | Timestamp representing insight repo deletion | [optional] 

## Methods

### NewDbInsightRepo

`func NewDbInsightRepo(id int32, insightId int32, repoId int32, fullName string, ) *DbInsightRepo`

NewDbInsightRepo instantiates a new DbInsightRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbInsightRepoWithDefaults

`func NewDbInsightRepoWithDefaults() *DbInsightRepo`

NewDbInsightRepoWithDefaults instantiates a new DbInsightRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbInsightRepo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbInsightRepo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbInsightRepo) SetId(v int32)`

SetId sets Id field to given value.


### GetInsightId

`func (o *DbInsightRepo) GetInsightId() int32`

GetInsightId returns the InsightId field if non-nil, zero value otherwise.

### GetInsightIdOk

`func (o *DbInsightRepo) GetInsightIdOk() (*int32, bool)`

GetInsightIdOk returns a tuple with the InsightId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightId

`func (o *DbInsightRepo) SetInsightId(v int32)`

SetInsightId sets InsightId field to given value.


### GetRepoId

`func (o *DbInsightRepo) GetRepoId() int32`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *DbInsightRepo) GetRepoIdOk() (*int32, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *DbInsightRepo) SetRepoId(v int32)`

SetRepoId sets RepoId field to given value.


### GetFullName

`func (o *DbInsightRepo) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *DbInsightRepo) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *DbInsightRepo) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetCreatedAt

`func (o *DbInsightRepo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbInsightRepo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbInsightRepo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbInsightRepo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *DbInsightRepo) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DbInsightRepo) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DbInsightRepo) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DbInsightRepo) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


