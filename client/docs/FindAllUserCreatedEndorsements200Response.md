# FindAllUserCreatedEndorsements200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbEndorsement**](DbEndorsement.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewFindAllUserCreatedEndorsements200Response

`func NewFindAllUserCreatedEndorsements200Response(data []DbEndorsement, meta PageMetaDto, ) *FindAllUserCreatedEndorsements200Response`

NewFindAllUserCreatedEndorsements200Response instantiates a new FindAllUserCreatedEndorsements200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindAllUserCreatedEndorsements200ResponseWithDefaults

`func NewFindAllUserCreatedEndorsements200ResponseWithDefaults() *FindAllUserCreatedEndorsements200Response`

NewFindAllUserCreatedEndorsements200ResponseWithDefaults instantiates a new FindAllUserCreatedEndorsements200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindAllUserCreatedEndorsements200Response) GetData() []DbEndorsement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindAllUserCreatedEndorsements200Response) GetDataOk() (*[]DbEndorsement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindAllUserCreatedEndorsements200Response) SetData(v []DbEndorsement)`

SetData sets Data field to given value.


### GetMeta

`func (o *FindAllUserCreatedEndorsements200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindAllUserCreatedEndorsements200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindAllUserCreatedEndorsements200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


