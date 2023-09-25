# DbFilteredUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **string** | User Login | [optional] 
**FullName** | **string** | Users fullname | 

## Methods

### NewDbFilteredUser

`func NewDbFilteredUser(fullName string, ) *DbFilteredUser`

NewDbFilteredUser instantiates a new DbFilteredUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbFilteredUserWithDefaults

`func NewDbFilteredUserWithDefaults() *DbFilteredUser`

NewDbFilteredUserWithDefaults instantiates a new DbFilteredUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *DbFilteredUser) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *DbFilteredUser) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *DbFilteredUser) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *DbFilteredUser) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetFullName

`func (o *DbFilteredUser) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *DbFilteredUser) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *DbFilteredUser) SetFullName(v string)`

SetFullName sets FullName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


