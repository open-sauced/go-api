# DeleteWorkspaceMembersDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]Array**](Array.md) | An array of member objects to delete from the workspace | 

## Methods

### NewDeleteWorkspaceMembersDto

`func NewDeleteWorkspaceMembersDto(members []Array, ) *DeleteWorkspaceMembersDto`

NewDeleteWorkspaceMembersDto instantiates a new DeleteWorkspaceMembersDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteWorkspaceMembersDtoWithDefaults

`func NewDeleteWorkspaceMembersDtoWithDefaults() *DeleteWorkspaceMembersDto`

NewDeleteWorkspaceMembersDtoWithDefaults instantiates a new DeleteWorkspaceMembersDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *DeleteWorkspaceMembersDto) GetMembers() []Array`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *DeleteWorkspaceMembersDto) GetMembersOk() (*[]Array, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *DeleteWorkspaceMembersDto) SetMembers(v []Array)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


