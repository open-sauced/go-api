# UserOnboardingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interests** | **[]string** | An array of interests | 
**Timezone** | **string** | User timezone in UTC | 

## Methods

### NewUserOnboardingDto

`func NewUserOnboardingDto(interests []string, timezone string, ) *UserOnboardingDto`

NewUserOnboardingDto instantiates a new UserOnboardingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserOnboardingDtoWithDefaults

`func NewUserOnboardingDtoWithDefaults() *UserOnboardingDto`

NewUserOnboardingDtoWithDefaults instantiates a new UserOnboardingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterests

`func (o *UserOnboardingDto) GetInterests() []string`

GetInterests returns the Interests field if non-nil, zero value otherwise.

### GetInterestsOk

`func (o *UserOnboardingDto) GetInterestsOk() (*[]string, bool)`

GetInterestsOk returns a tuple with the Interests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterests

`func (o *UserOnboardingDto) SetInterests(v []string)`

SetInterests sets Interests field to given value.


### GetTimezone

`func (o *UserOnboardingDto) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserOnboardingDto) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserOnboardingDto) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


