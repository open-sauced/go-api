# CreateIssueSummaryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SummaryLength** | **int32** | Summary Length | 
**Temperature** | **int32** | Summary Temperature | 
**Tone** | **string** | Summary Tone | 
**Language** | **string** | Summary Language | [default to "english"]
**IssueTitle** | **string** | Issue Title | 
**IssueDescription** | **string** | Issue Description | 
**IssueComments** | **string** | Issue Comments | 

## Methods

### NewCreateIssueSummaryDto

`func NewCreateIssueSummaryDto(summaryLength int32, temperature int32, tone string, language string, issueTitle string, issueDescription string, issueComments string, ) *CreateIssueSummaryDto`

NewCreateIssueSummaryDto instantiates a new CreateIssueSummaryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIssueSummaryDtoWithDefaults

`func NewCreateIssueSummaryDtoWithDefaults() *CreateIssueSummaryDto`

NewCreateIssueSummaryDtoWithDefaults instantiates a new CreateIssueSummaryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummaryLength

`func (o *CreateIssueSummaryDto) GetSummaryLength() int32`

GetSummaryLength returns the SummaryLength field if non-nil, zero value otherwise.

### GetSummaryLengthOk

`func (o *CreateIssueSummaryDto) GetSummaryLengthOk() (*int32, bool)`

GetSummaryLengthOk returns a tuple with the SummaryLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryLength

`func (o *CreateIssueSummaryDto) SetSummaryLength(v int32)`

SetSummaryLength sets SummaryLength field to given value.


### GetTemperature

`func (o *CreateIssueSummaryDto) GetTemperature() int32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *CreateIssueSummaryDto) GetTemperatureOk() (*int32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *CreateIssueSummaryDto) SetTemperature(v int32)`

SetTemperature sets Temperature field to given value.


### GetTone

`func (o *CreateIssueSummaryDto) GetTone() string`

GetTone returns the Tone field if non-nil, zero value otherwise.

### GetToneOk

`func (o *CreateIssueSummaryDto) GetToneOk() (*string, bool)`

GetToneOk returns a tuple with the Tone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTone

`func (o *CreateIssueSummaryDto) SetTone(v string)`

SetTone sets Tone field to given value.


### GetLanguage

`func (o *CreateIssueSummaryDto) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CreateIssueSummaryDto) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CreateIssueSummaryDto) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetIssueTitle

`func (o *CreateIssueSummaryDto) GetIssueTitle() string`

GetIssueTitle returns the IssueTitle field if non-nil, zero value otherwise.

### GetIssueTitleOk

`func (o *CreateIssueSummaryDto) GetIssueTitleOk() (*string, bool)`

GetIssueTitleOk returns a tuple with the IssueTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTitle

`func (o *CreateIssueSummaryDto) SetIssueTitle(v string)`

SetIssueTitle sets IssueTitle field to given value.


### GetIssueDescription

`func (o *CreateIssueSummaryDto) GetIssueDescription() string`

GetIssueDescription returns the IssueDescription field if non-nil, zero value otherwise.

### GetIssueDescriptionOk

`func (o *CreateIssueSummaryDto) GetIssueDescriptionOk() (*string, bool)`

GetIssueDescriptionOk returns a tuple with the IssueDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDescription

`func (o *CreateIssueSummaryDto) SetIssueDescription(v string)`

SetIssueDescription sets IssueDescription field to given value.


### GetIssueComments

`func (o *CreateIssueSummaryDto) GetIssueComments() string`

GetIssueComments returns the IssueComments field if non-nil, zero value otherwise.

### GetIssueCommentsOk

`func (o *CreateIssueSummaryDto) GetIssueCommentsOk() (*string, bool)`

GetIssueCommentsOk returns a tuple with the IssueComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueComments

`func (o *CreateIssueSummaryDto) SetIssueComments(v string)`

SetIssueComments sets IssueComments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


