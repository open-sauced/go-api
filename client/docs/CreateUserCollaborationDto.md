# CreateUserCollaborationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Collaboration Recipient Username | 
**Message** | **string** | Collaboration Request Message | 

## Methods

### NewCreateUserCollaborationDto

`func NewCreateUserCollaborationDto(username string, message string, ) *CreateUserCollaborationDto`

NewCreateUserCollaborationDto instantiates a new CreateUserCollaborationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserCollaborationDtoWithDefaults

`func NewCreateUserCollaborationDtoWithDefaults() *CreateUserCollaborationDto`

NewCreateUserCollaborationDtoWithDefaults instantiates a new CreateUserCollaborationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateUserCollaborationDto) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateUserCollaborationDto) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateUserCollaborationDto) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetMessage

`func (o *CreateUserCollaborationDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateUserCollaborationDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateUserCollaborationDto) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


