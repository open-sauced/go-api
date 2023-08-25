# GenerateCodeRefactorSuggestionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DescriptionLength** | **int32** | Suggestion Length | 
**Temperature** | **int32** | Description Temperature | 
**Language** | **string** | Suggestion Language | [default to "english"]
**Code** | **string** | Code | 

## Methods

### NewGenerateCodeRefactorSuggestionDto

`func NewGenerateCodeRefactorSuggestionDto(descriptionLength int32, temperature int32, language string, code string, ) *GenerateCodeRefactorSuggestionDto`

NewGenerateCodeRefactorSuggestionDto instantiates a new GenerateCodeRefactorSuggestionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCodeRefactorSuggestionDtoWithDefaults

`func NewGenerateCodeRefactorSuggestionDtoWithDefaults() *GenerateCodeRefactorSuggestionDto`

NewGenerateCodeRefactorSuggestionDtoWithDefaults instantiates a new GenerateCodeRefactorSuggestionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptionLength

`func (o *GenerateCodeRefactorSuggestionDto) GetDescriptionLength() int32`

GetDescriptionLength returns the DescriptionLength field if non-nil, zero value otherwise.

### GetDescriptionLengthOk

`func (o *GenerateCodeRefactorSuggestionDto) GetDescriptionLengthOk() (*int32, bool)`

GetDescriptionLengthOk returns a tuple with the DescriptionLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLength

`func (o *GenerateCodeRefactorSuggestionDto) SetDescriptionLength(v int32)`

SetDescriptionLength sets DescriptionLength field to given value.


### GetTemperature

`func (o *GenerateCodeRefactorSuggestionDto) GetTemperature() int32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GenerateCodeRefactorSuggestionDto) GetTemperatureOk() (*int32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GenerateCodeRefactorSuggestionDto) SetTemperature(v int32)`

SetTemperature sets Temperature field to given value.


### GetLanguage

`func (o *GenerateCodeRefactorSuggestionDto) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GenerateCodeRefactorSuggestionDto) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GenerateCodeRefactorSuggestionDto) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetCode

`func (o *GenerateCodeRefactorSuggestionDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GenerateCodeRefactorSuggestionDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GenerateCodeRefactorSuggestionDto) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


