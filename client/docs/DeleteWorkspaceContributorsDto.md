# DeleteWorkspaceContributorsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contributors** | [**[]interface{}**](Array.md) | An array of contributor objects to delete from the workspace | 

## Methods

### NewDeleteWorkspaceContributorsDto

`func NewDeleteWorkspaceContributorsDto(contributors []interface{}, ) *DeleteWorkspaceContributorsDto`

NewDeleteWorkspaceContributorsDto instantiates a new DeleteWorkspaceContributorsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteWorkspaceContributorsDtoWithDefaults

`func NewDeleteWorkspaceContributorsDtoWithDefaults() *DeleteWorkspaceContributorsDto`

NewDeleteWorkspaceContributorsDtoWithDefaults instantiates a new DeleteWorkspaceContributorsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContributors

`func (o *DeleteWorkspaceContributorsDto) GetContributors() []interface{}`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *DeleteWorkspaceContributorsDto) GetContributorsOk() (*[]interface{}, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *DeleteWorkspaceContributorsDto) SetContributors(v []interface{})`

SetContributors sets Contributors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


