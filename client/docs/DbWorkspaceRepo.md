# DbWorkspaceRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Workspace repo identifier | 
**RepoId** | **int32** | Repo identifier | 
**WorkspaceId** | **string** | Workspace identifier | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing workspace member creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing workspace member last updated | [optional] 
**DeletedAt** | Pointer to **time.Time** | Timestamp representing workspace member deletion | [optional] 

## Methods

### NewDbWorkspaceRepo

`func NewDbWorkspaceRepo(id string, repoId int32, workspaceId string, ) *DbWorkspaceRepo`

NewDbWorkspaceRepo instantiates a new DbWorkspaceRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbWorkspaceRepoWithDefaults

`func NewDbWorkspaceRepoWithDefaults() *DbWorkspaceRepo`

NewDbWorkspaceRepoWithDefaults instantiates a new DbWorkspaceRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbWorkspaceRepo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbWorkspaceRepo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbWorkspaceRepo) SetId(v string)`

SetId sets Id field to given value.


### GetRepoId

`func (o *DbWorkspaceRepo) GetRepoId() int32`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *DbWorkspaceRepo) GetRepoIdOk() (*int32, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *DbWorkspaceRepo) SetRepoId(v int32)`

SetRepoId sets RepoId field to given value.


### GetWorkspaceId

`func (o *DbWorkspaceRepo) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *DbWorkspaceRepo) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *DbWorkspaceRepo) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetCreatedAt

`func (o *DbWorkspaceRepo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbWorkspaceRepo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbWorkspaceRepo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbWorkspaceRepo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbWorkspaceRepo) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbWorkspaceRepo) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbWorkspaceRepo) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbWorkspaceRepo) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *DbWorkspaceRepo) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DbWorkspaceRepo) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DbWorkspaceRepo) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DbWorkspaceRepo) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


