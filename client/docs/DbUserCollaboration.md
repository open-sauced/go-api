# DbUserCollaboration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User Collaboration identifier | 
**UserId** | Pointer to **int32** | Collaboration Receipient User ID | [optional] 
**RequestUserId** | Pointer to **int32** | Collaboration Request User ID | [optional] 
**Message** | **string** | Collaboration Request Message | 
**Status** | **string** | Collaboration Status | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing user collaboration creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing user collaboration last updated | [optional] 
**DeletedAt** | Pointer to **time.Time** | Timestamp representing user collaboration deletion | [optional] 
**RequestEmailedAt** | Pointer to **time.Time** | Timestamp representing collaboration request email sent date | [optional] 
**CollaborationEmailedAt** | Pointer to **time.Time** | Timestamp representing collaboration acceptance email sent date | [optional] 

## Methods

### NewDbUserCollaboration

`func NewDbUserCollaboration(id string, message string, status string, ) *DbUserCollaboration`

NewDbUserCollaboration instantiates a new DbUserCollaboration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbUserCollaborationWithDefaults

`func NewDbUserCollaborationWithDefaults() *DbUserCollaboration`

NewDbUserCollaborationWithDefaults instantiates a new DbUserCollaboration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbUserCollaboration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbUserCollaboration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbUserCollaboration) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *DbUserCollaboration) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DbUserCollaboration) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DbUserCollaboration) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DbUserCollaboration) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRequestUserId

`func (o *DbUserCollaboration) GetRequestUserId() int32`

GetRequestUserId returns the RequestUserId field if non-nil, zero value otherwise.

### GetRequestUserIdOk

`func (o *DbUserCollaboration) GetRequestUserIdOk() (*int32, bool)`

GetRequestUserIdOk returns a tuple with the RequestUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUserId

`func (o *DbUserCollaboration) SetRequestUserId(v int32)`

SetRequestUserId sets RequestUserId field to given value.

### HasRequestUserId

`func (o *DbUserCollaboration) HasRequestUserId() bool`

HasRequestUserId returns a boolean if a field has been set.

### GetMessage

`func (o *DbUserCollaboration) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DbUserCollaboration) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DbUserCollaboration) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStatus

`func (o *DbUserCollaboration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DbUserCollaboration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DbUserCollaboration) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *DbUserCollaboration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbUserCollaboration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbUserCollaboration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbUserCollaboration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbUserCollaboration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbUserCollaboration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbUserCollaboration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbUserCollaboration) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *DbUserCollaboration) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DbUserCollaboration) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DbUserCollaboration) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DbUserCollaboration) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetRequestEmailedAt

`func (o *DbUserCollaboration) GetRequestEmailedAt() time.Time`

GetRequestEmailedAt returns the RequestEmailedAt field if non-nil, zero value otherwise.

### GetRequestEmailedAtOk

`func (o *DbUserCollaboration) GetRequestEmailedAtOk() (*time.Time, bool)`

GetRequestEmailedAtOk returns a tuple with the RequestEmailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestEmailedAt

`func (o *DbUserCollaboration) SetRequestEmailedAt(v time.Time)`

SetRequestEmailedAt sets RequestEmailedAt field to given value.

### HasRequestEmailedAt

`func (o *DbUserCollaboration) HasRequestEmailedAt() bool`

HasRequestEmailedAt returns a boolean if a field has been set.

### GetCollaborationEmailedAt

`func (o *DbUserCollaboration) GetCollaborationEmailedAt() time.Time`

GetCollaborationEmailedAt returns the CollaborationEmailedAt field if non-nil, zero value otherwise.

### GetCollaborationEmailedAtOk

`func (o *DbUserCollaboration) GetCollaborationEmailedAtOk() (*time.Time, bool)`

GetCollaborationEmailedAtOk returns a tuple with the CollaborationEmailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborationEmailedAt

`func (o *DbUserCollaboration) SetCollaborationEmailedAt(v time.Time)`

SetCollaborationEmailedAt sets CollaborationEmailedAt field to given value.

### HasCollaborationEmailedAt

`func (o *DbUserCollaboration) HasCollaborationEmailedAt() bool`

HasCollaborationEmailedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


