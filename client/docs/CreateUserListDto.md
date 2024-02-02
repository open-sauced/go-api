# CreateUserListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | List Name | 
**IsPublic** | **bool** | List Visibility | 
**Contributors** | [**[]interface{}**](Array.md) | An array of contributor objects | 

## Methods

### NewCreateUserListDto

`func NewCreateUserListDto(name string, isPublic bool, contributors []interface{}, ) *CreateUserListDto`

NewCreateUserListDto instantiates a new CreateUserListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserListDtoWithDefaults

`func NewCreateUserListDtoWithDefaults() *CreateUserListDto`

NewCreateUserListDtoWithDefaults instantiates a new CreateUserListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateUserListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserListDto) SetName(v string)`

SetName sets Name field to given value.


### GetIsPublic

`func (o *CreateUserListDto) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *CreateUserListDto) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *CreateUserListDto) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetContributors

`func (o *CreateUserListDto) GetContributors() []interface{}`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *CreateUserListDto) GetContributorsOk() (*[]interface{}, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *CreateUserListDto) SetContributors(v []interface{})`

SetContributors sets Contributors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


