# DbWorkspaceContributor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Workspace contributor identifier | 
**ContributorId** | **int32** | Contributor identifier | 
**WorkspaceId** | **string** | Workspace identifier | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing workspace member creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing workspace member last updated | [optional] 
**DeletedAt** | Pointer to **time.Time** | Timestamp representing workspace member deletion | [optional] 

## Methods

### NewDbWorkspaceContributor

`func NewDbWorkspaceContributor(id string, contributorId int32, workspaceId string, ) *DbWorkspaceContributor`

NewDbWorkspaceContributor instantiates a new DbWorkspaceContributor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbWorkspaceContributorWithDefaults

`func NewDbWorkspaceContributorWithDefaults() *DbWorkspaceContributor`

NewDbWorkspaceContributorWithDefaults instantiates a new DbWorkspaceContributor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbWorkspaceContributor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbWorkspaceContributor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbWorkspaceContributor) SetId(v string)`

SetId sets Id field to given value.


### GetContributorId

`func (o *DbWorkspaceContributor) GetContributorId() int32`

GetContributorId returns the ContributorId field if non-nil, zero value otherwise.

### GetContributorIdOk

`func (o *DbWorkspaceContributor) GetContributorIdOk() (*int32, bool)`

GetContributorIdOk returns a tuple with the ContributorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorId

`func (o *DbWorkspaceContributor) SetContributorId(v int32)`

SetContributorId sets ContributorId field to given value.


### GetWorkspaceId

`func (o *DbWorkspaceContributor) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *DbWorkspaceContributor) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *DbWorkspaceContributor) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetCreatedAt

`func (o *DbWorkspaceContributor) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbWorkspaceContributor) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbWorkspaceContributor) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbWorkspaceContributor) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbWorkspaceContributor) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbWorkspaceContributor) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbWorkspaceContributor) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbWorkspaceContributor) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *DbWorkspaceContributor) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DbWorkspaceContributor) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DbWorkspaceContributor) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DbWorkspaceContributor) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


