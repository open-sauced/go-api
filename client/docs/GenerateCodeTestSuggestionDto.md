# GenerateCodeTestSuggestionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DescriptionLength** | **float32** | Suggestion Length | 
**Temperature** | **float32** | Description Temperature | 
**Code** | **string** | Code | 

## Methods

### NewGenerateCodeTestSuggestionDto

`func NewGenerateCodeTestSuggestionDto(descriptionLength float32, temperature float32, code string, ) *GenerateCodeTestSuggestionDto`

NewGenerateCodeTestSuggestionDto instantiates a new GenerateCodeTestSuggestionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCodeTestSuggestionDtoWithDefaults

`func NewGenerateCodeTestSuggestionDtoWithDefaults() *GenerateCodeTestSuggestionDto`

NewGenerateCodeTestSuggestionDtoWithDefaults instantiates a new GenerateCodeTestSuggestionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptionLength

`func (o *GenerateCodeTestSuggestionDto) GetDescriptionLength() float32`

GetDescriptionLength returns the DescriptionLength field if non-nil, zero value otherwise.

### GetDescriptionLengthOk

`func (o *GenerateCodeTestSuggestionDto) GetDescriptionLengthOk() (*float32, bool)`

GetDescriptionLengthOk returns a tuple with the DescriptionLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLength

`func (o *GenerateCodeTestSuggestionDto) SetDescriptionLength(v float32)`

SetDescriptionLength sets DescriptionLength field to given value.


### GetTemperature

`func (o *GenerateCodeTestSuggestionDto) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GenerateCodeTestSuggestionDto) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GenerateCodeTestSuggestionDto) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.


### GetCode

`func (o *GenerateCodeTestSuggestionDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GenerateCodeTestSuggestionDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GenerateCodeTestSuggestionDto) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


