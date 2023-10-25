# DbUserOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | User organization identifier | 
**UserId** | **int32** | User identifier | 
**OrganizationId** | **int32** | Organization identifier | 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing top repo first index | [optional] 

## Methods

### NewDbUserOrganization

`func NewDbUserOrganization(id int32, userId int32, organizationId int32, ) *DbUserOrganization`

NewDbUserOrganization instantiates a new DbUserOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbUserOrganizationWithDefaults

`func NewDbUserOrganizationWithDefaults() *DbUserOrganization`

NewDbUserOrganizationWithDefaults instantiates a new DbUserOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbUserOrganization) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbUserOrganization) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbUserOrganization) SetId(v int32)`

SetId sets Id field to given value.


### GetUserId

`func (o *DbUserOrganization) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DbUserOrganization) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DbUserOrganization) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetOrganizationId

`func (o *DbUserOrganization) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DbUserOrganization) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DbUserOrganization) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetCreatedAt

`func (o *DbUserOrganization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbUserOrganization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbUserOrganization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbUserOrganization) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


