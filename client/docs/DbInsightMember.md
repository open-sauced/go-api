# DbInsightMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Insight Member identifier | 
**InsightId** | **int32** | Insight ID | 
**UserId** | Pointer to **int32** | User ID | [optional] 
**Name** | Pointer to **string** | User&#39;s Name | [optional] 
**Access** | **string** | Insight Member Access | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing team member creation | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing team member last updated | [optional] 
**DeletedAt** | Pointer to **time.Time** | Timestamp representing team member deletion | [optional] 
**InvitationEmailedAt** | Pointer to **time.Time** | Timestamp representing team member invitation email sent date | [optional] 
**InvitationEmail** | Pointer to **string** | Team Member Invitation Email | [optional] 

## Methods

### NewDbInsightMember

`func NewDbInsightMember(id string, insightId int32, access string, ) *DbInsightMember`

NewDbInsightMember instantiates a new DbInsightMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbInsightMemberWithDefaults

`func NewDbInsightMemberWithDefaults() *DbInsightMember`

NewDbInsightMemberWithDefaults instantiates a new DbInsightMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbInsightMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbInsightMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbInsightMember) SetId(v string)`

SetId sets Id field to given value.


### GetInsightId

`func (o *DbInsightMember) GetInsightId() int32`

GetInsightId returns the InsightId field if non-nil, zero value otherwise.

### GetInsightIdOk

`func (o *DbInsightMember) GetInsightIdOk() (*int32, bool)`

GetInsightIdOk returns a tuple with the InsightId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightId

`func (o *DbInsightMember) SetInsightId(v int32)`

SetInsightId sets InsightId field to given value.


### GetUserId

`func (o *DbInsightMember) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DbInsightMember) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DbInsightMember) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DbInsightMember) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetName

`func (o *DbInsightMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbInsightMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbInsightMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DbInsightMember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccess

`func (o *DbInsightMember) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *DbInsightMember) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *DbInsightMember) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetCreatedAt

`func (o *DbInsightMember) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbInsightMember) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbInsightMember) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbInsightMember) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbInsightMember) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbInsightMember) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbInsightMember) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbInsightMember) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *DbInsightMember) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DbInsightMember) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DbInsightMember) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DbInsightMember) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetInvitationEmailedAt

`func (o *DbInsightMember) GetInvitationEmailedAt() time.Time`

GetInvitationEmailedAt returns the InvitationEmailedAt field if non-nil, zero value otherwise.

### GetInvitationEmailedAtOk

`func (o *DbInsightMember) GetInvitationEmailedAtOk() (*time.Time, bool)`

GetInvitationEmailedAtOk returns a tuple with the InvitationEmailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationEmailedAt

`func (o *DbInsightMember) SetInvitationEmailedAt(v time.Time)`

SetInvitationEmailedAt sets InvitationEmailedAt field to given value.

### HasInvitationEmailedAt

`func (o *DbInsightMember) HasInvitationEmailedAt() bool`

HasInvitationEmailedAt returns a boolean if a field has been set.

### GetInvitationEmail

`func (o *DbInsightMember) GetInvitationEmail() string`

GetInvitationEmail returns the InvitationEmail field if non-nil, zero value otherwise.

### GetInvitationEmailOk

`func (o *DbInsightMember) GetInvitationEmailOk() (*string, bool)`

GetInvitationEmailOk returns a tuple with the InvitationEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationEmail

`func (o *DbInsightMember) SetInvitationEmail(v string)`

SetInvitationEmail sets InvitationEmail field to given value.

### HasInvitationEmail

`func (o *DbInsightMember) HasInvitationEmail() bool`

HasInvitationEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


