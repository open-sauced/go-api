# FindAllOrgsByUsername200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DbUserOrganization**](DbUserOrganization.md) |  | 
**Meta** | [**PageMetaDto**](PageMetaDto.md) |  | 

## Methods

### NewFindAllOrgsByUsername200Response

`func NewFindAllOrgsByUsername200Response(data []DbUserOrganization, meta PageMetaDto, ) *FindAllOrgsByUsername200Response`

NewFindAllOrgsByUsername200Response instantiates a new FindAllOrgsByUsername200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindAllOrgsByUsername200ResponseWithDefaults

`func NewFindAllOrgsByUsername200ResponseWithDefaults() *FindAllOrgsByUsername200Response`

NewFindAllOrgsByUsername200ResponseWithDefaults instantiates a new FindAllOrgsByUsername200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FindAllOrgsByUsername200Response) GetData() []DbUserOrganization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FindAllOrgsByUsername200Response) GetDataOk() (*[]DbUserOrganization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FindAllOrgsByUsername200Response) SetData(v []DbUserOrganization)`

SetData sets Data field to given value.


### GetMeta

`func (o *FindAllOrgsByUsername200Response) GetMeta() PageMetaDto`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FindAllOrgsByUsername200Response) GetMetaOk() (*PageMetaDto, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FindAllOrgsByUsername200Response) SetMeta(v PageMetaDto)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


