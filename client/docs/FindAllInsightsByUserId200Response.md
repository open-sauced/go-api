# FindAllInsightsByUserId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbInsight**](DbInsight.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewFindAllInsightsByUserId200Response

`func NewFindAllInsightsByUserId200Response(data []DbInsight, meta PageMetaDto, ) *FindAllInsightsByUserId200Response`

NewFindAllInsightsByUserId200Response instantiates a new FindAllInsightsByUserId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindAllInsightsByUserId200ResponseWithDefaults

`func NewFindAllInsightsByUserId200ResponseWithDefaults() *FindAllInsightsByUserId200Response`

NewFindAllInsightsByUserId200ResponseWithDefaults instantiates a new FindAllInsightsByUserId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindAllInsightsByUserId200Response) GetData() []DbInsight`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindAllInsightsByUserId200Response) GetDataOk() (*[]DbInsight, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindAllInsightsByUserId200Response) SetData(v []DbInsight)`

SetData sets Data field to given value.


### GetMeta

`func (o *FindAllInsightsByUserId200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindAllInsightsByUserId200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindAllInsightsByUserId200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


