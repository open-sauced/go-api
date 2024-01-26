# UpdateUserEmailPreferencesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayEmail** | **bool** | User Display Public Email | 
**ReceiveCollaboration** | **bool** | User Recieve Collaboration Requests | 
**ReceiveProductUpdates** | **bool** | User Recieve Email Product Updates | 

## Methods

### NewUpdateUserEmailPreferencesDto

`func NewUpdateUserEmailPreferencesDto(displayEmail bool, receiveCollaboration bool, receiveProductUpdates bool, ) *UpdateUserEmailPreferencesDto`

NewUpdateUserEmailPreferencesDto instantiates a new UpdateUserEmailPreferencesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserEmailPreferencesDtoWithDefaults

`func NewUpdateUserEmailPreferencesDtoWithDefaults() *UpdateUserEmailPreferencesDto`

NewUpdateUserEmailPreferencesDtoWithDefaults instantiates a new UpdateUserEmailPreferencesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayEmail

`func (o *UpdateUserEmailPreferencesDto) GetDisplayEmail() bool`

GetDisplayEmail returns the DisplayEmail field if non-nil, zero value otherwise.

### GetDisplayEmailOk

`func (o *UpdateUserEmailPreferencesDto) GetDisplayEmailOk() (*bool, bool)`

GetDisplayEmailOk returns a tuple with the DisplayEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayEmail

`func (o *UpdateUserEmailPreferencesDto) SetDisplayEmail(v bool)`

SetDisplayEmail sets DisplayEmail field to given value.


### GetReceiveCollaboration

`func (o *UpdateUserEmailPreferencesDto) GetReceiveCollaboration() bool`

GetReceiveCollaboration returns the ReceiveCollaboration field if non-nil, zero value otherwise.

### GetReceiveCollaborationOk

`func (o *UpdateUserEmailPreferencesDto) GetReceiveCollaborationOk() (*bool, bool)`

GetReceiveCollaborationOk returns a tuple with the ReceiveCollaboration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveCollaboration

`func (o *UpdateUserEmailPreferencesDto) SetReceiveCollaboration(v bool)`

SetReceiveCollaboration sets ReceiveCollaboration field to given value.


### GetReceiveProductUpdates

`func (o *UpdateUserEmailPreferencesDto) GetReceiveProductUpdates() bool`

GetReceiveProductUpdates returns the ReceiveProductUpdates field if non-nil, zero value otherwise.

### GetReceiveProductUpdatesOk

`func (o *UpdateUserEmailPreferencesDto) GetReceiveProductUpdatesOk() (*bool, bool)`

GetReceiveProductUpdatesOk returns a tuple with the ReceiveProductUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveProductUpdates

`func (o *UpdateUserEmailPreferencesDto) SetReceiveProductUpdates(v bool)`

SetReceiveProductUpdates sets ReceiveProductUpdates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


