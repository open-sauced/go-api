# UpdateWorkspaceMembersDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]interface{}**](Array.md) | An array of member objects and their associated role | 

## Methods

### NewUpdateWorkspaceMembersDto

`func NewUpdateWorkspaceMembersDto(members []interface{}, ) *UpdateWorkspaceMembersDto`

NewUpdateWorkspaceMembersDto instantiates a new UpdateWorkspaceMembersDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWorkspaceMembersDtoWithDefaults

`func NewUpdateWorkspaceMembersDtoWithDefaults() *UpdateWorkspaceMembersDto`

NewUpdateWorkspaceMembersDtoWithDefaults instantiates a new UpdateWorkspaceMembersDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *UpdateWorkspaceMembersDto) GetMembers() []interface{}`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *UpdateWorkspaceMembersDto) GetMembersOk() (*[]interface{}, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *UpdateWorkspaceMembersDto) SetMembers(v []interface{})`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


