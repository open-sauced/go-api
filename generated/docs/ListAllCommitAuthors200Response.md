# ListAllCommitAuthors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbCommitAuthors**](DbCommitAuthors.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewListAllCommitAuthors200Response

`func NewListAllCommitAuthors200Response(data []DbCommitAuthors, meta PageMetaDto, ) *ListAllCommitAuthors200Response`

NewListAllCommitAuthors200Response instantiates a new ListAllCommitAuthors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllCommitAuthors200ResponseWithDefaults

`func NewListAllCommitAuthors200ResponseWithDefaults() *ListAllCommitAuthors200Response`

NewListAllCommitAuthors200ResponseWithDefaults instantiates a new ListAllCommitAuthors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAllCommitAuthors200Response) GetData() []DbCommitAuthors`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAllCommitAuthors200Response) GetDataOk() (*[]DbCommitAuthors, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAllCommitAuthors200Response) SetData(v []DbCommitAuthors)`

SetData sets Data field to given value.


### GetMeta

`func (o *ListAllCommitAuthors200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAllCommitAuthors200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAllCommitAuthors200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


