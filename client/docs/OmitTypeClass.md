# OmitTypeClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Notification identifier | 
**UserId** | **int32** | User ID | 
**FromUserId** | Pointer to **int32** | From User ID | [optional] 
**Type** | **string** | User notification type | 
**Message** | Pointer to **string** | User notification message | [optional] 
**NotifiedAt** | Pointer to **time.Time** | Timestamp representing db-user-notification creation | [optional] 
**MetaId** | Pointer to **string** | Notification Source ID (highlight, user, invite) | [optional] 

## Methods

### NewOmitTypeClass

`func NewOmitTypeClass(id int32, userId int32, type_ string, ) *OmitTypeClass`

NewOmitTypeClass instantiates a new OmitTypeClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOmitTypeClassWithDefaults

`func NewOmitTypeClassWithDefaults() *OmitTypeClass`

NewOmitTypeClassWithDefaults instantiates a new OmitTypeClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OmitTypeClass) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OmitTypeClass) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OmitTypeClass) SetId(v int32)`

SetId sets Id field to given value.


### GetUserId

`func (o *OmitTypeClass) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OmitTypeClass) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OmitTypeClass) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetFromUserId

`func (o *OmitTypeClass) GetFromUserId() int32`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *OmitTypeClass) GetFromUserIdOk() (*int32, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *OmitTypeClass) SetFromUserId(v int32)`

SetFromUserId sets FromUserId field to given value.

### HasFromUserId

`func (o *OmitTypeClass) HasFromUserId() bool`

HasFromUserId returns a boolean if a field has been set.

### GetType

`func (o *OmitTypeClass) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OmitTypeClass) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OmitTypeClass) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *OmitTypeClass) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OmitTypeClass) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OmitTypeClass) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OmitTypeClass) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNotifiedAt

`func (o *OmitTypeClass) GetNotifiedAt() time.Time`

GetNotifiedAt returns the NotifiedAt field if non-nil, zero value otherwise.

### GetNotifiedAtOk

`func (o *OmitTypeClass) GetNotifiedAtOk() (*time.Time, bool)`

GetNotifiedAtOk returns a tuple with the NotifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedAt

`func (o *OmitTypeClass) SetNotifiedAt(v time.Time)`

SetNotifiedAt sets NotifiedAt field to given value.

### HasNotifiedAt

`func (o *OmitTypeClass) HasNotifiedAt() bool`

HasNotifiedAt returns a boolean if a field has been set.

### GetMetaId

`func (o *OmitTypeClass) GetMetaId() string`

GetMetaId returns the MetaId field if non-nil, zero value otherwise.

### GetMetaIdOk

`func (o *OmitTypeClass) GetMetaIdOk() (*string, bool)`

GetMetaIdOk returns a tuple with the MetaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaId

`func (o *OmitTypeClass) SetMetaId(v string)`

SetMetaId sets MetaId field to given value.

### HasMetaId

`func (o *OmitTypeClass) HasMetaId() bool`

HasMetaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


