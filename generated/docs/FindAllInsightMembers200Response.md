# FindAllInsightMembers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbInsightMember**](DbInsightMember.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewFindAllInsightMembers200Response

`func NewFindAllInsightMembers200Response(data []DbInsightMember, meta PageMetaDto, ) *FindAllInsightMembers200Response`

NewFindAllInsightMembers200Response instantiates a new FindAllInsightMembers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindAllInsightMembers200ResponseWithDefaults

`func NewFindAllInsightMembers200ResponseWithDefaults() *FindAllInsightMembers200Response`

NewFindAllInsightMembers200ResponseWithDefaults instantiates a new FindAllInsightMembers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindAllInsightMembers200Response) GetData() []DbInsightMember`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindAllInsightMembers200Response) GetDataOk() (*[]DbInsightMember, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindAllInsightMembers200Response) SetData(v []DbInsightMember)`

SetData sets Data field to given value.


### GetMeta

`func (o *FindAllInsightMembers200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindAllInsightMembers200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindAllInsightMembers200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


