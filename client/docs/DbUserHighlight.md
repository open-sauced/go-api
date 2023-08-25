# DbUserHighlight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | User Highlight identifier | 
**UserId** | **int32** | User identifier | 
**Url** | Pointer to **string** | Highlight Pull Request URL | [optional] 
**Title** | Pointer to **string** | Highlight Title | [optional] 
**Highlight** | **string** | Highlight Text | 
**Type** | **string** | Highlight type | 
**Pinned** | Pointer to **bool** | Whether the highlight is pinned to the top | [optional] 
**Featured** | Pointer to **bool** | Whether the highlight is featured or not | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp representing highlight creation date | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp representing highlight updated date | [optional] 
**DeletedAt** | Pointer to **time.Time** | Timestamp representing highlight deletion date | [optional] 
**ShippedAt** | Pointer to **time.Time** | Timestamp representing highlight shipped date | [optional] 
**FullName** | Pointer to **string** | Highlight Repo Full Name | [optional] 
**Name** | Pointer to **string** | Highlight User Full Name | [optional] 
**Login** | Pointer to **string** | Highlight User Login | [optional] 
**TaggedRepos** | **[]string** | An array of full-names of tagged repositories | 

## Methods

### NewDbUserHighlight

`func NewDbUserHighlight(id int32, userId int32, highlight string, type_ string, taggedRepos []string, ) *DbUserHighlight`

NewDbUserHighlight instantiates a new DbUserHighlight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbUserHighlightWithDefaults

`func NewDbUserHighlightWithDefaults() *DbUserHighlight`

NewDbUserHighlightWithDefaults instantiates a new DbUserHighlight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbUserHighlight) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbUserHighlight) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbUserHighlight) SetId(v int32)`

SetId sets Id field to given value.


### GetUserId

`func (o *DbUserHighlight) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DbUserHighlight) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DbUserHighlight) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetUrl

`func (o *DbUserHighlight) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DbUserHighlight) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DbUserHighlight) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DbUserHighlight) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTitle

`func (o *DbUserHighlight) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DbUserHighlight) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DbUserHighlight) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DbUserHighlight) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetHighlight

`func (o *DbUserHighlight) GetHighlight() string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *DbUserHighlight) GetHighlightOk() (*string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *DbUserHighlight) SetHighlight(v string)`

SetHighlight sets Highlight field to given value.


### GetType

`func (o *DbUserHighlight) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DbUserHighlight) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DbUserHighlight) SetType(v string)`

SetType sets Type field to given value.


### GetPinned

`func (o *DbUserHighlight) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *DbUserHighlight) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *DbUserHighlight) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *DbUserHighlight) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetFeatured

`func (o *DbUserHighlight) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *DbUserHighlight) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *DbUserHighlight) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *DbUserHighlight) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DbUserHighlight) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbUserHighlight) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbUserHighlight) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbUserHighlight) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbUserHighlight) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbUserHighlight) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbUserHighlight) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbUserHighlight) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *DbUserHighlight) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DbUserHighlight) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DbUserHighlight) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DbUserHighlight) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetShippedAt

`func (o *DbUserHighlight) GetShippedAt() time.Time`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *DbUserHighlight) GetShippedAtOk() (*time.Time, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *DbUserHighlight) SetShippedAt(v time.Time)`

SetShippedAt sets ShippedAt field to given value.

### HasShippedAt

`func (o *DbUserHighlight) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

### GetFullName

`func (o *DbUserHighlight) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *DbUserHighlight) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *DbUserHighlight) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *DbUserHighlight) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetName

`func (o *DbUserHighlight) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbUserHighlight) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbUserHighlight) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DbUserHighlight) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogin

`func (o *DbUserHighlight) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *DbUserHighlight) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *DbUserHighlight) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *DbUserHighlight) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetTaggedRepos

`func (o *DbUserHighlight) GetTaggedRepos() []string`

GetTaggedRepos returns the TaggedRepos field if non-nil, zero value otherwise.

### GetTaggedReposOk

`func (o *DbUserHighlight) GetTaggedReposOk() (*[]string, bool)`

GetTaggedReposOk returns a tuple with the TaggedRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedRepos

`func (o *DbUserHighlight) SetTaggedRepos(v []string)`

SetTaggedRepos sets TaggedRepos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


