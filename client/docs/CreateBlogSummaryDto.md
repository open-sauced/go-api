# CreateBlogSummaryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SummaryLength** | **int32** | Summary Length | 
**Temperature** | **int32** | Summary Temperature | 
**Tone** | **string** | Summary Tone | 
**Language** | **string** | Summary Language | [default to "english"]
**BlogTitle** | **string** | Blog Title | 
**BlogMarkdown** | **string** | Blog Markdown | 

## Methods

### NewCreateBlogSummaryDto

`func NewCreateBlogSummaryDto(summaryLength int32, temperature int32, tone string, language string, blogTitle string, blogMarkdown string, ) *CreateBlogSummaryDto`

NewCreateBlogSummaryDto instantiates a new CreateBlogSummaryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBlogSummaryDtoWithDefaults

`func NewCreateBlogSummaryDtoWithDefaults() *CreateBlogSummaryDto`

NewCreateBlogSummaryDtoWithDefaults instantiates a new CreateBlogSummaryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummaryLength

`func (o *CreateBlogSummaryDto) GetSummaryLength() int32`

GetSummaryLength returns the SummaryLength field if non-nil, zero value otherwise.

### GetSummaryLengthOk

`func (o *CreateBlogSummaryDto) GetSummaryLengthOk() (*int32, bool)`

GetSummaryLengthOk returns a tuple with the SummaryLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryLength

`func (o *CreateBlogSummaryDto) SetSummaryLength(v int32)`

SetSummaryLength sets SummaryLength field to given value.


### GetTemperature

`func (o *CreateBlogSummaryDto) GetTemperature() int32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *CreateBlogSummaryDto) GetTemperatureOk() (*int32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *CreateBlogSummaryDto) SetTemperature(v int32)`

SetTemperature sets Temperature field to given value.


### GetTone

`func (o *CreateBlogSummaryDto) GetTone() string`

GetTone returns the Tone field if non-nil, zero value otherwise.

### GetToneOk

`func (o *CreateBlogSummaryDto) GetToneOk() (*string, bool)`

GetToneOk returns a tuple with the Tone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTone

`func (o *CreateBlogSummaryDto) SetTone(v string)`

SetTone sets Tone field to given value.


### GetLanguage

`func (o *CreateBlogSummaryDto) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CreateBlogSummaryDto) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CreateBlogSummaryDto) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetBlogTitle

`func (o *CreateBlogSummaryDto) GetBlogTitle() string`

GetBlogTitle returns the BlogTitle field if non-nil, zero value otherwise.

### GetBlogTitleOk

`func (o *CreateBlogSummaryDto) GetBlogTitleOk() (*string, bool)`

GetBlogTitleOk returns a tuple with the BlogTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogTitle

`func (o *CreateBlogSummaryDto) SetBlogTitle(v string)`

SetBlogTitle sets BlogTitle field to given value.


### GetBlogMarkdown

`func (o *CreateBlogSummaryDto) GetBlogMarkdown() string`

GetBlogMarkdown returns the BlogMarkdown field if non-nil, zero value otherwise.

### GetBlogMarkdownOk

`func (o *CreateBlogSummaryDto) GetBlogMarkdownOk() (*string, bool)`

GetBlogMarkdownOk returns a tuple with the BlogMarkdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlogMarkdown

`func (o *CreateBlogSummaryDto) SetBlogMarkdown(v string)`

SetBlogMarkdown sets BlogMarkdown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


