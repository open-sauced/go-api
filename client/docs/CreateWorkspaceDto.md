# CreateWorkspaceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Workspace name | 
**Description** | **string** | Workspace description | 
**Members** | [**[]interface{}**](Array.md) | An array of new member objects | 
**Repos** | [**[]interface{}**](Array.md) | An array of repo objects to be added to the workspace | 

## Methods

### NewCreateWorkspaceDto

`func NewCreateWorkspaceDto(name string, description string, members []interface{}, repos []interface{}, ) *CreateWorkspaceDto`

NewCreateWorkspaceDto instantiates a new CreateWorkspaceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkspaceDtoWithDefaults

`func NewCreateWorkspaceDtoWithDefaults() *CreateWorkspaceDto`

NewCreateWorkspaceDtoWithDefaults instantiates a new CreateWorkspaceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateWorkspaceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWorkspaceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWorkspaceDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateWorkspaceDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateWorkspaceDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateWorkspaceDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMembers

`func (o *CreateWorkspaceDto) GetMembers() []interface{}`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *CreateWorkspaceDto) GetMembersOk() (*[]interface{}, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *CreateWorkspaceDto) SetMembers(v []interface{})`

SetMembers sets Members field to given value.


### GetRepos

`func (o *CreateWorkspaceDto) GetRepos() []interface{}`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *CreateWorkspaceDto) GetReposOk() (*[]interface{}, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *CreateWorkspaceDto) SetRepos(v []interface{})`

SetRepos sets Repos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


