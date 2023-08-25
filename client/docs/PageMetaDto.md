# PageMetaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** | The current page | 
**Limit** | **int32** | The number of items per page | 
**ItemCount** | **int32** | The number of items in the collection | 
**PageCount** | **int32** | The number of pages in the collection | 
**HasPreviousPage** | **bool** | Flag indicating if there is a previous page | 
**HasNextPage** | **bool** | Flag indicating if there is a next page | 

## Methods

### NewPageMetaDto

`func NewPageMetaDto(page int32, limit int32, itemCount int32, pageCount int32, hasPreviousPage bool, hasNextPage bool, ) *PageMetaDto`

NewPageMetaDto instantiates a new PageMetaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageMetaDtoWithDefaults

`func NewPageMetaDtoWithDefaults() *PageMetaDto`

NewPageMetaDtoWithDefaults instantiates a new PageMetaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PageMetaDto) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageMetaDto) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageMetaDto) SetPage(v int32)`

SetPage sets Page field to given value.


### GetLimit

`func (o *PageMetaDto) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PageMetaDto) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PageMetaDto) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetItemCount

`func (o *PageMetaDto) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *PageMetaDto) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *PageMetaDto) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.


### GetPageCount

`func (o *PageMetaDto) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *PageMetaDto) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *PageMetaDto) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.


### GetHasPreviousPage

`func (o *PageMetaDto) GetHasPreviousPage() bool`

GetHasPreviousPage returns the HasPreviousPage field if non-nil, zero value otherwise.

### GetHasPreviousPageOk

`func (o *PageMetaDto) GetHasPreviousPageOk() (*bool, bool)`

GetHasPreviousPageOk returns a tuple with the HasPreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreviousPage

`func (o *PageMetaDto) SetHasPreviousPage(v bool)`

SetHasPreviousPage sets HasPreviousPage field to given value.


### GetHasNextPage

`func (o *PageMetaDto) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *PageMetaDto) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *PageMetaDto) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


