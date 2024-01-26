# DbWorkspaceOrg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Workspace org identifier | 
**OrgId** | **int32** | Org identifier | 
**WorkspaceId** | **string** | Workspace identifier | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing workspace org creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing workspace org last updated | [optional] 
**DeletedAt** | Pointer to **time.Time** | Timestamp representing workspace org deletion | [optional] 

## Methods

### NewDbWorkspaceOrg

`func NewDbWorkspaceOrg(id string, orgId int32, workspaceId string, ) *DbWorkspaceOrg`

NewDbWorkspaceOrg instantiates a new DbWorkspaceOrg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbWorkspaceOrgWithDefaults

`func NewDbWorkspaceOrgWithDefaults() *DbWorkspaceOrg`

NewDbWorkspaceOrgWithDefaults instantiates a new DbWorkspaceOrg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbWorkspaceOrg) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbWorkspaceOrg) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbWorkspaceOrg) SetId(v string)`

SetId sets Id field to given value.


### GetOrgId

`func (o *DbWorkspaceOrg) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *DbWorkspaceOrg) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *DbWorkspaceOrg) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.


### GetWorkspaceId

`func (o *DbWorkspaceOrg) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *DbWorkspaceOrg) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *DbWorkspaceOrg) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetCreatedAt

`func (o *DbWorkspaceOrg) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbWorkspaceOrg) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbWorkspaceOrg) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbWorkspaceOrg) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbWorkspaceOrg) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbWorkspaceOrg) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbWorkspaceOrg) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbWorkspaceOrg) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *DbWorkspaceOrg) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DbWorkspaceOrg) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DbWorkspaceOrg) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DbWorkspaceOrg) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


