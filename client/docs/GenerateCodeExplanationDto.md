# GenerateCodeExplanationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DescriptionLength** | **int32** | Description Length | 
**Temperature** | **int32** | Description Temperature | 
**Language** | **string** | Description Language | [default to "english"]
**Code** | **string** | Code | 

## Methods

### NewGenerateCodeExplanationDto

`func NewGenerateCodeExplanationDto(descriptionLength int32, temperature int32, language string, code string, ) *GenerateCodeExplanationDto`

NewGenerateCodeExplanationDto instantiates a new GenerateCodeExplanationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCodeExplanationDtoWithDefaults

`func NewGenerateCodeExplanationDtoWithDefaults() *GenerateCodeExplanationDto`

NewGenerateCodeExplanationDtoWithDefaults instantiates a new GenerateCodeExplanationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptionLength

`func (o *GenerateCodeExplanationDto) GetDescriptionLength() int32`

GetDescriptionLength returns the DescriptionLength field if non-nil, zero value otherwise.

### GetDescriptionLengthOk

`func (o *GenerateCodeExplanationDto) GetDescriptionLengthOk() (*int32, bool)`

GetDescriptionLengthOk returns a tuple with the DescriptionLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLength

`func (o *GenerateCodeExplanationDto) SetDescriptionLength(v int32)`

SetDescriptionLength sets DescriptionLength field to given value.


### GetTemperature

`func (o *GenerateCodeExplanationDto) GetTemperature() int32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GenerateCodeExplanationDto) GetTemperatureOk() (*int32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GenerateCodeExplanationDto) SetTemperature(v int32)`

SetTemperature sets Temperature field to given value.


### GetLanguage

`func (o *GenerateCodeExplanationDto) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GenerateCodeExplanationDto) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GenerateCodeExplanationDto) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetCode

`func (o *GenerateCodeExplanationDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GenerateCodeExplanationDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GenerateCodeExplanationDto) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


