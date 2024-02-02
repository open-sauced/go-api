# DeleteWorkspaceReposDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repos** | [**[]interface{}**](Array.md) | An array of repo objects to be added to the workspace | 

## Methods

### NewDeleteWorkspaceReposDto

`func NewDeleteWorkspaceReposDto(repos []interface{}, ) *DeleteWorkspaceReposDto`

NewDeleteWorkspaceReposDto instantiates a new DeleteWorkspaceReposDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteWorkspaceReposDtoWithDefaults

`func NewDeleteWorkspaceReposDtoWithDefaults() *DeleteWorkspaceReposDto`

NewDeleteWorkspaceReposDtoWithDefaults instantiates a new DeleteWorkspaceReposDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepos

`func (o *DeleteWorkspaceReposDto) GetRepos() []interface{}`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *DeleteWorkspaceReposDto) GetReposOk() (*[]interface{}, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *DeleteWorkspaceReposDto) SetRepos(v []interface{})`

SetRepos sets Repos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


