# CreateUserHighlightDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Highlight PR URL | 
**Title** | Pointer to **string** | Highlight Title | [optional] 
**Highlight** | **string** | Highlight Text | 
**ShippedAt** | Pointer to **string** | Shipped Date | [optional] 
**Type** | **string** | Highlight type | [default to "pull_request"]
**TaggedRepos** | **[]string** | An array of full-names of tagged repositories | 

## Methods

### NewCreateUserHighlightDto

`func NewCreateUserHighlightDto(url string, highlight string, type_ string, taggedRepos []string, ) *CreateUserHighlightDto`

NewCreateUserHighlightDto instantiates a new CreateUserHighlightDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserHighlightDtoWithDefaults

`func NewCreateUserHighlightDtoWithDefaults() *CreateUserHighlightDto`

NewCreateUserHighlightDtoWithDefaults instantiates a new CreateUserHighlightDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CreateUserHighlightDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateUserHighlightDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateUserHighlightDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTitle

`func (o *CreateUserHighlightDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateUserHighlightDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateUserHighlightDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateUserHighlightDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetHighlight

`func (o *CreateUserHighlightDto) GetHighlight() string`

GetHighlight returns the Highlight field if non-nil, zero value otherwise.

### GetHighlightOk

`func (o *CreateUserHighlightDto) GetHighlightOk() (*string, bool)`

GetHighlightOk returns a tuple with the Highlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlight

`func (o *CreateUserHighlightDto) SetHighlight(v string)`

SetHighlight sets Highlight field to given value.


### GetShippedAt

`func (o *CreateUserHighlightDto) GetShippedAt() string`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *CreateUserHighlightDto) GetShippedAtOk() (*string, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *CreateUserHighlightDto) SetShippedAt(v string)`

SetShippedAt sets ShippedAt field to given value.

### HasShippedAt

`func (o *CreateUserHighlightDto) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

### GetType

`func (o *CreateUserHighlightDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateUserHighlightDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateUserHighlightDto) SetType(v string)`

SetType sets Type field to given value.


### GetTaggedRepos

`func (o *CreateUserHighlightDto) GetTaggedRepos() []string`

GetTaggedRepos returns the TaggedRepos field if non-nil, zero value otherwise.

### GetTaggedReposOk

`func (o *CreateUserHighlightDto) GetTaggedReposOk() (*[]string, bool)`

GetTaggedReposOk returns a tuple with the TaggedRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedRepos

`func (o *CreateUserHighlightDto) SetTaggedRepos(v []string)`

SetTaggedRepos sets TaggedRepos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


