# GeneratePullRequestDescriptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DescriptionLength** | **int32** | Description Length | 
**Temperature** | **int32** | Description Temperature | 
**Tone** | **string** | Description Tone | 
**Language** | **string** | Description Language | [default to "english"]
**Diff** | Pointer to **string** | PR Diff | [optional] 
**CommitMessages** | Pointer to **[]string** | PR Commit Messages | [optional] 

## Methods

### NewGeneratePullRequestDescriptionDto

`func NewGeneratePullRequestDescriptionDto(descriptionLength int32, temperature int32, tone string, language string, ) *GeneratePullRequestDescriptionDto`

NewGeneratePullRequestDescriptionDto instantiates a new GeneratePullRequestDescriptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneratePullRequestDescriptionDtoWithDefaults

`func NewGeneratePullRequestDescriptionDtoWithDefaults() *GeneratePullRequestDescriptionDto`

NewGeneratePullRequestDescriptionDtoWithDefaults instantiates a new GeneratePullRequestDescriptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptionLength

`func (o *GeneratePullRequestDescriptionDto) GetDescriptionLength() int32`

GetDescriptionLength returns the DescriptionLength field if non-nil, zero value otherwise.

### GetDescriptionLengthOk

`func (o *GeneratePullRequestDescriptionDto) GetDescriptionLengthOk() (*int32, bool)`

GetDescriptionLengthOk returns a tuple with the DescriptionLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLength

`func (o *GeneratePullRequestDescriptionDto) SetDescriptionLength(v int32)`

SetDescriptionLength sets DescriptionLength field to given value.


### GetTemperature

`func (o *GeneratePullRequestDescriptionDto) GetTemperature() int32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GeneratePullRequestDescriptionDto) GetTemperatureOk() (*int32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GeneratePullRequestDescriptionDto) SetTemperature(v int32)`

SetTemperature sets Temperature field to given value.


### GetTone

`func (o *GeneratePullRequestDescriptionDto) GetTone() string`

GetTone returns the Tone field if non-nil, zero value otherwise.

### GetToneOk

`func (o *GeneratePullRequestDescriptionDto) GetToneOk() (*string, bool)`

GetToneOk returns a tuple with the Tone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTone

`func (o *GeneratePullRequestDescriptionDto) SetTone(v string)`

SetTone sets Tone field to given value.


### GetLanguage

`func (o *GeneratePullRequestDescriptionDto) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GeneratePullRequestDescriptionDto) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GeneratePullRequestDescriptionDto) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetDiff

`func (o *GeneratePullRequestDescriptionDto) GetDiff() string`

GetDiff returns the Diff field if non-nil, zero value otherwise.

### GetDiffOk

`func (o *GeneratePullRequestDescriptionDto) GetDiffOk() (*string, bool)`

GetDiffOk returns a tuple with the Diff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiff

`func (o *GeneratePullRequestDescriptionDto) SetDiff(v string)`

SetDiff sets Diff field to given value.

### HasDiff

`func (o *GeneratePullRequestDescriptionDto) HasDiff() bool`

HasDiff returns a boolean if a field has been set.

### GetCommitMessages

`func (o *GeneratePullRequestDescriptionDto) GetCommitMessages() []string`

GetCommitMessages returns the CommitMessages field if non-nil, zero value otherwise.

### GetCommitMessagesOk

`func (o *GeneratePullRequestDescriptionDto) GetCommitMessagesOk() (*[]string, bool)`

GetCommitMessagesOk returns a tuple with the CommitMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessages

`func (o *GeneratePullRequestDescriptionDto) SetCommitMessages(v []string)`

SetCommitMessages sets CommitMessages field to given value.

### HasCommitMessages

`func (o *GeneratePullRequestDescriptionDto) HasCommitMessages() bool`

HasCommitMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


